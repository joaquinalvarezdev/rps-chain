import { config } from "dotenv";
import _ from "../environment";
import { RPSStargateClient } from "../../client/src/rps_stargateclient";
import { RPSExtension } from "../../client/src/modules/rps/queries";
import { describe, beforeAll, it, expect } from "@jest/globals";

config();

describe("Game", function () {
  let client: RPSStargateClient, rps: RPSExtension["rps"];

  beforeAll(async function () {
    client = await RPSStargateClient.connect(process.env.RPC_URL);
    rps = client.rpsQueryClient!.rps;
  });

  it("cannot get params", async function () {
    const params = await rps.getParams();
    expect(params?.ttl.toNumber()).toEqual(20);
  });
});