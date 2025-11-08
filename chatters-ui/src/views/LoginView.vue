<script setup>
import UserLogin from '@/components/UserLogin.vue'
</script>

<template>
  <div class="login-page">
    <div class="card">
      <h2>Sign in</h2>
      <UserLogin />
    </div>
  </div>
</template>

<style>
.login-page {
  min-height: calc(100vh - 64px);
  display:flex;
  align-items:center;
  justify-content:center;
  padding:24px;
}
.card {
  width: 100%;
  max-width: 380px;
  padding: 28px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 6px 22px rgba(12,20,40,0.08);
  display:flex;
  flex-direction:column;
  gap:12px;
}
.card input {
  padding:10px 12px;
  border:1px solid #e6e9ef;
  border-radius:6px;
  outline:none;
}
.card button {
  padding:10px 12px;
  border:none;
  border-radius:6px;
  cursor:pointer;
  background:#2b6ef6;
  color:white;
}
</style>
