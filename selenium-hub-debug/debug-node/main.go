package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var ID string
var Host string
var Port string

func registerNode() {
	// Define the URL

	url := "http://127.0.0.1:4444/grid/register"

	// Define the JSON payload as a raw string
	jsonPayload := `{
		"class": "org.openqa.grid.common.RegistrationRequest",
		"name": null,
		"description": null,
		"configuration": {
			"hubHost": "127.0.0.1",
			"hubPort": 4444,
			"id": "` + ID + `",
			"capabilities": [
				{
					"seleniumProtocol": "WebDriver",
					"browserName": "chrome",
					"maxInstances": 1,
					"platform": "LINUX"
				}
			],
			"downPollingLimit": 2,
			"hub": "http://127.0.0.1:4444/grid/register",
			"nodePolling": 5000,
			"nodeStatusCheckTimeout": 5000,
			"proxy": "org.openqa.grid.selenium.proxy.DefaultRemoteProxy",
			"register": true,
			"registerCycle": 5000,
			"unregisterIfStillDownAfter": 60000,
			"cleanUpCycle": null,
			"custom": {},
			"host": "` + Host + `",
			"maxSession": 5,
			"servlets": [],
			"withoutServlets": [],
			"browserTimeout": 0,
			"debug": false,
			"jettyMaxThreads": null,
			"log": null,
			"port": ` + Port + `,
			"role": "node",
			"timeout": 1800
		}
	}`

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonPayload)))
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	// Set headers as described in the log
	// req.Header.Set("Content-Type", "text/plain; charset=UTF-8") -> "ok"
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("User-Agent", "Apache-HttpClient/4.5.3 (Java/17.0.1)")
	req.Header.Set("Accept-Encoding", "gzip,deflate")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
}

func sendPingRequest() {
  targetURL := "http://localhost:4444/grid/api/proxy?id=" + ID

	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return
	}

	// Set headers
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("User-Agent", "Apache-HttpClient/4.5.3 (Java/17.0.1)")
	req.Header.Set("Accept-Encoding", "gzip,deflate")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return
	}
	defer resp.Body.Close()

	// Read and log the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}

	log.Printf("Received response: %s", body)
}

// RequestLogger is an HTTP handler that prints all the details of the request it receives
func RequestLogger(w http.ResponseWriter, r *http.Request) {
	// Print HTTP method, URL, and remote address
	fmt.Printf("Received %s request for %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)

  // Log the full query string
	if query := r.URL.Query(); len(query) > 0 {
		fmt.Printf("Query Parameters: %v\n", query)
	}

	// Optionally, you can print headers or other parts of the request
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}

  // Read and log the body
	// Make sure not to do this with large bodies or in production without limiting size
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Body: %s\n", body)

	defer r.Body.Close()

  // Rewind the body if you plan to use it again after this handler (middleware scenario)
	r.Body = io.NopCloser(bytes.NewReader(body))

	// Set the header Content-Type for JSON response
  response := []byte(`{"success": true, "message": "Node registered successfully"}`)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func startServer() {
	server := &http.Server{
		Addr:    "localhost:0",
		Handler: http.HandlerFunc(RequestLogger),
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	// defer listener.Close() // Don't close the listener here

	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("Server started on port %d\n", port)

  go func() {
    log.Fatal(server.Serve(listener))
  }()
  ID = fmt.Sprintf("http://localhost:%d", port)
  Host = "localhost"
  Port = fmt.Sprintf("%d", port)
}

func main() {
  startServer()

  registerNode()

  // Start the ticker to send a GET request every 10 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Start a goroutine to handle periodic GET requests
	go func() {
		for range ticker.C {
			sendPingRequest()
		}
	}()

  select {}
}
