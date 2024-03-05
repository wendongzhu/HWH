
<script setup lang="ts">
import {onMounted, ref} from "vue";

const alarmVisible = ref(false)

onMounted(()=>{

})

function alarmBtnFunc() {

}

function databaseBtnFunc() {

}

</script>

<template>
  <div id="aside">
    <el-row style="padding: 10px 10px 0 10px">
      <el-tooltip class="item" effect="dark" content="提示信息" placement="right-start">
      <el-button
          class="right-aside-button"
          color="#AAAAAA"
          icon="Bell"
          size="large"
          @click="alarmVisible = true"
      />
      </el-tooltip>
    </el-row>
    <el-row style="padding: 10px 10px 0 10px">
      <el-tooltip class="item" effect="dark" content="数据存储" placement="right-start">
      <el-button
          class="right-aside-button"
          color="#AAAAAA"
          icon="Coin"
          size="large"
          @click="databaseBtnFunc"
      />
    </el-tooltip>
    </el-row>

  </div>
  <el-drawer class="alarm-dialog" v-model="alarmVisible" :with-header="false">

  </el-drawer>
</template>


<style scoped lang="scss">
.right-aside-button{
  width: 30px;
  height: 30px;
  margin: 0;
  padding: 10px;
  border-radius: 5px;
}
.right-aside-button:hover {
  background: #409eff !important;
  color: white !important;
  font-weight: bold;
  border-color: #409eff !important;
}

.alarm-dialog{
  width:3000px;
  height: calc(100vh - 300px);
  border-radius: 10px;
  position: absolute;
  top:50px;
  right: 50px
}
</style>
