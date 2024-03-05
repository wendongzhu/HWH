<script setup lang="ts">
import {onMounted, ref, Ref} from "vue";
import {ElMessageBox} from "element-plus";
import {param} from "../../../wailsjs/go/models";
import {ReadExportFileAPI, ReadFileNameAPI} from "../../../wailsjs/go/end/Backend";

let winSize = ref({mainSize:{width: "",height: ""}})
const dialogVisible = ref(false)

const pageSize = ref(50)
const currentPage = ref(1)
const pageTotal = ref(0)
const tableLoading = ref(false)
const tableLoadingText = ref("Loading")

let selectFileName = ref()
let optionFileName = ref()

onMounted(()=>{
  getWinSize();
  window.addEventListener("resize", getWinSize);
  ReadFileNameFunc()
  ReadExportFileFunc()
})

function getWinSize() {
  winSize.value.mainSize.width = window.innerWidth - 102 + "px";
  winSize.value.mainSize.height = window.innerHeight - 92 + "px";
}

let dataList:Ref<param.ExportParam[]> = ref([])

function fileNameFunc(){

}

function ReadFileNameFunc(){
  ReadFileNameAPI().then((result: param.FileNameListParam) => {
    optionFileName.value = result.data
    if (optionFileName.value.length > 0){
      selectFileName.value = optionFileName.value[0].value
    }
    console.log("ReadSystemSettingFunc: ", optionFileName.value )
  })

}

function ReadExportFileFunc(){
  let selectParam:param.SelectDataParam = {
    file_name: selectFileName.value,
    page_size: pageSize.value,
    current_page: currentPage.value,

  }
  ReadExportFileAPI(selectParam).then((result: param.ExportListParam) => {
    dataList.value = result.data
    pageTotal.value = result.count
    console.log("ReadSystemSettingFunc: ", dataList)
  })

}


const queryButtonFunc = () => {
  tableLoadingText.value = "Loading"
  tableLoading.value = true
  ReadExportFileFunc()
  tableLoading.value = false
}



const handleClose = (done: () => void) => {
  ElMessageBox.confirm('Are you sure to close this dialog?')
      .then(() => {
        done()
      })
      .catch(() => {
        // catch error
      })
}

const handleSizeChange = () => {
  queryButtonFunc()
}
const handleCurrentChange = () => {
  queryButtonFunc()
}
</script>
<template>
  <div class="data-card">
    <div class="data-card-header">
      <div class="data-header-form">
        <el-form-item label="文件名" style="font-size: 14px;padding-left: 20px">
          <el-select v-model="selectFileName" clearable style="font-size: 14px;width: 300px;" :disabled="false" >
            <el-option v-for="item in optionFileName" :key="item.value" :label="item.label" :value="item.value" @click="fileNameFunc"/>
          </el-select>
        </el-form-item>
      </div>


      <div style="margin-left: auto; padding-right: 20px">
        <el-button icon="Refresh" style="width: 80px; margin-right: 10px" @click="ReadFileNameFunc">刷 新</el-button>
        <el-button icon="Search" style="width: 80px; margin-right: 10px" @click="queryButtonFunc">查 询</el-button>
      </div>

    </div>
    <div class="data-card-main">
      <el-table class="data-card-main-table"
                stripe
                v-loading="tableLoading"
                :element-loading-text="tableLoadingText"
                :data="dataList"
      >
        <el-table-column type="index" label="序号" width="80" />
        <el-table-column prop="model_group" label="模块组" width="180"/>
        <el-table-column prop="model_name" label="模块名" width="180" />
        <el-table-column prop="program_number" sortable label="程序编号" width="140" />
        <el-table-column prop="weld_time" sortable label="焊接时间" width="180" />
        <el-table-column prop="sampling_time" label="采样时间点" width="140"/>
        <el-table-column prop="extract_current" label="提取电流(A)" width="140"/>
        <el-table-column prop="extract_voltage" label="提取电压(V)" width="140"/>
        <el-table-column prop="extract_resistance" label="提取电阻(μΩ)" width="140"/>
        <el-table-column prop="peak_current" label="电流峰值(A)" width="140"/>
        <el-table-column prop="peak_voltage" label="电压峰值(V)" width="140"/>
        <el-table-column prop="peak_resistance" label="电阻峰值(μΩ)" width="140"/>
        <el-table-column prop="archive_series" label="归档序列号" width="180"/>
        <el-table-column  label="备注" auto/>
      </el-table>
      <div class="data-card-main-footer">
        <el-pagination
            class="pagination"
            small
            background
            v-model:page-size="pageSize"
            v-model:current-page="currentPage"
            :total="pageTotal"
            :page-sizes="[30, 50, 100, 150, 200]"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            style="margin-left: 20px; margin-top: 3px"
        />
      </div>
    </div>
  </div>

  <el-dialog class="data-curve-dialog"
             v-model="dialogVisible"
             title="焊接过程曲线"
             :show-close="false"
             :before-close="handleClose"
             style="border-radius: 10px;width: 1500px; height: 680px"
  >

    <el-row class="data-curve-card">
      <div class="data-curve-chart"  ref="curveCardChart" ></div>
    </el-row>
    <template #footer>
      <span class="dialog-footer">
        <el-button  @click="dialogVisible = false">关闭</el-button>
      </span>
    </template>
  </el-dialog>

</template>



<style scoped lang="scss" >
.data-card{
  margin: 15px;
  width: calc(100% - 30px);
  height: calc(100vh - 160px);
}
.data-card-header{
  width: 100%;
  height: 60px;
  border-radius: 10px;
  background: #DDD;
  display: flex;
  align-items: center;
  .data-header-form{
    display: flex;
    align-items: center;
    padding-top: 18px;
  }
}
.data-card-main{
  margin-top: 20px;
  width: 100%;
  height: calc(100vh - 200px);
  border-radius: 10px;
  background: #DDD;
}
.data-card-main-table{
  padding-left: 13px;
  padding-right: 13px;
  width: auto;
  height: calc(100% - 32px);
  background: #DDD;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
}
.data-card-main-footer{
  height: 30px;
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px;
  background: #DDD;
}


.el-table{
  --el-table-border-color: #bbb;
  --el-table-border: 1px solid var(--el-table-border-color);
  --el-table-text-color: #333;
  --el-table-header-text-color: #333;
  --el-table-row-hover-bg-color: #999;
  --el-table-current-row-bg-color: red;
  --el-table-header-bg-color: #DDD;
  --el-table-fixed-box-shadow: #333;
  --el-table-bg-color: #BBB;
  --el-table-tr-bg-color: #CCC;
  --el-table-expanded-cell-bg-color: red;
  --el-table-fixed-left-column: inset 10px 0 10px -10px rgba(0, 0, 0, 0.15);
  --el-table-fixed-right-column: inset -10px 0 10px -10px rgba(0, 0, 0, 0.15);
}



</style>