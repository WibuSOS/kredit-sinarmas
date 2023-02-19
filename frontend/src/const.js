// import * as dotenv from 'dotenv'; // see https://github.com/motdotla/dotenv#how-do-i-use-dotenv-with-import
// dotenv.config();
// require('dotenv').config();

// URLs
// const API_URL = process.env.API_URL;
const API_URL = "http://localhost:8080";
const LOGIN_URL = `${API_URL}/login`;

// DIRECTORIES
const ASSETS_DIR = "/assets";
const ICONS_DIR = `${ASSETS_DIR}/icons`;

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
	ICONS_DIR,
	RUPIAH
}
