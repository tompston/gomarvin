

/**
 * NOTE: function not used currently
 * util function for creating headers. Headers specified in the
 * API object will be sent on each request.
 * @param custom_headers
 * custom object of optional headers, where key and value are strings only
 * @returns {Headers} headers used in the fetch request
 */
export function createHeaders(custom_headers: object, use_default: boolean = true): Headers {
    // let headers = new Headers(API.init_headers);
    let headers = new Headers();

    if (Object.keys(custom_headers).length != 0) {
        for (const [k, v] of Object.entries(custom_headers)) {
            // some weird thing is appended in some cases, so avoid it
            if (k !== "isTrusted") {
                headers.append(k, v);
            }
        }
    }
    return headers;
}
