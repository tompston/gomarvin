interface _Client {
    headers: Object
    host_url: string
    api_prefix: string
}

export class API_Client {

    headers: Object
    host_url: string
    api_prefix: string
    url: string

    constructor({ headers, host_url, api_prefix }: _Client) {
        this.headers = headers
        this.host_url = host_url
        this.api_prefix = api_prefix
        this.url = `${host_url}${api_prefix}`
    }

    Log(x: string): string {
        return x
    }

    SayWoof(): void {
        console.log("Im saying Woooooof!")
    }

    PurgeThisOnBuild() {
        console.log("is this included in the build, if not used?")
        // yup
    }
}

// client.SayWoof()
// client.Log("something")