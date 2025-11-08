<script setup>

import { ref } from 'vue'
import { useAuthStore } from '@/stores/counter'

const auth = useAuthStore()
const username = ref('')
const password = ref('')

const onSubmit = async () => {
  try {
    await auth.login(username.value, password.value)
  } catch (e) {
    console.error('Login failed', auth.error)
  }
}

async function createUser(){
  if (!username.value || !password.value) return alert('username+password')
  await api.post('/users', { username: username.value, password: password.value })
  alert('created — now login')
}

</script>

<template>
  <form @submit.prevent="onSubmit" class="login-form">
    <input v-model="username" placeholder="Username" />
    <input v-model="password" placeholder="Password" type="password" />
    <button :disabled="auth.loading">Login</button>
    <button @click="createUser">Create</button>
  </form>
</template>
