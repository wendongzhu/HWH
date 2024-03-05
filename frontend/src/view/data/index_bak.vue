<script setup lang="ts">
import {onMounted, ref} from "vue";
import * as XLSX from "xlsx";
import {ElLoading, ElMessageBox} from "element-plus";
import * as echarts from "echarts";

let winSize = ref({mainSize:{width: "",height: ""}})
const dialogVisible = ref(false)
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

const pageSize = ref(100)
const currentPage = ref(1)
const pageTotal = ref(0)
const tableLoading = ref(false)
const tableLoadingText = ref("Loading")

onMounted(()=>{
  getWinSize();
  window.addEventListener("resize", getWinSize);

  getDateTimeFunc()
  getDeviceFunc()

})

function getWinSize() {
  winSize.value.mainSize.width = window.innerWidth - 102 + "px";
  winSize.value.mainSize.height = window.innerHeight - 92 + "px";
}

let dataList = ref([])
let selectLineValue = ref([])
let selectLine = ref()
let selectStation = ref("")
let selectDevice = ref("")

let optionLine = ref()
let optionStation = ref()
let optionDevice = ref()

let selectDatetime = ref()
let datetime = ref()
const selectShortcuts = [
  {
    text: '最近一周',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
      return [start, end]
    },
  },
  {
    text: '最近一月',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      return [start, end]
    },
  },
  {
    text: '最近三月',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
      return [start, end]
    },
  },
]
const defaultSelectDatetime = ref()

export interface reqDeviceParam {
  line: string[],
  station: string,
  name: string,
}
interface deviceLabel {
  value: string,
  label: string,
}
export interface resDeviceResult {
  line: deviceLabel,
  station: deviceLabel,
  name: deviceLabel,
}
interface selectDatetime {
  start_date: string,
  start_time: string,
  end_date:string
  end_time: string,
}
export interface reqSelectParam {
  page_size: number,
  current_page: number,
  line: string[],
  station: string,
  device_name: string,
  date_time: selectDatetime,
}

let deviceParam:reqDeviceParam
let deviceList:resDeviceResult
let resDataList:any

async function getDataFunc() {
  let selectParam:reqSelectParam = {
    page_size: pageSize.value,
    current_page: currentPage.value,
    line: selectLineValue.value,
    station: selectStation.value,
    device_name: selectDevice.value,
    date_time: datetime.value,
  }
  // console.log("selectParam", selectParam)
  // resDataList = await HwhDataApi(selectParam)
  //TODO: 数据查询接口
  dataList.value = resDataList.data
  pageTotal.value = resDataList.count
}
async function getDeviceFunc() {
  deviceParam = {
    line: [],
    station: "",
    name: "",
  }
  //TODO: 设备查询接口
  // deviceList = await DeviceListApi(deviceParam)
  optionLine.value = deviceList.line
  optionStation.value = deviceList.station
  optionDevice.value = deviceList.name

  selectLineValue.value = optionLine.value[0]["value"]
  selectLine.value = optionLine.value[0]["label"]
  selectStation.value = optionStation.value[0]["label"]
  selectDevice.value = optionDevice.value[0]["label"]
}
function getDateTimeFunc(){
  const end = new Date()
  const start = new Date()
  start.setTime(start.getTime() - 3600 * 1000 * 24)
  selectDatetime.value = [start, end]
  datetimeToStr()
}
const lineFunc = async () => {
  selectLineValue.value = selectLine.value
  deviceParam.line = selectLine.value
  //TODO: 设备查询接口
  // deviceList = await DeviceListApi(deviceParam)
  optionStation.value = deviceList.station
  optionDevice.value = deviceList.name

  selectStation.value = optionStation.value[0]["label"]
  selectDevice.value = optionDevice.value[0]["label"]
}
const stationFunc = async () => {
  deviceParam.station = selectStation.value
  //TODO: 设备查询接口
  // deviceList = await DeviceListApi(deviceParam)
  optionDevice.value = deviceList.name
  selectDevice.value = optionDevice.value[0]["label"]
}

function datetimeToStr(){
  let data: any
  data = selectDatetime.value
  datetime.value = {
    "start_date": `${timeToStr(data[0].getFullYear(), 4)}-${timeToStr(data[0].getMonth() + 1)}-${timeToStr(data[0].getDate())}`,
    "start_time": `${timeToStr(data[0].getHours())}:${timeToStr(data[0].getMinutes())}:${timeToStr(data[0].getSeconds())}`,
    "end_date": `${timeToStr(data[1].getFullYear(), 4)}-${timeToStr(data[1].getMonth() + 1)}-${timeToStr(data[1].getDate())}`,
    "end_time": `${timeToStr(data[1].getHours())}:${timeToStr(data[1].getMinutes())}:${timeToStr(data[1].getSeconds())}`,
  }
}
const timeToStr = (data:any, total = 2, str = '0') => {
  return data.toString().padStart(total, str)
}
const queryButtonFunc = () => {
  tableLoadingText.value = "Loading"
  tableLoading.value = true
  getDataFunc()
  tableLoading.value = false
}
const settingButtonFunc = () => {

}
const exportDataButtonFunc = async () => {
  let selectParam: reqSelectParam = {
    page_size: pageSize.value,
    current_page: currentPage.value,
    line: selectLineValue.value,
    station: selectStation.value,
    device_name: selectDevice.value,
    date_time: datetime.value,
  }

}

const exportLogButtonFunc = async () => {
  let selectParam: reqSelectParam = {
    page_size: pageSize.value,
    current_page: currentPage.value,
    line: selectLineValue.value,
    station: selectStation.value,
    device_name: selectDevice.value,
    date_time: datetime.value,
  }

}

let timerSetTimeout:any
const openDetail = (row:any) => {
  dialogVisible.value = true

  SKTCurveData.value = row["properties"]["skt"]
  IKVCurveData.value = row["properties"]["ikv"]
  UKVCurveData.value = row["properties"]["ukv"]
  RKVCurveData.value = row["properties"]["rkv"]
  QKVCurveData.value = row["properties"]["qkv"]

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
            <el-form-item label="模块组" style="font-size: 14px;padding-left: 20px">
              <el-select v-model="selectStation" clearable style="font-size: 14px;width: 120px;" :disabled="false" >
                <el-option v-for="item in optionStation" :key="item.value" :label="item.label" :value="item.value" @click="stationFunc"/>
              </el-select>
            </el-form-item>
            <el-form-item label="模块名" style="font-size: 14px;padding-left: 20px">
              <el-select v-model="selectDevice" clearable style="font-size: 14px;width: 200px;" >
                <el-option v-for="item in optionDevice" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="日期" style="font-size: 14px;width: 400px; margin-right: 20px;padding-left: 20px">
              <el-date-picker
                  v-model="selectDatetime"
                  type="datetimerange"
                  :shortcuts="selectShortcuts"
                  :default-time="defaultSelectDatetime"
                  @change="datetimeToStr"
                  range-separator="To"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
              />
            </el-form-item >
          </div>


          <div style="margin-left: auto; padding-right: 20px">
            <el-button icon="Setting" circle @click="settingButtonFunc" />
            <el-button icon="Search" circle @click="queryButtonFunc"/>
            <el-button icon="ArrowDownBold" circle @click="exportDataButtonFunc" />
          </div>

        </div>
        <div class="data-card-main">
          <el-table class="data-card-main-table"
            stripe
            v-loading="tableLoading"
            :element-loading-text="tableLoadingText"
            :data="dataList"
          >
            //  @row-click="openDetail"
            <el-table-column type="index"  label="序号" width="60" />
            <el-table-column prop="name"  label="模块组" width="200"/>
            <el-table-column prop="area"  label="模块名" width="100" />
            <el-table-column prop="properties.spot_time" sortable label="焊点时间" width="180"/>
            <el-table-column prop="properties.control_mode" label="控制模式" width="100" />
            <el-table-column prop="properties.type_id" label="类型ID" width="100" />
            <el-table-column prop="properties.spot_number" label="焊点编号" width="100" />
            <el-table-column prop="properties.program_number" label="程序编号" width="100" />
            <el-table-column prop="properties.spot_name" label="焊点名称" width="auto"  min-width="200"/>
            <el-table-column prop="properties.pre_time" sortable label="预热时间(ms)" width="140"/>
            <el-table-column prop="properties.real_time" sortable label="焊接时间(ms)" width="140"/>
            <el-table-column prop="properties.keep_time" sortable label="维持时间(ms)" width="140"/>
            <el-table-column prop="properties.set_current" label="设定电流(A)" width="140"/>
            <el-table-column prop="properties.real_mean_current" label="实际电流(A)" width="140"/>
            <el-table-column prop="properties.real_mean_voltage" label="实际电压(V)" width="140"/>
            <el-table-column prop="properties.real_energy" label="能量(J)" width="100"/>
            <el-table-column prop="properties.spot_quality" label="焊接质量" width="100"/>
            <el-table-column prop="properties.q_threshold" label="Q阈值" width="100"/>
            <el-table-column prop="properties.q_value" label="Q值" width="100"/>
            <el-table-column prop="properties.r_max" label="R最大值" width="100"/>
            <el-table-column prop="properties.spatter" label="飞溅时间点" width="120"/>
            <el-table-column prop="properties.milling_cycle_counter" label="循环计数" width="100"/>
            <el-table-column prop="properties.milling_weld_spot_counter" label="修模计数" width="100"/>
            <el-table-column prop="properties.archive_series" label="归档序列号" width="180"/>
            <el-table-column fixed="right" >
              <template #header>
                操作
              </template>
              <template #default="scope">
                <el-button
                    size="small"
                    type="danger"
                    @click="openDetail(scope.row)"
                >Delete</el-button
                >
              </template>
            </el-table-column>
          </el-table>
          <div class="data-card-main-footer">
            <el-pagination
                class="pagination"
                small
                background
                v-model:page-size="pageSize"
                v-model:current-page="currentPage"
                :total="pageTotal"
                :page-sizes="[50, 100, 300, 500, 1000]"
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