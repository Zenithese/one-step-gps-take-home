<template>
  <div :class="currentRouteClass">
    <router-view />
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const currentRouteClass = computed(() => {
  if (route.path === "/login") {
    return "auth-bg login-animation";
  } else if (route.path === "/signup") {
    return "auth-bg signup-animation";
  }
  return "bg-default";
});
</script>

<style scoped>
.auth-bg {
  background: linear-gradient(135deg, seashell, #0088cc);
  background-size: 200% 200%;
  height: 100vh;
}

.login-animation {
  animation: login-gradient 2s linear forwards;
}

.signup-animation {
  animation: signup-gradient 2s linear forwards;
}

.bg-default {
  height: 100vh;
}

/* Gradient moves from bottom-right to top-left */
@keyframes login-gradient {
  0% {
    background-position: 100% 100%;
  }
  100% {
    background-position: 0% 0%;
  }
}

/* Gradient moves from top-left to bottom-right */
@keyframes signup-gradient {
  0% {
    background-position: 0% 0%;
  }
  100% {
    background-position: 100% 100%;
  }
}
</style>