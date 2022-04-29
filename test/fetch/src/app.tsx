import * as F from "../../../gomarvin.gen";

export function App() {
  async function Test() {
    /** -- user test */
    console.log("GetUsers()", await (await F.GetUsers()).json());
    console.log("GetUser(1)", await (await F.GetUserById(1)).json());
    console.log(
      "CreateUser (INVALID)",
      await (
        await F.CreateUser({
          username: "qweqwe",
          email: "string",
          age: 20,
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
          age: 20,
          password: "very-long-and-good-password",
        })
      ).json()
    );

    /** -- comment test */
    console.log("GetComment(1)", await (await F.GetComment(1)).json());
    console.log("GetComments()", await (await F.GetComments()).json());
    console.log("CreateComment()", await (await F.CreateComment()).json());
    console.log("DeleteComment()", await (await F.DeleteComment()).json());
    console.log("UpdateComment()", await (await F.UpdateComment()).json());
    //
  }

  async function FetchComments() {
    console.log(
      "GetComment(1)",
      await (await F.CommentEndpoints.GetComment(1)).json()
    );
  }

  async function FetchGetUsersEndpoint() {
    let res = await F.GetUserById(1);
    let users = await res.json();
    console.log(users);
  }

  async function FetchCreateUserEndpoint() {
    let res = await F.CreateUser({
      username: "qweqwe",
      email: "qwe@qwe.com",
      age: 20,
      password: "very-long-and-good-password",
    });

    let users = await res.json();
    console.log(users);
  }

  return (
    <>
      <div class="flex-center">
        <div onClick={Test} class="test-btn">
          Test
        </div>
      </div>

      {/*  */}
      <div class="flex-center--">
        <div onClick={FetchGetUsersEndpoint} class="test-btn">
          FetchGetUsersEndpoint
        </div>
        <div onClick={FetchCreateUserEndpoint} class="test-btn">
          FetchCreateUserEndpoint
        </div>
        <div onClick={FetchComments} class="test-btn">
          FetchComments
        </div>

        {/* FetchComments */}
      </div>
    </>
  );
}
