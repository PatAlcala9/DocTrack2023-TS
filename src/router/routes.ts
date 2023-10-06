import { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('src/layouts/DoctrackLayout.vue'),
    children: [
      { path: '', component: () => import('pages/LoginPage.vue') },
      { path: '/incominginquire', component: () => import('pages/inquiry/IncomingInquirePage.vue') },
      { path: '/outgoinginquire', component: () => import('pages/inquiry/OutgoingInquirePage.vue') },
      { path: '/qrscanner', component: () => import('pages/inquiry/QRScannerPage.vue') },
      { path: '/qrresult', component: () => import('pages/inquiry/QRResultPage.vue') },

      { path: '/dashboard', component: () => import('pages/IndexPage.vue') },
      { path: '/register', component: () => import('pages/RegisterPage.vue') },
      { path: '/notready', component: () => import('pages/NotReadyPage.vue') },
      { path: '/unsupported', component: () => import('pages/UnsupportedPage.vue') },

      { path: '/incoming', component: () => import('pages/logged/IncomingDashboardPage.vue') },
      { path: '/incomingadd', component: () => import('pages/logged/IncomingAddPage.vue') },
      { path: '/incomingmain', component: () => import('pages/logged/IncomingMainPage.vue') },
      { path: '/outgoing', component: () => import('pages/logged/OutgoingDashboardPage.vue') },
      { path: '/releasing', component: () => import('pages/logged/ReleasingPage.vue') },
      { path: '/inventory', component: () => import('pages/logged/InventoryPage.vue') },
      { path: '/inventoryfile', component: () => import('pages/logged/InventoryFileDocPage.vue') },
      { path: '/complaint', component: () => import('pages/logged/ComplaintDashboardPage.vue') },
      { path: '/complaintmain', component: () => import('pages/logged/ComplaintMainPage.vue') },
      { path: '/complaintadd', component: () => import('pages/logged/ComplaintAddPage.vue') },
      { path: '/complaintinspect', component: () => import('pages/logged/ComplaintInspectionPage.vue') },
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
