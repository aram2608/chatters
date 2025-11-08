<script setup>
import ChatChannels from '../components/ChatChannels.vue';
import ChatMessages from '../components/ChatMessages.vue';
</script>

<template>
  <div class="app-shell">
    <aside class="sidebar">
      <ChatChannels />
    </aside>

    <section class="main">
      <ChatMessages />
    </section>
  </div>
</template>

<style>
.app-shell {
  display: grid;
  grid-template-columns: 280px 1fr;
  gap: 20px;
  padding: 20px;
}
.sidebar { background: white; padding: 12px; border-radius:8px; box-shadow:0 4px 18px rgba(0,0,0,0.04); }
.main { background: white; padding: 12px; border-radius:8px; box-shadow:0 4px 18px rgba(0,0,0,0.04); min-height:60vh; }

@media (max-width: 800px) {
  .app-shell { grid-template-columns: 1fr; }
}
</style>
