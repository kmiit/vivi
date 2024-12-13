<template>
  <div id="fdiv">
    <div id="logo">
      <!-- Logo -->
      <p>Vivi</p>
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
        <el-avatar>
          <img v-if="avatarUrl !== ''"
            :src="avatarUrl" />
          <!--<p v-else>{{ User.username[0] }}</p>-->
        </el-avatar>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'

const searchbox = ref('')
const avatarUrl = ''
</script>

<style scoped>
#fdiv {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
}

#onRight {
  display: flex;
  align-items: center;
}

#searchbox {
  margin-right: 10px;
}
</style>
