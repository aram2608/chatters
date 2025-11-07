<script setup>
import { ref } from 'vue'
import api from '../lib/api'

const username = ref('')
const password = ref('')

async function createUser(){
  if (!username.value || !password.value) return alert('username+password')
  await api.post('/users', { username: username.value, password: password.value })
  alert('created — now login')
}

async function login(){
  try {
    const res = await api.post('/login', { username: username.value, password: password.value })
    const user = res.data.user
    localStorage.setItem('user_id', user.id)
    localStorage.setItem('username', user.username)
    alert('logged in')
    location.reload()
  } catch (error) {
    return alert('login failed')
  }
}
</script>

<template>
  <div class="p-4">
    <input v-model="username" placeholder="username" />
    <input v-model="password" type="password" placeholder="password" />
    <button @click="createUser">Create</button>
    <button @click="login">Login</button>
  </div>
</template>
