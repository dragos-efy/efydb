const KEY = "token";

export const getToken = () => {
    try {
        const token = localStorage.getItem(KEY);
        return token ? token: null;
    } catch (error) {
        console.error(error);
        return null;
    }
}

export const setToken = (token: string) => {
    if (!token) return;
    try {
        localStorage.setItem(KEY, token)
    } catch (error) {
        console.error(error);
    }
}