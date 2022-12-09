const KEY = "token";
const ROLE = "role"

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

export const clearToken = () => {
    try {
        localStorage.removeItem(KEY);
        localStorage.removeItem(ROLE);
    } catch (error) {
        console.error(error);
    }
}

export const getRole = () => {
    let role: number = 0;
    try {
        role = parseInt(localStorage.getItem(ROLE)!);
    } catch (error) {
        console.error(error);
    }
    return role;
}

export const setRole = (role: number) => {
    try {
        localStorage.setItem(ROLE, role.toString());
    } catch (error) {
        console.error(error);
    }
}