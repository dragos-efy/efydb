import { getToken } from "./token";

const BASE_URL = "http://127.0.0.1:8001"

const fetchJson = async (path: string, options: any) => {
    let response = await fetch(new URL(path, BASE_URL), {
        ...options,
        headers: {
            "Authorization": getToken(),
            ...options.headers
        }
    });
    return await response.json();
}

export default fetchJson;