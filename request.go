package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)


func main() {
	// get the API key from the .env file
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        fmt.Println(err)
    }

	// Webex Integration API Key
	apiKey := os.Getenv("WEBEX_API_KEY")

	// Webex target meeting number
	meetingNum := os.Getenv("WEBEX_MEETING_NUMBER")

	// Make a GET request to a web server
	// body
	// Content-Type: application/json
	// Authorization: Bearer MjZkYzYwNmUtM2UxZC00YTVlLTk2MDctMWFkNWViMGE1MTcxYzQ0Y2I2MzYtOTEz_PF84_40c5c1ec-675d-4892-8f53-cf458f994ae6
	// target
	// https://webexapis.com/v1/meetings?meetingNumber=789679268

	// Create a new request
	req, err := http.NewRequest("GET", "https://webexapis.com/v1/meetings?meetingNumber=" + meetingNum, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Add a header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + apiKey)

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// Read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// get body["items"][0]["id"]
	var bodyMap map[string]interface{}
	json.Unmarshal(body, &bodyMap)
	meetingId := bodyMap["items"].([]interface{})[0].(map[string]interface{})["id"].(string)

	// Close the response
	res.Body.Close()


	// Make a GET request to a web server
	// body
	// Content-Type: application/json
	// meetingId = meetingId
	// Authorization: Bearer MjZkYzYwNmUtM2UxZC00YTVlLTk2MDctMWFkNWViMGE1MTcxYzQ0Y2I2MzYtOTEz_PF84_40c5c1ec-675d-4892-8f53-cf458f994ae6
	// target
	// https://webexapis.com/v1/meetings?meetingNumber=789679268

	// Create a new request
	req, err = http.NewRequest("GET", "https://webexapis.com/v1/meetingParticipants?meetingId=" + meetingId, nil)

	// Add a header
	req.Header.Add("Authorization", "Bearer " + apiKey)

	// Send the request
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// Read the response
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// get body["items"][i]["email"] and print
	var bodyMap2 map[string]interface{}
	json.Unmarshal(body, &bodyMap2)
	for _, item := range bodyMap2["items"].([]interface{}) {
		email := item.(map[string]interface{})["email"].(string)
		fmt.Println(email)
	}

	// Close the response
	res.Body.Close()
}