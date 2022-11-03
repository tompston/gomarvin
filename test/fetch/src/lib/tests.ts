import * as F from "../../../../gomarvin.gen";

export async function Test() {
  /** -- user test */
  console.log("GetUsers()", await (await F.GetUsers()).json());
  console.log("GetUser(1)", await (await F.GetUserById(1)).json());
  console.log(
    "CreateUser (INVALID)",
    await (
      await F.CreateUser({
        username: "qweqwe",
        email: "string",
        password: "string",
      })
    ).json()
  );
  console.log(
    "CreateUser (VALID)",
    await (
      await F.CreateUser({
        username: "qweqwe",
        email: "qwe@qwe.com",
        password: "very-long-and-good-password",
      })
    ).json()
  );

  /** -- comment test */
  //// Not included in the Build, if not used
  // console.log("GetComment(1)", await (await F.GetComment(1)).json());
  // console.log("GetComments()", await (await F.GetComments()).json());
  // console.log("CreateComment()", await (await F.CreateComment()).json());
  // console.log("DeleteComment()", await (await F.DeleteComment()).json());
  // console.log("UpdateComment()", await (await F.UpdateComment()).json());
  //
}

export async function FetchCreateUserEndpoint() {
  let res = await F.CreateUser({
    username: "qweqwe",
    email: "qwe@qwe.com",
    password: "very-long-and-good-password",
  });

  let users = await res.json();
  console.log(users);
}
