package randomorg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RandomIntegers(n int, min int, max int) []int {
	if apiKey == "" {
		ReadApiKeyFromEnv()
	}

	request := &request{Key: apiKey, N: n, Min: min, Max: max}
	reqBytes, _ := json.Marshal(request)
	reqTemplate := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "generateIntegers",
		"params": %s,
		"id": 42
		}`, string(reqBytes))

	resp, _ := http.Post(
		endpoint,
		"application/json",
		bytes.NewBuffer([]byte(reqTemplate)))

	respBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	response := &Response{}
	json.Unmarshal(respBytes, response)

	return response.Ints()
}
