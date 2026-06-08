<script setup>

import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()

const equipamento = ref(null)
const componentes = ref([])
const carregando = ref(true)

async function carregarEquipamento() {

  try {

    // DADOS DO EQUIPAMENTO
    const { data } = await axios.get(
      `http://192.168.0.7:8080/qrcode/equipamento/${route.params.id}`
    )

    equipamento.value = data

    // COMPONENTES VINCULADOS
    const respostaComponentes = await axios.get(
      `http://192.168.0.7:8080/qrcode/equipamento/${route.params.id}/componentes`
    )

    componentes.value = respostaComponentes.data

    console.log('Equipamento:', data)
    console.log('Componentes:', respostaComponentes.data)

  } catch (error) {

    console.error(error)

  } finally {

    carregando.value = false
  }
}

onMounted(() => {

  carregarEquipamento()

})

</script>
<template>

  <div class="pagina-publica">

    <div
      v-if="carregando"
      class="card-equipamento"
    >
      Carregando...
    </div>

    <div
      v-else-if="equipamento"
      class="card-equipamento"
    >

      <h1>
        {{ equipamento.nome }}
      </h1>

      <div class="campo">
        <strong>Marca:</strong>
        {{ equipamento.marca }}
      </div>

      <div class="campo">
        <strong>Modelo:</strong>
        {{ equipamento.modelo }}
      </div>

      <div class="campo">
        <strong>Setor:</strong>
        {{ equipamento.setor }}
      </div>

      <div class="campo">
        <strong>Descrição Técnica:</strong>
        {{ equipamento.descricao_tecnica }}
      </div>
<hr>

<h2>Componentes vinculados</h2>

<div
  v-if="componentes.length === 0"
>
  Nenhum componente vinculado.
</div>

<div
  v-for="comp in componentes"
  :key="comp.id"
  class="componente"
>
  <strong>
    {{ comp.nome }}
  </strong>

  <div>
    {{ comp.marca }}
    {{ comp.modelo }}
  </div>

  <div>
    Quantidade: {{ comp.quantidade }}
  </div>
</div>
    </div>

    <div
      v-else
      class="card-equipamento"
    >
      Equipamento não encontrado.
    </div>

  </div>

</template>

<style scoped>

.pagina-publica {

  min-height: 100vh;

  display: flex;
  justify-content: center;
  align-items: center;

  background: #f5f5f5;

  padding: 20px;
}

.card-equipamento {

  width: 100%;
  max-width: 700px;

  background: white;

  padding: 30px;

  border-radius: 12px;

  box-shadow: 0 4px 15px rgba(0,0,0,.1);
}

h1 {

  margin-bottom: 20px;

  color: #243447;
}

.campo {

  margin-bottom: 12px;

  font-size: 16px;
}
.componente {

  padding: 12px;

  margin-top: 10px;

  border: 1px solid #ddd;

  border-radius: 8px;

  background: #fafafa;
}

hr {

  margin: 25px 0;
}

h2 {

  margin-bottom: 15px;
}
</style>