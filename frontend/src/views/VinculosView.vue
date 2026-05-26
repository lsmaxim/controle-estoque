<script setup>

import { ref, onMounted } from 'vue'
import Swal from 'sweetalert2'
import api from '../services/api'

// LISTAS
const equipamentos = ref([])
const componentes = ref([])
const vinculos = ref([])

// CAMPOS
const produto_pai_id = ref('')
const componente_id = ref('')
const quantidade = ref(1)

// BUSCA
const busca = ref('')

// CARREGAR PRODUTOS
async function carregarProdutos() {

  try {

    const resposta = await api.get('/produtos')

    // FILTRA EQUIPAMENTOS
    equipamentos.value = resposta.data.filter(
      produto => produto.tipo === 'EQUIPAMENTO'
    )

    // FILTRA COMPONENTES
    componentes.value = resposta.data.filter(
      produto => produto.tipo === 'COMPONENTE'
    )

  } catch (erro) {

    console.error('Erro ao carregar produtos:', erro)
  }
}

// CARREGAR VÍNCULOS
async function carregarVinculos() {

  try {

    const resposta = await api.get('/vinculos')

    vinculos.value = resposta.data

  } catch (erro) {

    console.error('Erro ao carregar vínculos:', erro)
  }
}

// CRIAR VÍNCULO
async function criarVinculo() {

  // VALIDAÇÕES
  if (
    !produto_pai_id.value ||
    !componente_id.value ||
    quantidade.value < 1
  ) {

    Swal.fire({
      icon: 'error',
      title: 'Erro',
      text: 'Preencha todos os campos corretamente.'
    })

    return
  }

  try {

    // ENVIA PARA API
    await api.post('/vinculos', {

      produto_pai_id: produto_pai_id.value,

      componente_id: componente_id.value,

      quantidade: quantidade.value
    })

    // LIMPA CAMPOS
    produto_pai_id.value = ''
    componente_id.value = ''
    quantidade.value = 1

    // SUCESSO
    Swal.fire({
      icon: 'success',
      title: 'Sucesso',
      text: 'Vínculo criado com sucesso!'
    })

    // RECARREGA
    carregarVinculos()

  } catch (erro) {

    console.error('Erro ao criar vínculo:', erro)

    Swal.fire({
      icon: 'error',
      title: 'Erro',
      text: 'Erro ao criar vínculo.'
    })
  }
}

// BUSCAR VÍNCULOS
async function buscarVinculos() {

  try {

    // SEM FILTRO
    if (!busca.value) {

      carregarVinculos()

      return
    }

    // BUSCA
    const resposta = await api.get('/vinculos/buscar', {

      params: {
        q: busca.value
      }
    })

    vinculos.value = resposta.data

  } catch (erro) {

    console.error('Erro ao buscar vínculos:', erro)
  }
}

// EXCLUIR VÍNCULO
async function excluirVinculo(id) {

  // POPUP
  const resultado = await Swal.fire({

    icon: 'warning',

    title: 'Excluir vínculo',

    text: 'Informe o motivo da exclusão.',

    input: 'text',

    inputPlaceholder: 'Digite o motivo da exclusão',

    confirmButtonText: 'Excluir',

    cancelButtonText: 'Cancelar',

    showCancelButton: true,

    allowOutsideClick: false,

    inputValidator: (value) => {

      if (!value || value.trim() === '') {

        return 'Informe o motivo da exclusão!'
      }
    }
  })

  // CANCELADO
  if (!resultado.isConfirmed) {
    return
  }

  try {

    // DELETE
    await api.delete(`/vinculos/${id}`, {

      data: {
        observacao: resultado.value
      }
    })

    // SUCESSO
    Swal.fire({
      icon: 'success',
      title: 'Sucesso',
      text: 'Vínculo excluído com sucesso!'
    })

    // RECARREGA
    carregarVinculos()

  } catch (erro) {

    console.error('Erro ao excluir vínculo:', erro)

    const mensagem =
      erro.response?.data?.erro ||
      'Erro ao excluir vínculo.'

    Swal.fire({
      icon: 'error',
      title: 'Erro',
      text: mensagem
    })
  }
}

// EXECUTA AO ABRIR
onMounted(() => {

  carregarProdutos()

  carregarVinculos()
})
</script>

<template>

  <div class="pagina">

    <!-- CARD CADASTRO -->
    <div class="card">

      <h1>Vincular Componentes</h1>

      <div class="grid">

        <!-- EQUIPAMENTO -->
        <div class="campo">

          <label>Equipamento</label>

          <select v-model="produto_pai_id">

            <option disabled value="">
              Selecione
            </option>

            <option
              v-for="equipamento in equipamentos"
              :key="equipamento.id"
              :value="equipamento.id"
            >

              {{ equipamento.nome }}

            </option>

          </select>

        </div>

        <!-- COMPONENTE -->
        <div class="campo">

          <label>Componente</label>

          <select v-model="componente_id">

            <option disabled value="">
              Selecione
            </option>

            <option
              v-for="componente in componentes"
              :key="componente.id"
              :value="componente.id"
            >

              {{ componente.nome }}

            </option>

          </select>

        </div>

        <!-- QUANTIDADE -->
        <div class="campo">

          <label>Quantidade</label>

          <input
            type="number"
            min="1"
            v-model="quantidade"
          />

        </div>

      </div>

      <!-- BOTÃO -->
      <button
        class="botao"
        @click="criarVinculo"
      >

        Vincular

      </button>

    </div>

    <!-- BUSCA -->
    <div class="card">

      <div class="campo">

        <label>Buscar equipamento ou ID</label>

        <input
          v-model="busca"
          placeholder="Digite nome ou ID"
        />

      </div>

      <button
        class="botao"
        @click="buscarVinculos"
      >

        Pesquisar

      </button>

    </div>

    <!-- TABELA -->
    <div class="card">

      <h2>Vínculos cadastrados</h2>

      <table>

        <thead>

          <tr>

            <th>ID</th>
            <th>Equipamento</th>
            <th>Componente</th>
            <th>Quantidade</th>
            <th>Ações</th>

          </tr>

        </thead>

        <tbody>

          <tr
            v-for="vinculo in vinculos"
            :key="vinculo.id"
          >

            <td>{{ vinculo.id }}</td>

            <td>{{ vinculo.produto_pai_nome }}</td>

            <td>{{ vinculo.componente_nome }}</td>

            <td>{{ vinculo.quantidade }}</td>

            <td>

              <!-- REMOVER -->
              <button
                class="botao-excluir"
                @click="excluirVinculo(vinculo.id)"
              >

                Remover

              </button>

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

.card {
  background-color: white;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}

.grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
  margin-bottom: 20px;
}

.campo {
  display: flex;
  flex-direction: column;
}

label {
  margin-bottom: 8px;
  font-weight: bold;
}

input,
select {
  padding: 12px;
  border: 1px solid #dcdcdc;
  border-radius: 8px;
}

.botao {
  background-color: #2c3e50;
  color: white;
  border: none;
  padding: 14px 24px;
  border-radius: 8px;
  cursor: pointer;
}

.botao-excluir {
  background-color: #c0392b;
  color: white;
  border: none;
  padding: 8px 15px;
  border-radius: 8px;
  cursor: pointer;
}

.botao-editar {
  background-color: #2980b9;
  color: white;
  border: none;
  padding: 8px 15px;
  border-radius: 8px;
  cursor: pointer;
  margin-left: 10px;
}

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

</style>