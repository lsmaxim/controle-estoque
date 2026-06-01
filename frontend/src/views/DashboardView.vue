<script setup>

import { ref, onMounted } from 'vue'
import api from '../services/api'

// =======================
// DADOS
// =======================
const produtos = ref([])
const categorias = ref([])
const movimentacoes = ref([])
const agendamentos = ref([])

// =======================
// CARREGAR DADOS
// =======================
async function carregarDados() {

  try {

    // PRODUTOS
    const produtosResponse =
      await api.get('/produtos')

    produtos.value =
      produtosResponse.data

    // CATEGORIAS
    const categoriasResponse =
      await api.get('/categorias')

    categorias.value =
      categoriasResponse.data

    // MOVIMENTAÇÕES
    const movResponse =
      await api.get('/movimentacoes')

    movimentacoes.value =
      movResponse.data

    // AGENDAMENTOS
    const agendaResponse =
      await api.get('/agendamentos')

    agendamentos.value =
      agendaResponse.data

  } catch (erro) {

    console.error(
      'Erro ao carregar dashboard:',
      erro
    )
  }
}

// =======================
// FORMATAR DATA
// =======================
function formatarData(data) {

  if (!data) return '-'

  return data
}

// =======================
// INIT
// =======================
onMounted(() => {

  carregarDados()
})

</script>

<template>

  <div class="pagina">

    <!-- TOPO -->
    <div class="topo">

      <div>

        <h1>Dashboard</h1>

        <p>
          Visão geral do sistema
        </p>

      </div>

    </div>

    <!-- CARDS -->
    <div class="cards-resumo">

      <!-- PRODUTOS -->
      <div class="card-info">

        <h3>Total Produtos</h3>

        <span>
          {{ produtos.length }}
        </span>

      </div>

      <!-- ESTOQUE -->
      <div class="card-info critico">

        <h3>Estoque Crítico</h3>

        <span>
          {{
            produtos.filter(
              p => p.quantidade <= p.quantidade_minima
            ).length
          }}
        </span>

      </div>

      <!-- EQUIPAMENTOS -->
      <div class="card-info">

        <h3>Equipamentos</h3>

        <span>
          {{
            produtos.filter(
              p => p.tipo === 'EQUIPAMENTO'
            ).length
          }}
        </span>

      </div>

      <!-- CATEGORIAS -->
      <div class="card-info">

        <h3>Categorias</h3>

        <span>
          {{ categorias.length }}
        </span>

      </div>

    </div>

    <!-- GRID -->
    <div class="grid-dashboard">

      <!-- MOVIMENTAÇÕES -->
      <div class="card-dashboard">

        <h2>
          Últimas movimentações
        </h2>

        <div
          class="mov-item"
          v-for="mov in movimentacoes.slice(0, 6)"
          :key="mov.id"
        >

          <div class="mov-info">

            <strong>
              {{ mov.produto }}
            </strong>

            <span class="mov-data">

              {{ formatarData(mov.data_movimentacao) }}

            </span>

          </div>

          <div class="mov-lateral">

            <span
              :class="[
                'badge',
                mov.tipo === 'ENTRADA'
                  ? 'entrada'
                  : 'saida'
              ]"
            >
              {{ mov.tipo }}
            </span>

            <strong
              :class="
                mov.tipo === 'ENTRADA'
                  ? 'texto-entrada'
                  : 'texto-saida'
              "
            >
              {{ mov.quantidade }}
            </strong>

          </div>

        </div>

      </div>

      <!-- AGENDAMENTOS -->
      <div class="card-dashboard">

        <h2>
          Próximos agendamentos
        </h2>

        <div
          class="agendamento"
          v-for="agenda in agendamentos.slice(0, 6)"
          :key="agenda.id"
        >

          <!-- ESQUERDA -->
          <div class="agenda-info">

            <strong>
              {{ agenda.titulo }}
            </strong>

            <p>
              {{
                new Date(
                  agenda.data_agendamento
                ).toLocaleDateString('pt-BR')
              }}
            </p>

          </div>

          <!-- DIREITA -->
          <span
  :class="[
    'status-agenda',

    new Date(agenda.data_agendamento) < new Date() &&
    agenda.status !== 'FINALIZADO'
      ? 'status-atrasado'

      : agenda.status === 'FINALIZADO'
        ? 'status-finalizado'
        : 'status-pendente'
  ]"
>
  {{
    new Date(agenda.data_agendamento) < new Date() &&
    agenda.status !== 'FINALIZADO'
      ? 'ATRASADO'
      : agenda.status
  }}
</span>

        </div>

      </div>

    </div>

  </div>

</template>

<style scoped>

.pagina {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

/* =======================
TOPO
======================= */

.topo h1 {
  color: #2c3e50;
}

.topo p {
  color: #7f8c8d;
  margin-top: 5px;
}

/* =======================
CARDS
======================= */

.cards-resumo {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.card-info {
  background: white;
  padding: 25px;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}

.card-info h3 {
  color: #7f8c8d;
  margin-bottom: 10px;
}

.card-info span {
  font-size: 34px;
  font-weight: bold;
  color: #2c3e50;
}

.card-info.critico span {
  color: #e74c3c;
}

/* =======================
GRID
======================= */

.grid-dashboard {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 20px;
}

.card-dashboard {
  background: white;
  padding: 26px;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}



.mov-item {

  display: flex;

  justify-content: space-between;

  align-items: center;

  padding: 15px 0;

  border-bottom: 1px solid #ececec;
}

.mov-info {

  display: flex;

  flex-direction: column;
}

.mov-data {

  font-size: 13px;

  color: #7f8c8d;

  margin-top: 5px;
}

.mov-lateral {

  display: flex;

  align-items: center;

  gap: 12px;
}

.texto-entrada {

  color: #27ae60;
}

.texto-saida {

  color: #e74c3c;
}

/* =======================
BADGES
======================= */

.badge {
  padding: 6px 12px;
  border-radius: 20px;
  color: white;
  font-size: 12px;
  font-weight: bold;
}

.badge.entrada {
  background-color: #27ae60;
}

.badge.saida {
  background-color: #e74c3c;
}

/* =======================
AGENDAMENTOS
======================= */

.agendamento {

  display: flex;

  justify-content: space-between;

  align-items: center;

  border-bottom: 1px solid #ececec;

  padding: 14px 0;
}

.agenda-info {

  display: flex;

  flex-direction: column;
}

.agenda-info p {

  margin-top: 5px;

  color: #7f8c8d;

  font-size: 14px;
}

/* =======================
STATUS AGENDAMENTO
======================= */

.status-agenda {

  padding: 6px 12px;

  border-radius: 20px;

  font-size: 12px;

  font-weight: bold;

  color: white;
}

/* FINALIZADO */

.status-finalizado {

  background-color: #27ae60;
}

/* PENDENTE */

.status-pendente {
  background-color: #ff7606;
}
.status-atrasado {

  background-color: #e73b27;
}
/* =======================
RESPONSIVO
======================= */

@media (max-width: 1200px) {

  .cards-resumo {
    grid-template-columns: repeat(2, 1fr);
  }

  .grid-dashboard {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 700px) {

  .cards-resumo {
    grid-template-columns: 1fr;
  }

  .agendamento {

    flex-direction: column;

    align-items: flex-start;

    gap: 10px;
  }
}

</style>