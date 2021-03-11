package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

var (
	// contracts is a global variable
	contract *Contract

	// Various configuration + test Data
	dir, _  = os.Getwd()
	pactDir = fmt.Sprintf("%s/../consumer/pacts", dir)
	logDir  = fmt.Sprintf("%s/log", dir)
)

// createPact setups the Pact client
func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "prodiver",
		LogDir:   logDir,
	}
}

func TestProvider_Success(t *testing.T) {
	pact := createPact()

	// Map test descriptions to message producer (handlers)
	functionMappings := dsl.MessageHandlers{
		"a contract": func(m dsl.Message) (interface{}, error) {
			if contract != nil {
				return contract, nil
			} else {
				return Contract{}, nil
			}
		},
	}

	stateMappings := dsl.StateHandlers{
		"a contract exists": func(s dsl.State) error {
			contract = &Contract{
				Name: "BIG",
			}
			return nil
		},
	}

	// Verify the Provider with local Pact Files
	pact.VerifyMessageProvider(t, dsl.VerifyMessageRequest{
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/consumer-producer.json", pactDir))},
		MessageHandlers: functionMappings,
		StateHandlers:   stateMappings,
	})
}
