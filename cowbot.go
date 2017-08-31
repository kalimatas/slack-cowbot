package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	clientId     string
	clientSecret string
)

func init() {
	clientId = os.Getenv("CLIENT_ID")
	if "" == clientId {
		panic("CLIENT_ID is not set")
	}

	clientSecret = os.Getenv("CLIENT_SECRET")
	if "" == clientSecret {
		panic("CLIENT_SECRET is not set")
	}
}

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	if "" == r.URL.Query().Get("code") {
		http.Error(w, "No code provided", http.StatusInternalServerError)
		fmt.Println("No code provided")
		return
	}

	req, _ := http.NewRequest("GET", "https://slack.com/api/oauth.access", nil)
	req.URL.Query().Add("code", r.URL.Query().Get("code"))
	req.URL.Query().Add("client_id", clientId)
	req.URL.Query().Add("client_secret", clientSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Cannot perform access request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody, _ := json.Marshal(body)
	fmt.Fprintf(w, string(jsonBody))
}

func cowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, "Got command from Slack!")
	for k, v := range r.Form {
		fmt.Fprintf(w, "param: %s, value: %s\n", k, v[0])
	}
}

func main() {
	http.HandleFunc("/oauth", oauthHandler)
	http.HandleFunc("/", cowHandler)
	log.Fatalln(http.ListenAndServe(":4390", nil))
}
