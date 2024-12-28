<template>
  <section class="mapContainer">
    <GoogleMap />
  </section>
  <Drawer>
    <ButtonComponent additionalClass="logout-button absolute top-[100px] right-[-120px] z-[1100]" @click="logoutUser">
      Logout
    </ButtonComponent>
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

const authStore = useAuthStore();
  const router = useRouter();
  
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
</style>