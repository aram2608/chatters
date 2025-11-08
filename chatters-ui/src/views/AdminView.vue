<script setup>
import ChatChannels from '@/components/ChatChannels.vue';
import ChatMessages from '@/components/ChatMessages.vue';
import ChatUsers from '@/components/ChatUsers.vue';
import ChatUsers from '@/components/ChatUsers.vue';
</script>

<template>
    <div class="app-shell">

    <section class="sidebar">
      <ChatChannels />
    </section>

    <section class="main">
      <ChatMessages />
    </section>

    <section class="main">
      <ChatUsers />
    </section>
  </div>
</template>