<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import Button from './components/Button.vue'
import LoginButton from "@/components/LoginButton.vue";
import Footer from "@/components/Footer.vue";
</script>

<template>
  <header>
    <RouterLink to="/">
      <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="50" height="50" />
    </RouterLink>

    <LoginButton/>
  </header>

  <RouterView />
  <footer>
    <Footer/>
  </footer>
</template>

<style scoped>

header {
  padding: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  height: 70px;
  display: flex;
  justify-content: space-between;
}

footer {
  position: absolute;
  bottom: 0;
  width: 100%;
  display: flex;
  justify-content: center;
}

</style>
