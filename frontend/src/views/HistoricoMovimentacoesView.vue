<script setup>

import { ref, onMounted, computed } from 'vue'
import api from '../services/api'
import Swal from 'sweetalert2'

// =======================
// LISTA
// =======================
const movimentacoes = ref([])

// =======================
// FILTROS
// =======================
const filtroBusca = ref('')
const filtroTipo = ref('')
const filtroDataInicial = ref('')
const filtroDataFinal = ref('')

// =======================
// CARREGAR MOVIMENTAÇÕES
// =======================
async function carregarMovimentacoes() {

  try {

    const resposta = await api.get('/movimentacoes')

    movimentacoes.value = resposta.data

  } catch (erro) {

    console.error('Erro ao carregar movimentações:', erro)

    Swal.fire({
      icon: 'error',
      title: 'Erro ao carregar movimentações'
    })
  }
}

// =======================
// FILTROS
// =======================
const movimentacoesFiltradas = computed(() => {

  return movimentacoes.value.filter(mov => {

    // BUSCA
    const busca = filtroBusca.value.toLowerCase()

    const produto =
      mov.produto?.toLowerCase() || ''

    const observacao =
      mov.observacao?.toLowerCase() || ''

    const correspondeBusca =

      produto.includes(busca) ||
      observacao.includes(busca)

    // TIPO
    const correspondeTipo =

      !filtroTipo.value ||
      mov.tipo === filtroTipo.value

    // DATA
    const dataMov =
      mov.data_movimentacao?.split('T')[0]

    const correspondeDataInicial =

      !filtroDataInicial.value ||
      dataMov >= filtroDataInicial.value

    const correspondeDataFinal =

      !filtroDataFinal.value ||
      dataMov <= filtroDataFinal.value

    return (

      correspondeBusca &&
      correspondeTipo &&
      correspondeDataInicial &&
      correspondeDataFinal
    )
  })
})

// =======================
// LIMPAR FILTROS
// =======================
function limparFiltros() {

  filtroBusca.value = ''
  filtroTipo.value = ''
  filtroDataInicial.value = ''
  filtroDataFinal.value = ''
}

// =======================
// INIT
// =======================
onMounted(() => {

  carregarMovimentacoes()
})

</script>

<template>

  <div class="pagina">

    <!-- TOPO -->
    <div class="topo">

      <div>

        <h1>
          Histórico de Movimentações
        </h1>

        <p>
          Controle de entradas e saídas
        </p>

      </div>

    </div>

    <!-- CARDS -->
    <div class="cards-resumo">

      <!-- TOTAL -->
      <div class="card-info">

        <h3>Total</h3>

        <span>
          {{ movimentacoes.length }}
        </span>

      </div>

      <!-- ENTRADAS -->
      <div class="card-info entrada">

        <h3>Entradas</h3>

        <span>
          {{
            movimentacoes.filter(
              m => m.tipo === 'ENTRADA'
            ).length
          }}
        </span>

      </div>

      <!-- SAIDAS -->
      <div class="card-info saida">

        <h3>Saídas</h3>

        <span>
          {{
            movimentacoes.filter(
              m => m.tipo === 'SAIDA'
            ).length
          }}
        </span>

      </div>

    </div>

    <!-- FILTROS -->
    <div class="card-filtros">

      <h2>Filtros</h2>

      <div class="grid-filtros">

        <!-- BUSCA -->
        <div class="campo">

          <label>Buscar</label>

          <input
            type="text"
            v-model="filtroBusca"
            placeholder="Produto ou observação"
          />

        </div>

        <!-- TIPO -->
        <div class="campo">

          <label>Tipo</label>

          <select v-model="filtroTipo">

            <option value="">
              Todos
            </option>

            <option value="ENTRADA">
              ENTRADA
            </option>

            <option value="SAIDA">
              SAÍDA
            </option>

          </select>

        </div>

        <!-- DATA INICIAL -->
        <div class="campo">

          <label>Data inicial</label>

          <input
            type="date"
            v-model="filtroDataInicial"
          />

        </div>

        <!-- DATA FINAL -->
        <div class="campo">

          <label>Data final</label>

          <input
            type="date"
            v-model="filtroDataFinal"
          />

        </div>

      </div>

      <!-- BOTÕES -->
      <div class="acoes-filtros">

        <button
          class="botao-limpar"
          @click="limparFiltros"
        >
          Limpar filtros
        </button>

      </div>

    </div>

    <!-- TABELA -->
    <div class="card-tabela">

      <div class="cabecalho-tabela">

        <h2>
          Movimentações
        </h2>

      </div>

      <table>

        <thead>

          <tr>

            <th>ID</th>
            <th>Produto</th>
            <th>Tipo</th>
            <th>Quantidade</th>
            <th>Observação</th>
            <th>Data</th>

          </tr>

        </thead>

        <tbody>

          <tr
            v-for="mov in movimentacoesFiltradas"
            :key="mov.id"
          >

            <td>
              {{ mov.id }}
            </td>

            <td>
              {{ mov.produto }}
            </td>

            <td>

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

            </td>

            <td>
              {{ mov.quantidade }}
            </td>

            <td>
              {{ mov.observacao }}
            </td>
            
            <td>
              {{ mov.data_movimentacao }}
            </td>
           

          </tr>

        </tbody>

      </table>

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

.topo {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

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
  grid-template-columns: repeat(3, 1fr);
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

.card-info.entrada span {
  color: #27ae60;
}

.card-info.saida span {
  color: #e74c3c;
}

/* =======================
CARDS
======================= */

.card-filtros,
.card-tabela {
  background: white;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}

/* =======================
GRID FILTROS
======================= */

.grid-filtros {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.campo {
  display: flex;
  flex-direction: column;
}

label {
  margin-bottom: 8px;
  font-weight: bold;
  color: #555;
}

input,
select {
  padding: 12px;
  border: 1px solid #dcdcdc;
  border-radius: 8px;
  font-size: 14px;
}

/* =======================
BOTÕES
======================= */

.acoes-filtros {
  margin-top: 20px;
}

.botao-limpar {
  background-color: #7f8c8d;
  color: white;
  border: none;
  padding: 12px 20px;
  border-radius: 8px;
  cursor: pointer;
}

.botao-limpar:hover {
  background-color: #636e72;
}

/* =======================
TABELA
======================= */

.cabecalho-tabela {
  margin-bottom: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  padding: 14px;
  border-bottom: 1px solid #ececec;
  text-align: left;
}

th {
  background-color: #2c3e50;
  color: white;
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
RESPONSIVO
======================= */

@media (max-width: 1200px) {

  .grid-filtros {
    grid-template-columns: repeat(2, 1fr);
  }

  .cards-resumo {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 700px) {

  .grid-filtros,
  .cards-resumo {
    grid-template-columns: 1fr;
  }

  .topo {
    flex-direction: column;
    align-items: flex-start;
    gap: 15px;
  }
}

</style>