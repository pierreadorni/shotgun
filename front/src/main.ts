import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import './assets/main.css'

import { OhVueIcon, addIcons } from "oh-vue-icons";
import { MdAddcircle } from "oh-vue-icons/icons";

addIcons(MdAddcircle);


const app = createApp(App)

app.use(router)
app.component("v-icon", OhVueIcon);

app.mount('#app')
