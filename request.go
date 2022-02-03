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

	// webex Integration API Key
	apiKey := os.Getenv("WEBEX_API_KEY")

	// webex target meeting number
	meetingNum := os.Getenv("WEBEX_MEETING_NUMBER")

	// make a GET request to a web server
	// create a new request
	req, err := http.NewRequest("GET", "https://webexapis.com/v1/meetings?meetingNumber=" + meetingNum, nil)
	if err != nil {
		fmt.Println(err)
	}

	// add a header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + apiKey)

	// send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// get body["items"][0]["id"]
	var bodyMap map[string]interface{}
	json.Unmarshal(body, &bodyMap)
	meetingId := bodyMap["items"].([]interface{})[0].(map[string]interface{})["id"].(string)

	// close the response
	res.Body.Close()


	// make a GET request to a web server
	// create a new request
	req, err = http.NewRequest("GET", "https://webexapis.com/v1/meetingParticipants?meetingId=" + meetingId, nil)

	// add a header
	req.Header.Add("Authorization", "Bearer " + apiKey)

	// send the request
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// read the response
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

	// close the response
	res.Body.Close()
}