<template>
  <div class="signup-container">
    <h1>Sign Up</h1>
    <form @submit.prevent="handleSignup">
      <div class="form-group">
        <label for="username">Username</label>
        <input
          id="username"
          v-model="form.username"
          type="text"
          placeholder="Choose a username"
          required
        />
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input
          id="password"
          v-model="form.password"
          type="password"
          placeholder="Create a password"
          required
        />
      </div>
      <div class="form-group">
        <label for="confirm-password">Confirm Password</label>
        <input
          id="confirm-password"
          v-model="form.confirmPassword"
          type="password"
          placeholder="Re-enter your password"
          required
        />
      </div>
      <button type="submit">Sign Up</button>
    </form>
    <p>Already have an account? <router-link to="/login">Login</router-link></p>
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
  confirmPassword: "",
});

const errorMessage = ref("");

const handleSignup = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    errorMessage.value = "Passwords do not match.";
    return;
  }

  try {
    await authStore.signup(form.value);
    router.push("/");
  } catch {
    errorMessage.value = "Signup failed. Please try again.";
  }
};
</script>

<style scoped>
.signup-container {
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
