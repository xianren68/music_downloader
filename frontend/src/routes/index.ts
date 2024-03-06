import {createRouter, createWebHistory} from 'vue-router'
const routes = [
    {
        path: '/',
        name: 'search',
        component: () => import('@/view/Search.vue'),
        meta:{
            iconName:'icon-icon-sousuo',
        }
    },
    {
        path:'/download',
        name:'download',
        component:()=>import('@/view/Download.vue'),
        meta:{
            iconName:'icon-chuanshu',
        }
    },
    {
        path:'/setting',
        name:'setting',
        component:()=>import('@/view/Setting.vue'),
        meta:{
            iconName:'icon-shezhi',
        }
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})
export default router