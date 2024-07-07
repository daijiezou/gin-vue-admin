import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [{
    path: '/',
    redirect: '/login'
},
    {
        path: '/init',
        name: 'Init',
        component: () => import('@/view/init/index.vue')
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/view/login/index.vue')
    },
    {
        path: '/:catchAll(.*)',
        meta: {
            closeTab: true,
        },
        component: () => import('@/view/error/index.vue')
    },
    // {
    //     path: '/workplace/detail/:workplaceId',
    //     name: "workplaceDetail",
    //     component: () => import('@/view/workerspace/workplace/workplace_detail.vue')
    // }


]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router
