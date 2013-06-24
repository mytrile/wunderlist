package wunderlist

import (
	"encoding/json"
)

type Task struct {
	Id                string `json:"id"`
	Type              string `json:"type"`
	Title             string `json:"title"`
	Note              string `json:"note"`
	DueDate           string `json:"due_date"`
	Version           int    `json:"versions"`
	Position          int    `json:"position"`
	Starred           bool   `json:"starred"`
	RecurrenceType    string `json:"recurren_type"`
	RecurrenceCount   int    `json:"recurrenc_count"`
	RecurringParentId string `json:"recurring_parent_id"`
	TaskId            string `json:"task_id"`
	OwnerId           string `json:"owner_id"`
	ParentId          string `json:"parent_id"`
	ListId            string `json:"list_id"`
	AssigneeId        string `json:"assignee_id"`
	LocalIdentifier   string `json:"local_identifier"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	CompletedAt       string `json:"completed_at"`
	DeletedAt         string `json:"deleted_at"`
	CreatedById       string `json:"created_by_id"`
	UpdatedById       string `json:"updated_at"`
}

func (c *Client) Tasks() ([]Task, error) {
	data := []Task{}
	body, _ := c.GetRequest("me/tasks")
	err := json.Unmarshal(body, &data)
	return data, err
}

func (c *Client) AddTask(list_id, title, starred, due_date string) (Task, error) {
	data := Task{}
	params := map[string]string{"list_id": list_id, "title": title, "starred": starred, "due_date": due_date}
	body, _ := c.PostRequest("me/tasks", params)
	err := json.Unmarshal(body, &data)
	return data, err
}
