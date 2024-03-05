<script setup lang="ts">
import {EventsOn} from "../../../wailsjs/runtime";
import {onMounted, onUnmounted, ref} from "vue";
import { ReadTimerListAPI, SelectTimerAPI} from "../../../wailsjs/go/end/Backend";
import {param} from "../../../wailsjs/go/models";

let defaultFooterWidth = ref({width: ""})
let currentTime = ref("")
let timesInterval:number


onMounted(()=>{
  getFooterWidth()
  window.addEventListener("resize", getFooterWidth);

  timesInterval = setInterval(getTimesInterval, 1000);
})

onUnmounted(()=>{
  clearInterval(timesInterval);
})
function getFooterWidth() {
  defaultFooterWidth.value.width = window.innerWidth - 180 - 80 + "px";
}


function getTimesInterval() {
  let yearStr = ""
  let monthStr = ""
  let dayStr = ""
  let hoursStr = ""
  let minutesStr = ""
  let secondsStr = ""

  let year = new Date().getFullYear(); //获取当前时间的年份
  let month = new Date().getMonth() + 1; //获取当前时间的月份
  let day = new Date().getDate(); //获取当前时间的天数
  let hours = new Date().getHours(); //获取当前时间的小时
  let minutes = new Date().getMinutes(); //获取当前时间的分数
  let seconds = new Date().getSeconds(); //获取当前时间的秒数
  //当小于 10 的是时候，在前面加 0
  if (month < 10) {
    monthStr = "0" + month
  }else {

  }
  if (day < 10) {
    dayStr = "0" + day.toString()
  }else {
    dayStr = day.toString()
  }
  if (hours < 10) {
    hoursStr = "0" + hours.toString()
  }else {
    hoursStr = hours.toString()
  }
  if (minutes < 10) {
    minutesStr = "0" + minutes.toString()
  }else {
    minutesStr = minutes.toString()
  }
  if (seconds < 10) {
    secondsStr = "0" + seconds.toString()
  }else {
    secondsStr = seconds.toString()
  }
  //拼接格式化当前时间
  currentTime.value = year + "-" + monthStr + "-" + dayStr + " " + hoursStr + ":" + minutesStr + ":" + secondsStr;
}
</script>

<template>
<div>
  <el-row :gutter="0">
    <!--==========头部==========-->


    <!--==========中间==========-->
    <div class="footer" :style="defaultFooterWidth">
    <!--中间预留-->
      <el-col style="height: 100%">
        <div style="color: #2B2D30; font-weight: bold"> © 2024 Harms+Wende. All rights reserved.</div>
      </el-col>
    </div>

    <!--==========时间==========-->
    <div class="footer-time" >
      <font size="2" font-weight="bold">{{ currentTime }}</font>
    </div>
</el-row>
</div>
</template>



<style scoped>
.footer{
  padding-top: 8px;
  font-weight: bold;
  color: #333333;
}

.footer-time{
  margin-left: 50px;
  padding: 9px 0 9px 0;
  font-weight: bold;
  color: #333333;
}
</style>