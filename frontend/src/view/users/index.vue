<script setup lang="ts">
import {onMounted, ref} from "vue";

let winSize = ref({mainSize:{width: "",height: ""}})

onMounted(()=>{
  getWinSize();
  window.addEventListener("resize", getWinSize);



})

function getWinSize() {
  winSize.value.mainSize.width = window.innerWidth - 102 + "px";
  winSize.value.mainSize.height = window.innerHeight - 92 + "px";
}
</script>

<template>
<div> <el-tabs type="border-card" :style="winSize.mainSize">
  <el-tab-pane label="User">diagnosis</el-tab-pane>
  <el-tab-pane label="Config">Config</el-tab-pane>
  <el-tab-pane label="Role">Role</el-tab-pane>
  <el-tab-pane label="Task">Task</el-tab-pane>
</el-tabs></div>
</template>

<style scoped>

</style>