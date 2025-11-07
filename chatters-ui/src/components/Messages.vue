<script setup>
import { ref, onMounted } from 'vue'
import api from '../lib/api'

const messages = ref([])
const channelId = ref(null)
const text = ref('')
const username = localStorage.getItem('username')
const user_id = Number(localStorage.getItem('user_id'))

async function loadMessages(cid){
  if (!cid) return
  channelId.value = cid
  const res = await api.get(`/messages?channel_id=${cid}`)
  messages.value = res.data.messages
}

async function send(){
  if (!channelId.value) return alert('select a channel')
  if (!text.value) return
  await api.post('/messages', { channel_id: Number(channelId.value), user_id, user_name: username, text: text.value })
  text.value = ''
  await loadMessages(channelId.value)
}

window.addEventListener('channel-selected', (e) => { loadMessages(e.detail) })

onMounted(() => {
})
</script>

<template>
  <div class="p-4">
    <h3>Messages</h3>
    <div v-if="channelId">Channel: {{ channelId }}</div>
    <div class="space-y-2">
      <div v-for="m in messages" :key="m.id" class="border p-2">
        <div>{{ m.text }}</div>
        <div class="text-xs">{{ m.user_name }} • #{{ m.id }}</div>
      </div>
    </div>

    <div class="mt-3">
      <input v-model="text" placeholder="message..." />
      <button @click="send">Send</button>
    </div>
  </div>
</template>
