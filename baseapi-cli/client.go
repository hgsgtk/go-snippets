package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var apiDomain string
var redirectURL = "http://127.0.0.1:18888"

func getClient(clientID string, clientSecret string, state string) (*http.Client, error) {
	endPoint := oauth2.Endpoint{
		AuthURL:  apiDomain + "/oauth/authorize",
		TokenURL: apiDomain + "/oauth/token",
	}
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"read_users", "read_items", "write_items", "read_orders", "write_orders"},
		Endpoint:     endPoint,
		RedirectURL:  redirectURL,
	}
	var token *oauth2.Token

	// self-signed invalid certification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx := context.Background()
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
		if err := open.Start(url); err != nil {
			return nil, err
		}
		token, err = conf.Exchange(ctx, <-code)
		if err != nil {
			return nil, err
		}

		file, err := os.Create("access_token.json")
		if err != nil {
			return nil, err
		}
		json.NewEncoder(file).Encode(token)
	} else if err == nil {
		token = &oauth2.Token{}
		json.NewDecoder(file).Decode(token)
	} else {
		return nil, err
	}

	client := oauth2.NewClient(ctx, conf.TokenSource(ctx, token))
	return client, nil
}

func getUser(client *http.Client) {
	resp, err := client.Get(apiDomain + "/users/me")
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

func getItems(client *http.Client) {
	resp, err := client.Get(apiDomain + "/items")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	items, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(items))
}

func getItem(client *http.Client, itemId int) {
	resp, err := client.Get(apiDomain + "/items/detail/" + string(itemId))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	item, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(item))
}

// ex.
// values := url.Values{
//	"title": {"test item posted by golang api client"},
//	"detail": {"test item posted by golang api client"},
//	"price": {"300"},
//	"stock": {"100"},
//	"visible": {"0"},
//	"identifier": {"snNDoWbCC1"},
// }
func postItem(client *http.Client, values url.Values) {
	resp, err := client.PostForm(apiDomain+"/items/add", values)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	addItem, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(addItem))
}

func deleteItem(client *http.Client, itemId int) {
	values := url.Values{
		"item_id": {string(itemId)}, // todo: need to pass item_id
	}
	resp, err := client.PostForm(apiDomain+"/items/delete", values)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

func getOrders(client *http.Client) {
	resp, err := client.Get(apiDomain + "/orders")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	orders, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(orders))
}

func getOrder(client *http.Client, uniqueKey string) {
	resp, err := client.Get(apiDomain + "/orders/detail/" + uniqueKey)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	order, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(order))
}

func getDeliveryCompanies(client *http.Client) {
	resp, err := client.Get(apiDomain + "/delivery_companies")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	deliv, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(deliv))
}

// e.g
// values := url.Values{
// "order_item_id": {"yours"},
// "status": {"dispatched"},
// "add_comment": {"お買い上げ頂き誠にありがとうございました。"},
// "atobarai_status": {"shipping"},
// "delivery_company_id": {"1"},
// "tracking_number": {"12345678901234"},
//}
func postOrderEditStatus(client *http.Client, values url.Values) {
	resp, err := client.PostForm(apiDomain+"/orders/edit_status", values)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	editedOrder, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(editedOrder))
}

type item struct {
	Title string
	Price int
	Stock int
}

// See also https://docs.thebase.in/docs/api/items/add
func addNewItem(client *http.Client, item item) error {
	resp, err := client.PostForm(apiDomain+"/items/add", url.Values{
		"title": {"2021-02-15 メンテナンス明け記念商品"},
		"price": {"5000"},
		"stock": {"10"},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(os.Stdout, string(resb))
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	apiDomain = os.Getenv("API_DOMAIN")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	state := os.Getenv("STATE")

	// oauth2 authenticated client
	client, err := getClient(clientID, clientSecret, state)
	if err != nil {
		_, _ = fmt.Fprint(os.Stdout, "error while get authenticated client")
		os.Exit(1)
	}

	if err := addNewItem(client, item{
		Title: "2021-02-15 メンテナンス明け記念商品",
		Price: 5000,
		Stock: 10,
	}); err != nil {
		_, _ = fmt.Fprint(os.Stdout, "error while add new item")
		os.Exit(1)
	}
}
