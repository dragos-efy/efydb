const TOKEN = 'token';
const USERNAME = 'username';
const ROLE = 'role';

const save = (key: string, value: string) => {
	try {
		localStorage.setItem(key, value);
	} catch (error) {
		console.error(error);
	}
};

export const get = (key: string) => {
	try {
		const token = localStorage.getItem(key);
		return token ? token : null;
	} catch (error) {
		console.error(error);
		return null;
	}
};

export const getToken = () => {
	return get(TOKEN);
};

export const setToken = (token: string) => {
	if (!token) return;
	save(TOKEN, token);
};

export const setUsername = (username: string) => {
    save(USERNAME, username);
}

export const getUsername = () => {
	return get(USERNAME);
};

export const logoutUser = () => {
	try {
		localStorage.removeItem(TOKEN);
		localStorage.removeItem(ROLE);
        localStorage.removeItem(USERNAME);
	} catch (error) {
		console.error(error);
	}
};

export const getRole = () => {
	let role: number = 0;
	role = parseInt(get(ROLE)!);
	return role;
};

export const setRole = (role: number) => {
    save(ROLE, role.toString());
};

export const setUserInfo = (userInfo: any) => {
    setRole(userInfo.role);
    setToken(userInfo.token);
    setUsername(userInfo.name);
}