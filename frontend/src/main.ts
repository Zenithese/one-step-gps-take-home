import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from "pinia";
import App from './App.vue'
import router from './router'
import piniaLogger from './stores/plugin/logger'

const app = createApp(App)

const pinia = createPinia();
pinia.use(piniaLogger); 
app.use(pinia);
app.use(router)

app.mount('#app')
