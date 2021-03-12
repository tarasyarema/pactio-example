//+build mage

package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/sh"
)

var (
	commit string = "v1.0.0"
	url    string
	token  string
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	url = os.Getenv("PACT_BROKER_BASE_URL")
	token = os.Getenv("PACT_BROKER_TOKEN")

	if url == "" || token == "" {
		panic("Setup pactflow broker env variables: 'PACT_BROKER_BASE_URL' and 'PACT_BROKER_TOKEN'")
	}
}

// Publishes the consumer pacts to pactflow.io
func Publish() error {
	return sh.RunV(
		"pact-broker",
		"publish",
		"consumer/pacts",
		"--consumer-app-version",
		commit,
		"--tag",
		"prod",
		"--broker-base-url",
		url,
		"--broker-token",
		token,
	)
}

// Created a version for the consumer
func ConsumerVersion() error {
	return sh.RunV(
		"pact-broker",
		"create-version-tag",
		"--pacticipant",
		"consumer",
		"--version",
		commit,
		"--tag",
		"prod",
		"--broker-base-url",
		url,
		"--broker-token",
		token,
	)
}

// Created a version for the producer
func ProducerVersion() error {
	return sh.RunV(
		"pact-broker",
		"create-version-tag",
		"--pacticipant",
		"producer",
		"--version",
		commit,
		"--tag",
		"prod",
		"--broker-base-url",
		url,
		"--broker-token",
		token,
	)
}

// Checks if the consumer can deploy
func CanIDeploy() error {
	return sh.RunV(
		"pact-broker",
		"can-i-deploy",
		"--pacticipant",
		"consumer",
		"-l",
		"--broker-base-url",
		url,
		"--broker-token",
		token,
	)
}
