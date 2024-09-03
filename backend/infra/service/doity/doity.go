package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SyncWithDoity synchronizes events with Doity
func SyncWithDoity(eventData map[string]interface{}) error {
	url := "https://api.doity.com.br/v1/events"

	jsonData, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_DOITY_API_KEY")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to sync with Doity: %s", resp.Status)
	}

	return nil
}
