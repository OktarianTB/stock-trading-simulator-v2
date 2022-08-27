package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeGetRequest(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("request resulted in non 200 status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
