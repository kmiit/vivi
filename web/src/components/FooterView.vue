<template>
  <el-menu
    mode="horizontal"
    :ellipsis="false"
    background-color="#545c64"
    :default-active="activeIndex"
    :router="true"
    text-color="#fff"
    active-text-color="#ffd04b"
    id="foobar">
      <el-menu-item index="/" class="fooitem">Main Page</el-menu-item>
      <el-menu-item index="/search" class="fooitem">Search</el-menu-item>
      <el-menu-item index="/files" class="fooitem">Files</el-menu-item>
      <el-menu-item index="/self" class="fooitem">Me</el-menu-item>
    </el-menu>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

const activeIndex = ref('/');
// 监听路由变化并动态更新 `default-active` 的值
watch(route, () => {
  activeIndex.value = route.path
}, { immediate: true });
</script>

<style>
#foobar {
  display: flex;
  justify-content: space-evenly;
}
.fooitem {
  width: 25%;
  flex: auto;
  text-align: center;
}
</style>