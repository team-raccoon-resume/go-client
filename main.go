package main

import (
	"errors"
	"net/http"
	"io"
	// "bytes"
	"encoding/json"

	"github.com/charmbracelet/log"
)

type Intro struct {
	Markdown string `json:"markdown"`
}

type ResumeClient struct {
	httpClient *http.Client
	endpoint string
}

func NewResumeClient(url string) *ResumeClient {
	return &ResumeClient{
		httpClient: &http.Client{},
		endpoint: url,
	}
}

func (c* ResumeClient) newRequest(requestType, requestPath string) (*http.Request, error) {
	log.Debug("Starting a new request", "Type", requestType, "Path", c.endpoint + requestPath)
	req, err := http.NewRequest(requestType, c.endpoint + requestPath, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// func (c *ResumeClient) RequestWithJsonBody() ( ,error)

func (c *ResumeClient) apiCall(requestType, requestPath string) (*http.Response, error) {
	log.Debug("Starting an api call", "Type", requestType, "Path", requestPath)
	req, err := c.newRequest(requestType, requestPath)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return resp, errors.New("Status code received: " + resp.Status)
	}
	return resp, nil
}

func (c *ResumeClient) GetIntros() (map[string]Intro, error) {
	path := "/intros"

    resp, err := c.apiCall("GET", path)
    if err != nil {
        log.Fatal("Error in retrieving data", "Error:", err)
    }

    defer resp.Body.Close()
    bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

    // Convert response body to Todo struct
	var intros map[string]Intro
    err = json.Unmarshal(bodyBytes, &intros)
	if err != nil {
		return nil, err
	}

	return intros, nil
}


func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Test")

	resumeClient := NewResumeClient("http://host.docker.internal:8000")

	intros, err := resumeClient.GetIntros()
	if err != nil {
        log.Fatal("Error in retrieving data", "Error:", err)
	}
	
	log.Debug("Just the DevOps position", "DevOps", intros["devops"].Markdown)
}