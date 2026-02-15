package notion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// PollDatabase checks your Notion database for new "Command" entries
func (c *NotionClient) PollDatabase() {
	fmt.Println("🛰️ PHANE Notion Listener: Online and watching...")
	
	for {
		// This URL queries the database for entries
		url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", c.DatabaseID)
		
		req, _ := http.NewRequest("POST", url, nil)
		req.Header.Set("Authorization", "Bearer "+c.Token)
		req.Header.Set("Notion-Version", "2022-06-28")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Connection error: %v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Logic to parse the results and trigger local skills
		// For example: if a new row says "Run Backup", the Pi executes it.
		
		resp.Body.Close()
		time.Sleep(30 * time.Second) // Poll every 30 seconds to stay within 1GB RAM limits
	}
}
