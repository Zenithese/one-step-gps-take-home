<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { Loader } from "@googlemaps/js-api-loader";
import { useDeviceStore } from "@/stores/deviceStore";
import { calculateCenter } from "@/utils/GoogleMap/calculateCenter";
import { fitMapToMarkers } from "@/utils/GoogleMap/fitMapToMarkers";

const deviceStore = useDeviceStore();
const devices = ref(deviceStore.devices);
const mapRef = ref<HTMLElement | null>(null);
const map = ref<google.maps.Map | null>(null);
const clickedInfoWindows = ref<string[]>([]);
const markers = ref<google.maps.Marker[]>([]);

const apiKey = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;

const initializeMap = () => {
  if (mapRef.value) {
    const markerPositions = devices.value.map((device) => device.position);
    map.value = new google.maps.Map(mapRef.value, {
      center: calculateCenter(markerPositions),
      zoom: 2,
      mapTypeControlOptions: {
        position: google.maps.ControlPosition.TOP_CENTER,
        style: google.maps.MapTypeControlStyle.HORIZONTAL_BAR,
      },
    });

    // google.maps.event.addListenerOnce(map.value, "idle", () => {
    //   map.value?.panBy(-100, 0);
    // });

    renderMarkers();
    fitMapToMarkers(map.value, markerPositions);
  }
}

const clearMarkers = () => {
  markers.value.forEach((marker) => marker.setMap(null));
  markers.value = [];
};

const renderMarkers = () => {
  clearMarkers();

  console.log(window.location.origin)

  devices.value.forEach((device) => {
    const marker = new google.maps.Marker({
      position: device.position,
      map: map.value,
      title: device.name,
      icon:  {
      url: deviceStore.preferences.devicePhotos[device.id] ? (
        `data:image/svg+xml;charset=UTF-8,${encodeURIComponent(`
        <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40">
          <circle cx="20" cy="20" r="20" fill="${device.driveState === "off" ? "#fb3c3c" : "#2aee32"}" />
          <image x="4" y="4" width="32" height="32" href="${deviceStore.preferences.devicePhotos[device.id]}" clip-path="circle(16px at 16px 16px)" />
        </svg>
      `)}`
    ) : (
      `data:image/svg+xml;charset=UTF-8,${encodeURIComponent(`
        <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40">
          <circle cx="20" cy="20" r="20" fill="${device.driveState === "off" ? "#fb3c3c" : "#2aee32"}" />
          <circle cx="20" cy="20" r="16" fill="white" />
          <svg x="10" y="10" width="20" height="20" viewBox="0 -960 960 960" xmlns="http://www.w3.org/2000/svg" fill="${deviceStore.ui.colors[device.id].hex}">
            <path d="M480-480q-60 0-102-42t-42-102q0-60 42-102t102-42q60 0 102 42t42 102q0 60-42 102t-102 42ZM192-192v-96q0-23 12.5-43.5T239-366q55-32 116.29-49 61.29-17 124.5-17t124.71 17Q666-398 721-366q22 13 34.5 34t12.5 44v96H192Zm72-72h432v-24q0-5.18-3.03-9.41-3.02-4.24-7.97-6.59-46-28-98-42t-107-14q-55 0-107 14t-98 42q-5 4-8 7.72-3 3.73-3 8.28v24Zm216.21-288Q510-552 531-573.21t21-51Q552-654 530.79-675t-51-21Q450-696 429-674.79t-21 51Q408-594 429.21-573t51 21Zm-.21-72Zm0 360Z"/>
          </svg>
        </svg>
      `)}`
    ),
      scaledSize: new google.maps.Size(40, 40),
    },
    });

    const infoWindow = new google.maps.InfoWindow();

    const updateInfoWindowContent = (deviceId: string) => {
      infoWindow.setContent(`
        <div class="info-window-content" id="${deviceId}">
          <h3>${device.name}</h3>
          <p><strong>Active State:</strong> ${device.activeState}</p>
          <p><strong>Drive State:</strong> ${device.driveState}</p>
        </div>
      `);

      google.maps.event.addListenerOnce(infoWindow, "domready", () => {
        const allInfoWindows = document.querySelectorAll(".gm-style-iw");
        allInfoWindows.forEach((iwOuter) => {
          const infoWindowContent = iwOuter.querySelector(
            ".info-window-content"
          );
          const id = infoWindowContent?.id;

          if (id && clickedInfoWindows.value.includes(id)) {
            (iwOuter as HTMLElement).style.border = "2px solid blue";
          } else {
            (iwOuter as HTMLElement).style.border = "none";
          }
        });
      });
    };

    marker.addListener("mouseover", () => {
      if (!clickedInfoWindows.value.includes(device.id)) {
        updateInfoWindowContent(device.id);
        infoWindow.open(map.value, marker);
      }
    });

    marker.addListener("mouseout", () => {
      if (!clickedInfoWindows.value.includes(device.id)) {
        infoWindow.close();
      }
    });

    marker.addListener("click", () => {
      if (clickedInfoWindows.value.includes(device.id)) {
        clickedInfoWindows.value = clickedInfoWindows.value.filter(
          (id) => id !== device.id
        );
        infoWindow.close();
      } else {
        clickedInfoWindows.value.push(device.id);
        updateInfoWindowContent(device.id);
        infoWindow.open(map.value, marker);
      }
    });

    infoWindow.addListener("closeclick", () => {
      clickedInfoWindows.value = clickedInfoWindows.value.filter(
        (id) => id !== device.id
      );
    });

    markers.value.push(marker);
  });
};

onMounted(async () => {
  const loader = new Loader({
    apiKey,
    version: "weekly",
    libraries: ["places"],
  });

  await loader.load();
  await deviceStore.loadPreferences();
  await deviceStore.loadDevices();

  initializeMap();
});

watch(
  () => deviceStore.devices,
  (newDevices) => {
    devices.value = newDevices;
    if (deviceStore.devices.length !== markers.value.length) {
      initializeMap();
    }
  },
  { deep: true }
);
</script>

<template>
  <div ref="mapRef" style="height: 100%; width: 100%"></div>
</template>

<style scoped>
.gm-style-iw {
  background-color: lightblue !important;
  border: 2px solid blue !important;
  border-radius: 10px !important;
  box-shadow: 0px 5px 15px rgba(0, 0, 0, 0.3) !important;
}

.gm-style-iw-t::after {
  border-color: lightblue !important;
}
</style>
