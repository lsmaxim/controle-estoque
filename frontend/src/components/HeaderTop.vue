<script setup>



import {
  ref,
  onMounted,
  onUnmounted
} from 'vue'

import api from '../services/api'

/* =========================
RELÓGIO
========================= */

/*
  ref cria variável reativa
*/

const horario = ref('')

/* =========================
DARK MODE
========================= */

const darkMode = ref(false)
/* =========================
NOTIFICAÇÕES
========================= */

const notificacoes = ref([])

const mostrarNotificacoes = ref(false)
/* =========================
INTERVALO
========================= */

let intervalo = null

/* =========================
ATUALIZA HORÁRIO
========================= */

function atualizarHorario() {

  /*
    Cria data atual
  */

  const agora = new Date()

  /*
    Formata horário
  */

  horario.value = agora.toLocaleTimeString(
    'pt-BR'
  )
}



function alternarTema() {

  /*
    Inverte estado
  */

  darkMode.value = !darkMode.value

  /*
    Adiciona/remove classe dark
  */

  document.body.classList.toggle(
    'dark',
    darkMode.value
  )

  /*
    Salva preferência
  */

  localStorage.setItem(
    'darkMode',
    darkMode.value
  )
}



async function carregarNotificacoes() {

  try {

    /*
      Busca produtos
    */

    const response = await api.get(
      '/produtos'
    )

    const produtos = response.data

    /*
      Filtra estoque crítico
    */

    const estoqueCritico = produtos.filter(

      produto =>

        produto.quantidade <=
        produto.quantidade_minima
    )

    

    notificacoes.value = estoqueCritico.map(
      produto => ({

        id: produto.id,

        titulo: 'Estoque crítico',

        mensagem:
          `${produto.nome} está abaixo do mínimo`,

        tipo: 'estoque'
      })
    )

  } catch (error) {

    console.error(
      'Erro notificações:',
      error
    )
  }
}
onMounted(() => {


carregarNotificacoes()
  

  const temaSalvo = localStorage.getItem(
    'darkMode'
  )

  

  if (temaSalvo === 'true') {

    darkMode.value = true

    document.body.classList.add('dark')
  }

  /*
    Atualiza imediatamente
  */

  atualizarHorario()

  /*
    Atualiza a cada 1 segundo
  */

  intervalo = setInterval(() => {

    atualizarHorario()

  }, 1000)
})



onUnmounted(() => {

  

  clearInterval(intervalo)
})

</script>

<template>

  
  <header class="topbar">

    
    <div class="topbar-esquerda">

    
      <h1>Painel do Sistema</h1>

      
      <span>
        Controle de estoque e equipamentos
      </span>

    </div>

    <!-- DIREITA -->
    <div class="topbar-direita">
     

        <div class="relogio">

        {{ horario }}

      </div>

      <!-- NOTIFICAÇÃO -->
     <div class="notificacao-area">

  <!-- BOTÃO -->
  <button
    class="botao-topo"
    @click="
      mostrarNotificacoes =
      !mostrarNotificacoes
    "
  >

    🔔

    <!-- BADGE -->
    <span
      v-if="notificacoes.length > 0"
      class="badge"
    >
      {{ notificacoes.length }}
    </span>

  </button>

  <!-- DROPDOWN -->
  <div
    v-if="mostrarNotificacoes"
    class="dropdown-notificacoes"
  >

    <h3>
      Notificações
    </h3>

    <!-- LISTA -->
    <div
      v-for="notificacao in notificacoes"
      :key="notificacao.id"
      class="item-notificacao"
    >

      <strong>
        {{ notificacao.titulo }}
      </strong>

      <span>
        {{ notificacao.mensagem }}
      </span>

    </div>

    <!-- SEM NOTIFICAÇÕES -->
    <div
      v-if="notificacoes.length === 0"
      class="sem-notificacao"
    >

      Nenhuma notificação

    </div>

  </div>

</div>

      <!-- DARK MODE -->
      <button
        class="botao-topo"
        @click="alternarTema"
      >

        {{ darkMode ? '☀️' : '🌙' }}

      </button>

      <!-- USUÁRIO -->
      <div class="usuario">

        <div class="avatar">

          X

        </div>

        <div class="usuario-info">

          <strong>Administrador</strong>

          <span>TI</span>

        </div>

      </div>

    </div>

  </header>

</template>

<style scoped>

/* =========================
TOPBAR
========================= */

.topbar {

  background: var(--card-bg);

  border-radius: 14px;

  padding: 18px 25px;

  display: flex;

  justify-content: space-between;

  align-items: center;

  box-shadow: var(--shadow);

  margin-bottom: 25px;

  transition: 0.3s;
}

/* =========================
ESQUERDA
========================= */

.topbar-esquerda {

  display: flex;

  flex-direction: column;
}

.topbar-esquerda h1 {

  margin: 0;

  font-size: 24px;

  color: var(--text);
}

.topbar-esquerda span {

  color: var(--text-secondary);

  margin-top: 5px;
}

/* =========================
DIREITA
========================= */

.topbar-direita {

  display: flex;

  align-items: center;

  gap: 15px;
}

/* =========================
BUSCA
========================= */

.busca-global input {

  width: 250px;

  padding: 12px 15px;

  border: 1px solid var(--border);

  border-radius: 10px;

  outline: none;

  font-size: 14px;

  background: var(--card-bg);

  color: var(--text);
}

/* =========================
RELÓGIO
========================= */

.relogio {

  background: var(--hover);

  padding: 10px 16px;

  border-radius: 10px;

  font-weight: bold;

  color: var(--text);
}

/* =========================
BOTÕES
========================= */

.botao-topo {

  width: 42px;

  height: 42px;

  border: none;

  border-radius: 10px;

  cursor: pointer;

  background: var(--hover);

  font-size: 18px;

  transition: 0.3s;

  color: var(--text);
}

.botao-topo:hover {

  opacity: 0.8;
}

/* =========================
USUÁRIO
========================= */

.usuario {

  display: flex;

  align-items: center;

  gap: 10px;
}

/* AVATAR */

.avatar {

  width: 42px;

  height: 42px;

  border-radius: 50%;

  background: var(--sidebar);

  color: white;

  display: flex;

  align-items: center;

  justify-content: center;

  font-weight: bold;
}

/* INFO */

.usuario-info {

  display: flex;

  flex-direction: column;
}

.usuario-info strong {

  color: var(--text);
}

.usuario-info span {

  color: var(--text-secondary);

  font-size: 13px;
}

/* =========================
RESPONSIVO
========================= */

@media (max-width: 900px) {

  .topbar {

    flex-direction: column;

    gap: 20px;

    align-items: flex-start;
  }

  .topbar-direita {

    width: 100%;

    flex-wrap: wrap;
  }

  .busca-global input {

    width: 100%;
  }
}
/* =========================
NOTIFICAÇÕES
========================= */

.notificacao-area {

  position: relative;
}

/* BADGE */

.badge {

  position: absolute;

  top: -5px;

  right: -5px;

  background: #e74c3c;

  color: white;

  width: 20px;

  height: 20px;

  border-radius: 50%;

  font-size: 11px;

  display: flex;

  align-items: center;

  justify-content: center;

  font-weight: bold;
}

/* DROPDOWN */

.dropdown-notificacoes {

  position: absolute;

  top: 55px;

  right: 0;

  width: 320px;

  background: var(--card-bg);

  border-radius: 12px;

  box-shadow: var(--shadow);

  padding: 15px;

  z-index: 999;

  border: 1px solid var(--border);
}

/* TÍTULO */

.dropdown-notificacoes h3 {

  margin-top: 0;

  margin-bottom: 15px;

  color: var(--text);
}

/* ITEM */

.item-notificacao {

  display: flex;

  flex-direction: column;

  padding: 12px;

  border-radius: 10px;

  background: var(--hover);

  margin-bottom: 10px;
}

.item-notificacao strong {

  color: var(--text);

  margin-bottom: 5px;
}

.item-notificacao span {

  color: var(--text-secondary);

  font-size: 14px;
}

/* SEM NOTIFICAÇÃO */

.sem-notificacao {

  text-align: center;

  color: var(--text-secondary);

  padding: 20px 0;
}
</style>