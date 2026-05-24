package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteRequest(url string, queryParams map[string]string,c *fiber.Ctx) (map[string]interface{}, error) {
	// Create a new request with DELETE method
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", c.Get("Authorization"))
	// Set query parameters if any
	q := request.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()
	
	// Perform the request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode the response
	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		return nil, err
	}

	// Check for errors in the response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", jsonResponse["error"])
	}

	return jsonResponse, nil
}