import { getToken } from "./token";

const BASE_URL = "http://127.0.0.1:8080"

const fetchJson = async (path: string, options: any) => {
    let response = await fetch(new URL(path, BASE_URL), {
        ...options,
        headers: {
            "Authorization": getToken(),
            "Content-Type": "application/json",
        }
    });
    return await response.json();
}

export const fetchFormJson = async (path: string, options: any) => {
    let response = await fetch(new URL(path, BASE_URL), {
        ...options,
        headers: {
            "Authorization": getToken(),
        }
    });
    return await response.json();
}

export default fetchJson;