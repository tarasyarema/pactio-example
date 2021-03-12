package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
)

var (
	// these are pactflow related
	version = "v1.2.0"
	tags    = []string{"prod"}

	// Various configuration + test Data for local testing
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

// TestProducerPactflow verify the Provider with pactflow.io pact files
func TestProducerPactflow(t *testing.T) {
	pact := createPact()

	// Map test descriptions to message producer (handlers)
	functionMappings := dsl.MessageHandlers{
		"a contract": func(m dsl.Message) (interface{}, error) {
			return Handler()
		},
	}

	stateMappings := dsl.StateHandlers{}

	_, err := pact.VerifyMessageProvider(t, dsl.VerifyMessageRequest{
		Tags:                       tags,
		BrokerURL:                  url,
		BrokerToken:                token,
		PublishVerificationResults: true,
		ProviderVersion:            version,
		MessageHandlers:            functionMappings,
		StateHandlers:              stateMappings,
	})

	if err != nil {
		t.Fatal(err)
	}
}

// TestProducerLocal test against local contracts
func TestProducerLocal(t *testing.T) {
	pact := createPact()

	// Map test descriptions to message producer (handlers)
	functionMappings := dsl.MessageHandlers{
		"a contract": func(m dsl.Message) (interface{}, error) {
			return Handler()
		},
	}

	stateMappings := dsl.StateHandlers{}

	// Verify the Provider with pactflow.io Pact Files
	_, err := pact.VerifyMessageProvider(t, dsl.VerifyMessageRequest{
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/consumer-producer.json", pactDir))},
		MessageHandlers: functionMappings,
		StateHandlers:   stateMappings,
	})

	if err != nil {
		t.Fatal(err)
	}
}
