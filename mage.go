//+build mage

package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/sh"
)

var (
	consumerVersion string = "v2.0.0"
	producerVersion string = "v1.1.0"

	url   string
	token string
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
		consumerVersion,
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
		consumerVersion,
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
		producerVersion,
		"--tag",
		"prod",
		"--broker-base-url",
		url,
		"--broker-token",
		token,
	)
}

// Checks if the consumer can deploy
func ConsumerCanIDeploy() error {
	return sh.RunV(
		"pact-broker",
		"can-i-deploy",
		"--pacticipant",
		"consumer",
		"--version",
		consumerVersion,
		"--to",
		"prod",
		"--broker-base-url",
		url,
		"--broker-token",
		token,
	)
}

// Checks if the producer can deploy
func ProducerCanIDeploy() error {
	return sh.RunV(
		"pact-broker",
		"can-i-deploy",
		"--pacticipant",
		"producer",
		"--version",
		producerVersion,
		"--to",
		"prod",
		"--broker-base-url",
		url,
		"--broker-token",
		token,
	)
}
