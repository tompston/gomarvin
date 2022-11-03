import * as F from "../../../gomarvin.gen";
// import { CommentEndpoints } from "../../../gomarvin.gen";
import { Test, FetchCreateUserEndpoint } from "./lib/tests";
import { API_Client } from "./lib/class_purge_on_build";


export function App() {

  const client = new API_Client(
    {
      host_url: "http://localhost:4444",
      api_prefix: "/api/v1",
      headers: {
        "Content-type": "application/json;charset=UTF-8",
      },
    }
  )
  client.SayWoof()


  // fetch GetUserById endpoint
  async function FetchGetUsersById() {
    const res = await F.GetUserById(10);
    console.log(res);
  }




  // append optional string to the existing endpoint url
  async function FetchEndpointWithAppendedUrl() {
    const res = await F.GetUserById(10, { append_url: "?name=jim" });
    console.log(res);
  }

  // define custom options for the fetch request
  async function FetchEndpointWithCustomOptions() {
    const res = await F.GetUserById(10, { options: { method: "POST" } });
    console.log(res);
  }

  // Use both optional values
  // - append a string to the fetch url
  // - define a new options object used in the fetch request
  async function FetchWithAppendedUrlAndCustomOptions() {
    const res = await F.GetUserById(10, {
      options: { method: "DELETE" },
      append_url: "?name=jim",
    });
    console.log(res);
  }

  return (
    <>
      <div class="flex-center test-btn-grid">
        {/* -- Imported funcs */}
        <div onClick={Test} class="test-btn">
          Test
        </div>
        <div onClick={FetchCreateUserEndpoint} class="test-btn">
          FetchCreateUserEndpoint
        </div>
        {/* -- local funcs */}
        <div onClick={FetchGetUsersById} class="test-btn">
          FetchGetUsersById
        </div>
        <div onClick={FetchEndpointWithAppendedUrl} class="test-btn">
          FetchEndpointWithAppendedUrl
        </div>
        <div
          onClick={FetchEndpointWithCustomOptions}
          class="test-btn should-return-err"
        >
          FetchEndpointWithCustomOptions
        </div>
        <div
          onClick={FetchWithAppendedUrlAndCustomOptions}
          class="test-btn should-return-err"
        >
          FetchWithAppendedUrlAndCustomOptions
        </div>

      </div>
    </>
  );
}
