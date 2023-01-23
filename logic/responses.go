package logic

import "fmt"

type AssetsResponses struct {
	Assets []AssetData
}

type AssetData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	WorkspaceID string `json:"workspaceID"`
	Note        string `json:"note"`
	Archived    string `json:"archived"`
	Address     string `json:"address"`
}

func (d AssetData) Info() string {
	return fmt.Sprintf("ID: %s\nWorkspaceID: %s\nName: %s\nNote: %s\n", d.ID, d.WorkspaceID, d.Name, d.Note)
}
