// import { CommentEndpoints } from "../../../gomarvin.gen";
import { Test, FetchCreateUserEndpoint } from "./lib/tests";

// import the generated file
import * as F from "../../../chi_with_modules/public/gomarvin.gen" 
// or just import a single fetch function
import { GetUserById } from "../../../chi_with_modules/public/gomarvin.gen"

// either use the default client created from
// the settings of the config file, or create a new one
// (useful when switching environments)
const defaultClient = F.defaultClient

// api client when deployed
const productionClient: F.Client = {
  host_url: "http://example.com",
  api_prefix: "/api/v1",
  headers: {
    "Content-type": "application/json;charset=UTF-8",
  },
}

const DEV_MODE = true

// switch to productionClient if DEV_MODE is false
const client = DEV_MODE ? defaultClient : productionClient

// fetch GetUserById endpoint
async function FetchGetUsersById() {
  const res = await F.GetUserById(client, 10);
  console.log(res);
}

// append optional string to the existing endpoint url
async function FetchEndpointWithAppendedUrl() {
  const res = await F.GetUserById(client, 10, { append_url: "?name=jim" });
  console.log(res);
}

// define custom options for the fetch request
async function FetchEndpointWithCustomOptions() {
  const res = await F.GetUserById(client, 10, { options: { method: "POST" } });
  console.log(res);
}

// Use both optional values
// - append a string to the fetch url
// - define a new options object used in the fetch request
async function FetchWithAppendedUrlAndCustomOptions() {
  const res = await F.GetUserById(client, 10, {
    options: { method: "DELETE" },
    append_url: "?name=jim",
  });
  console.log(res);
}


export function App() {
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
