<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const relatorios = ref([])



async function baixarRelatorio() {

  const token = localStorage.getItem('token')

  const response = await api.get(
    '/relatorios/estoque',
    {
      responseType: 'blob',
      headers: {
        Authorization: `Bearer ${token}`
      }
    }
  )

  const url = window.URL.createObjectURL(
    new Blob([response.data])
  )

  const link = document.createElement('a')

  link.href = url
  link.download = 'relatorio_estoque.pdf'

  document.body.appendChild(link)

  link.click()

  link.remove()
}
onMounted(async () => {
  try {
    const response = await api.get('/produtos')
    relatorios.value = response.data

  } catch (error) {
    console.error(error)
  }
})
</script>

<template>

<div class="pagina">

  <div class="card">

    <h1>Relatórios do Sistema</h1>
<button @click="baixarRelatorio" class="baixar-relatorio">
              📄 Relatório de Estoque
        </button>
        <p> </p>
    <table>

      <thead>
        <tr>
          <th>ID</th>
          <th>Nome</th>
          <th>Tipo</th>
          <th>Marca</th>
          <th>Quantidade</th>
        </tr>
      </thead>

      <tbody>

        <tr
          v-for="item in relatorios"
          :key="item.id"
        >
          <td>{{ item.id }}</td>
          <td>{{ item.nome }}</td>
          <td>{{ item.tipo }}</td>
          <td>{{ item.marca }}</td>
          <td>{{ item.quantidade }}</td>
        </tr>

      </tbody>

    </table>

  </div>

</div>

</template>

<style scoped>

.pagina {
  padding: 30px;
}

.card {
  background: white;
  padding: 25px;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}

h1 {
  margin-bottom: 20px;
  color: #2c3e50;
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
  text-align: left;
}

tr:hover {
  background-color: #f8f9fa;
}
.baixar-relatorio {
  margin-top: 15px;
  padding: 8px 16px;
  background-color: #2980b9;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>