# [pact.io](https://pact.io) example

## Initial setup

- Only general requisites is `make`, but it can be skipped running
the commands directly.
- Run `git fetch --all --tags`

### Consumer

1. Use LTS `node` version
2. Install deps
    ```bash
    cd consumer
    yarn i
    ```

### Producer

1. Setup all pact.io binaries (you can follow [these steps](https://github.com/pact-foundation/pact-go#installation))
2. Have `go` installed (any `>=1.14` should work)
3. Install modules (may be not needed as it will auto-fetch on first test)
    ```bash
    cd producer
    go get -d -v ./...
    ```

## Testing

From the root directory run `make` to run all the test. Look at the terminal to
see the output from pact.io.

Use the following tags to test various scenarios:
- `git checkout tags/v1-stable`: intial state, where the flow is ok, `v1` is deployed
- `git checkout tags/v2-break`: the consumer changes the contract and we see that the current state makes the producer tests fail, hence `v2` is not deployed
- `git checkout tags/v2-stable`: the producer changes the contracts so that version `v2` can be deployed
