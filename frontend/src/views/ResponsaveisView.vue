<script setup>

import { ref, onMounted } from 'vue'
import api from '../services/api'

const usuarios = ref([])
const equipamentos = ref([])

const usuarioId = ref('')
const produtoId = ref('')

async function carregarUsuarios() {

  const { data } = await api.get('/usuarios')

  usuarios.value = data
}

async function carregarEquipamentos() {

  const { data } = await api.get('/produtos')

  equipamentos.value = data.filter(
    p => p.tipo === 'EQUIPAMENTO'
  )
}

onMounted(() => {

  carregarUsuarios()
  carregarEquipamentos()
})

</script>

<template>

  <div class="container">

    <h1>Responsáveis</h1>

    <select v-model="usuarioId">

      <option value="">
        Selecione o usuário
      </option>

      <option
        v-for="u in usuarios"
        :key="u.id"
        :value="u.id"
      >
        {{ u.nome }}
      </option>

    </select>

    <select v-model="produtoId">

      <option value="">
        Selecione o equipamento
      </option>

      <option
        v-for="p in equipamentos"
        :key="p.id"
        :value="p.id"
      >
        {{ p.nome }}
      </option>

    </select>

    <button @click="vincular">
      Vincular
    </button>

  </div>

</template>