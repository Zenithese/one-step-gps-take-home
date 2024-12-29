import { defineStore } from "pinia";
import { fetchDevices } from "@/services/deviceService";
import { fetchPreferences, updatePreferences } from "@/services/preferenceService";
import { type Device } from "@/types/Device";
import { type Preferences } from "@/types/Preferences";
import { convertToCamelCase } from "@/utils/Device/convertToCamelCase";
import { convertToSnakeCase } from "@/utils/Device/convertToSnakeCase";

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
    allDevices: [] as Device[],
    devices: [] as Device[],
    preferences: {
      deviceOrder: [] as string[],
      hiddenDevices: [] as string[],
      devicePhotos: {} as Record<string, string>,
    } as Preferences,
    ui: {
        hiddenDevicesVisible: false,
    }
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
            activeState: device.active_state,
            driveState: device.latest_device_point.device_state.drive_status,
            speed: device.latest_device_point.speed,
            address: await getAddressFromLatLng(latLng),
            rssi: device.latest_device_point.params.rssi,
            hdop: device.latest_device_point.params.hdop,
            gpsLev: device.latest_device_point.params.gpslev,
          };
        })
      );

      this.allDevices = devices;
      this.devices = this.applyPreferencesToDevices(devices);
    },

    async loadPreferences() {
      try {
        const preferences = await fetchPreferences();
        this.preferences = convertToCamelCase(preferences);
        this.devices = this.applyPreferencesToDevices(this.devices);
      } catch (error) {
        console.error("Failed to fetch preferences:", error);
      }
    },

    async savePreferences() {
      try {
        await updatePreferences(convertToSnakeCase(this.preferences));
      } catch (error) {
        console.error("Failed to update preferences:", error);
      }
    },

    reorderDevices(newOrder: string[]) {
      this.preferences.deviceOrder = newOrder;
      this.devices = this.applyPreferencesToDevices(this.devices);
      this.savePreferences();
    },

    toggleHiddenDevice(deviceId: string) {
      const hiddenIndex = this.preferences.hiddenDevices.indexOf(deviceId);
      if (hiddenIndex === -1) {
        this.preferences.hiddenDevices.push(deviceId);
      } else {
        this.preferences.hiddenDevices.splice(hiddenIndex, 1);
      }
      this.devices = this.applyPreferencesToDevices(this.allDevices);
      this.savePreferences();
    },

    updateDevicePhoto(deviceId: string, photo: string) {
      this.preferences.devicePhotos[deviceId] = photo;
      this.savePreferences();
    },

    toggleUiHiddenDevicesVisible() {
      this.ui.hiddenDevicesVisible = !this.ui.hiddenDevicesVisible;
      this.devices = this.applyPreferencesToDevices(this.allDevices);
    },

    applyPreferencesToDevices(devices: Device[]): Device[] {
      const hiddenDevicesSet = new Set(this.ui.hiddenDevicesVisible ? undefined : this.preferences.hiddenDevices);
      const reorderedDevices: Device[] = [];

      for (const deviceId of this.preferences.deviceOrder) {
        const device = devices.find((d) => d.id === deviceId);
        if (device && !hiddenDevicesSet.has(deviceId)) {
          reorderedDevices.push(device);
        }
      }

      const remainingDevices = devices.filter(
        (device) => !hiddenDevicesSet.has(device.id) && !reorderedDevices.some((d) => d.id === device.id)
      );
      return [...reorderedDevices, ...remainingDevices];
    },
  },
});