package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBaseUrl = "https://f-api.ir/%s"

func GetARandomFact() (*Fact, error) {
	url := fmt.Sprintf(apiBaseUrl, "api/facts/random")
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s , Something happened on fetching data.", response.Status)
	}

	bodyBytes, readBytesError := io.ReadAll(response.Body)

	if readBytesError != nil {
		return nil, readBytesError
	}

	var fetchedFact Fact
	decodeError := json.Unmarshal(bodyBytes, &fetchedFact)

	if decodeError != nil {
		return nil, decodeError
	}

	return &fetchedFact, nil
}
