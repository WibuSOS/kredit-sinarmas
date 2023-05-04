import { createContext, useContext, useReducer } from 'react';

function userReducer(state, action) {
	switch (action.type) {
		case 'login': {
			localStorage.setItem("loggedIn", 1);
			localStorage.setItem("name", action.payload.name);
			return { ...state, loggedIn: 1, user: action.payload };
		}
		case 'logout': {
			localStorage.clear();
			return { ...state, loggedIn: 0, user: { name: null } };
		}
		default: {
			console.log(action.type + "not found");
		}
	}
}

const Store = createContext();
const useStore = () => useContext(Store);
export { Store, useStore };

export default function UserContext({ children }) {
	const [state, dispatch] = useReducer(userReducer, { loggedIn: 0, user: null }, () => {
		const loggedIn = localStorage.getItem("loggedIn");
		const user = { name: localStorage.getItem("name") };
		return { loggedIn, user };
	});
	const value = { state, dispatch };

	return <Store.Provider value={value}>{children}</Store.Provider>;
}
