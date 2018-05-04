package main

import (
	"github.com/garyburd/go-oauth/oauth"
	"os"
)
func main() {
	oauthClient := &oauth.Client{
		Credentials: oauth.Credentials{
			Token: os.Getenv("")
		}
	}
	url := "https://www.hatena.com/oauth/initiate"

}
