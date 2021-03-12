const {
  Matchers,
  MessageConsumerPact,
  asynchronousBodyHandler,
} = require("@pact-foundation/pact");
const { like } = Matchers;
const path = require("path");
const { handleContract } = require("../consumer")

describe("unit tests", function () {
    it("ok", async () => {
        const response = await handleContract({
            name: 'something',
        })
        expect(response).toBe(true)
    })

    it("bad", async () => {
        const response = await handleContract()
        expect(response).toBe(false)
    })
})

describe("Pact.io test", () => {
  const contractPact = new MessageConsumerPact({
    consumer: "consumer",
    dir: path.resolve(process.cwd(), "pacts"),
    pactfileWriteMode: "update",
    provider: "producer",
    logLevel: "info",
  });

  describe("receive a contract", () => {
    it("handles it", () => {
      return contractPact
        .expectsToReceive("a contract")
        .withContent({
          name: like("2pac"),
        })
        .verify(asynchronousBodyHandler(handleContract));
    });
  });
});
