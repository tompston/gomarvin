// @ts-ignore
import * as gomarvin from "./build/gomarvin.ts";
import {
  assertEquals,
  assert,
  // @ts-ignore
} from "https://deno.land/std@0.177.0/testing/asserts.ts";

// deno test --allow-net ./test/client.ts

// @ts-ignore
// Deno.test("name of the test", async () => {});

const _t = gomarvin.TestEndpoints;
const response_ok = { status: 200, message: "", data: {} };
const response_ok_with_links = {
  status: 200,
  message: "",
  data: {},
  links: {},
};

const client = gomarvin.defaultClient;

const no_headers_client: gomarvin.Client = {
  host_url: "http://localhost:4444",
  api_prefix: "/api/v1",
  headers: {},
};

// @ts-ignore
Deno.test("Home view works", async () => {
  const data = await (await fetch("http://localhost:4444/")).json();
  assertEquals(data, {
    status: 200,
    message: "Hello There!",
    data: null,
  });
});

// @ts-ignore
Deno.test(
  "api return err when empty body is passed to create user",
  async () => {
    // @ts-ignore
    const data = await (await gomarvin.CreateUser(client, null)).json();
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
    assertEquals(get_many, response_ok_with_links);
    // get_one
    const get_one = await (await crud.GetCrud(client, 1)).json();
    assertEquals(get_one, response_ok);
    // create
    const create = await (await crud.CreateCrud(client)).json();
    assertEquals(create, response_ok);
    // update
    const update = await (await crud.UpdateCrud(client)).json();
    assertEquals(update, response_ok);
    // delete
    const del = await (await crud.DeleteCrud(client)).json();
    assertEquals(del, response_ok);
  }
);

// @ts-ignore
Deno.test("Create user with valid data returns status ok", async () => {
  const data = await (
    await gomarvin.CreateUser(client, {
      username: "test",
      email: "tom@gmail.com",
      password: "test123123",
    })
  ).json();
  assertEquals(data, response_ok);
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

// @ts-ignore
Deno.test(
  "Create user fails if password does not pass validations",
  async () => {
    const data = await (
      await gomarvin.CreateUser(client, {
        username: "test",
        email: "qwe@qwe.com",
        password: "short",
      })
    ).json();
    assertEquals(data, {
      status: 400,
      message: "Payload validation failed!",
      data: {
        errors: [
          {
            failed_field: "Password",
            message: "Password should be at least 10 characters long",
          },
        ],
      },
    });
  }
);

// @ts-ignore
Deno.test("Body type - type validation", async () => {
  // @ts-ignore
  const data = await (await _t.BodyTypes(client, null)).json();
  assertEquals(data, {
    status: 400,
    message: "Payload validation failed!",
    data: {
      errors: [
        {
          failed_field: "String",
          message: "String is required",
        },
        {
          failed_field: "Int",
          message: "Int is required",
        },
        {
          failed_field: "IntSixtyfour",
          message: "IntSixtyfour is required",
        },
        {
          failed_field: "IntThirtytwo",
          message: "IntThirtytwo is required",
        },
        {
          failed_field: "IntSixteen",
          message: "IntSixteen is required",
        },
        {
          failed_field: "IntEight",
          message: "IntEight is required",
        },
        {
          failed_field: "FloatSixtyfour",
          message: "FloatSixtyfour is required",
        },
        {
          failed_field: "FloatThirtytwo",
          message: "FloatThirtytwo is required",
        },
        {
          failed_field: "TimeTime",
          message: "TimeTime is required",
        },
        {
          failed_field: "Bool",
          message: "Bool is required",
        },
        {
          failed_field: "Uint",
          message: "Uint is required",
        },
        {
          failed_field: "UintSixtyfour",
          message: "UintSixtyfour is required",
        },
        {
          failed_field: "UintThirtytwo",
          message: "UintThirtytwo is required",
        },
        {
          failed_field: "UintSixteen",
          message: "UintSixteen is required",
        },
        {
          failed_field: "UintEight",
          message: "UintEight is required",
        },
      ],
    },
  });
});

// @ts-ignore
Deno.test("Url With string param", async () => {
  const data = await (await _t.UrlWithString(client, "qwe")).json();
  assertEquals(data, response_ok);
});

// @ts-ignore
Deno.test("Submitting optional empty body (optional body field) works)", async () => {
  const data = await (await _t.OptionalBodyField(client, {})).json();
  assertEquals(data, response_ok);
});

// @ts-ignore
Deno.test("Url With int param", async () => {
  const data = await (await _t.UrlWithInt(client, 123)).json();
  assertEquals(data, response_ok);
});

// @ts-ignore
Deno.test("UndefinedValidation field returns correct err", async () => {
  const data = await (
    await _t.UndefinedValidation(client, { undefined_validation_field: "" })
  ).json();
  assertEquals(data, {
    status: 400,
    message: "Payload validation failed!",
    data: {
      errors: [
        {
          failed_field: "UndefinedValidationField",
          message: "UndefinedValidationField is required",
        },
      ],
    },
  });
});

// @ts-ignore
Deno.test("Min validation fails when needed", async () => {
  // Min validation fails
  const t4 = await (await _t.MinValidation(client, { min_ten: "qwe" })).json();
  assertEquals(t4, {
    status: 400,
    message: "Payload validation failed!",
    data: {
      errors: [
        {
          failed_field: "MinTen",
          message: "MinTen should be at least 10 characters long",
        },
      ],
    },
  });

  // Min validation passes
  const t5 = await (
    await _t.MinValidation(client, { min_ten: "qweqweqweqweqweqweqweqwe" })
  ).json();
  assertEquals(t5, response_ok);
});

// @ts-ignore
Deno.test("Max validation fails when needed", async () => {
  // Max validation fails
  const t6 = await (
    await _t.MaxValidation(client, { max_ten: "qweqweqweqweqweqweqweqwe" })
  ).json();
  assertEquals(t6, {
    status: 400,
    message: "Payload validation failed!",
    data: {
      errors: [
        {
          failed_field: "MaxTen",
          message: "MaxTen should be less than 10 characters long",
        },
      ],
    },
  });

  // Max validation passes
  const t7 = await (await _t.MaxValidation(client, { max_ten: "qwe" })).json();
  assertEquals(t7, response_ok);
});

// @ts-ignore
Deno.test("Can append a string to the fetch url", async () => {
  // Max validation fails
  const res = await gomarvin.GetUserById(client, 10, {
    append_url: "?name=jim",
  });
  const data = await res.json();
  // check if fetched url includes ?name=jim
  assert(res.url.includes("?name=jim"));
});

// @ts-ignore
Deno.test("Send an empty body with no headers", async () => {
  // Max validation fails
  const res = await gomarvin.CreateUser(
    no_headers_client,
    {
      username: "test",
      email: "qwe@qwe.com",
      password: "short",
    },
    { options: { body: "", method: "POST" } }
  );
  const data = await res.json();
  assertEquals(data, {
    status: 400,
    message: "Payload validation failed!",
    data: {
      code: 422,
      message: "Unprocessable Entity",
    },
  });
});
