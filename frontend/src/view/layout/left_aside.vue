<script setup lang="ts">
import {useRouter} from "vue-router";
import {ref} from "vue";

let $router = useRouter()

let btnStatus = ref(["#529b2e", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA"])

function collectionFunc(){
  $router.push("/collection")
  btnStatus.value = ["#529b2e", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA"]
}
function paramFunc(){
  $router.push("/param")
  btnStatus.value = ["#AAAAAA", "#529b2e", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA"]
}
function inspectorFunc() {
  $router.push("/inspector")
  btnStatus.value = ["#AAAAAA", "#AAAAAA", "#529b2e", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA"]
}
function dataFunc(){
  $router.push("/data")
  btnStatus.value = ["#AAAAAA", "#AAAAAA", "#AAAAAA", "#529b2e", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA"]
}
function settingFunc() {
  $router.push("/setting")
  btnStatus.value = ["#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#529b2e", "#AAAAAA", "#AAAAAA", "#AAAAAA"]
}
function  diagnosisFunc() {
  $router.push("/diagnosis")
  btnStatus.value = ["#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#529b2e", "#AAAAAA", "#AAAAAA"]
}
function  usersFunc() {
  $router.push("/users")
  btnStatus.value = ["#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#529b2e", "#AAAAAA"]
}
function helpFunc(){
  $router.push("/help")
  btnStatus.value = ["#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#AAAAAA", "#529b2e"]
}
</script>

<template>
  <div class="left-aside">
    <el-row style="padding: 10px 10px 0 10px">
      <el-tooltip class="item" effect="dark" content="数据采集" placement="right-start">
      <el-button
          class="left-aside-button"
          :color=btnStatus[0]
          icon="Menu"
          size="large"
          @click="collectionFunc" />
      </el-tooltip>
    </el-row>
    <el-row style="padding: 10px 10px 0 10px">
      <el-tooltip class="item" effect="dark" content="数据管理" placement="right-start">
        <el-button
            class="left-aside-button"
            :color=btnStatus[3]
            icon="Histogram"
            size="large"
            @click="dataFunc" />
      </el-tooltip>
    </el-row>
    <el-row style="padding: 10px 10px 0 10px">
      <el-tooltip class="item" effect="dark" content="系统设置" placement="right-start">
      <el-button
          class="left-aside-button"
          :color=btnStatus[4]
          icon="Tools"
          size="large"
          @click="settingFunc" />
      </el-tooltip>
    </el-row>

<!--    <el-row style="padding: 10px 10px 0 10px">-->
<!--      <el-tooltip class="item" effect="dark" content="帮助信息" placement="right-start">-->
<!--      <el-button-->
<!--          class="left-aside-button"-->
<!--          :color=btnStatus[7]-->
<!--          icon="HelpFilled"-->
<!--          size="large"-->
<!--          @click="helpFunc" />-->
<!--      </el-tooltip>-->
<!--    </el-row>-->
  </div>

</template>




<style scoped lang="scss">
.left-aside{

}
.left-aside-button{
  width: 30px;
  height: 30px;
  margin: 0;
  padding: 0;
  border-radius: 5px;
}
.left-aside-button:hover {
  background: #409eff !important;
  color: white !important;
  font-weight: bold;
  border-color: #409eff !important;
}

</style>
