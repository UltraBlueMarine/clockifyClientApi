package logic

import (
	"fmt"
)

type Clients struct {
	clients []ClientData
}

type ClientData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	WorkspaceID string `json:"workspaceID"`
	Note        string `json:"note"`
	Archived    bool   `json:"archived"`
	Address     string `json:"address"`
}

type InsertedData struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

type UpdatedData struct {
	Archived bool   `json:"archived"`
	Name     string `json:"name"`
	Note     string `json:"note"`
}

func (d ClientData) Info() string {
	return fmt.Sprintf("ID: %s\nWorkspaceID: %s\nName: %s\nNote: %s\n", d.ID, d.WorkspaceID, d.Name, d.Note)
}
