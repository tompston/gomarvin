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
    "X-Custom-Header": "hello world",
  };

  // https://developer.mozilla.org/en-US/docs/Web/API/Headers/Headers
  function testHeaders(custom_headers: object = {}) {
    // let headers = new Headers(init_headers_test);
    // console.log(Object.entries(custom_headers));
    // console.log(`custom_headers --> ${custom_headers}`);
    // for (const [k, v] of Object.entries(custom_headers)) {
    //   console.log(k + ": " + v);
    // }
    // console.log(
    //   `Object.keys(custom_headers).length -> ${
    //     Object.keys(custom_headers).length
    //   }`
    // );

    // if (Object.keys(custom_headers).length != 0) {
    //   for (const [k, v] of Object.entries(custom_headers)) {
    //     // some weird thing is appended if object is empty
    //     if (k !== "isTrusted") {
    //       headers.append(k, v);
    //     }
    //   }
    // }
    let headers = createHeaders(custom_headers);

    for (let x of headers.entries()) {
      console.log(x[0] + ": " + x[1]);
    }

    // const url = `${F.API.url}/comment/`;
    // return await fetch(url, {
    //   method: "GET",
    //   headers: API.init_headers,
    // });
  }

  /**
   *
   * @param custom_headers custom object of optional headers, where key and value are strings only
   * @returns {Headers} headers used in the fetch request
   */
  function createHeaders(custom_headers: object): Headers {
    let headers = new Headers(init_headers_test);

    if (Object.keys(custom_headers).length != 0) {
      for (const [k, v] of Object.entries(custom_headers)) {
        // some weird thing is appended if object is empty, so avoid it
        if (k !== "isTrusted") {
          headers.append(k, v);
        }
      }
    }
    return headers;
  }

  let customheadersFetch = () =>
    testHeaders({
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
        <div onClick={customheadersFetch} class="test-btn">
          customheadersFetch
        </div>

        {/* customheadersFetch */}

        {/* FetchComments */}
      </div>
    </>
  );
}
