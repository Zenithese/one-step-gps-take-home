import axios from "axios";
import { type DeviceResultListItem } from "@/types/DeviceResultListItem";

// const API_URL = process.env.VUE_APP_API_URL;

export async function fetchDevices(): Promise<DeviceResultListItem[]> {
  const response = await axios.get<{ result_list: DeviceResultListItem[] }>(`http://localhost:8080/api/fetch-track-data`);
  return response.data.result_list;
}