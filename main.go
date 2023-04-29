package main

import (
	"net/http"
	"io"
	// "bytes"
	"encoding/json"

	"github.com/charmbracelet/log"
)

type Intro struct {
	Markdown string `json:"markdown"`
}

func getIntros() {
    resp, err := http.Get("http://host.docker.internal:8000/intros")
    if err != nil {
        log.Fatal("Error in retrieving data", "Error:", err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := io.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
	log.Debug("API Response", "Body:", bodyString)

    // Convert response body to Todo struct
	var intros map[string]Intro
    json.Unmarshal(bodyBytes, &intros)
	log.Debug("API Response as a map", "Intros:", intros)
	log.Debug("Just the DevOps position", "DevOps", intros["devops"].Markdown)
    // fmt.Printf("API Response as struct %+v\n", todoStruct)
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Test")

	getIntros()
}