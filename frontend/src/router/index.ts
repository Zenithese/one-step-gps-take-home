// import { createRouter, createWebHistory } from "vue-router";
// import DashboardView from "@/views/DashboardView.vue";

// const routes = [
//   { path: "/", component: DashboardView },
// ];

// const router = createRouter({
//   history: createWebHistory(),
//   routes,
// });

// export default router;

import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from '@/stores/authStore';

const routes = [
  { path: '/login', component: () => import('@/views/LoginView.vue') },
  { path: '/signup', component: () => import('@/views/SignupView.vue') },
  { path: '/', component: () => import('@/views/DashboardView.vue'), meta: { requiresAuth: true } },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();

  if (!authStore.isAuthenticated) {
    await authStore.checkSession();
  }

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login');
  } else {
    next();
  }
});


export default router;