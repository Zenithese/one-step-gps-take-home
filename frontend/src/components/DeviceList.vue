<template>
    <ul>
        <div
            v-for="(device, index) in deviceStore.devices" 
            :key="device.id" 
            class="flex flex-col"
            :class="{ 
                'opacity-50': deviceStore.ui.hiddenDevicesVisible && deviceStore.preferences.hiddenDevices.includes(device.id),
                'hoverable': hoverable, }"
            @contextmenu.prevent="showContextMenu($event, device)"
            @mousemove="updateTooltipPosition($event)"
            @mouseenter="showTooltip = true"
            @mouseleave="showTooltip = false"
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
                @mousedown="onMouseDown"
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
                            v-if="deviceStore.preferences.devicePhotos[device.id]"
                            :src="deviceStore.preferences.devicePhotos[device.id]"
                            alt="Device Photo"
                            width="16px"
                            class="rounded-full"
                        />
                        <div v-else v-html="awaitedColoredSvg[device.id]" class="relative right-[2px] w-[16px]"></div>
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
                <section class="min-w-[75px] flex items-center justify-between">
                <strong class="min-w-[50px] text-end">{{ Math.round(convertKmhToMph(device.speed)) }} mph</strong>
                <img
                    :src="`/device-icons/ellipsis-vertical-outline.svg`"
                    :alt="`ellipsis-vertical-outline`"
                    width="16px"
                    @click.stop="showContextMenu($event, device)"
                    @mousedown.stop
                    @mouseenter="showTooltip = false"
                    @mouseleave="showTooltip = true"
                    class="cursor-pointer w-6 h-6"
                />
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

    <!-- Context menu -->
    <div
      v-if="contextMenu.visible"
      :style="{ top: contextMenu.y + 'px', left: contextMenu.x + 'px' }"
      class="context-menu"
      @mouseover="showTooltip = false"
    >
      <ul>
        <li @click="hide()">{{ contextMenu.text }}</li>
        <li @click="uploadPhoto()" class="text-nowrap">Upload Photo</li>
      </ul>
    </div>

    <!-- Tooltip -->
    <div v-if="showTooltip && hoverable" class="tooltip" :style="{ top: `${tooltip.y}px`, left: `${tooltip.x}px` }">
        Drag me
    </div>
  </template>

<script setup lang="ts">
import { useDeviceStore } from "@/stores/deviceStore";
import { rssiInterpreter } from "@/utils/Device/rssiInterpreter";
import { driveStateFormatter } from "@/utils/Device/driveStateFormatter";
import { convertKmhToMph } from "@/utils/Device/convertKmhToMph";
import { powerStatus } from "@/utils/Device/powerStatus";
import { reactive, ref, watch } from "vue";
import { type Device } from "@/types/Device";
import { changeSvgColor } from "@/utils/helpers/changeSvgColor";

const deviceStore = useDeviceStore();

let dragStartIndex: number | null = null;
const hoverIndex = ref<number | null>(null);
const hoverable = ref(true);
const awaitedColoredSvg = reactive<Record<string, string | null>>({});
const showTooltip = ref(false);
const tooltip = ref({ x: 0, y: 0 });
const contextMenu = ref({
  visible: false,
  x: 0,
  y: 0,
  target: {} as Device,
  text: "",
});

const onMouseDown = () => {
    if (hoverable.value) {
        hoverable.value = false;
    };
    if (contextMenu.value.visible) {
        contextMenu.value.visible = false;
    };
};

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

  if (!hoverable.value) {
    hoverable.value = true;
  }

  if (showTooltip.value) {
    showTooltip.value = false;
  }

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

const uploadPhoto = async () => {
  const targetDevice = contextMenu.value.target as Device;

  if (!targetDevice) {
    console.error("No device selected for photo upload.");
    return;
  }

  const input = document.createElement("input");
  input.type = "file";
  input.accept = "image/*";

  input.addEventListener("change", async (event: Event) => {
    const file = (event.target as HTMLInputElement)?.files?.[0];

    if (file) {
      const reader = new FileReader();
      reader.onload = async () => {
        const base64Photo = reader.result as string;
        deviceStore.updateDevicePhoto(targetDevice.id, base64Photo);
        await deviceStore.savePreferences();
        console.log(`Photo uploaded for device: ${targetDevice.name}`);
      };

      reader.onerror = () => {
        console.error("Error reading the uploaded file.");
      };

      reader.readAsDataURL(file);
    }
  });

  input.click();
};

const hide = () => {
    const target = contextMenu.value.target! as Device;
    if (target) {
        const deviceId = target.id;
        deviceStore.toggleHiddenDevice(deviceId);
        contextMenu.value.visible = false;
    }
};

const showContextMenu = (event: MouseEvent, device: Device) => {
  event.preventDefault();
  contextMenu.value = {
    visible: true,
    x: event.clientX,
    y: event.clientY,
    target: device,
    text: deviceStore.preferences.hiddenDevices.includes(device.id) ? "Show" : "Hide",
  };
};

document.addEventListener("click", () => {
  contextMenu.value.visible = false;
});

const updateTooltipPosition = (event: MouseEvent) => {
  tooltip.value.x = event.clientX + 10;
  tooltip.value.y = event.clientY + 10;
};

const loadColoredSvgs = async () => {
  for (const device of deviceStore.devices) {
    const color = deviceStore.ui.colors[device.id]?.hex || "#a3a3a3";
    const svg = await changeSvgColor(`/device-icons/person-outline.svg`, color);
    awaitedColoredSvg[device.id] = svg;
  }
};

watch(
  () => deviceStore.devices,
  async () => {
    await loadColoredSvgs();
  },
  { deep: true, immediate: true }
);
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

.hoverable {
  position: relative;
}

.hoverable:hover {
  background-color: #f5f5f5;
  border-radius: 4px;
}

.hoverable:hover::before {
  content: '';
  position: absolute;
  top: 0;
  left: -1rem;
  height: 100%;
  width: calc(100% + 1rem);
  background-color: #f5f5f5;
  border-radius: 4px;
  z-index: -1;
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

.context-menu {
  position: absolute;
  background: white;
  border: 1px solid black;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  z-index: 2000;
  border-radius: 4px;
}

.context-menu ul {
  margin: 0;
  padding: 0;
  list-style: none;
}

.context-menu li {
  padding: 8px 12px;
  cursor: pointer;
}

.context-menu li:hover {
  background: #f5f5f5;
  border-radius: 4px;
}

.tooltip {
  position: absolute;
  background-color: #333;
  color: #fff;
  padding: 5px 10px;
  border-radius: 4px;
  font-size: 12px;
  z-index: 1000;
  pointer-events: none;
  white-space: nowrap;
}
</style>
