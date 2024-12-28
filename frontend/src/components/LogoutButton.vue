<template>
    <button :class="['logout-button', additionalClass]" @click="logoutUser">
      Logout
    </button>
  </template>
  
  <script setup lang="ts">
  import { useAuthStore } from '@/stores/authStore';
  import { useRouter } from 'vue-router';

  defineProps({
  additionalClass: {
    type: String,
    default: '',
  },
});
  
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
  .logout-button {
    background-color: #ff4d4f;
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
    transition: background-color 0.3s;
  }
  
  .logout-button:hover {
    background-color: #e04343;
  }
  
  .logout-button:active {
    background-color: #cc3a3a;
  }
  </style>