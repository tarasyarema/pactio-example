test-local: test-consumer test-producer-local

test: test-consumer test-producer

test-consumer:
	cd consumer && yarn test

test-producer:
	cd producer && make test

test-producer-local:
	cd producer && make test-local
