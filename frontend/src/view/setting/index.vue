<script setup lang="ts">

import {
  ReadSystemSettingAPI,
  WriteSystemSettingAPI,
} from "../../../wailsjs/go/end/Backend";
import {param} from "../../../wailsjs/go/models";
import {onBeforeMount, onMounted, ref} from "vue";
import { ElMessage } from 'element-plus'
import {EventsOn} from "../../../wailsjs/runtime";

let winSize = ref({mainSize:{width: "",height: ""}})

let modelFile = ref({
  ArchiveFilePath:"",
  ArchiveFileName:"",
  ExportFilePath: "",
  ExportFileName: "",
  SamplingTime: 0,
})


onBeforeMount(()=>{

})

onMounted(()=>{
  getWinSize();
  window.addEventListener("resize", getWinSize);
  ReadSystemSettingFunc()
})

function getWinSize() {
  winSize.value.mainSize.width = window.innerWidth - 102 + "px";
  winSize.value.mainSize.height = window.innerHeight - 92 + "px";
}

function ReadSystemSettingFunc(){
  ReadSystemSettingAPI().then((result: param.ModelFileParam) => {
    console.log("ReadSystemSettingFunc: ", result)
    modelFile.value= result
  })

}

function WriteSystemSettingFunc(){
  let writeData:param.ModelFileParam
  writeData = modelFile.value
  WriteSystemSettingAPI(writeData).then((result: param.ModelFileParam) => {
    modelFile.value = result
    ElMessage({
      message: '参数保持成功！',
      grouping: true,
      type: 'success',
      offset: 5,
    })
  })
}


</script>

<template>
      <div class="system-main" :style="winSize.mainSize">
        <el-row class="system-param-row">
          <el-form :model="modelFile" label-width="150px" style="min-width: 600px;">
            <el-form-item label="归档文件路径:">
              <el-input
                  v-model="modelFile.ArchiveFilePath"
                  placeholder="Please input"
                  class="input-with-select"
              >
<!--                <template #prepend>-->
<!--                  <el-button icon="Folder" />-->
<!--                </template>-->
                <template #append>
                  <el-button icon="Document" @click="WriteSystemSettingFunc">保存</el-button>
                </template>
              </el-input>
            </el-form-item>
          </el-form>
        </el-row>
        <el-row class="system-param-row">
          <el-form :model="modelFile" label-width="150px" style="min-width: 600px;">
            <el-form-item label="归档文件名称:">
              <el-input
                  v-model="modelFile.ArchiveFileName"
                  placeholder="Please input"
                  class="input-with-select"
              >
                <template #append>
                  <el-button icon="Document" @click="WriteSystemSettingFunc">保存</el-button>
                </template>
              </el-input>
            </el-form-item>
          </el-form>
        </el-row>
        <el-row class="system-param-row">
          <el-form :model="modelFile" label-width="150px" style="min-width: 600px;">
            <el-form-item label="保存文件路径:">
              <el-input
                  v-model="modelFile.ExportFilePath"
                  placeholder="Please input"
                  class="input-with-select"
              >
<!--                <template #prepend>-->
<!--                  <el-button icon="Folder" />-->
<!--                </template>-->
                <template #append>
                  <el-button icon="Document" @click="WriteSystemSettingFunc">保存</el-button>
                </template>
              </el-input>
            </el-form-item>
          </el-form>
        </el-row>
        <el-row class="system-param-row">
          <el-form :model="modelFile" label-width="150px" style="min-width: 600px;">
            <el-form-item label="保存文件名称:">
              <el-input
                  v-model="modelFile.ExportFileName"
                  placeholder="Please input"
                  class="input-with-select"
              >
                <template #append>
                  <el-button icon="Document" @click="WriteSystemSettingFunc">保存</el-button>
                </template>
              </el-input>
            </el-form-item>
          </el-form>
        </el-row>
        <el-row class="system-param-row">
          <el-form :model="modelFile" label-width="150px" style="min-width: 600px;">
            <el-form-item label="采样时间点:">
              <el-input
                  v-model.number="modelFile.SamplingTime"
                  placeholder="Please input"
                  class="input-with-select"
              >
                <template #append>
                  <el-button icon="Document" @click="WriteSystemSettingFunc">保存</el-button>
                </template>
              </el-input>
            </el-form-item>
          </el-form>
        </el-row>

      </div>

</template>

<style scoped lang="scss" >
.system-main{
  color: #eeeeee;
  border-radius: 8px;
}
.system-param-row{
  padding-top: 25px;
}


</style>