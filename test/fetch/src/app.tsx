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
   * interface used in the fetch request with optional parameters
   * @param {RequestInit} [options] 
   * If default options need to be edited, provide a custom options object
   * @param {string} [append_url] 
   * extend the url with custom params (like "?name=jim")
   */
  interface OptionalFetchParams {
    options?: RequestInit;
    append_url?: string;
  }

  async function RandomEndpointWithInterface(
    opt?: OptionalFetchParams
  ): Promise<Response> {
    const appended_url = opt?.append_url ? opt?.append_url : "";
    const url = `${F.API.url}/user${appended_url}`;
    console.log(url);

    console.log("options -->" + opt?.options);
    if (opt?.options) {
      return await fetch(url, opt?.options);
    } else {
      return await fetch(url, {
        method: "GET",
        headers: F.API.init_headers,
      });
    }
  }

  /**
   * @param {RequestInit} [options]
   * If default options need to be edited, provide a custom options object
   */
  async function RandomEndpoint(
    options?: RequestInit,
    append_url?: string
  ): Promise<Response> {
    const appended_url = append_url ? append_url : "";
    const url = `${F.API.url}/user${appended_url}`;
    console.log(url);

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

  // ----- Functioons that don't have optional named params
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

  // ----- Functioons that have optional named params
  async function randomEndpointCustomOptionsInterface() {
    const res = await RandomEndpointWithInterface({
      // append_url: "/qwe",
      options: { method: "DELETE" },
    });
    // console.log(res);
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
        <div onClick={randomEndpointCustomOptionsInterface} class="test-btn">
          randomEndpointCustomOptionsInterface
        </div>

        {/* randomEndpointCustomOptionsqq */}
      </div>
    </>
  );
}
