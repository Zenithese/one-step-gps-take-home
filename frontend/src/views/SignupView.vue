<template>
  <div class="signup-container">
    <div class="signup-form">
      <h1 class="text-center mb-4">Sign Up</h1>
      <form @submit.prevent="handleSignup">
        <InputField
          id="username"
          v-model="form.username"
          placeholder="Choose a username"
          required
        />
        <InputField
          id="password"
          v-model="form.password"
          type="password"
          placeholder="Create a password"
          required
        />
        <InputField
          id="confirm-password"
          v-model="form.confirmPassword"
          type="password"
          placeholder="Re-enter your password"
          required
        />
        <ButtonComponent type="submit" additionalClass="centered">
          Sign Up
        </ButtonComponent>
      </form>
      <p class="text-xs mt-4">
        Already have an account? 
        <router-link to="/login" class="text-[#0088cc]">Login</router-link>
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
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(135deg, forestgreen, #0088cc);
}

.signup-form {
  max-width: 400px;
  margin: auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 12px;
  background-color: white;
  box-shadow: 0px 8px 15px rgba(0, 0, 0, 0.2);
}

.text-center {
  text-align: center;
}

.mb-4 {
  margin-bottom: 16px;
}

.text-xs {
  font-size: 12px;
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