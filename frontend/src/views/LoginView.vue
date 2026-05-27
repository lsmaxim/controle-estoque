<script setup>

import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'

const router = useRouter()

const email = ref('')
const senha = ref('')

async function fazerLogin() {

  try {

    const resposta = await api.post('/login', {

      email: email.value,
      senha: senha.value
    })

    // SALVA TOKEN
    localStorage.setItem(
      'token',
      resposta.data.token
    )

    // REDIRECIONA
    router.push('/')

  } catch (erro) {

    alert('Usuário ou senha inválidos')
  }
}

</script>

<template>

  <div class="login-container">

    <form
      class="login-box"
      @submit.prevent="fazerLogin"
    >
      <img src="/img/LOGOTIPO-XB.PNG.ico" alt="Logo">
      <h1>Login</h1>

      <input
        v-model="email"
        type="email"
        placeholder="E-mail"
      >

      <input
        v-model="senha"
        type="password"
        placeholder="Senha"
      >

      <button type="submit">
        Entrar
      </button>

    </form>

  </div>

</template>

<style scoped>

.login-container {

  height: 100vh;

  display: flex;

  justify-content: center;

  align-items: center;

  background: #f5f6fa;
}

.login-box {

  width: 320px;

  background: white;

  padding: 30px;

  border-radius: 12px;

  box-shadow: 0 2px 10px rgba(0,0,0,0.08);

  display: flex;

  flex-direction: column;

  gap: 15px;
}

.login-box h1 {

  text-align: center;

  color: #2c3e50;
}

.login-box input {

  padding: 12px;

  border: 1px solid #dcdde1;

  border-radius: 8px;
}

.login-box button {

  padding: 12px;

  border: none;

  border-radius: 8px;

  background: #2c3e50;

  color: white;

  font-weight: bold;

  cursor: pointer;
}
.logo {

  display: flex;

  justify-content: center;
}

.logo img {

  width: 100px;

  margin-bottom: 20px;
}
</style>