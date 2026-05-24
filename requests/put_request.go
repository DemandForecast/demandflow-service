package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
    "github.com/gofiber/fiber/v2"
	"net/http"
)

func PutRequest(url string, data interface{},c *fiber.Ctx) (map[string]interface{}, error) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Authorization", c.Get("Authorization"))
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
    fmt.Println(jsonResponse)
    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("%s", jsonResponse["error"])
    }
  
    return jsonResponse, nil
}



func PatchRequest(url string,c *fiber.Ctx) error {
    request, err := http.NewRequest("PATCH", url, nil)
    if err != nil {
        return err
    }
   // request.Header.Set("Authorization", c.Get("Authorization"))
    client := http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return err
    }
    defer response.Body.Close()
    if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
        return fmt.Errorf("Error: %s", response.Status)
    }
    return nil
}
