import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import OrdersView from '@/views/OrdersView.vue'
import ItemsView from '@/views/ItemsView.vue'
import InventoryView from '@/views/InventoryView.vue'
import UsersView from '@/views/UsersView.vue'
import { useAuthStore } from '../stores/authStores'
import { refreshingToken } from "@/api/auth"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/orders',
      name: 'orders',
      component: OrdersView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/items',
      name: 'items',
      component: ItemsView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/inventory',
      name: 'inventory',
      component: InventoryView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/users',
      name: 'users',
      component: UsersView,
      meta: {
        requiresAuth: true,
      },
    },
  ],
})

router.beforeEach(async (to) => {
  const auth = useAuthStore()

  const tokenIsValid = auth.validateToken()
  if (to.meta.requiresAuth && !tokenIsValid) {
    const response = await refreshingToken({
      refreshToken: auth.refreshToken
    })
    if(response.status === 200){
      auth.setTokens(
        response.data.data.token,
        response.data.data.refreshToken
      );
      return // Retry the original navigation after refreshing the token
    }
    return '/login'
  }
})

export default router
