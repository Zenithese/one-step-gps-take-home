<script setup lang="ts">
import { onMounted, ref } from "vue";
import { Loader } from "@googlemaps/js-api-loader";

const mapRef = ref<HTMLElement | null>(null);
const map = ref<google.maps.Map | null>(null);

const apiKey = "";

onMounted(async () => {
  const loader = new Loader({
    apiKey,
    version: "weekly",
    libraries: ["places"],
  });

  await loader.load();

  if (mapRef.value) {
    map.value = new google.maps.Map(mapRef.value, {
      center: { lat: 37.7749, lng: -122.4194 },
      zoom: 12,
    });
  }
});
</script>

<template>
  <div ref="mapRef" style="height: 500px; width: 500px;"></div>
</template>