package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PostRequest(url string, data interface{},c fiber.Ctx) (map[string]interface{}, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Authorization", c.Get("Authorization"))
    request.Header.Set("organizationId", c.Get("organizationId"))
    client := http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()
    var jsonResponse map[string]interface{}
    err = json.NewDecoder(response.Body).Decode(&jsonResponse)
    if err != nil {
        return nil, err
    }
    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("%s", jsonResponse["error"])
    }
  
    return jsonResponse, nil
}
