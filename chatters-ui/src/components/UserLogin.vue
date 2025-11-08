<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../lib/api'

const username = ref('')
const password = ref('')
const router = useRouter()
const errorMsg = ref('')
const loading = ref(false)

async function createUser(){
  if (!username.value || !password.value) return alert('username+password')
  await api.post('/users', { username: username.value, password: password.value })
  alert('created — now login')
}

async function login() {
  errorMsg.value = ''
  loading.value = true
  try {
    const { data } = await api.post('/login', {
      username: username.value,
      password: password.value,
    })

    localStorage.setItem('user', JSON.stringify(data.user))

    await router.push('/')
  } catch (err) {
    console.error(err)
    errorMsg.value = err.response?.data?.message || 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <form @submit.prevent="login" class="login-form">
    <input v-model="username" placeholder="Username" />
    <input v-model="password" placeholder="Password" type="password" />
    <button :disabled="loading"> {{ loading ? 'Logging in…' : 'Login' }} </button>
    <button @click="createUser">Create</button>

    <p v-if="errorMsg" class="error">{{ errorMsg }}</p>
  </form>
</template>
