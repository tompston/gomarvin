import * as F from "../../../gomarvin.gen";
import { Test, FetchCreateUserEndpoint } from "./lib/tests";

export function App() {
  // https://developer.mozilla.org/en-US/docs/Web/API/Headers/Headers
  function testHeaders(custom_headers: object = {}) {
    let headers = F.createHeaders(custom_headers);
    for (let x of headers.entries()) {
      console.log(x[0] + ": " + x[1]);
    }
  }

  let customheadersFetch = () =>
    testHeaders({
      "new-header": "this-is-a-new-header",
    });

  /**
   * @param {RequestInit} [options]
   * If default options need to be edited, provide a custom options object
   */
  async function RandomEndpoint(options?: RequestInit): Promise<Response> {
    const url = `${F.API.url}/user`;

    // console.log("options -->" + options);
    if (options) {
      return await fetch(url, options);
    } else {
      return await fetch(url, {
        method: "GET",
        headers: F.API.init_headers,
      });
    }
  }

  async function randomEndpointDefaultOptions() {
    const res = await RandomEndpoint();
    console.log(res);
  }
  async function randomEndpointCustomOptions() {
    const res = await RandomEndpoint({ method: "get" });
    console.log(res);
  }

  async function IncorrectOptionsTest() {
    const res = await F.GetUserById(10, { method: "POST" });
    console.log(res);
  }

  return (
    <>
      <div class="flex-center">
        <div onClick={Test} class="test-btn">
          Test
        </div>
      </div>

      <div>
        <div onClick={FetchCreateUserEndpoint} class="test-btn">
          FetchCreateUserEndpoint
        </div>
        <div onClick={testHeaders} class="test-btn">
          testHeaders
        </div>
        <div onClick={customheadersFetch} class="test-btn">
          customheadersFetch
        </div>
        <div onClick={randomEndpointDefaultOptions} class="test-btn">
          randomEndpointDefaultOptions
        </div>
        <div onClick={randomEndpointCustomOptions} class="test-btn">
          randomEndpointCustomOptions
        </div>
        <div onClick={IncorrectOptionsTest} class="test-btn">
          IncorrectOptionsTest
        </div>
      </div>
    </>
  );
}
