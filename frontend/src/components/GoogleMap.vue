<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { Loader } from "@googlemaps/js-api-loader";
import { useDeviceStore } from "@/stores/deviceStore";

const deviceStore = useDeviceStore();
const devices = ref(deviceStore.devices);
const mapRef = ref<HTMLElement | null>(null);
const map = ref<google.maps.Map | null>(null);
const clickedInfoWindows = ref<string[]>([]);
const markers = ref<Map<string, google.maps.Marker>>(new Map());

const apiKey = "";

const renderMarkers = () => {
  markers.value.forEach((marker) => marker.setMap(null));
  markers.value.clear();

  devices.value.forEach((device) => {
    const marker = new google.maps.Marker({
      position: device.position,
      map: map.value,
      title: device.name,
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

    markers.value.set(device.id, marker);
  });
};

const calculateCenter = (points: { lat: number; lng: number }[]): { lat: number; lng: number } => {
  const sum = points.reduce(
    (acc, point) => {
      acc.lat += point.lat;
      acc.lng += point.lng;
      return acc;
    },
    { lat: 0, lng: 0 }
  );

  return {
    lat: sum.lat / points.length,
    lng: sum.lng / points.length,
  };
};

const fitMapToMarkers = (map: google.maps.Map, markers: { lat: number; lng: number }[]) => {
  const bounds = new google.maps.LatLngBounds();
  markers.forEach((marker) => bounds.extend(marker));
  map.fitBounds(bounds);
};

onMounted(async () => {
  const loader = new Loader({
    apiKey,
    version: "weekly",
    libraries: ["places"],
  });

  await loader.load();
  await deviceStore.loadDevices();

  if (mapRef.value) {
    const markerPositions = devices.value.map((device) => device.position);
    map.value = new google.maps.Map(mapRef.value, {
      center: calculateCenter(markerPositions),
      zoom: 2,
    });

    renderMarkers();
    fitMapToMarkers(map.value, markerPositions);
  }
});

watch(
  () => deviceStore.devices,
  (newDevices) => {
    devices.value = newDevices;
    renderMarkers();
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
