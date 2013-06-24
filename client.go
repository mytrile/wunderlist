package wunderlist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	api_base = "https://api.wunderlist.com"
)

type Client struct {
	token  string
	client *http.Client
}

type AuthResponse struct {
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	Email             string            `json:"email"`
	Avatar            string            `json:"avatar"`
	Token             string            `json:"token"`
	EmailConfirmed    string            `json:"email_confirmed"`
	ConfirmationState bool              `json:"confirmation_state"`
	Product           string            `json:"product"`
	Channel           string            `json:"channel"`
	TermsAcceptedAt   string            `json:"terms_accepted_at"`
	CreateAt          string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	Settings          map[string]string `json:"settings"`
}

func login(email, password string) (string, error) {
	data := AuthResponse{}
	values := make(url.Values)
	values.Set("email", email)
	values.Set("password", password)
	loginUrl := fmt.Sprintf("%s/login", api_base)
	r, _ := http.PostForm(loginUrl, values)
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)
	return data.Token, err
}

func NewClient(email, password string) *Client {
	token, _ := login(email, password)
	token = fmt.Sprintf("Bearer %s", token)
	return &Client{token, new(http.Client)}
}

func (c *Client) PostRequest(urlString string, params map[string]string) ([]byte, error) {
	var p string
	for key, value := range params {
		p += fmt.Sprintf("%s=%s&", key, value)
	}
	url := fmt.Sprintf("%s/%s", api_base, urlString)
	req, err := http.NewRequest("POST", url, strings.NewReader(p))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", c.token)
	resp, err := c.client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (c *Client) GetRequest(urlString string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api_base, urlString)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", c.token)
	resp, err := c.client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	if resp.StatusCode != 200 {
		return make([]byte, 0), errors.New(fmt.Sprintf("Got a %v status code on fetch of %v", resp.StatusCode, urlString))
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
