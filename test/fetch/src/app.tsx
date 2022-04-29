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

  let appendable_headers = {
    "asd-key": "asd-value",
    aaaaaa: "bbbbbb",
  };

  // object passed to new Headers()
  let init_headers_test = {
    "Content-type": "application/json;charset=UTF-8",
    "init-test": "fake-header",
  };

  // https://developer.mozilla.org/en-US/docs/Web/API/Headers/Headers
  function testHeaders(custom_headers: object = {}) {
    // let headers = new Headers();
    // let headers = new Headers(F.API.init_headers);
    let headers = new Headers(init_headers_test);

    // console.log(headers);
    for (const [k, v] of Object.entries(appendable_headers)) {
      headers.append(k, v);
    }

    // console.log(Object.entries(custom_headers));

    for (const [k, v] of Object.entries(custom_headers)) {
      headers.append(k, v);
    }

    for (let x of headers.entries()) {
      console.log(x[0] + ": " + x[1]);
    }

    // console.log(headers);
    // console.log(F.API.init_headers);
  }

  let customheadersFetch = testHeaders({
    "new-header": "this-is-a-new-header",
  });

  return (
    <>
      <div class="flex-center">
        <div onClick={Test} class="test-btn">
          Test
        </div>
      </div>

      {/*  */}
      <div>
        <div onClick={FetchGetUsersEndpoint} class="test-btn">
          FetchGetUsersEndpoint
        </div>
        <div onClick={FetchCreateUserEndpoint} class="test-btn">
          FetchCreateUserEndpoint
        </div>
        <div onClick={FetchComments} class="test-btn">
          FetchComments
        </div>
        <div onClick={testHeaders} class="test-btn">
          testHeaders
        </div>

        {/* FetchComments */}
      </div>
    </>
  );
}
