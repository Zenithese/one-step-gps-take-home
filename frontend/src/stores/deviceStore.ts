import { defineStore } from "pinia";
import { fetchDevices } from "@/services/deviceService";
import { type Device } from "@/types/Device";

export const useDeviceStore = defineStore("deviceStore", {
  state: () => ({
    devices: [] as Device[],
  }),
  actions: {
    async loadDevices() {
        console.log("loadDevices");
      const devices = (await fetchDevices()).map((device) => ({
        id: device.device_id,
        name: device.display_name,
        position: device.latest_device_point.device_state.adjusted_lat_lng,
        activeState: device.active_status,
        driveState: device.latest_device_point.device_state.drive_status,
      }));
      console.log("devices:", devices);
      this.devices = devices;
    },
  },
});
