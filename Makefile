all: test-consumer test-producer

test-consumer:
	cd consumer && yarn test

test-producer:
	cd producer && make test
