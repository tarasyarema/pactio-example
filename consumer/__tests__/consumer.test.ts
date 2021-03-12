import {
  Matchers,
  MessageConsumerPact,
  asynchronousBodyHandler,
} from "@pact-foundation/pact"
import path from "path"
import { Contract, handleContract } from "../consumer"

describe("unit tests", function () {
  it("ok", async () => {
    const contract: Contract =  {
      name: 'some name',
      email: 'name@email.com',
    }

    const response: boolean = await handleContract(contract)
    expect(response).toBeTruthy()
  })

  it("bad", async () => {
    const response = await handleContract({} as Contract)
    expect(response).toBeFalsy()
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
        name: Matchers.like("2pac"),
        email: Matchers.like("tupac@domain.com"),
      })
      .verify(asynchronousBodyHandler(handleContract));
    });
  });
});
