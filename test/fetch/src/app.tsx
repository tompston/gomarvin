import * as F from "../../../gomarvin.gen";
import { CommentEndpoints } from "../../../gomarvin.gen";
import { Test, FetchCreateUserEndpoint } from "./lib/tests";

export function App() {
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

  // Fetch a singe endpoint from the Comment module
  async function FetchCommentById() {
    const res = await CommentEndpoints.GetComment(20);
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
        <div onClick={FetchCommentById} class="test-btn">
          FetchCommentById
        </div>
      </div>
    </>
  );
}
