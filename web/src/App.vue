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
import FooterView from '@/components/FooterView.vue'
import HeaderView from '@/components/HeaderView.vue'
</script>

<style>
#header {
  top: 0;
  left: 0;
  position: fixed;
  width: 100%;
}

#main-container {
  border-radius: 0;
  padding: 0;
  margin: 0;
}

#footer {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
}
</style>