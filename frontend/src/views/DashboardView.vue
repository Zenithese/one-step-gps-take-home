<template>
  <section class="mapContainer">
    <GoogleMap />
  </section>
  <Drawer>
    <DeviceList />
  </Drawer>
</template>

<script setup lang="ts">
import GoogleMap from "../components/GoogleMap.vue";
import DeviceList from "@/components/DeviceList.vue";
import Drawer from '@/components/DrawerComponent.vue';
</script>

<style scoped>
.mapContainer {
  height: 100vh;
  width: 100vw;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
</style>