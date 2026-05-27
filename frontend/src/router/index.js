import { createRouter, createWebHistory } from 'vue-router'

import DashboardView from '../views/DashboardView.vue'
import CategoriasView from '../views/CategoriasView.vue'
import ProdutosView from '../views/ProdutosView.vue'
import VinculosView from '../views/VinculosView.vue'
import CalendarioView from '../views/CalendarioView.vue'
import RelatorioView from '../views/RelatorioView.vue'
import HistoricoMovimentacoesView from '../views/HistoricoMovimentacoesView.vue'
import LoginView from '../views/LoginView.vue'

const routes = [

  // LOGIN
  {
    path: '/login',
    component: LoginView
  },

  // DASHBOARD
  {
    path: '/',
    component: DashboardView
  },

  {
    path: '/categorias',
    component: CategoriasView
  },

  {
    path: '/produtos',
    component: ProdutosView
  },

  {
    path: '/vinculos',
    component: VinculosView
  },

  {
    path: '/calendario',
    component: CalendarioView
  },

  {
    path: '/relatorios',
    component: RelatorioView
  },

  {
    path: '/movimentacoes',
    component: HistoricoMovimentacoesView
  }
]

const router = createRouter({

  history: createWebHistory(),

  routes
})

// ===========================
// PROTEÇÃO LOGIN
// ===========================

router.beforeEach((to, from, next) => {

  const token =
    localStorage.getItem('token')

  // SE NÃO ESTIVER LOGADO
  if (!token && to.path !== '/login') {

    return next('/login')
  }

  next()
})

export default router