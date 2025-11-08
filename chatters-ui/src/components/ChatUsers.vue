<script setup>
import { ref, onMounted } from 'vue'
import api from '../lib/api'

const users = ref([])

async function loadUsers() {
    const res = await api.get("/users")
    users.value = res.data.users
}
</script>

<template>
    <div class="p-4">
        <h3>Users</h3>
        <div v-for="u in users" :key="u.id" class="border p-2">
        <div class="text-xs">{{ u.username }} • #{{ u.id }}</div>
      </div>
    </div>
</template>