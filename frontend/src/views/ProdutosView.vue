<script setup>

import { ref, onMounted, computed } from 'vue'
import api from '../services/api'
import Swal from 'sweetalert2'

// =======================
// LISTA DE PRODUTOS
// =======================
const produtos = ref([])

// =======================
// LISTA DE CATEGORIAS
// =======================
const categorias = ref([])

// =======================
// FILTROS
// =======================
const filtroBusca = ref('')
const filtroTipo = ref('')
const filtroCategoria = ref('')
const filtroEstoqueBaixo = ref(false)

// =======================
// CAMPOS DO FORMULÁRIO
// =======================
const tipo = ref('EQUIPAMENTO')

const nome = ref('')
const categoria_id = ref('')

const marca = ref('')
const modelo = ref('')

const quantidade = ref(1)
const quantidade_minima = ref(0)

const setor = ref('')

const descricao = ref('')
const observacao = ref('')

// =======================
// CONTROLE DE EDIÇÃO
// =======================
const modoEdicao = ref(false)
const idEditando = ref(null)

// =======================
// PRODUTOS FILTRADOS
// =======================
const produtosFiltrados = computed(() => {

  return produtos.value.filter(produto => {

    const busca = filtroBusca.value.toLowerCase()

    const nomeProduto = produto.nome?.toLowerCase() || ''
    const marcaProduto = produto.marca?.toLowerCase() || ''
    const modeloProduto = produto.modelo?.toLowerCase() || ''

    // BUSCA TEXTO
    const correspondeBusca =

      nomeProduto.includes(busca) ||
      marcaProduto.includes(busca) ||
      modeloProduto.includes(busca)

    // FILTRO TIPO
    const correspondeTipo =

      !filtroTipo.value ||
      produto.tipo === filtroTipo.value

    // FILTRO CATEGORIA
    const correspondeCategoria =

      !filtroCategoria.value ||
      produto.categoria_id == filtroCategoria.value

    // FILTRO ESTOQUE
    const correspondeEstoque =

      !filtroEstoqueBaixo.value ||
      produto.quantidade <= produto.quantidade_minima

    return (

      correspondeBusca &&
      correspondeTipo &&
      correspondeCategoria &&
      correspondeEstoque
    )
  })
})

// =======================
// CARREGAR PRODUTOS
// =======================
// =======================
// CARREGAR PRODUTOS
// =======================
async function carregarProdutos() {

  try {

    const resposta = await api.get('/produtos')

    produtos.value = resposta.data

  } catch (erro) {

    console.error('Erro ao carregar produtos:', erro)
  }
}

// =======================
// CARREGAR CATEGORIAS
// =======================
async function carregarCategorias() {

  try {

    const resposta = await api.get('/categorias')

    categorias.value = resposta.data

  } catch (erro) {

    console.error('Erro ao carregar categorias:', erro)
  }
}

// =======================
// SALVAR PRODUTO
// =======================
async function salvarProduto() {

  if (!nome.value.trim() || !categoria_id.value) {

    Swal.fire({
      icon: 'warning',
      title: 'Preencha os campos obrigatórios'
    })

    return
  }

  try {

    const payload = {

      categoria_id: categoria_id.value,

      tipo: tipo.value,

      nome: nome.value,

      marca: marca.value,

      modelo: modelo.value,

      quantidade: quantidade.value,

      quantidade_minima: quantidade_minima.value,

      setor: setor.value,

      descricao_tecnica: descricao.value,

      observacao: observacao.value
    }

    if (modoEdicao.value) {

      await api.put(`/produtos/${idEditando.value}`, payload)

    } else {

      await api.post('/produtos', payload)
    }

    Swal.fire({
      icon: 'success',
      title: 'Sucesso',
      text: 'Produto salvo com sucesso!'
    })

    limparFormulario()

    carregarProdutos()

  } catch (erro) {

    console.error('Erro ao salvar produto:', erro)

    Swal.fire({
      icon: 'error',
      title: 'Erro ao salvar produto'
    })
  }
}

// =======================
// EDITAR
// =======================
function editarProduto(produto) {

  modoEdicao.value = true

  idEditando.value = produto.id

  tipo.value = produto.tipo
  nome.value = produto.nome
  categoria_id.value = produto.categoria_id
  marca.value = produto.marca
  modelo.value = produto.modelo
  quantidade.value = produto.quantidade
  quantidade_minima.value = produto.quantidade_minima
  setor.value = produto.setor
  descricao.value = produto.descricao_tecnica
  observacao.value = produto.observacao

  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}

// =======================
// LIMPAR FORMULÁRIO
// =======================
function limparFormulario() {

  modoEdicao.value = false

  idEditando.value = null

  tipo.value = 'EQUIPAMENTO'
  nome.value = ''
  categoria_id.value = ''
  marca.value = ''
  modelo.value = ''
  quantidade.value = 1
  quantidade_minima.value = 0
  setor.value = ''
  descricao.value = ''
  observacao.value = ''
}

// =======================
// MOVIMENTAR ESTOQUE
// =======================
async function movimentarEstoque(produto) {

  const { value: formValues } = await Swal.fire({

    title: `Movimentar: ${produto.nome}`,

    html: `

      <select id="tipoMov" class="swal2-input">

        <option value="ENTRADA">
          ENTRADA
        </option>

        <option value="SAIDA">
          SAÍDA
        </option>

      </select>

      <input
        id="quantidadeMov"
        type="number"
        class="swal2-input"
        placeholder="Quantidade"
      >

      <input
        id="obsMov"
        type="text"
        class="swal2-input"
        placeholder="Observação"
      >
    `,

    focusConfirm: false,

    showCancelButton: true,

    confirmButtonText: 'Salvar',

    cancelButtonText: 'Cancelar',

    preConfirm: () => {

      return {

        tipo:
          document.getElementById('tipoMov').value,

        quantidade: Number(
          document.getElementById('quantidadeMov').value
        ),

        observacao:
          document.getElementById('obsMov').value
      }
    }
  })

  if (!formValues) return

  if (!formValues.quantidade || formValues.quantidade <= 0) {

    Swal.fire({
      icon: 'warning',
      title: 'Quantidade inválida'
    })

    return
  }

  try {

    await api.post('/movimentacoes', {

      produto_id: produto.id,

      tipo: formValues.tipo,

      quantidade: formValues.quantidade,

      observacao: formValues.observacao
    })

    Swal.fire({

      icon: 'success',

      title: 'Movimentação realizada'
    })

    carregarProdutos()

  } catch (erro) {

    console.error(erro)

    Swal.fire({

      icon: 'error',

      title: 'Erro ao movimentar estoque'
    })
  }
}

// =======================
// EXCLUIR
// =======================
async function excluirProduto(id) {

  const resultado = await Swal.fire({

    icon: 'warning',

    title: 'Excluir produto',

    text: 'Informe o motivo da exclusão.',

    input: 'text',

    inputPlaceholder: 'Digite o motivo da exclusão',

    confirmButtonText: 'Excluir',

    cancelButtonText: 'Cancelar',

    showCancelButton: true,

    inputValidator: (value) => {

      if (!value || value.trim() === '') {

        return 'Informe o motivo da exclusão!'
      }
    }
  })

  if (!resultado.isConfirmed) return

  const observacao = resultado.value

  try {

    await api.delete(`/produtos/${id}`, {

      data: {
        observacao: observacao
      }
    })

    Swal.fire({
      icon: 'success',
      title: 'Sucesso',
      text: 'Produto excluído com sucesso!'
    })

    carregarProdutos()

  } catch (erro) {

    console.error('Erro ao excluir produto:', erro)

    Swal.fire({
      icon: 'error',
      title: 'Erro',
      text: 'Erro ao excluir produto.'
    })
  }
}

//Gerar QR Code
const gerarQRCode = async (id) => {
   
  try {

    const response = await api.get(
      `/equipamentos/${id}/qrcode`
    )

    console.log(response.data)

    alert('QR Code gerado com sucesso')

  } catch (error) {

    console.error(error)

    alert('Erro ao gerar QR Code')

  }

  
}

const historico = ref([])
const modalHistorico = ref(false)

async function abrirHistorico(id) {

  console.log("Produto clicado:", id)

  try {

    const response = await api.get(
      `/produtos/${id}/historico`
    )

    console.log(response.data)

    historico.value = Array.isArray(response.data)
      ? response.data
      : []

    modalHistorico.value = true

  } catch (error) {
    console.error(error)
  }
}

// =======================
// INIT
// =======================
onMounted(() => {

  carregarProdutos()
  carregarCategorias()
})

</script>

<template>

  <div class="pagina">

    <!-- CARDS -->
    <div class="cards-resumo">

      <div class="card-info">

        <h3>Total Produtos</h3>

        <span>
          {{ produtos.length }}
        </span>

      </div>

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

      <div class="card-info">

        <h3>Categorias</h3>

        <span>
          {{ categorias.length }}
        </span>

      </div>

    </div>

    <!-- FORMULÁRIO -->
    <div class="card-formulario">

      <h1>Cadastro de Produtos</h1>

      <div class="bloco">

        <h2>Dados principais</h2>

        <div class="grid">

          <div class="campo">

            <label>Tipo</label>

            <select v-model="tipo">

              <option value="EQUIPAMENTO">
                EQUIPAMENTO
              </option>

              <option value="COMPONENTE">
                COMPONENTE
              </option>

              <option value="SETOR">
                SETOR
              </option>

              <option value="DVR">
                DVR
              </option>

            </select>

          </div>

          <div class="campo">

            <label>Categoria</label>

            <select v-model="categoria_id">

              <option disabled value="">
                Selecione
              </option>

              <option
                v-for="categoria in categorias"
                :key="categoria.id"
                :value="categoria.id"
              >
                {{ categoria.nome }}
              </option>

            </select>

          </div>

          <div class="campo">
            <label>Nome</label>
            <input type="text" v-model="nome" />
          </div>

          <div class="campo">
            <label>Marca</label>
            <input type="text" v-model="marca" />
          </div>

          <div class="campo">
            <label>Modelo</label>
            <input type="text" v-model="modelo" />
          </div>

          <div class="campo">
            <label>Setor</label>
            <input type="text" v-model="setor" />
          </div>

        </div>

      </div>

      <!-- ESTOQUE -->
      <div class="bloco">

        <h2>Controle de estoque</h2>

        <div class="grid">

          <div class="campo">
            <label>Quantidade</label>
            <input type="number" v-model="quantidade" />
          </div>

          <div class="campo">
            <label>Quantidade mínima</label>
            <input type="number" v-model="quantidade_minima" />
          </div>

        </div>

      </div>

      <!-- INFO -->
      <div class="bloco">

        <h2>Informações técnicas</h2>

        <textarea
          v-model="descricao"
          rows="3"
        ></textarea>

      </div>

      <!-- OBS -->
      <div class="bloco">

        <h2>Observações</h2>

        <textarea
          v-model="observacao"
          rows="3"
        ></textarea>

      </div>

      <button
        type="button"
        class="botao"
        @click="salvarProduto"
      >
        {{ modoEdicao ? 'Atualizar Produto' : 'Cadastrar Produto' }}
      </button>

      <button
        v-if="modoEdicao"
        class="botao-cancelar"
        @click="limparFormulario"
      >
        Cancelar
      </button>

    </div>

    <!-- FILTROS -->
    <div class="card-filtros">

      <h2>Filtros</h2>

      <div class="grid-filtros">

        <div class="campo">

          <label>Buscar</label>

          <input
            type="text"
            v-model="filtroBusca"
            placeholder="Nome, marca ou modelo"
          />

        </div>

        <div class="campo">

          <label>Tipo</label>

          <select v-model="filtroTipo">

            <option value="">Todos</option>

            <option value="EQUIPAMENTO">
              EQUIPAMENTO
            </option>

            <option value="COMPONENTE">
              COMPONENTE
            </option>

            <option value="SETOR">
              SETOR
            </option>

            <option value="DVR">
              DVR
            </option>

          </select>

        </div>

        <div class="campo">

          <label>Categoria</label>

          <select v-model="filtroCategoria">

            <option value="">Todas</option>

            <option
              v-for="categoria in categorias"
              :key="categoria.id"
              :value="categoria.id"
            >
              {{ categoria.nome }}
            </option>

          </select>

        </div>

        <div class="campo-checkbox">

          <input
            type="checkbox"
            v-model="filtroEstoqueBaixo"
          />

          <label>
            Mostrar somente estoque crítico
          </label>

        </div>

      </div>

    </div>

    <!-- TABELA -->
    <div class="card-tabela">

      <h2>Produtos cadastrados</h2>

      <table>

        <thead>

          <tr>

            <th>ID</th>
            <th>Nome</th>
            <th>Tipo</th>
            <th>Categoria</th>
            <th>Quantidade</th>
            <th>Ações</th>

          </tr>

        </thead>

        <tbody>

          <tr
            v-for="produto in produtosFiltrados"
            :key="produto.id"
            :class="{
              'linha-critica':
                produto.quantidade <= produto.quantidade_minima
            }"
          >

            <td>{{ produto.id }}</td>
            <td>{{ produto.nome }}</td>
            <td>{{ produto.tipo }}</td>
            <td>{{ produto.categoria_nome }}</td>
            <td>{{ produto.quantidade }}</td>

            <td>

              <button
                class="botao-movimentar"
                @click="movimentarEstoque(produto)"
              >
                Movimentar
              </button>

              <button
                class="botao-editar"
                @click="editarProduto(produto)"
              >
                Editar
              </button>

              <button
                class="botao-excluir"
                @click="excluirProduto(produto.id)"
              >
                Excluir
              </button>
              <button 
              @click="gerarQRCode(produto.id)" 
              class="botao-gerarQRCode">
              Gerar QR
              </button>
             <button
  class="botao-historico"
  @click="abrirHistorico(produto.id)"
>
  Histórico
</button>
            </td>

          </tr>

        </tbody>

      </table>

    </div>

  </div>
<!-- MODAL HISTÓRICO -->
<div
  v-if="modalHistorico"
  class="modal-overlay"
>

  <div class="modal-historico">

    <h2>Histórico do Produto</h2>

    <table>

      <thead>
        <tr>
          <th>Data</th>
          <th>Usuário</th>
          <th>Ação</th>
          <th>Descrição</th>
        </tr>
      </thead>

      <tbody>

        <tr
          v-for="item in historico"
          :key="item.id"
        >
          <td>{{ item.data_registro }}</td>
          <td>{{ item.usuario }}</td>
          <td>{{ item.acao }}</td>
          <td>{{ item.descricao }}</td>
        </tr>

        <tr v-if="historico.length === 0">
          <td colspan="4">
            Nenhum histórico encontrado
          </td>
        </tr>

      </tbody>

    </table>

    <button
      class="botao"
      @click="modalHistorico = false"
    >
      Fechar
    </button>

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
  font-size: 15px;
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
CARDS PRINCIPAIS
======================= */

.card-formulario,
.card-tabela,
.card-filtros {
  background-color: white;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}

h1,
h2 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.bloco {
  margin-bottom: 35px;
}

.grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
}

.grid-filtros {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.campo {
  display: flex;
  flex-direction: column;
}

.campo-checkbox {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 35px;
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

textarea {
  width: 100%;
  min-height: 90px;
  padding: 12px;
  border: 1px solid #dcdcdc;
  border-radius: 8px;
  font-size: 14px;
}

/* =======================
BOTÕES
======================= */

.botao {
  background-color: #2c3e50;
  color: white;
  border: none;
  padding: 14px 24px;
  border-radius: 8px;
  cursor: pointer;
}

.botao:hover {
  background-color: #1f2d3a;
}

.botao-movimentar {
  background-color: #f39c12;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
  margin-right: 5px;
}

.botao-movimentar:hover {
  background-color: #d68910;
}

.botao-editar {
  background-color: #3498db;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
  margin-right: 5px;
}

.botao-excluir {
  background-color: #ff1900;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
}

.botao-cancelar {
  background-color: #95a5a6;
  color: white;
  border: none;
  padding: 10px 18px;
  border-radius: 8px;
  cursor: pointer;
  margin-left: 10px;
}
.botao-gerarQRCode {
  background-color: #27ae60;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
  margin-left: 9px;
}
/* =======================
TABELA
======================= */

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  padding: 14px;
  border-bottom: 1px solid #ececec;
}

th {
  background-color: #2c3e50;
  color: white;
}

tr:hover {
  background-color: #f8f9fa;
}

/* =======================
ESTOQUE CRÍTICO
======================= */

.linha-critica {
  background-color: #fff5f5;
}

.linha-critica td {
  color: #ee1800;
  font-weight: bold;
}

/* =======================
RESPONSIVO
======================= */

@media (max-width: 1200px) {

  .grid-filtros,
  .cards-resumo {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 700px) {

  .grid,
  .grid-filtros,
  .cards-resumo {
    grid-template-columns: 1fr;
  }
}
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999;
}

.modal-historico {
  background: white;
  width: 80%;
  max-width: 900px;
  padding: 20px;
  border-radius: 10px;
  max-height: 80vh;
  overflow-y: auto;
}

.botao-historico {
  background-color: #8e44ad;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
  margin-left: 9px;
}

.botao-historico:hover {
  background-color: #6c3483;
}
</style>