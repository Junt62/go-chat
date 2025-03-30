import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
import ChatView from '@/views/ChatView.vue'
import ChatPage from '@/components/ChatPage.vue'
import { useUserStore } from '@/stores/user'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import SettingsView from '@/views/SettingsView.vue'
import DiscoveryView from '@/views/DiscoveryView.vue'
import CreateView from '@/views/CreateView.vue'

const routes = [
  {
    path: 'chat/:chatId',
    component: ChatPage,
    props: true,
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/chat',
      name: 'chat',
      component: ChatView,
      meta: { requiresAuth: true },
    },
    {
      path: '/create',
      name: 'create',
      component: CreateView,
      meta: { requiresAuth: true },
    },
    {
      path: '/discovery',
      name: 'discovery',
      component: DiscoveryView,
      meta: { requaresAuth: true },
    },
    {
      path: '/settings',
      name: 'settings',
      component: SettingsView,
      meta: { requiresAuth: true },
    },
  ],
})

router.beforeEach((to, from, next) => {
  const userStore = useUserStore()

  if (to.meta.requiresAuth && !userStore.token) {
    next('/login')
    // } else if (userStore.isTokenExpired()) {
    //   next('/login')
  } else {
    next()
  }
})

export default router
