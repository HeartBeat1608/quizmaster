import axios from 'axios';

const BASE_URL: string = process.env.BASE_URL || 'http://localhost:3000';

export const api = axios.create({
	baseURL: BASE_URL,
	headers: {
		'Content-Type': 'application/json'
	}
});

api.interceptors.request.use((config) => {
	if (!localStorage) return config;

	const token = localStorage.getItem('token');
	if (!token) return config;

	config.headers.Authorization = `Bearer ${token}`;

	return config;
});
