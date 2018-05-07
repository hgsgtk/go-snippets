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
	"crypto/tls"
)

var redirectURL = "http://127.0.0.1:18888"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	apiDomain := os.Getenv("API_DOMAIN")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	state := os.Getenv("STATE")

	endPoint := oauth2.Endpoint{
		AuthURL: apiDomain + "/oauth/authorize",
		TokenURL: apiDomain + "/oauth/token",
	}
	conf := &oauth2.Config{
		ClientID: clientID,
		ClientSecret: clientSecret,
		Scopes: []string{"read_users", "read_items", "write_items", "read_orders"},
		Endpoint: endPoint,
		RedirectURL: redirectURL,
	}
	var token *oauth2.Token

	// self-signed invalid certification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx := context.TODO()
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

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
		token, err = conf.Exchange(ctx, <-code)
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

	client := oauth2.NewClient(ctx, conf.TokenSource(ctx, token))


	// GET users/me
	resp, err := client.Get(apiDomain + "/users/me")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//users, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(users))

	// GET items
	resp, err = client.Get(apiDomain + "/items")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//items, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(items))

	// GET items/detail/:item_id
	// todo: need to pass item_id
	resp, err = client.Get(apiDomain + "/items/detail/11057201")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//item, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(item))

	// POST items/add
	//values := url.Values{
	//	"title": {"test item posted by golang api client"},
	//	"detail": {"test item posted by golang api client"},
	//	"price": {"300"},
	//	"stock": {"100"},
	//	"visible": {"0"},
	//	"identifier": {"snNDoWbCC1"},
	//}
	//resp, err = client.PostForm(apiDomain + "/items/add", values)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp.Status)
	//defer resp.Body.Close()
	//addItem, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(addItem))

	// POST /items/delete
	//values = url.Values{
	//	"item_id": {"11151305"}, // todo: need to pass item_id
	//}
	//resp, err = client.PostForm(apiDomain + "/items/delete", values)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp.Status)
	//defer resp.Body.Close()
	//result, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(result))

	// GET /orders
	resp, err = client.Get(apiDomain + "/orders")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	orders, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(orders))

	// GET /orders/detail/:unique_key
	//resp, err = client.Get(apiDomain + "/orders/detail/25EE43F1549E92FB")
	//	//if err != nil {
	//	//	panic(err)
	//	//}
	//	//defer resp.Body.Close()
	//	//order, err := ioutil.ReadAll(resp.Body)
	//	//if err != nil {
	//	//	panic(err)
	//	//}
	//	//fmt.Println(string(order))

	// GET /delivery_companies
	resp, err = client.Get(apiDomain + "/delivery_companies")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//deliv, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(deliv))
}
