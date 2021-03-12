package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
)

var (
	// contracts is a global variable
	contract *Contract

	// Various configuration + test Data
	dir, _  = os.Getwd()
	pactDir = fmt.Sprintf("%s/../consumer/pacts", dir)
	logDir  = fmt.Sprintf("%s/log", dir)

	// init related
	url   string
	token string
)

func init() {
	err := godotenv.Load("../.env")

	if err != nil {
		panic("Error loading .env file")
	}

	url = os.Getenv("PACT_BROKER_BASE_URL")
	token = os.Getenv("PACT_BROKER_TOKEN")

	if url == "" || token == "" {
		panic("Setup pactflow broker env variables: 'PACT_BROKER_BASE_URL' and 'PACT_BROKER_TOKEN'")
	}
}

// createPact setups the Pact client
func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "producer",
		LogDir:   logDir,
	}
}

func TestProvider_Success(t *testing.T) {
	pact := createPact()

	// Map test descriptions to message producer (handlers)
	functionMappings := dsl.MessageHandlers{
		"a contract": func(m dsl.Message) (interface{}, error) {
			return &Contract{
				Name:  stringToPtr("2pac Shakur"),
				Email: stringToPtr("tupac@shakur.com"),
			}, nil
		},
	}

	stateMappings := dsl.StateHandlers{}

	// Verify the Provider with pactflow.io Pact Files
	_, err := pact.VerifyMessageProvider(t, dsl.VerifyMessageRequest{
		// Use this if you want to test without the Pact Broker
		// PactURLs:                   []string{filepath.ToSlash(fmt.Sprintf("%s/consumer-producer.json", pactDir))},
		Tags:                       []string{"prod"},
		BrokerURL:                  url,
		BrokerToken:                token,
		PublishVerificationResults: true,
		ProviderVersion:            "v1.0.0",
		MessageHandlers:            functionMappings,
		StateHandlers:              stateMappings,
	})

	if err != nil {
		t.Fatal(err)
	}
}
