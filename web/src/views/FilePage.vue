<template>
  <el-card v-for="item in data">
    <template #header>
      <div class="card-header">
        <span>{{ item.name }}</span>
      </div>
    </template>
    <p class="text item">file name: {{ item.fullName }}</p>
    <p class="text item">extend name: {{ item.ext }}</p>
    <p class="text item">file id: {{ item.id }}</p>
    <p class="text item">parent id: {{ item.parent }}</p>
    <template #footer>is dir: {{ item.isDir }}</template>
  </el-card>
</template>

<script lang="ts">
import Getdata from "@/scripts/get"

let data_ = await Getdata()
export default {
  data (){
    return {
      data: data_,
    }
}}
</script>
