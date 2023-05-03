import { createContext, useContext, useReducer } from 'react';

function userReducer(state, action) {
	switch (action.type) {
		case 'login': {
			localStorage.setItem("loggedIn", 1);
			return { ...state, loggedIn: 1 };
		}
		case 'logout': {
			localStorage.clear();
			return { ...state, loggedIn: 0 };
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
	const [state, dispatch] = useReducer(userReducer, { loggedIn: 0 }, () => {
		const loggedIn = localStorage.getItem("loggedIn");
		return { loggedIn };
	});
	const value = { state, dispatch };

	return <Store.Provider value={value}>{children}</Store.Provider>;
}
