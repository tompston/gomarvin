// @ts-check


export const qwe = "asd"

interface _Client {
    headers: Object
    host_url: string
    api_prefix: string
}

function Client({ headers, host_url, api_prefix }: _Client) {
    this.headers = headers
    this.host_url = host_url
    this.api_prefix = api_prefix
    this.url = `${host_url}${api_prefix}`
}

// Animal.prototype.name

Client.prototype.myUrl = function myUrl() {
    console.log(`url :: ${this.url}`)
}

/**
 * @name Client#qwe
 * @function
 */
Client.prototype.printThis = function printThis(x: string): string {
    return `${x}`
}

const api = Client({
    host_url: "http://localhost:4444",
    api_prefix: "/api/v1",
    headers: {
        "Content-type": "application/json;charset=UTF-8",
    },
})


let asd = api.config()

let print_my_value = api.printThis()
// console.log(print_my_value)


// function Log(x: string): string {
//     return x
// }
// Client.prototype.Log = Log
// let test_1 = api.Log()


/**
 * Find out why the info of types is lost when using prototypes.
 * 
 */