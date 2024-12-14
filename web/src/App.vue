<template>
  <div id="fcontainer">
    <el-container class="main-container">
      <el-header id="header">
        <HeaderView/>
      </el-header>
      <el-main id="main-content">
        <RouterView/>
      </el-main>
      <el-footer id="footer">
         <FooterView/>
      </el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import HeaderView from '@/components/HeaderView.vue'
import FooterView from '@/components/FooterView.vue'
</script>

<style>
html, body, #app {
  padding: 0;
}

.main-container {
  height: 100vh;
}

#footer {
  padding: 0;
}
</style>