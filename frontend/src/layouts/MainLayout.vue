<script setup>

// IMPORTA ROUTERLINK
// RouterLink é usado para navegar entre páginas
// sem recarregar o sistema.

import { RouterLink } from 'vue-router'
import HeaderTop from '../components/HeaderTop.vue'

</script>

<template>

  <!--
    CONTAINER PRINCIPAL
    Toda aplicação ficará aqui dentro
  -->
  <div class="layout">

    <!--
      SIDEBAR
      Menu lateral do sistema
    -->
    <aside class="sidebar">

      <!--
        ÁREA DA LOGO
        Ao clicar:
        volta para dashboard
      -->
      <RouterLink
        to="/"
        class="logo-area"
      >

        <!-- LOGO -->
        <img
          src="/img/LOGOTIPO-XB.PNG.ico"
          alt="Logo"
          class="logo"
        />

        <!-- NOME -->
        <h2>Xamego TI</h2>

      </RouterLink>

      <!--
        MENU
      -->
      <nav class="menu">

        <!-- DASHBOARD -->
        <RouterLink
          to="/"
          class="menu-item"
        >
          📊 Dashboard
        </RouterLink>

        <!-- PRODUTOS -->
        <RouterLink
          to="/produtos"
          class="menu-item"
        >
          📦 Produtos
        </RouterLink>

        <!-- CATEGORIAS -->
        <RouterLink
          to="/categorias"
          class="menu-item"
        >
          🗂 Categorias
        </RouterLink>

        <!-- VÍNCULOS -->
        <RouterLink
          to="/vinculos"
          class="menu-item"
        >
          🔗 Vincular Componentes
        </RouterLink>

        <!-- MOVIMENTAÇÕES -->
        <RouterLink
          to="/movimentacoes"
          class="menu-item"
        >
          📈 Movimentações
        </RouterLink>

        <!-- CALENDÁRIO -->
        <RouterLink
          to="/calendario"
          class="menu-item"
        >
          📅 Agendamentos
        </RouterLink>

        <!-- RELATÓRIOS -->
        <RouterLink
          to="/relatorios"
          class="menu-item"
        >
          📄 Relatórios
        </RouterLink>

      </nav>

    </aside>

    <!--
      CONTEÚDO PRINCIPAL
    -->
    <main class="conteudo">

  <!-- HEADER -->
  <HeaderTop />

  <!-- PÁGINAS -->
  <router-view />

</main>

  </div>

</template>

<style scoped>

/* =======================
LAYOUT PRINCIPAL
======================= */

.layout {

  /*
    FLEX:
    coloca sidebar e conteúdo
    lado a lado
  */

  display: flex;

  /*
    Altura total da tela
  */

  min-height: 100vh;

  /*
    Cor de fundo do sistema
  */

  background: #f4f6f9;
}

/* =======================
SIDEBAR
======================= */

.sidebar {

  /*
    Largura fixa
  */

  width: 200px;

  /*
    Fundo escuro
  */

  background: #243447;

  /*
    Cor texto
  */

  color: white;

  /*
    Espaçamento interno
  */

  padding: 25px 20px;

  /*
    Sidebar fixa
  */

  position: fixed;

  /*
    Altura total
  */

  height: 100vh;

  /*
    Scroll automático
  */

  overflow-y: auto;
}

/* =======================
LOGO
======================= */

.logo-area {

  /*
    FLEX vertical
  */

  display: flex;

  flex-direction: column;

  align-items: center;

  /*
    Remove sublinhado
  */

  text-decoration: none;

  color: white;

  margin-bottom: 40px;

  /*
    Remove hover estranho
  */

  background: transparent !important;
}

.logo-area:hover {
  background: transparent !important;
}

/* LOGO */

.logo {

  width: 90px;

  margin-bottom: 12px;
}

/* TEXTO */

.logo-area h2 {

  margin: 0;

  font-size: 34px;
}

/* =======================
MENU
======================= */

.menu {

  display: flex;

  flex-direction: column;

  gap: 10px;
}

/* ITEM MENU */

.menu-item {

  color: white;

  text-decoration: none;

  padding: 14px 16px;

  border-radius: 10px;

  transition: 0.3s;
}

/*
  HOVER:
  efeito ao passar mouse
*/

.menu-item:hover {

  background: rgba(255,255,255,0.12);
}

/*
  LINK ATIVO
  Vue adiciona automaticamente:
  router-link-active
*/

.router-link-active {

  background: rgba(255,255,255,0.18);

  font-weight: bold;
}

/* =======================
CONTEÚDO
======================= */

.conteudo {

  /*
    Empurra conteúdo
    para o lado da sidebar
  */

  margin-left: 260px;

  /*
    Ocupa resto da tela
  */

  flex: 1;

  padding: 30px;
}

/* =======================
RESPONSIVO
======================= */

@media (max-width: 900px) {

  .sidebar {

    width: 100%;

    height: auto;

    position: relative;
  }

  .conteudo {

    margin-left: 0;
  }

  .layout {

    flex-direction: column;
  }
}

</style>