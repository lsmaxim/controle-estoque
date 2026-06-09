<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const chamados = ref([])

async function carregarChamados() {

  try {

    const { data } = await api.get('/chamados')

    chamados.value = data

  } catch (err) {

    console.error(err)
  }
}

async function fecharChamado(id) {

  if (!confirm('Deseja fechar este chamado?')) {
    return
  }

  try {

    await api.put(`/chamados/${id}/fechar`)

    carregarChamados()

  } catch (err) {

    console.error(err)
    alert('Erro ao fechar chamado')
  }
}

onMounted(() => {

  carregarChamados()
})
</script>

<template>

  <div class="pagina">

    <h1>Chamados</h1>

    <table>

      <thead>
        <tr>
          <th>ID</th>
          <th>Título</th>
          <th>Solicitante</th>
          <th>Setor</th>
          <th>Prioridade</th>
          <th>Status</th>
          <th>Ações</th>
        </tr>
      </thead>

      <tbody>

        <tr
          v-for="chamado in chamados"
          :key="chamado.id"
        >

          <td>{{ chamado.id }}</td>

          <td>{{ chamado.titulo }}</td>

          <td>{{ chamado.solicitante }}</td>

          <td>{{ chamado.setor }}</td>

          <td>{{ chamado.prioridade }}</td>

          <td>

            <span
              class="status"
              :class="chamado.status.toLowerCase()"
            >
              {{ chamado.status }}
            </span>

          </td>

          <td>

            <button
              v-if="chamado.status !== 'FECHADO'"
              @click="fecharChamado(chamado.id)"
            >
              Fechar
            </button>

          </td>

        </tr>

      </tbody>

    </table>

  </div>

</template>

<style scoped>

.pagina {
  padding: 20px;
}

table {

  width: 100%;

  border-collapse: collapse;

  background: white;
}

th,
td {

  padding: 12px;

  border: 1px solid #ddd;
}

th {

  background: #243447;

  color: white;
}

.status {

  padding: 5px 10px;

  border-radius: 6px;

  font-size: 12px;

  font-weight: bold;
}

.aberto {

  background: #ffeeba;
}

.fechado {

  background: #c3e6cb;
}

button {

  padding: 8px 12px;

  border: none;

  border-radius: 6px;

  cursor: pointer;
}
</style>