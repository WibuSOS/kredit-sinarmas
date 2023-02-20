import { createContext, useContext, useReducer } from 'react';

function userReducer(state, action) {
	switch (action.type) {
		case 'set': {
			return { ...state, user: action.payload };
		}
		case 'delete': {
			localStorage.clear();
			return { ...state, user: null };
		}
		case 'setDefault': {
			return { ...state, ...action.payload };
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
	const [state, dispatch] = useReducer(userReducer, { user: null }, () => {
		const id = localStorage.getItem("id");
		const username = localStorage.getItem("username");
		const name = localStorage.getItem("name");
		const token = localStorage.getItem("token");
		return { user: { id, username, name, token } };
	});
	const value = { state, dispatch };

	return <Store.Provider value={value}>{children}</Store.Provider>;
}
