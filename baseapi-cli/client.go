package main

import (
	"github.com/joho/godotenv"
	"os"
	"golang.org/x/oauth2"
	"net/http"
	"io"
	"context"
	"github.com/skratchdot/open-golang/open"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

var redirectURL = "http://127.0.0.1:18888"

var Endpoint = oauth2.Endpoint{
	AuthURL: "https://api.thebase.in/1/oauth/authorize",
	TokenURL: "https://api.thebase.in/1/oauth/token",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	state := os.Getenv("STATE")

	conf := &oauth2.Config{
		ClientID: clientID,
		ClientSecret: clientSecret,
		Scopes: []string{"read_users"},
		Endpoint: Endpoint,
		RedirectURL: redirectURL,
	}
	var token *oauth2.Token

	file, err := os.Open("access_token.json")
	if os.IsNotExist(err) {
		url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)

		code := make(chan string)
		var server *http.Server
		server = &http.Server{
			Addr: ":18888",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				io.WriteString(w, "<html><script>window.open('about:blank', '_self').close()</script></html>")
				w.(http.Flusher).Flush()
				code <- r.URL.Query().Get("code")
				server.Shutdown(context.Background())
			}),
		}
		go server.ListenAndServe()
		open.Start(url)
		token, err = conf.Exchange(oauth2.NoContext, <-code)
		if err != nil {
			panic(err)
		}

		file, err := os.Create("access_token.json")
		if err != nil {
			panic(err)
		}
		json.NewEncoder(file).Encode(token)
	} else if err == nil {
		token = &oauth2.Token{}
		json.NewDecoder(file).Decode(token)
	} else {
		panic(err)
	}
	client := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))

	resp, err := client.Get("https://api.thebase.in/1/users/me")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	users, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(users))

}
