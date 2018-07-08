package main

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"
	"net/url"

	plus "google.golang.org/api/plus/v1"

	"golang.org/x/oauth2"

	uuid "github.com/satori/go.uuid"
)

const (
	defaultSessionID     = "default"
	oauthTokenSessionKey = "oauth_token"

	googleProfileSessionKey = "google_profile"
	oauthFlowRedirectKey    = "redirect"
)

func init() {
	// Gob encoding for gorilla/sessions
	gob.Register(&oauth2.Token{})
	gob.Register(&Profile{})
}

type Profile struct {
	ID, DisplayName, ImageURL string
}

func loginHandler(w http.ResponseWriter, r *http.Request) *appError {
	sessionID := uuid.Must(uuid.NewV4()).String()

	oauthFlowSession, err := SessionStore.New(r, sessionID)
	if err != nil {
		return appErrorf(err, "could not create oauth sesson: %v", err)
	}
	oauthFlowSession.Options.MaxAge = 10 * 60 // 10 minutes

	redirectURL, err := validateRedirectURL(r.FormValue("redirect"))
	if err != nil {
		return appErrorf(err, "invalid redirect URL: %v", err)
	}
	oauthFlowSession.Values[oauthFlowRedirectKey] = redirectURL

	if err := oauthFlowSession.Save(r, w); err != nil {
		return appErrorf(err, "could not save session: %v", err)
	}

	// Use the sessionID for the "state" parameter
	url := OAuthConfig.AuthCodeURL(sessionID, oauth2.ApprovalForce, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}

func validateRedirectURL(path string) (string, error) {
	if path == "" {
		return "/", nil
	}

	parsedURL, err := url.Parse(path)
	if err != nil {
		return "/", err
	}
	if parsedURL.IsAbs() {
		return "/", errors.New("URL must not be absolute")
	}
	return path, nil
}

func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) *appError {
	oauthFlowSession, err := SessionStore.Get(r, r.FormValue("state"))
	if err != nil {
		return appErrorf(err, "invalid state parameter. try logging in again.")
	}

	redirectURL, ok := oauthFlowSession.Values[oauthFlowRedirectKey].(string)
	if !ok {
		return appErrorf(err, "invalid state parameter. try logging in again.")
	}

	code := r.FormValue("code")
	tok, err := OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return appErrorf(err, "could not get auth token: %v", err)
	}

	session, err := SessionStore.New(r, defaultSessionID)
	if err != nil {
		return appErrorf(err, "could not get default session: %v", err)
	}

	ctx := context.Background()
	profile, err := fetchProfile(ctx, tok)
	if err != nil {
		return appErrorf(err, "could not fetch Google profile: %v", err)
	}

	session.Values[oauthTokenSessionKey] = tok
	session.Values[googleProfileSessionKey] = stripProfile(profile)
	if err := session.Save(r, w); err != nil {
		return appErrorf(err, "could not save session: %v", err)
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
	return nil
}

func fetchProfile(ctx context.Context, tok *oauth2.Token) (*plus.Person, error) {
	client := oauth2.NewClient(ctx, OAuthConfig.TokenSource(ctx, tok))
	plusService, err := plus.New(client)
	if err != nil {
		return nil, err
	}
	return plusService.People.Get("me").Do()
}

func logoutHandler(w http.ResponseWriter, r *http.Request) *appError {
	session, err := SessionStore.New(r, defaultSessionID)
	if err != nil {
		return appErrorf(err, "could not get default session: %v", err)
	}
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		return appErrorf(err, "could not save session: %v", err)
	}
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/"
	}
	http.Redirect(w, r, redirectURL, http.StatusFound)
	return nil
}

func profileFromSession(r *http.Request) *Profile {
	session, err := SessionStore.Get(r, defaultSessionID)
	if err != nil {
		return nil
	}
	tok, ok := session.Values[oauthTokenSessionKey].(*oauth2.Token)
	if !ok || !tok.Valid() {
		return nil
	}
	profile, ok := session.Values[googleProfileSessionKey].(*Profile)
	if !ok {
		return nil
	}
	return profile
}

func stripProfile(p *plus.Person) *Profile {
	return &Profile{
		ID:          p.Id,
		DisplayName: p.DisplayName,
		ImageURL:    p.Image.Url,
	}
}
