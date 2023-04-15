// @ts-ignore
import * as gomarvin from "./build/gomarvin.gen.ts";
import {
  assertEquals,
  assert,
} from "https://deno.land/std@0.177.0/testing/asserts.ts";

// deno test --allow-net ./test/client.ts

const client = gomarvin.defaultClient;

// @ts-ignore
Deno.test(
  "api return empty links object when pagination setting is true",
  async () => {
    const data = await (await gomarvin.GetUsers(client)).json();
    assertEquals(data, {
      status: 200,
      message: "",
      data: {},
      links: {},
    });
  }
);

// @ts-ignore
Deno.test(
  "api return err when empty body is passed to create user",
  async () => {
    // @ts-ignore
    const data = await (await gomarvin.CreateUser(client, null)).json();
    // console.log(JSON.stringify(data, null));
    assertEquals(data, {
      status: 400,
      message: "Payload validation failed!",
      data: {
        errors: [
          { failed_field: "Username", message: "Username is required" },
          { failed_field: "Email", message: "Email is required" },
          { failed_field: "Password", message: "Password is required" },
        ],
      },
    });
  }
);
