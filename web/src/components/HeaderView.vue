<template>
  <div id="fdiv">
    <div id="logo">
      <!-- Logo -->
      <router-link :to="'/'">Vivi</router-link>
    </div>
    <div id="onRight">
      <div id="searchbox">
        <!-- Search box-->
        <el-input
          v-model="searchbox"
          style="width: 240px"
          placeholder="Input hint"
          :suffix-icon="Search">
        </el-input>
      </div>
      <div id="avatar">
        <!-- Avatar -->
        <router-link :to="'/self'">
        <el-avatar>
          <img v-if="avatarUrl !== ''"
            :src="avatarUrl"/>
          <!--<p v-else>{{ User.username[0] }}</p>-->
        </el-avatar>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
import '@/css/header.css'

const searchbox = ref('')
const avatarUrl = ''
</script>
