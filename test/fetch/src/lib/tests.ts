// import * as F from "../../../../gomarvin.gen";
import * as F from "../../../../chi_with_modules/public/gomarvin.gen"

const client = F.defaultClient

export async function Test() {
  /** -- user test */
  console.log("GetUsers()", await (await F.GetUsers(client)).json());
  console.log("GetUser(1)", await (await F.GetUserById(client, 1)).json());

  console.log(
    "CreateUser (INVALID)",
    await (
      await F.CreateUser(client, {
        username: "qweqwe",
        email: "string",
        password: "string",
      })
    ).json()
  );
  console.log(
    "CreateUser (VALID)",
    await (
      await F.CreateUser(client, {
        username: "qweqwe",
        email: "qwe@qwe.com",
        password: "very-long-and-good-password",
      })
    ).json()
  );

  /** -- comment test */
  //// Not included in the Build, if not used
  console.log("GetComment(1)", await (await F.GetComment(client, 1)).json());
  console.log("GetComments()", await (await F.GetComments(client)).json());
  console.log("CreateComment()", await (await F.CreateComment(client)).json());
  console.log("DeleteComment()", await (await F.DeleteComment(client)).json());
  console.log("UpdateComment()", await (await F.UpdateComment(client)).json());
  //
}

export async function FetchCreateUserEndpoint() {
  let res = await F.CreateUser(client, {
    username: "qweqwe",
    email: "qwe@qwe.com",
    password: "very-long-and-good-password",
  });

  let users = await res.json();
  console.log(users);
}
