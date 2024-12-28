<template>
    <ul>
        <div
            v-for="(device, index) in deviceStore.devices" 
            :key="device.id" 
            class="flex flex-col"
        >
            <div 
                v-if="index === 0" 
                class="drop-space"
                :class="{ 'drag-over': hoverIndex === -1 }"
                @dragleave="onDragLeave"
                @dragover.prevent="(event) => {}"
                @drop="() => onDrop(-1)"
            />
            <div
                id="device"
                class="relative"
                :class="{
                    'drag-hidden': dragStartIndex === index,
                }"
                draggable="true"
                @dragstart="onDragStart(index)"
                @dragover.prevent="(event) => onDragOver(index, event)"
                @drop="() => onDrop(index)"
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
                @dragover.prevent="() => {}"
                @drop="() => onDrop(index)"
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
    const rect = (event.currentTarget as HTMLElement)?.getBoundingClientRect();
    if (!rect) return;
    const middle = rect.top + rect.height / 2;
    if (!index && event.clientY < middle) {
        hoverIndex.value = -1;
    } else {
        hoverIndex.value = event.clientY < middle ? index - 1 : index;
    }
};

const onDragLeave = (event: DragEvent) => {
    const rect = (event.currentTarget as HTMLElement)?.getBoundingClientRect();
    if (!rect) return;
    if (event.clientX > rect.width) hoverIndex.value = null;
};

const onDrop = (index: number) => {
  if (dragStartIndex === null) return;

  const devices = [...deviceStore.devices];
  const draggedItem = devices[dragStartIndex];
  devices.splice(dragStartIndex, 1);

  if (index === -1) {
    deviceStore.devices = [draggedItem, ...devices];
  } else {
    const idx = hoverIndex.value ?? index;
    const newIndex = dragStartIndex > idx ? idx + 1 : idx;
    deviceStore.devices = [...devices.slice(0, newIndex), draggedItem, ...devices.slice(newIndex)];
  }

  const newOrder = deviceStore.devices.map((device) => device.id);
  deviceStore.reorderDevices(newOrder);

  dragStartIndex = null;
  hoverIndex.value = null;
};
</script>

<style scoped>
ul {
  list-style-type: none;
  display: flex;
  flex-direction: column;
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
  
}

.drag-hidden {
  display: none;
}

.drag-over {
  height: 117px;
  background: 
        linear-gradient(90deg, #1E90FF 50%, transparent 0) repeat-x,
        linear-gradient(90deg, #1E90FF 50%, transparent 0) repeat-x,
        linear-gradient(0deg, #1E90FF 50%, transparent 0) repeat-y,
        linear-gradient(0deg, #1E90FF 50%, transparent 0) repeat-y;
    background-size: 4px 2px, 4px 2px, 2px 4px, 2px 4px;
    background-position: 0 0, 0 100%, 0 0, 100% 0;
  position: relative;
  animation: linearGradientMove .3s infinite linear;
}

@keyframes linearGradientMove {
    100% {
        background-position: 4px 0, -4px 100%, 0 -4px, 100% 4px;
    }
}

.drag-over::after{
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(30, 144, 255, 0.1);
}
</style>
