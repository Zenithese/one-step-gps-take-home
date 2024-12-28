<template>
  <div class="input-field">
    <label class="input-label" :for="id">{{ label }}</label>
    <input
      :id="id"
      class="input-box"
      :type="type"
      :value="modelValue"
      @input="onInput"
      :placeholder="placeholder"
    />
    <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';

defineProps({
  id: String,
  label: String,
  type: {
    type: String,
    default: 'text'
  },
  modelValue: String,
  errorMessage: String,
  placeholder: String
});

const emit = defineEmits(['update:modelValue']);

const onInput = (event: Event) => {
  const target = event.target as HTMLInputElement | null;
  if (target) {
    emit('update:modelValue', target.value);
  }
};
</script>

<style scoped>
.input-field {
  margin-bottom: 1rem;
}

.input-label {
  display: block;
  font-size: small;
}

.input-box {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 1rem;
}

.input-box:focus {
  border-color: #007bff;
  outline: none;
}

.error-message {
  color: red;
  font-size: 0.875rem;
  margin-top: 0.5rem;
}
</style>