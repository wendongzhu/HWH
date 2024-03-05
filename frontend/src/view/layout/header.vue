
<script setup lang="ts">
import {WindowFullscreen,WindowMinimise,WindowQuit} from '../../../wailsjs/go/main/App';
import {onMounted, ref} from "vue";


let defaultTittleWidth = ref({width: ""})
let iconFullScreen = ref("FullScreen")
let isFullScreen = ref(true)

onMounted(()=>{
  window.addEventListener("resize", getTittleSize);
  getTittleSize();
})

function getTittleSize() {
  defaultTittleWidth.value.width = window.innerWidth - 200 + "px";
}

function Fullscreen(){
  WindowFullscreen()
  getTittleSize()
  isFullScreen.value = !isFullScreen.value
  if (isFullScreen) {
    iconFullScreen.value = "FullScreen"
  }else {
    iconFullScreen.value = "CopyDocument"
  }
}
function Minimise(){
  console.log("min -1")
  WindowMinimise();
  console.log("min -2")
}
function Quit(){
  WindowQuit();
}



</script>

<template>
  <div id="header">
    <el-row :gutter="0">
      <div class="header-logo-avatar" >
        <el-avatar :size="30" shape="square" fit="fill">
          <img src="../../assets/logo/hwh_logo.ico"  alt=""/>
        </el-avatar>
      </div>
      <div class="header-title" :style="defaultTittleWidth">
        <font size="4" font-weight="bold">Harms+Wende 数采软件</font>
      </div>
      <div >
        <el-button-group class="header-button-group" >
          <el-button
              class="header-button min-button"
              style="width: 50px; height: 50px; margin: 0; padding: 0;border-radius: 0"
              color="#BBBBBB"
              icon="Minus"
              @click="Minimise" />
          <el-button
              class="header-button max-button"
              style="width: 50px; height: 50px; margin: 0; padding: 0;border-radius: 0"
              color="#BBBBBB"
              :icon="iconFullScreen"
              @click="Fullscreen"/>

          <el-button
              class="header-button quit-button"
              style="width: 50px; height: 50px; margin: 0; padding: 0;border-radius: 0"
              color="#BBBBBB"
              icon="Close"
              @click="Quit"/>
        </el-button-group>

      </div>
    </el-row>
  </div>
</template>


<style lang="scss" scoped>
.header-logo-avatar {
  padding: 10px 10px 10px 10px;
}
.header-button-group{
  width: 150px;
  margin: 0;
  padding-left: 0;
}


.header-title{
  padding-top: 9px;
  padding-bottom: 9px;
  font-weight: bold;
  color: #333333;
  //background: #1baeae;
}

/*按钮悬浮*/
.min-button:hover {
  background: #F6BE4F !important;
  color: white !important;
  font-weight: bold;
  border-color: #F6BE4F !important;
}
.max-button:hover {
  background: #62C655 !important;
  color: white !important;
  font-weight: bold;
  border-color: #62C655 !important;
}
.quit-button:hover {
  background: #EE6B60 !important;
  color: white !important;
  font-weight: bold;
  border-color: #EE6B60 !important;
}
</style>
