package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"encoding/json"

	sc "github.com/kalimatas/slack-cowbot"
)

var port string = "80"
var token string

type cowResponse struct {
	Type string `json:"response_type"`
	Text string `json:"text"`
}

func init() {
	token = os.Getenv("COWSAY_TOKEN")
	if "" == token {
		panic("COWSAY_TOKEN is not set!")
	}

	if "" != os.Getenv("PORT") {
		port = os.Getenv("PORT")
	}
}

func cowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if token != r.FormValue("token") {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	replacer := strings.NewReplacer("\r", "")
	text := replacer.Replace(r.FormValue("text"))

	balloonWithCow, err := sc.Cowsay(text)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := cowResponse{
		Type: "in_channel",
		Text: fmt.Sprintf("```%s```", balloonWithCow),
	}
	jsonResp, _ := json.Marshal(resp)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func main() {
	http.HandleFunc("/", cowHandler)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
