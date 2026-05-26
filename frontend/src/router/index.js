import { createRouter, createWebHistory } from 'vue-router'

import DashboardView from '../views/DashboardView.vue'
import CategoriasView from '../views/CategoriasView.vue'
import ProdutosView from '../views/ProdutosView.vue'
import VinculosView from '../views/VinculosView.vue'
import CalendarioView from '../views/CalendarioView.vue'
import RelatorioView from '../views/RelatorioView.vue'
import HistoricoMovimentacoesView from '../views/HistoricoMovimentacoesView.vue'

const routes = [

  // DASHBOARD
  {
    path: '/',
    component: DashboardView
  },

  // CATEGORIAS
  {
    path: '/categorias',
    component: CategoriasView
  },

  // PRODUTOS
  {
    path: '/produtos',
    component: ProdutosView
  },

  // VÍNCULOS
  {
    path: '/vinculos',
    component: VinculosView
  },

  // CALENDÁRIO
  {
    path: '/calendario',
    component: CalendarioView
  },

  // RELATÓRIOS
  {
    path: '/relatorios',
    component: RelatorioView
  },

  // MOVIMENTAÇÕES
  {
    path: '/movimentacoes',
    component: HistoricoMovimentacoesView
  }

]

const router = createRouter({

  history: createWebHistory(),

  routes
})

export default router