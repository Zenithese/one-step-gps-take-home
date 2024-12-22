<template>
  <div>
    <h2>Device List</h2>
    <ul>
      <li v-for="device in deviceStore.devices" :key="device.id">
        <strong>{{ device.name }}</strong> - Position: {{ device.position }} -
        Status: {{ device.activeState }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { useDeviceStore } from "@/stores/deviceStore";
const deviceStore = useDeviceStore();
</script>

<style scoped>
div {
  background-color: aliceblue;
}
</style>