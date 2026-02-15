package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// NotionClient handles communication with the Notion API
type NotionClient struct {
	Token      string
	DatabaseID string
}

// PostToNotion sends a message or log to a Notion database
func (c *NotionClient) PostToNotion(content string) error {
	url := "https://api.notion.com/v1/pages"
	
	// Basic structure for a Notion page with a title
	data := map[string]interface{}{
		"parent": map[string]string{"database_id": c.DatabaseID},
		"properties": map[string]interface{}{
			"Name": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"text": map[string]string{"content": content},
					},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Notion-Version", "2022-06-28")
	req.Header.Set("Content-Type", "application/json")

	client := &http.http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notion api error: %s", resp.Status)
	}

	return nil
}
