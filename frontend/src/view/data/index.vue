<script setup lang="ts">
import {onMounted, ref, Ref} from "vue";
import {ElMessageBox} from "element-plus";
import {param} from "../../../wailsjs/go/models";
import {ReadArchiveFileAPI, ReadArchiveFileNameAPI} from "../../../wailsjs/go/end/Backend";
import * as echarts from "echarts";

let winSize = ref({mainSize:{width: "",height: ""}})

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

let dataList:Ref<param.HWHArchiveParam[]> = ref([])

function fileNameFunc(){

}

function ReadFileNameFunc(){
  ReadArchiveFileNameAPI().then((result: param.FileNameListParam) => {
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
  ReadArchiveFileAPI(selectParam).then((result: param.HWHArchiveDataListParam) => {
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


const dialogVisible = ref(false)
let timerSetTimeout:any
let curveOption:any;
let curveChart:any;
let curveCardChart = ref()
const IKVCurveData:any  = ref([])
const UKVCurveData:any = ref([])
const RKVCurveData:any = ref([])
const SKTCurveData:any = ref([])
const QKVCurveData:any = ref([])
let xAxisData:any = ref([])

curveOption = {
  grid:{
    x:45,
    y:25,
    x2:10,
    y2:35,
  },
  legend: {
    textStyle: {
      fontWeight: "bold",
      fontFamily: "Arial",
    },
    data: ['SKT(skt)', '电流(A)', '电压(v)', '电阻(μΩ)', 'Q参考(μΩ)']
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  xAxis: [
    {
      name: "时间(ms)",
      nameLocation: "center",
      nameTextStyle: {
        fontWeight: "bold",
        color: "#666",
        fontSize: 12,
        padding: [10, 10, 10, 10]
      },
      type: 'category',
      // inverse: true,
      splitLine: {
        lineStyle: {
          color: "#888"
        }
      },
      axisLabel: {
        fontSize: 12,
        color: "#888"
      },
    }
  ],
  yAxis: [
    {
      type: "value",
      nameLocation: "middle",
      nameTextStyle: {
        fontWeight: "bold",
        color: "#666",
        fontSize: 12,
        align: "center",
        verticalAlign: "bottom",
        padding: [10, 10, 10, 10]
      },
      axisLine: {
        show: true,
      },
      axisTick: {
        inside: true
      },
      axisLabel: {
        fontSize: 12,
        color: "#888"
      },
      splitLine: {
        lineStyle: {
          type: "dashed",
          color: "#888"
        }
      }
    },{
      axisLine: {
        show: true,
      },
    }
  ],
  series: [
    {
      name: 'SKT(skt)',
      type: 'line',
      smooth: true,
      showSymbol: false,
      itemStyle:{
        color: "#FF9800"
      },
    },
    {
      name: '电流(A)',
      type: 'line',
      smooth: true,
      showSymbol: false,
      itemStyle:{
        color: "#409EFF"
      },
    },
    {
      name: '电压(v)',
      type: 'line',
      smooth: true,
      showSymbol: false,
      itemStyle:{
        color: "#E91E63"
      },
    },
    {
      name: '电阻(μΩ)',
      type: 'line',
      smooth: true,
      showSymbol: false,
      itemStyle:{
        color: "#4CAF50"
      },
    },
    {
      name: 'Q参考(μΩ)',
      type: 'line',
      smooth: true,
      showSymbol: false,
      itemStyle:{
        color: "#00BCD4"
      },
    }
  ]
}
const openDetail = (row:any) => {
  dialogVisible.value = true

  SKTCurveData.value = row["skt"]
  IKVCurveData.value = row["ikv"]
  UKVCurveData.value = row["ukv"]
  RKVCurveData.value = row["rkv"]
  QKVCurveData.value = row["qkv"]

  let xData = []
  for (let i in IKVCurveData.value) {
    xData.push(parseInt(i))
  }
  xAxisData.value = xData

  timerSetTimeout = setTimeout(()=>{
    InitCurveChart()
    window.addEventListener("resize", () => curveChart.resize())
  },10)
}

function InitCurveChart(){
  curveOption.xAxis[0].data = xAxisData.value
  curveOption.series[0].data = SKTCurveData.value
  curveOption.series[1].data = IKVCurveData.value
  curveOption.series[2].data = UKVCurveData.value
  curveOption.series[3].data = RKVCurveData.value
  curveOption.series[4].data = QKVCurveData.value
  curveChart = echarts.init(curveCardChart.value);
  curveOption && curveChart.setOption(curveOption);
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
                @row-click="openDetail"
                :data="dataList"
      >
        <el-table-column type="index" label="序号" width="80" />
        <el-table-column prop="model_group" label="模块组" width="180"/>
        <el-table-column prop="model_name" label="模块名" width="180" />
        <el-table-column prop="spot_time" sortable label="焊点时间" width="180"/>
        <el-table-column prop="control_mode" label="控制模式" width="100" />
        <el-table-column prop="type_id" label="类型ID" width="100" />
        <el-table-column prop="spot_number" label="焊点编号" width="100" />
        <el-table-column prop="program_number" label="程序编号" width="100" />
        <el-table-column prop="spot_name" label="焊点名称" width="auto"  min-width="200"/>
        <el-table-column prop="pre_time" sortable label="预热时间(ms)" width="140"/>
        <el-table-column prop="real_time" sortable label="焊接时间(ms)" width="140"/>
        <el-table-column prop="keep_time" sortable label="维持时间(ms)" width="140"/>
        <el-table-column prop="set_current" label="设定电流(A)" width="140"/>
        <el-table-column prop="real_mean_current" label="实际电流(A)" width="140"/>
        <el-table-column prop="real_mean_voltage" label="实际电压(V)" width="140"/>
        <el-table-column prop="real_energy" label="能量(J)" width="100"/>
        <el-table-column prop="spot_quality" label="焊接质量" width="100"/>
        <el-table-column prop="q_threshold" label="Q阈值" width="100"/>
        <el-table-column prop="q_value" label="Q值" width="100"/>
        <el-table-column prop="r_max" label="R最大值" width="100"/>
        <el-table-column prop="spatter" label="飞溅时间点" width="120"/>
        <el-table-column prop="milling_cycle_counter" label="循环计数" width="100"/>
        <el-table-column prop="milling_weld_spot_counter" label="修模计数" width="100"/>
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
             style="border-radius: 10px;width: 1140px; height: 600px"
  >

    <el-row class="data-curve-card" >
      <div class="data-curve-chart" ref="curveCardChart" ></div>
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

.data-curve-card{
  width: 1100px;
  height: 420px;
}
.data-curve-chart{
  width: 100%;
  height: 100%;
}

</style>