<template>
  <transition name="slide">
    <section v-show="{ sidebar: true, hidden: !visible }" :class="{ sidebar: true, hidden: !visible }">
      <button class="toggleButton" @click="toggleSidebar">
        {{ visible ? "Hide" : "Show" }}
      </button>
      <slot />
    </section>
  </transition>
</template>

<script setup lang="ts">
import { ref } from "vue";

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

.hidden {
  transform: translateX(-80%);
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
  top: 20px;
  right: 0;
  z-index: 1100;
  background-color: white;
  border: none;
  padding: 10px 15px;
}
</style>