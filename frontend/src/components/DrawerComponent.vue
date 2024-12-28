<template>
  <transition name="slide">
    <section v-show="{ sidebar: true, ishidden: !visible }" :class="{ sidebar: true, ishidden: !visible }">
      <button class="toggleButton" @click="toggleSidebar">
        {{ visible ? "Hide" : "Show" }}
      </button>
      <LogoutButton additionalClass="absolute top-[70px] right-[-97px] z-[1100]"/>
      <slot />
    </section>
  </transition>
</template>

<script setup lang="ts">
import { ref } from "vue";
import LogoutButton from "./LogoutButton.vue";

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
  width: 300px;
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
  right: -70px;
  z-index: 1100;
  background-color: white;
  border: none;
  border-radius: 2px;
  padding: 10px 15px;
}
</style>