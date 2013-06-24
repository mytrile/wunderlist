package wunderlist

import (
	"encoding/json"
)

type Account struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Email             string `json:"email"`
	Avatar            string `json:"avatar"`
	Token             string `json:"token"`
	EmailConfirmed    bool   `json:"email_confirmed"`
	ConfirmationState string `json:"confirmation_state"`
	Product           string `json:"product"`
	Channel           string `json:"channel"`
	TermsAcceptedAt   string `json:"terms_accepted_at"`
	CreateAt          string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type Settings struct {
	AccountLocale             string `json:"account_locale"`
	Background                string `json:"background"`
	InboxName                 string `json:"inbox"`
	MigratedWunderlistOneUser string `json:"migrated_wunderlist_one_user"`
	Wl1InboxId                string `json:"wl1:inbox_id"`
	WunderlistTimezoneOffset  string `json:"wunderlist_timezone_offset"`
}

type Friend struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	CreateAt  string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ShareContact struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type Share struct {
	Id              string        `json:"id"`
	Type            string        `json:"type"`
	ResourceType    string        `json:"resource_type"`
	ResourceTitle   string        `json:"resource_title"`
	ResourceId      string        `json:"resource_id"`
	LocalIdentifier string        `json:"local_identifier"`
	Sender          *ShareContact `json:"sender"`
	Recipient       *ShareContact `json:"recipient"`
	AcceptedAt      string        `json:"accepted_at"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
}

// Account returns struct with user account information
func (c *Client) Account() (Account, error) {
	data := Account{}
	body, _ := c.GetRequest("me")
	err := json.Unmarshal(body, &data)
	return data, err
}

func (c *Client) Settings() (Settings, error) {
	data := Settings{}
	body, _ := c.GetRequest("me/settings")
	err := json.Unmarshal(body, &data)
	return data, err
}

func (c *Client) Friends() ([]Friend, error) {
	data := []Friend{}
	body, _ := c.GetRequest("me/friends")
	err := json.Unmarshal(body, &data)
	return data, err
}

func (c *Client) Shares() ([]Share, error) {
	data := []Share{}
	body, _ := c.GetRequest("me/shares")
	err := json.Unmarshal(body, &data)
	return data, err
}
