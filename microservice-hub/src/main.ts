import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import {createPinia} from 'pinia'

const pinia = createPinia();
createApp(App).use(pinia).use(vuetify).use(router).mount('#app')
