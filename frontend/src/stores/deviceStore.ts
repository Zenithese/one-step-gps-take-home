import { defineStore } from "pinia";
import { fetchDevices } from "@/services/deviceService";
import { type Device } from "@/types/Device";

const getAddressFromLatLng = async (latLng: { lat: number; lng: number }): Promise<string | null> => {
  const geocoder = new google.maps.Geocoder();

  return new Promise((resolve, reject) => {
    geocoder.geocode({ location: latLng }, (results, status) => {
      if (status === "OK" && results?.[0]) {
        resolve(results[0].formatted_address);
      } else {
        console.error("Geocoding failed: ", status);
        reject(null);
      }
    });
  });
};

export const useDeviceStore = defineStore("deviceStore", {
  state: () => ({
    devices: [] as Device[],
  }),
  actions: {
    async loadDevices() {
      const devices = await Promise.all(
        (await fetchDevices()).map(async (device) => {
          const latLng = device.latest_device_point.device_state.adjusted_lat_lng;

          return {
            id: device.device_id,
            name: device.display_name,
            position: latLng,
            activeState: device.active_status,
            driveState: device.latest_device_point.device_state.drive_status,
            speed: device.latest_device_point.speed,
            address: await getAddressFromLatLng(latLng),
          };
        })
      );
      this.devices = devices;
    },
  },
});