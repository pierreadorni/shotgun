import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import EventView from '../views/EventView.vue'
import NewEventView from "@/views/NewEventView.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/event/:id',
            name: 'event',
            component: EventView
        },
        {
            path: '/event/new',
            name: 'newEvent',
            component: NewEventView
        }
    ]
})

export default router
