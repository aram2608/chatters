<script setup>
import { useAuthStore } from '@/stores/counter'
const auth = useAuthStore()

function handleLogout() {
  auth.logout()
}
</script>

<template>
  <div id="app">
    <header class="site-header">
      <nav>
        <router-link to="/">Home</router-link>
        <router-link to="/login">Login</router-link>
      </nav>
      <div class="header">
        <button @click="handleLogout">Logout</button>
      </div>
    </header>

    <main>
      <router-view />
    </main>
  </div>
</template>

<style>
#app { font-family: system-ui, Arial; min-height:100vh; background:#f6f8fa; color:#111; }
.site-header { padding: 12px 20px; background: white; box-shadow: 0 1px 4px rgba(0,0,0,.06); }
.site-header nav a { margin-right: 12px; text-decoration: none; color: #333; }
</style>
