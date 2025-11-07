<script setup>
import { ref, onMounted } from 'vue'
import api from '../lib/api'

const channels = ref([])
const selected = ref(null)
const newName = ref('')

async function load(){
  const res = await api.get('/channels')
  channels.value = res.channels
}

async function createChannel(){
  if (!newName.value) return
  await api.post('/channels', { name: newName.value })
  newName.value = ''
  load()
}

function choose(id){
  selected.value = id
  window.dispatchEvent(new CustomEvent('channel-selected', { detail: id }))
}

onMounted(load)
</script>

<template>
  <div class="p-4">
    <h3>Channels</h3>
    <div>
      <input v-model="newName" placeholder="channel name" />
      <button @click="createChannel">Create</button>
    </div>
    <ul>
      <li v-for="c in channels" :key="c.id">
        <button @click="choose(c.id)">{{ c.name }}</button>
      </li>
    </ul>
  </div>
</template>
