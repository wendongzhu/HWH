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
<div>
  <el-tabs type="border-card" :style="winSize.mainSize">
    <el-tab-pane label="电流监控">电流监控</el-tab-pane>
    <el-tab-pane label="电压监控">电压监控</el-tab-pane>
    <el-tab-pane label="电阻监控">电阻监控</el-tab-pane>
    <el-tab-pane label="飞溅监控">飞溅监控</el-tab-pane>
    <el-tab-pane label="压力监控">压力监控</el-tab-pane>
    <el-tab-pane label="行程监控">行程监控</el-tab-pane>
    <el-tab-pane label="位置监控">位置监控</el-tab-pane>
    <el-tab-pane label="设备监控">位置监控</el-tab-pane>
  </el-tabs>
</div>

</template>

<style scoped>

</style>