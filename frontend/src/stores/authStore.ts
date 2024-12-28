import { defineStore } from "pinia";
import { signup, login, checkSession, logout } from "@/services/authServices";

export const useAuthStore = defineStore("authStore", {
  state: () => ({
    isAuthenticated: false,
  }),
  actions: {
    async signup({ username, password }: { username: string; password: string }) {
      try {
        await signup({ username, password });
        this.isAuthenticated = true;
      } catch (error) {
        console.error("Signup failed:", error);
        throw error;
      }
    },

    async login({ username, password }: { username: string; password: string }) {
      try {
        await login({ username, password });
        this.isAuthenticated = true;
      } catch (error) {
        console.error("Login failed:", error);
        throw error;
      }
    },

    async checkSession() {
      try {
        await checkSession();
        this.isAuthenticated = true;
      } catch (error) {
        console.error("Session check failed:", error);
      }
    },

    async logout() {
      try {
        await logout();
        this.isAuthenticated = false;
      } catch (error) {
        console.error("Logout failed:", error);
        throw error;
      }
    },
  },
});