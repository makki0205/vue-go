import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from './components/pages/index.vue'
import NotFound from './components/pages/notFound.vue'

Vue.component('index', Index)

Vue.use(VueRouter)

const routes = [
    { path: "/", component: Index },
    { path: "*", component: NotFound },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
    router
}).$mount("#app")
