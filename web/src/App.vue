<template>
  <el-container>
    <el-main>
      <RouterView/>
    </el-main>
  </el-container>
  <template id="main">
    <el-menu
      :default-active="activeIndex"
      class="el-menu"
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
    >
      <el-menu-item index="1">
        <router-link to="/">Main Page</router-link>
      </el-menu-item>
      <el-menu-item index="2">
        <router-link to="/files">Files</router-link>
      </el-menu-item>
      <el-menu-item index="3">
        <router-link to="/self">Me</router-link>
      </el-menu-item>
    </el-menu>
  </template>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import { ref } from 'vue'
  
const activeIndex = ref('1')
</script>
<style>
#main {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}
.el-menu {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
}
</style>