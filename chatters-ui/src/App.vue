<script setup>
import Login from './components/Login.vue'
import Channels from './components/Channels.vue'
import Messages from './components/Messages.vue'
</script>

<template>
  <div>
    <Login />
    <div style="display:flex; gap:20px;">
      <Channels />
      <Messages />
    </div>
  </div>
</template>
