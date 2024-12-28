<template>
    <div class="login-container">
      <h1>Login</h1>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">Username</label>
          <input
            id="username"
            v-model="form.username"
            type="text"
            placeholder="Enter your username"
            required
          />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input
            id="password"
            v-model="form.password"
            type="password"
            placeholder="Enter your password"
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>
      <p>
        Don't have an account? <router-link to="/signup">Sign up</router-link>
      </p>
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from "vue";
  import { useAuthStore } from "@/stores/authStore";
  import { useRouter } from "vue-router";
  
  const router = useRouter();
  const authStore = useAuthStore();
  
  const form = ref({
    username: "",
    password: "",
  });
  
  const errorMessage = ref("");
  
  const handleLogin = async () => {
    try {
      await authStore.login(form.value);
      router.push("/");
    } catch {
      errorMessage.value = "Invalid credentials. Please try again.";
    }
  };
  </script>
  
  <style scoped>
  .login-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 8px;
  }
  .error-message {
    color: red;
    margin-top: 10px;
  }
  </style>