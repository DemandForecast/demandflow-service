package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2"
)

func GetAllRequest(url string, target interface{},c fiber.Ctx) error {
	// Create a new HTTP client with a custom header
	client := &http.Client{}

	// Create a new request with the provided URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// // Set the Authorization header here
	// req.Header.Set("Authorization", c.Get("Authorization"))
	req.Header.Set("Authorization", c.Get("Authorization"))
	req.Header.Set("organizationId", "COM-001")
	// Perform the request with the custom client and request
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check the status code
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error: %s", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response body into the target interface
	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return nil
}


func GetRequest(url string, target interface{} , c fiber.Ctx) error {
	// Create a new HTTP client with a custom header
	client := &http.Client{}

	// Create a new request with the provided URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set your custom header here
	 req.Header.Set("Authorization", c.Get("Authorization"))
	 req.Header.Set("organizationId", c.Get("organizationId"))
	// Perform the request with the custom client and request
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check the status code
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error: %s", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response body into the target interface
	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}



func GetAllRequestFromWMS(url string, target interface{},c fiber.Ctx) error {
	// Create a new HTTP client with a custom header
	client := &http.Client{}

	// Create a new request with the provided URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// // Set the Authorization header here
	// req.Header.Set("Authorization", c.Get("Authorization"))
	req.Header.Set("Authorization", c.Get("Authorization"))
	req.Header.Set("companyId", c.Get("organizationId"))
	// Perform the request with the custom client and request
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check the status code
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error: %s", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response body into the target interface
	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return nil
}