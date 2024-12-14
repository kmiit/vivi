<template>
  <el-menu
    mode="horizontal"
    background-color="#545c64"
    :default-active="activeIndex"
    text-color="#fff"
    active-text-color="#ffd04b"
    id="foobar">
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

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

const activeIndex = ref('1');
// 监听路由变化并动态更新 `default-active` 的值
watch(route, () => {
  if (route.path === '/') {
    activeIndex.value = '1';
  } else if (route.path === '/files') {
    activeIndex.value = '2';
  } else if (route.path === '/self') {
    activeIndex.value = '3';
  }
}, { immediate: true });
</script>

<style>
#foobar {
  justify-content: space-between;
}
el-menu-item {
  text-align: center;
}
</style>