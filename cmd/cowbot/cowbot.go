package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"encoding/json"

	"github.com/kalimatas/slack-cowbot"
)

var port string = "80"

type cowResponse struct {
	Type string `json:"response_type"`
	Text string `json:"text"`
}

func init() {
	if "" != os.Getenv("COWSAY_PORT") {
		port = os.Getenv("COWSAY_PORT")
	}
}

func cowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := cowResponse{
		Type: "in_channel",
		Text: slack_cowbot.BuildBalloonWithCow(strings.Split(r.FormValue("text"), "\n")),
	}
	jsonResp, _ := json.Marshal(resp)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func main() {
	http.HandleFunc("/", cowHandler)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
