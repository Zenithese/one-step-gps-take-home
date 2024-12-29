<template>
  <section class="mapContainer">
    <GoogleMap />
  </section>
  <Drawer>
    <template #buttons>
      <ButtonComponent additionalClass="floating-button top-[60px]" @click="toggleHiddenVisibility">
        <img
          :src="`/device-icons/eye-${deviceStore.ui.hiddenDevicesVisible ? '' : 'off-'}outline.svg`"
          :alt="`hidden-visibility-${deviceStore.ui.hiddenDevicesVisible ? 'on' : 'off'}`"
          width="20px"
        />
      </ButtonComponent>
      <ButtonComponent additionalClass="logout-button floating-button top-[110px]" @click="logoutUser">
      <img
        :src="`/device-icons/log-out-outline.svg`"
        :alt="`log-out`"
        width="20px"
      />
    </ButtonComponent>
    </template>
    <DeviceList />
  </Drawer>
</template>

<script setup lang="ts">
import GoogleMap from "../components/GoogleMap.vue";
import DeviceList from "@/components/DeviceList.vue";
import Drawer from '@/components/DrawerComponent.vue';
import ButtonComponent from "@/components/ButtonComponent.vue";
import { useAuthStore } from '@/stores/authStore';
import { useRouter } from 'vue-router';
import { useDeviceStore } from "@/stores/deviceStore";

const authStore = useAuthStore();
  const router = useRouter();
  const deviceStore = useDeviceStore();

  const toggleHiddenVisibility = () => {
    deviceStore.toggleUiHiddenDevicesVisible();
  };
  
  const logoutUser = async () => {
    try {
      authStore.logout();
      router.push('/login');
    } catch (error) {
      console.error('Error logging out:', error);
    }
  };
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

.logout-button {
  --button-bg-color: #ff4d4f;
  --button-hover-bg-color: #e04343;
  --button-active-bg-color: #cc3a3a;
}

.floating-button {
  position: absolute;
  right: -60px;
  z-index: 1100;
}
</style>