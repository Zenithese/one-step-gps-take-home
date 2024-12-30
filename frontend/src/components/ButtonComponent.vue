<template>
  <button :class="additionalClass" :type="type || 'button'" @click="handleClick">
    <slot />
  </button>
</template>

<script setup lang="ts">
const props = defineProps<{
  additionalClass?: string;
  type?: "button" | "submit" | "reset";
  onClick?: () => void;
}>();

const handleClick = () => {
  if (props.onClick) {
    props.onClick();
  }
};
</script>

<style scoped>
button {
  background-color: var(--button-bg-color, #0088cc);
  color: var(--button-text-color, white);
  padding: 10px 15px;
  border: none;
  border-radius: 2px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

button:hover {
  background-color: var(--button-hover-bg-color, #0084c6);
}

button:active {
  background-color: var(--button-active-bg-color, #0084c6);
}
</style>
