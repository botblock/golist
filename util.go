package golist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// All unusual responses from the botblock api!
var UnusualResponses []int = []int{400, 401, 404, 429, 500}

func isUnusualResponse(code int) bool {
	for _, i := range UnusualResponses {
		if i == code {
			return true
		}
	}

	return false
}

// Simply fetches from the botblock api in an easy way!
func Fetch(method string, url string, structure interface{}, body map[string]interface{}) error {
	client := &http.Client{}

	marshalledBody, _ := json.Marshal(body)
	jsonBody := bytes.NewBuffer(marshalledBody)

	req, err := http.NewRequest(method, "https://botblock.org/api"+url, jsonBody)
	_, hasBody := body["bot_id"]

	if hasBody {
		req.Header.Add("Content-Type", "application/json")
	}

	if err != nil {
		return errors.New("UnexpectedError: Failed making a request")
	}

	res, err := client.Do(req)

	if err != nil {
		return errors.New("UnexpectedError: Failed making a request")
	}

	data, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		return errors.New("UnexpectedError: Failed reading request data")
	}

	fmt.Println(string(data), res.StatusCode)
	marshallErr := json.Unmarshal(data, structure)

	if marshallErr != nil {
		return errors.New("UnexpectedError: Failed while marshalling json: " + marshallErr.Error())
	}

	unusual := isUnusualResponse(res.StatusCode)

	if unusual {
		return errors.New("BotblockApiError: Botblock api sent an unusual api response as " + string(data) + " with status code as " + strconv.Itoa(res.StatusCode) + "!")
	}

	return nil
}
