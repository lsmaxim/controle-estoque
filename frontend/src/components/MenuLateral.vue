<template>

  <aside class="menu">

    <router-link
  to="/"
  class="logo-area"
    >

  <img 
    
    src="/img/LOGOTIPO-XB.PNG.ico"
    alt="Logo"
    class="logo"
  />

    <h2>Xamego TI</h2>

</router-link>

    <nav>

      <RouterLink to="/categorias">
        Categorias
      </RouterLink>

      <RouterLink to="/produtos">
        Produtos
      </RouterLink>

      <RouterLink to="/vinculos">
        Vincular Componentes
      </RouterLink>
      <RouterLink to="/relatorios">
        Relatórios
      </RouterLink>
      <router-link to="/calendario">
        Agendamento de Tarefas
       </router-link>

    </nav>

  </aside>

</template>

<style scoped>

.menu {
  width: 250px;
  min-height: 100vh;
  background-color: #2c3e50;
  padding: 30px 20px;
  color: white;
}

.logo-area {
  text-align: center;
  margin-bottom: 40px;
}

.logo {
  width: 90px;
  margin-bottom: 15px;
}

nav {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

a {
  color: white;
  text-decoration: none;
  padding: 12px;
  border-radius: 8px;
  transition: 0.2s;
}

a:hover {
  background-color: rgba(255,255,255,0.1);
}

.router-link-active {
  background-color: rgba(255,255,255,0.2);
}
.logo-area {
  text-decoration: none;
  color: white;
}
.logo-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  text-decoration: none;

  color: white;

  margin-bottom: 40px;

  background: transparent !important;

  padding: 0 !important;

  width: 100%;
}
.logo-area:hover {
  background: transparent !important;
}
.logo {
  width: 90px;
  margin-bottom: 10px;
}

.logo-area h2 {
  margin: 0;
  text-align: center;
}
</style>