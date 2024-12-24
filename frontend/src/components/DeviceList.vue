<template>
    <ul>
      <div v-for="device in deviceStore.devices" :key="device.id" class="relative">
        <section :class="`container status-${powerStatus(device.driveState)}`">
          <strong>{{ device.name }}</strong>
          <div>
            <div class="min-w-4">
              <img 
                :src="`/device-icons/signal-icons/connection-${rssiInterpreter(Number(device.rssi))}.svg`" 
                :alt="`/device-icons/connection-${rssiInterpreter(Number(device.rssi))}`"
                width="16px"
              />
            </div>
            <div class="grow">{{ driveStateFormatter(device.driveState) }}</div>
          </div>
          <div class="items-center">
            <div class="min-w-4">
              <img
                :src="`/device-icons/power-${powerStatus(device.driveState)}.svg`"
                :alt="`power-${powerStatus(device.driveState)}`"
                width="12px"
              />
            </div>
            <div class="grow">{{ powerStatus(device.driveState) }}</div>
          </div>
          <div>
            <div class="min-w-4">
              <img
                src="/device-icons/person-outline.svg"
                alt="person-outline"
                width="10px"
              />
            </div>
            <div class="grow">{{ device.name }}</div>
          </div>
          <div>
            <div class="min-w-4">
              <img
                src="/device-icons/location-dot.svg"
                alt="location-dot"
                width="10px"
              />
            </div>
            <div class="grow">{{ device.address }}</div>
          </div>
        </section>
        <section class="min-w-[55px] flex items-center"><strong>{{ Math.round(convertKmhToMph(device.speed)) }} mph</strong></section>
      </div>
    </ul>
</template>

<script setup lang="ts">
import { useDeviceStore } from "@/stores/deviceStore";
import { rssiInterpreter } from "@/utils/Device/rssiInterpreter";
import { driveStateFormatter } from "@/utils/Device/driveStateFormatter";
import { convertKmhToMph } from "@/utils/Device/convertKmhToMph";
const deviceStore = useDeviceStore();
const powerStatus = (driveState: string) => driveState === "off" ? "off" : "on";
</script>

<style scoped>
ul {
  list-style-type: none;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 1rem;
  padding-right: 0;
  padding-top: .5rem;
  padding-bottom: .5rem;
  height: 100%;
  overflow: auto;
  font-size: small;
}

ul > div {
    display: flex;
    gap: .5rem;
}

.container > div {
  display: flex;
  gap: 4px;
}

img {
  padding-top: 2px;
  margin: auto;
}

li {
    display: flex;
}

.container::before {
  content: "";
  display: block;
  width: 8px;
  height: 90%;
  position: absolute;
  left: -1rem;
  top: 5%;
}


.status-on::before {
  background-color: #2aee32;
}

.status-off::before {
  background-color: #fb3c3c;
}
</style>
