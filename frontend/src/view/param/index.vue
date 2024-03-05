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
  <div class="help-main" :style="winSize.mainSize">
    <el-col style="height: 100%">
      <div style="padding-top: 400px; color: #2B2D30; font-weight: bold"> © 2024 Harms+Wende All rights reserved.</div>
    </el-col>

  </div>
</template>

<script>

</script>


<style>

</style>