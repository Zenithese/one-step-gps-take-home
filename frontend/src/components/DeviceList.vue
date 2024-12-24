<template>
    <ul>
        <div
            v-for="(device, index) in deviceStore.devices" 
            :key="device.id" 
            class="flex flex-col">
            <div
                id="device"
                class="relative"
                :class="{
                    'drag-hidden': dragStartIndex === index,
                }"
                draggable="true"
                @dragstart="onDragStart(index)"
                @dragover.prevent="(event) => onDragOver(index, event)"
            >
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
                <section class="min-w-[55px] flex items-center">
                <strong>{{ Math.round(convertKmhToMph(device.speed)) }} mph</strong>
                </section>
            </div>
            <div
                class="drop-space"
                :class="{ 
                    'drag-over': hoverIndex === index, 
                    'drag-hidden': dragStartIndex === index 
                }"
                @dragleave="onDragLeave"
                @dragover.prevent="(event) => {}"
                @drop="(event) => onDrop(index)"
            />
        </div>
    </ul>
  </template>

<script setup lang="ts">
import { useDeviceStore } from "@/stores/deviceStore";
import { rssiInterpreter } from "@/utils/Device/rssiInterpreter";
import { driveStateFormatter } from "@/utils/Device/driveStateFormatter";
import { convertKmhToMph } from "@/utils/Device/convertKmhToMph";
import { ref } from "vue";

const deviceStore = useDeviceStore();
const powerStatus = (driveState: string) => (driveState === "off" ? "off" : "on");

let dragStartIndex: number | null = null;
const hoverIndex = ref<number | null>(null);

const onDragStart = (index: number) => {
  dragStartIndex = index;
};

const onDragOver = (index: number, event: DragEvent) => {
    const rect = (event.currentTarget as HTMLElement).getBoundingClientRect();
    if (!rect) return;
    const middle = rect.top + rect.height / 2;
    hoverIndex.value = event.clientY < middle ? index - 1 : index;
};

const onDragLeave = (event: DragEvent) => {
    const rect = (event.currentTarget as HTMLElement).getBoundingClientRect();
    if (!rect) return;
    if (event.clientX > rect.width) hoverIndex.value = null;
};

const onDrop = (index: number) => {
  if (dragStartIndex === null) return;

  const devices = [...deviceStore.devices];
  const draggedItem = devices[dragStartIndex];
  devices.splice(dragStartIndex, 1);
  const newIndex = dragStartIndex > index ? index + 1 : index;
  const newOrder = [...devices.slice(0, newIndex), draggedItem, ...devices.slice(newIndex)];

  deviceStore.devices = newOrder;
  dragStartIndex = null;
  hoverIndex.value = null;
};
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

ul > div > div {
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

.drop-space {
    width: calc(100% + 1rem);
    height: 10px;
    position: relative;
    left: -1rem;
    
}

.status-on::before {
  background-color: #2aee32;
}

.status-off::before {
  background-color: #fb3c3c;
}

ul > div {
  display: flex;
  gap: 0.5rem;
  
}

.drag-hidden {
  display: none;
}

.drag-over {
    height: 117px;
    transition: height 0.2s ease;
    background-color: aqua;
}
</style>
