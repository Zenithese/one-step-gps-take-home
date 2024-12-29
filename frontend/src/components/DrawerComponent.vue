<template>
  <transition name="slide">
    <section v-show="{ sidebar: true, ishidden: !visible }" :class="{ sidebar: true, ishidden: !visible }">
      <ButtonComponent additionalClass="toggleButton" @click="toggleSidebar">
        <img
          :src="`/device-icons/${visible ? 'close-outline' : 'menu-outline'}.svg`"
          :alt="`${visible ? 'menu-outline' : 'close-outline'}`"
          width="20px"
        />
      </ButtonComponent>
      <slot name="buttons" v-if="visible" />
      <slot />
    </section>
  </transition>
</template>

<script setup lang="ts">
import { ref } from "vue";
import ButtonComponent from "./ButtonComponent.vue";

const visible = ref(true);

const toggleSidebar = () => {
  visible.value = !visible.value;
};
</script>

<style scoped>
.sidebar {
  position: absolute;
  top: 0;
  left: 0;
  width: 325px;
  height: 100vh;
  background-color: white;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  transition: transform 0.3s ease-in-out;
}

.ishidden {
  transform: translateX(-100%);
}

.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease-in-out;
}

.slide-enter,
.slide-leave-to {
  transform: translateX(0);
}

.toggleButton {
  position: absolute;
  top: 10px;
  right: -60px;
  z-index: 1100;
  --button-bg-color: white;
  --button-hover-bg-color: #f5f5f5;
  --button-active-bg-color: #f5f5f5;
}
</style>