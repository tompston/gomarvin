// @ts-ignore
import * as gomarvin from "./build/gomarvin.gen.ts";
import {
  assertEquals,
  assert,
  // @ts-ignore
} from "https://deno.land/std@0.177.0/testing/asserts.ts";

// deno test --allow-net ./test/client.ts

const ok_response = { status: 200, message: "", data: {} };
const ok_response_with_links = {
  status: 200,
  message: "",
  data: {},
  links: {},
};

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

// @ts-ignore
Deno.test(
  "Module generate with the crud option works as expected",
  async () => {
    const crud = gomarvin.CrudEndpoints;
    // get_many
    const get_many = await (await crud.GetCruds(client)).json();
    assertEquals(get_many, ok_response_with_links);
    // get_one
    const get_one = await (await crud.GetCrud(client, 1)).json();
    assertEquals(get_one, ok_response);
    // create
    const create = await (await crud.CreateCrud(client)).json();
    assertEquals(create, ok_response);
    // update
    const update = await (await crud.UpdateCrud(client)).json();
    assertEquals(update, ok_response);
    // delete
    const del = await (await crud.DeleteCrud(client)).json();
    assertEquals(del, ok_response);
  }
);

// @ts-ignore
Deno.test("Create user with valid data", async () => {
  const data = await (
    await gomarvin.CreateUser(client, {
      username: "test",
      email: "tom@gmail.com",
      password: "test123123",
    })
  ).json();
  assertEquals(data, ok_response);
});

// @ts-ignore
Deno.test("Create user fails if incorrect email is provided", async () => {
  const data = await (
    await gomarvin.CreateUser(client, {
      username: "test",
      email: "tom@gmail",
      password: "test123123",
    })
  ).json();
  console.log(JSON.stringify(data, null));
  assertEquals(data, {
    status: 400,
    message: "Payload validation failed!",
    data: {
      errors: [
        {
          failed_field: "Email",
          message: "Email must be a valid email address",
        },
      ],
    },
  });
});

// @ts-ignore
Deno.test("Create user fails no body is provided", async () => {
  const data = await // @ts-ignore
  (await gomarvin.CreateUser(client, null)).json();
  console.log(JSON.stringify(data, null));
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
});
