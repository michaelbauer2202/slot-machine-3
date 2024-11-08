package randomorg

import "os"

const endpoint = "https://api.random.org/json-rpc/4/invoke"

var apiKey string

func ReadApiKeyFromEnv() {
	apiKey = os.Getenv("RANDOM_ORG_API_KEY")
}
