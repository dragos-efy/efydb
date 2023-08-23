import { getToken } from "./token";
import { PUBLIC_API_URL } from "$env/static/public";

const fetchJson = async (path: string, options: any) => {
    let response = await fetch(new URL(path, PUBLIC_API_URL), {
        ...options,
        headers: {
            "Authorization": getToken(),
            "Content-Type": "application/json",
        }
    });
    return await response.json();
}

export const fetchFormJson = async (path: string, options: any) => {
    let response = await fetch(new URL(path, PUBLIC_API_URL), {
        ...options,
        headers: {
            "Authorization": getToken(),
        }
    });
    return await response.json();
}

export default fetchJson;