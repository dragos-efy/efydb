const BASE_URL = "http://127.0.0.1:8001"

const fetchJson = async (path: string, options: any) => {
    let response = await fetch(new URL(path, BASE_URL), options);
    return await response.json();
}

export default fetchJson;