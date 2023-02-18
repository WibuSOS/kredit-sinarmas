// URLs
const API_URL = process.env.API_URL;
const LOGIN_URL = `${API_URL}/login`;

// DIRECTORIES
const ASSETS_DIR = "/assets";
const FIGURES_DIR = `${ASSETS_DIR}/figures`;

const RUPIAH = number => {
	return new Intl.NumberFormat("id-ID", {
		style: "currency",
		currency: "IDR"
	}).format(number);
};

export {
	API_URL,
	LOGIN_URL,
	ASSETS_DIR,
	RUPIAH
}
