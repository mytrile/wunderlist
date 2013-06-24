package wunderlist

import (
	"encoding/json"
)

type Reminder struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	Date            string `json:"date"`
	TaskId          string `json:"task_id"`
	UserId          string `json:"user_id"`
	LocalIdentifier string `json:"local_identifier"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (c *Client) Reminders() ([]Reminder, error) {
	data := []Reminder{}
	body, _ := c.GetRequest("me/reminders")
	err := json.Unmarshal(body, &data)
	return data, err
}

func (c *Client) AddReminder(task_id, date string) (Reminder, error) {
	data := Reminder{}
	params := map[string]string{"task_id": task_id, "date": date}
	body, _ := c.PostRequest("me/reminders", params)
	err := json.Unmarshal(body, &data)
	return data, err
}
