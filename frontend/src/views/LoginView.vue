<template>
  <div class="login-container">
    <div class="login-form">
      <h1 class="text-center mb-4">Login</h1>
      <form @submit.prevent="handleLogin">
        <InputField
          id="username"
          v-model="form.username"
          placeholder="Username"
          required
        />
        <InputField
          id="password"
          v-model="form.password"
          type="password"
          placeholder="Password"
          required
        />
        <ButtonComponent type="submit" additionalClass="centered">
          Login
        </ButtonComponent>
      </form>
      <p class="text-xs mt-4">
        Don't have an account? <router-link to="/signup" class="text-[#0088cc]">Sign up</router-link>
      </p>
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from "vue";
  import { useAuthStore } from "@/stores/authStore";
  import { useRouter } from "vue-router";
  import ButtonComponent from "@/components/ButtonComponent.vue";
  import InputField from "@/components/InputFieldComponent.vue";
  
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
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(135deg, #0088cc, forestgreen);
}

.login-form {
  max-width: 400px;
  margin: auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 12px;
  background-color: white;
  box-shadow: 0px 8px 15px rgba(0, 0, 0, 0.2);
}
  .error-message {
    color: red;
    margin-top: 10px;
  }

  .centered {
    display: block;
    margin: 20px auto 0 auto;
  }
  </style>