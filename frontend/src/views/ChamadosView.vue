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

import Swal from 'sweetalert2'

async function fecharChamado(id) {

  const confirmacao = await Swal.fire({

    title: 'Fechar chamado?',
    text: 'Deseja realmente fechar este chamado?',
    icon: 'question',

    showCancelButton: true,

    confirmButtonText: 'Sim, fechar',
    cancelButtonText: 'Cancelar'
  })

  if (!confirmacao.isConfirmed) {
    return
  }

  const resultado = await Swal.fire({

    title: 'Solução aplicada',
    text: 'Descreva a solução utilizada para resolver o chamado.',

    input: 'textarea',
    inputPlaceholder: 'Informe a solução...',

    showCancelButton: true,

    confirmButtonText: 'Finalizar',
    cancelButtonText: 'Cancelar',

    inputValidator: (value) => {

      if (!value) {
        return 'Informe a solução do chamado'
      }
    }
  })

  if (!resultado.isConfirmed) {
    return
  }

  try {

    await api.put(`/chamados/${id}/fechar`, {
      solucao: resultado.value
    })

    await Swal.fire({

      icon: 'success',
      title: 'Sucesso!',
      text: 'Chamado fechado com sucesso.',
      timer: 1500,
      showConfirmButton: false
    })

    carregarChamados()

  } catch (err) {

    console.error(err)

    Swal.fire({

      icon: 'error',
      title: 'Erro',
      text: 'Não foi possível fechar o chamado.'
    })
  }
}

async function novoChamado() {

  const { value: formValues } = await Swal.fire({

    title: 'Novo Chamado',

    html: `
      <input
        id="titulo"
        class="swal2-input"
        placeholder="Título"
      >

      <input
        id="solicitante"
        class="swal2-input"
        placeholder="Solicitante"
      >

      <input
        id="setor"
        class="swal2-input"
        placeholder="Setor"
      >

      <select
        id="equipamento"
        class="swal2-input"
      >
        <option value="">
          Selecione um equipamento
        </option>

        ${equipamentos.value.map(e => `
          <option value="${e.id}">
            ${e.nome}
          </option>
        `).join('')}
      </select>

      <select
        id="prioridade"
        class="swal2-input"
      >
        <option value="BAIXA">
          BAIXA
        </option>

        <option value="MEDIA">
          MÉDIA
        </option>

        <option value="ALTA">
          ALTA
        </option>
      </select>

      <textarea
        id="descricao"
        class="swal2-textarea"
        placeholder="Descrição do problema"
      ></textarea>
    `,

    focusConfirm: false,

    showCancelButton: true,

    confirmButtonText: 'Salvar',

    preConfirm: () => {

      return {

        titulo:
          document.getElementById('titulo').value,

        solicitante:
          document.getElementById('solicitante').value,

        setor:
          document.getElementById('setor').value,

        equipamento_id:
          parseInt(
            document.getElementById('equipamento').value
          ) || null,

        prioridade:
          document.getElementById('prioridade').value,

        descricao:
          document.getElementById('descricao').value
      }
    }
  })

  if (!formValues) return

  try {

    console.log(formValues)

    await api.post('/chamados', formValues)

    await Swal.fire({

      icon: 'success',
      title: 'Chamado criado'
    })

    carregarChamados()

  } catch (err) {

    console.error(err)

    Swal.fire({

      icon: 'error',
      title: 'Erro ao criar chamado'
    })
  }
}
const equipamentos = ref([])

async function carregarEquipamentos() {

  const { data } = await api.get('/produtos')

  equipamentos.value = data.filter(
    p => p.tipo === 'EQUIPAMENTO'
  )
}
onMounted(() => {

  carregarChamados()
  carregarEquipamentos()
})
</script>

<template>




  <div class="pagina">

  
    <h1>Chamados</h1>
<button
    class="btn-novo"
    @click="novoChamado"
  >
    + Novo Chamado
  </button>
  <p></p>
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
.btn-novo {

  background: #007bff;

  color: white;
}
</style>