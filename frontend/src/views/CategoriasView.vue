<script setup>

import { ref, onMounted } from 'vue'
import api from '../services/api'
import Swal from 'sweetalert2'

const categorias = ref([])
const novaCategoria = ref('')

async function carregarCategorias() {

 const resposta = await api.get(
  '/categorias'
)
  categorias.value = resposta.data
}

async function cadastrarCategoria() {

  if (!novaCategoria.value.trim()) {
    return
  }

  await api.post(
  '/categorias',
  {
    nome: novaCategoria.value
  }
)
Swal.fire({
    icon: 'success',
    title: 'Sucesso',
    text: 'Categoria cadastrada com sucesso!'
  })
  

  novaCategoria.value = ''

  carregarCategorias()
}

onMounted(() => {


  carregarCategorias()

})

</script>

<template>

  <div>
    <h1>Categorias</h1>

    <form @submit.prevent="cadastrarCategoria">

      <input
        type="text"
        v-model="novaCategoria"
        placeholder="Nova categoria"
      />

      <button type="submit">
        Cadastrar
      </button>

    </form>

    <table>

      <thead>

        <tr>
          <th>ID</th>
          <th>Nome</th>
        </tr>

      </thead>

      <tbody>

        <tr
          v-for="categoria in categorias"
          :key="categoria.id"
        >

          <td>{{ categoria.id }}</td>

          <td>{{ categoria.nome }}</td>

        </tr>

      </tbody>

    </table>

  </div>

</template>

<style scoped>

h1 {
  margin-bottom: 20px;
}

form {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

input {
  flex: 1;
  padding: 10px;
}

button {
  background-color: #2c3e50;
  color: white;
  border: none;
  padding: 10px 16px;
  cursor: pointer;
}

table {
  width: 100%;
  border-collapse: collapse;
  background-color: white;
}

th,
td {
  border: 1px solid #ddd;
  padding: 12px;
}

th {
  background-color: #2c3e50;
  color: white;
}

</style>