import axios from "axios";
import { type Preferences } from "@/types/Preferences";

axios.defaults.withCredentials = true;
const API_URL = import.meta.env.VITE_API_URL;

export async function updatePreferences(preferences: Preferences) {
  await axios.put(`${API_URL}/preferences`, preferences);
}

export async function fetchPreferences() {
  const response = await axios.get(`${API_URL}/preferences/fetch`);
  return response.data;
}