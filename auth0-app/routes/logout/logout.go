package logout

import (
	"net/http"
	"os"
	"net/url"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("AUTH0_DOMAIN")

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Url.Path += "/v2/logout"
	params := url.Values{}
	params.Add("returnTo", os.Getenv("http://localhost:3000"))
	params.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	Url.RawQuery = params.Encode()

	http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
}