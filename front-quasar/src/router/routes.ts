import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/IndexPage.vue') }],
  },
  {
    path: '/events',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        component: () => import('pages/EventsPage.vue'),
      },
      {
        path: ':id',
        component: () => import('pages/EventPage.vue'),
      },
      {
        path: 'new',
        component: () => import('pages/NewEventPage.vue'),
      },
    ],
  },
  {
    path: '/my-events',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/MyEventsPage.vue') }],
  },
  {
    path: '/login',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/LoginPage.vue') }],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
