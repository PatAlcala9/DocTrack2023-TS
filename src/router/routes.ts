import { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('src/layouts/DoctrackLayout.vue'),
    children: [
      { path: '', component: () => import('pages/LoginPage.vue') },
      { path: '/received', component: () => import('pages/inquiry/ReceivedPage.vue') },
      { path: '/released', component: () => import('pages/inquiry/ReleasedPage.vue') },

      { path: '/dashboard', component: () => import('pages/IndexPage.vue') },
      { path: '/register', component: () => import('pages/RegisterPage.vue') },

      { path: '/incoming', component: () => import('pages/logged/IncomingPage.vue') },
      { path: '/outgoing', component: () => import('pages/logged/OutgoingPage.vue') }
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
]

export default routes
