package wunderlist

import (
	"encoding/json"
)

type List struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	OwnerId         string `json:"owner_id"`
	Position        int    `json:"position"`
	Version         int    `json:"version"`
	LocalIdentifier string `json:"local_identifier"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (c *Client) Lists() ([]List, error) {
	data := []List{}
	body, _ := c.GetRequest("me/lists")
	err := json.Unmarshal(body, &data)
	return data, err
}

func (c *Client) AddList(title string) (List, error) {
	data := List{}
	params := map[string]string{"title": title}
	body, _ := c.PostRequest("me/lists", params)
	err := json.Unmarshal(body, &data)
	return data, err
}
