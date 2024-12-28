import axios from "axios";
import { type User } from "@/types/User";

axios.defaults.withCredentials = true;
const API_URL = import.meta.env.VITE_API_URL;

export async function signup({ username, password }: User) {
  await axios.post(`${API_URL}/signup`, { username, password });
}

export async function login({ username, password }: User) {
  await axios.post(`${API_URL}/login`, { username, password });
}

export async function checkSession() {
  await axios.get(`${API_URL}/check-session`, { withCredentials: true });
}

export async function logout() {
  await axios.post(`${API_URL}/logout`, {}, { withCredentials: true });
}
