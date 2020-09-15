import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from '../components/Home.vue'
import TitleModal from '../components/TitleModal.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/Modal',
        name: 'TitleModal',
        component: TitleModal
    }
];

const router = new VueRouter({
    mode: 'history',
    routes
})

export default router;