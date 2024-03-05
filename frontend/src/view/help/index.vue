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

<!--<template>-->
<!--  <el-container>-->
<!--    &lt;!&ndash;==========Collection Header==========&ndash;&gt;-->
<!--    <el-header class="collection-index-header" :style="winSize.headerSize">-->
<!--      <el-button-group   class="collection-header-button" style="float:left">-->
<!--        <el-button-->
<!--            icon="VideoPlay"-->
<!--            :type="BtnStatus.StartBtnStatus"-->
<!--            :color="BtnStatus.StartBtnColor"-->
<!--            @click="StartBtnFunc"-->
<!--        >开始采集</el-button>-->
<!--        <el-button-->
<!--            icon="VideoPause"-->
<!--            :type="BtnStatus.StopBtnStatus"-->
<!--            :color="BtnStatus.StopBtnColor"-->
<!--            @click="StopBtnFunc"-->
<!--        >停止采集</el-button>-->
<!--        <el-button-->
<!--            icon="Document"-->
<!--            :type="BtnStatus.SaveBtnStatus"-->
<!--            :color="BtnStatus.SaveBtnColor"-->
<!--            @click="SaveBtnFunc"-->
<!--        >记录数据</el-button>-->

<!--      </el-button-group>-->

<!--      &lt;!&ndash;    TODO:  测试信息太长情况是需要截取部分&ndash;&gt;-->
<!--      <el-button-group   class="collection-header-text" style="float:right">-->
<!--        <font size="2" style="padding-left: 30px">测试人员：{{ TestInformation.TestUser }}</font>-->
<!--        <font size="2" style="padding-left: 30px">测试地点：{{ TestInformation.TestAddress }}</font>-->
<!--        <font size="2" style="padding-left: 30px">测试单位：{{ TestInformation.TestCompany }}</font>-->
<!--      </el-button-group>-->
<!--    </el-header>-->
<!--    <el-container>-->
<!--      <el-container>-->
<!--        &lt;!&ndash;==========Collection Main==========&ndash;&gt;-->
<!--        <el-main class="collection-main" :style="winSize.mainSize">-->
<!--          <el-row>-->
<!--            <div id="collectionEcharts" :style="winSize.mainEchartsSize"></div>-->
<!--          </el-row>-->
<!--          <el-row :gutter="10" :style="winSize.mainFormSize">-->
<!--            <el-col span="4px" style="height: 25px" >-->
<!--              <el-form :model="CollectionData.FormData" label-width="120px">-->
<!--                <el-form-item label="缓存点数:">-->
<!--                  <el-input style="height: 25px" v-model="CollectionData.FormData.CachePoint" />-->
<!--                </el-form-item>-->
<!--              </el-form>-->
<!--            </el-col>-->
<!--            <el-col span="4px" style="height: 25px" >-->
<!--              <el-form :model="CollectionData.FormData" label-width="120px">-->
<!--                <el-form-item label="开始时间:">-->
<!--                  <el-input  style="height: 25px; background-color: #73767a" v-model="CollectionData.FormData.StartTime" />-->
<!--                </el-form-item>-->
<!--              </el-form>-->
<!--            </el-col>-->
<!--            <el-col span="4px" style="height: 25px" >-->
<!--              <el-form :model="CollectionData.FormData" label-width="120px">-->
<!--                <el-form-item label="采集盒温度:">-->
<!--                  <el-input style="height: 25px" v-model="CollectionData.FormData.Temperature" />-->
<!--                </el-form-item>-->
<!--              </el-form>-->
<!--            </el-col>-->
<!--          </el-row>-->

<!--        </el-main>-->
<!--        &lt;!&ndash;==========Collection Footer==========&ndash;&gt;-->
<!--        <el-footer class="collection-footer" :style="winSize.footerSize">-->
<!--          <el-card class="collection-footer-card" :style="winSize.footerCardSize">-->
<!--            <el-table-->
<!--                :header-cell-style="{padding:'4px'}"-->
<!--                :row-style="{height:'18px'}"-->
<!--                :cell-style="{padding:'0px'}"-->
<!--                class="collection-footer-table"-->
<!--                :data="CollectionTableData"-->
<!--                :style="winSize.footerTableSize"-->
<!--                border >-->
<!--              <el-table-column prop="Name" label="名称" auto align="center" />-->
<!--              <el-table-column prop="Data[0]" label="第1路" auto align="center" />-->
<!--              <el-table-column prop="Data[1]" label="第2路" auto align="center" />-->
<!--              <el-table-column prop="Data[2]" label="第3路" auto align="center" />-->
<!--              <el-table-column prop="Data[3]" label="第4路" auto align="center" />-->
<!--              <el-table-column prop="Data[4]" label="第5路" auto align="center" />-->
<!--              <el-table-column prop="Data[5]" label="第6路" auto align="center" />-->
<!--              <el-table-column prop="Data[6]" label="第7路" auto align="center" />-->
<!--              <el-table-column prop="Data[7]" label="第8路" auto align="center" />-->
<!--            </el-table>-->
<!--          </el-card>-->

<!--        </el-footer>-->
<!--      </el-container>-->

<!--      &lt;!&ndash;==========Collection Right Aside==========&ndash;&gt;-->
<!--      <el-aside class="collection-right-aside" :style="winSize.rightAsideSize">-->
<!--        <el-scrollbar :style="winSize.rightAsideSize">-->
<!--          <div class="collection-right-aside-tree">-->
<!--            <el-tree-->
<!--                :data="treeOptions"-->
<!--                :props="defaultProps"-->
<!--                :default-expanded-keys="[1,2,3,4,5]"-->
<!--                :default-checked-keys="[0,1,2,3,4,5,6,7,8]"-->
<!--                show-checkbox-->
<!--                @check="checkChange"-->
<!--                ref="tree"-->
<!--                node-key="id"-->
<!--                style="background: #2B2D30; color: #c8c9cc"-->
<!--            />-->
<!--          </div>-->
<!--        </el-scrollbar>-->
<!--      </el-aside>-->
<!--    </el-container>-->

<!--  </el-container>-->
<!--</template>-->

<!--<script>-->
<!--import {StartCollectionDataAPI,StopCollectionDataAPI,SaveCollectionDataAPI} from "../../../wailsjs/go/end/Backend";-->
<!--import {CollectionBtnStatusAPI,CollectionDataAPI} from "../../../wailsjs/go/end/Backend";-->
<!--import {EventsOn} from "../../../wailsjs/runtime/runtime.js";-->

<!--export default {-->
<!--  data() {-->
<!--    return {-->
<!--      myCharts: "",-->
<!--      // ========== Window Size Data ========== //-->
<!--      winSize:{-->
<!--        headerSize: {-->
<!--          width: "",-->
<!--        },-->
<!--        mainSize: {-->
<!--          width: "",-->
<!--          height: "",-->
<!--        },-->
<!--        footerSize: {-->
<!--          width: "",-->
<!--        },-->
<!--        rightAsideSize: {-->
<!--          width: "",-->
<!--          height: "",-->
<!--        },-->
<!--        mainEchartsSize:{-->
<!--          width: "",-->
<!--          height: "",-->
<!--        },-->
<!--        mainFormSize:{-->
<!--          width: "",-->
<!--        },-->
<!--        footerCardSize: {-->
<!--          width: "",-->
<!--        },-->
<!--        footerTableSize: {-->
<!--          width: "",-->
<!--        },-->
<!--      },-->

<!--      // ========== Collection Header Data ========== //-->
<!--      BtnStatus:{-->
<!--        StartBtnStatus: "",-->
<!--        StopBtnStatus: "",-->
<!--        SaveBtnStatus: "",-->
<!--        StartBtnColor: "#2B2D30",-->
<!--        StopBtnColor: "#2B2D30",-->
<!--        SaveBtnColor: "#2B2D30",-->
<!--      },-->
<!--      TestInformation: {-->
<!--        TestUser: "XXX",-->
<!--        TestAddress: "北京市通州区",-->
<!--        TestCompany: "北京XXX技术有限公司",-->
<!--      },-->

<!--      // ========== Collection Main Data ========== //-->
<!--      CachePointCount: 0,-->
<!--      SystemConfig: {-->
<!--        Version: "",-->
<!--        Description: "",-->
<!--        DisplayParam: {-->
<!--          TestUser: "",-->
<!--          TestAddress: "",-->
<!--          TestCompany: "",-->
<!--          ProbePipeSetlList: [{-->
<!--            ID: "",-->
<!--            ChannelName: "",-->
<!--            PowerStatus: "",-->
<!--            VoltageThreshold: "",-->
<!--            CurrentThreshold: "",-->
<!--            PipeNumber: "",-->
<!--          }],-->
<!--          TestDescription: "",-->
<!--        },-->
<!--        TestInfo: {-->
<!--          ColorList: [{-->
<!--            ID: "",-->
<!--            Name: "",-->
<!--            Color: "",-->
<!--          }],-->
<!--          SymbolList: [{-->
<!--            ID: "",-->
<!--            Name: "",-->
<!--            Symbol: "",-->
<!--          }],-->
<!--          ScalingRatio: "1:1",-->
<!--          CursorList: [{-->
<!--            ID: "",-->
<!--            Name: "",-->
<!--            Color: "",-->
<!--            LineType: "",-->
<!--            LineWidth: "",-->
<!--          }],-->
<!--        },-->
<!--        SystemParam: {-->
<!--          PortName: "",-->
<!--          TimeInterval: "",-->
<!--          FilePrefix: "",-->
<!--          WorkPath: "",-->
<!--        }-->
<!--      },-->
<!--      MainFormData: {-->
<!--        CachePoint: 100,-->
<!--        StartTime: "2023-03-04 13:30:10.543",-->
<!--        CollectionBoxTemperature: "30.5",-->
<!--      },-->


<!--      CollectionData: {-->
<!--        TableData: {-->
<!--          Route1: {-->
<!--            Name: "",-->
<!--            Data: [],-->
<!--          },-->
<!--          Route2: {-->
<!--            Name: "",-->
<!--            Data: [],-->
<!--          },-->
<!--          Route3: {-->
<!--            Name: "",-->
<!--            Data: [],-->
<!--          },-->
<!--          Route4: {-->
<!--            Name: "",-->
<!--            Data: [],-->
<!--          },-->
<!--          Route5: {-->
<!--            Name: "",-->
<!--            Data: [],-->
<!--          },-->
<!--          Route6: {-->
<!--            Name: "",-->
<!--            Data: [],-->
<!--          },-->
<!--          Route7: {-->
<!--            Name: "",-->
<!--            Data: [],-->
<!--          },-->
<!--        },-->
<!--        CurveData:{-->
<!--          Temperature: [],-->
<!--          CurrentData: [[],[],[],[],[],[],[],[]],-->
<!--          VoltageData: [[],[],[],[],[],[],[],[]],-->
<!--          Count: [[],[],[],[],[],[],[],[]],-->
<!--          CountRate: [[],[],[],[],[],[],[],[]],-->
<!--          PowerSwitch: [[],[],[],[],[],[],[],[]],-->
<!--        },-->
<!--        FormData: {-->
<!--          Temperature: 0,-->
<!--          StartTime: "",-->
<!--          CachePoint: 0,-->
<!--        }-->
<!--      },-->
<!--      Colors: ['#c93756','#177cb0','#16a951','#8d4bbb','#ffb61e','#1baeae'],-->

<!--      // ========== Collection Footer Data ========== //-->
<!--      CollectionTableData: [{-->
<!--        Name: "供电电流(mA)",-->
<!--        Data: [0,0,0,0,0,0,0,0]-->
<!--      },{-->
<!--        Name: "供电电压(V)",-->
<!--        Data: [0,0,0,0,0,0,0,0]-->
<!--      }, {-->
<!--        Name: "计数",-->
<!--        Data: [0,0,0,0,0,0,0,0]-->
<!--      }, {-->
<!--        Name: "计数率",-->
<!--        Data: [0,0,0,0,0,0,0,0]-->
<!--      }, {-->
<!--        Name: "电源开关",-->
<!--        Data: ["","","","","","","",""]-->
<!--      },{-->
<!--        Name: "电流状态",-->
<!--        Data: ["","","","","","","",""]-->
<!--      },{-->
<!--        Name: "电压状态",-->
<!--        Data:["","","","","","","",""]-->
<!--      }] ,-->

<!--      // ========== Collection Right Aside Data ========== //-->
<!--      // 控制popover弹出框是否展示-->
<!--      popoverVisible: false,-->
<!--      // 需把数据整理成以下结构-->
<!--      // tree数据(children的id第一位为父级id，用于在select中清除某一点，可找到其父级去掉全选)-->
<!--      treeOptions: [-->
<!--        {-->
<!--          id: '0',-->
<!--          name: '环境参数',-->
<!--          children: [-->
<!--            {-->
<!--              id: '0-1',-->
<!--              name: '测试盒温度'-->
<!--            }-->
<!--          ]-->
<!--        }, {-->
<!--          id: '1',-->
<!--          name: '第1路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '1-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '1-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '1-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '1-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '1-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        },{-->
<!--          id: '2',-->
<!--          name: '第2路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '2-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '2-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '2-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '2-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '2-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        }, {-->
<!--          id: '3',-->
<!--          name: '第3路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '3-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '3-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '3-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '3-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '3-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        }, {-->
<!--          id: '4',-->
<!--          name: '第4路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '4-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '4-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '4-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '4-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '4-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        }, {-->
<!--          id: '5',-->
<!--          name: '第5路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '5-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '5-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '5-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '5-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '5-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        }, {-->
<!--          id: '6',-->
<!--          name: '第6路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '6-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '6-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '6-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '6-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '6-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        }, {-->
<!--          id: '7',-->
<!--          name: '第7路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '7-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '7-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '7-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '7-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '7-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        }, {-->
<!--          id: '8',-->
<!--          name: '第8路',-->
<!--          children: [-->
<!--            {-->
<!--              id: '8-1',-->
<!--              name: '供电电压'-->
<!--            }, {-->
<!--              id: '8-2',-->
<!--              name: '供电电流'-->
<!--            }, {-->
<!--              id: '8-3',-->
<!--              name: '计数'-->
<!--            }, {-->
<!--              id: '8-4',-->
<!--              name: '计数率'-->
<!--            }, {-->
<!--              id: '8-5',-->
<!--              name: '供电开关'-->
<!--            }-->
<!--          ]-->
<!--        }-->
<!--      ],-->
<!--      // select选择框选项-->
<!--      typeOption: [-->
<!--        {-->
<!--          id: '1',-->
<!--          name: '父级1'-->
<!--        },-->
<!--        {-->
<!--          id: '1-1',-->
<!--          name: '类型1'-->
<!--        }, {-->
<!--          id: '1-2',-->
<!--          name: '类型2'-->
<!--        }, {-->
<!--          id: '1-3',-->
<!--          name: '类型3'-->
<!--        }, {-->
<!--          id: '1-4',-->
<!--          name: '类型4'-->
<!--        }, {-->
<!--          id: '1-5',-->
<!--          name: '类型5'-->
<!--        }, {-->
<!--          id: '2',-->
<!--          name: '父级2'-->
<!--        }, {-->
<!--          id: '2-1',-->
<!--          name: '类型6'-->
<!--        }, {-->
<!--          id: '2-2',-->
<!--          name: '类型7'-->
<!--        }-->
<!--      ],-->
<!--      defaultProps: {-->
<!--        children: 'children',-->
<!--        label: 'name'-->
<!--      },-->
<!--      // select的值-->
<!--      typeValue: [],-->


<!--    };-->
<!--  },-->
<!--  mounted() {-->
<!--    // 页面打开时加载-->
<!--    this.CollectionBtnStatusFunc();-->
<!--    this.CollectionDataFunc();-->

<!--    // 先初始化图表，再更新数据-->
<!--    this.initEcharts(this.CollectionData.CurveData);-->


<!--    // 实时获取后端数据-->
<!--    EventsOn("test",data => {-->
<!--      this.CollectionData = data-->


<!--      // Collection Table Data-->
<!--      this.CollectionTableData[0]=this.CollectionData.TableData.Route1-->
<!--      this.CollectionTableData[1]=this.CollectionData.TableData.Route2-->
<!--      this.CollectionTableData[2]=this.CollectionData.TableData.Route3-->
<!--      this.CollectionTableData[3]=this.CollectionData.TableData.Route4-->
<!--      this.CollectionTableData[4]=this.CollectionData.TableData.Route5-->
<!--      this.CollectionTableData[5]=this.CollectionData.TableData.Route6-->
<!--      this.CollectionTableData[6]=this.CollectionData.TableData.Route7-->


<!--      this.updateEcharts(this.CollectionData.CurveData);-->
<!--    })-->

<!--    this.getWindowSize();-->

<!--  },-->
<!--  watch: {-->

<!--  },-->
<!--  methods: {-->
<!--    // ========== Window Size Function ========== //-->
<!--    getWindowSize() {-->
<!--      this.winSize.headerSize.width = window.innerWidth - 40 - 2 -40 -2+ "px"-->
<!--      this.winSize.mainSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 -2 + "px"-->
<!--      this.winSize.mainSize.height = window.innerHeight - 70 - 2 -60 - 250 - 2 + "px"-->
<!--      this.winSize.mainEchartsSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2 - 2 + "px"-->
<!--      this.winSize.mainEchartsSize.height = window.innerHeight - 70 - 2 -60 - 250 - 2  -2 - 40+ "px"-->
<!--      this.winSize.mainFormSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2 - 2 + "px"-->
<!--      this.winSize.footerSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2+ "px"-->
<!--      this.winSize.footerCardSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2 - 22 + "px"-->
<!--      this.winSize.footerTableSize.width = window.innerWidth - 40 - 2 -40 - 2 - 180 - 2 - 40 + "px"-->
<!--      this.winSize.rightAsideSize.height = window.innerHeight - 70 - 2 -60  - 2  -2 + "px"-->

<!--    },-->

<!--    // ========== Collection Header Function ==========//-->
<!--    CollectionDataFunc(){-->
<!--      CollectionDataAPI().then(result =>{-->
<!--        this.CollectionData = result-->

<!--        // Collection Table Data-->
<!--        this.CollectionTableData[0]=this.CollectionData.TableData.Route1-->
<!--        this.CollectionTableData[1]=this.CollectionData.TableData.Route2-->
<!--        this.CollectionTableData[2]=this.CollectionData.TableData.Route3-->
<!--        this.CollectionTableData[3]=this.CollectionData.TableData.Route4-->
<!--        this.CollectionTableData[4]=this.CollectionData.TableData.Route5-->
<!--        this.CollectionTableData[5]=this.CollectionData.TableData.Route6-->
<!--        this.CollectionTableData[6]=this.CollectionData.TableData.Route7-->
<!--      })-->
<!--    },-->
<!--    CollectionBtnStatusFunc(){-->
<!--      CollectionBtnStatusAPI().then(result =>{-->
<!--        this.BtnStatus = result-->
<!--      })-->
<!--    },-->

<!--    StartBtnFunc(){-->
<!--      StartCollectionDataAPI().then(result =>{-->
<!--        // delete this.TemperatureList[0:]-->
<!--        this.BtnStatus = result-->
<!--      })-->
<!--    },-->
<!--    StopBtnFunc(){-->
<!--      StopCollectionDataAPI().then(result =>{-->
<!--        this.BtnStatus = result-->
<!--      })-->
<!--    },-->
<!--    SaveBtnFunc(){-->
<!--      SaveCollectionDataAPI().then(result =>{-->
<!--        this.BtnStatus = result-->
<!--        console.log(result)-->
<!--      })-->
<!--    },-->
<!--    // ========== Collection Main Function ========== //-->

<!--    initEcharts(es_data){-->
<!--      this.myChart = this.$echarts.init(document.getElementById("collectionEcharts"));-->
<!--      const option = {-->
<!--        toolbox: {-->
<!--          show: true,-->
<!--        },-->
<!--        tooltip: {-->
<!--          trigger: "axis",-->
<!--          axisPointer: {-->
<!--            type: 'line',-->
<!--            snap: false,-->
<!--            lineStyle: {-->
<!--              color: '#16a951'-->
<!--            },-->
<!--          },-->
<!--        },-->
<!--        grid: {-->
<!--          left: "3%",-->
<!--          right: "4%",-->
<!--          bottom: "3%",-->
<!--          containLabel: true,-->
<!--        },-->
<!--        legend: {-->
<!--          show: true,-->
<!--          textStyle: {-->
<!--            color: '#eeeeee',-->
<!--          },-->
<!--          data: ["供电电流", "供电电压", "采集盒温度", "计数", "计数率", "电源开关"],-->
<!--        },-->
<!--        xAxis: {-->
<!--          name: "时间(ms)",-->
<!--          nameLocation: 'end',-->
<!--          type: "category",-->
<!--          axisTick: {-->
<!--            show: false-->
<!--          },-->
<!--          nameTextStyle: {-->
<!--            fontSize: 14,-->
<!--            lineHeight: 60,-->
<!--            color: '#eeeeee',-->
<!--            align: "center",-->
<!--            verticalAlign: "left",-->
<!--          },-->
<!--        },-->
<!--        yAxis:[-->
<!--          {-->
<!--            //供电电流 第一个y轴位置在左侧-->
<!--            position:'left',-->
<!--            offset: 0,-->
<!--            type : 'value',-->
<!--            // 名字-->
<!--            name: '供电电流(mA)',-->
<!--            nameLocation: 'middle',-->
<!--            nameGap: 23,-->
<!--            min: 0, // 坐标Y轴刻度最小值-->
<!--            max: 800, // 坐标Y轴刻度最大值-->
<!--            interval: 50, // 坐标Y轴刻度间隔-->
<!--            axisLine: { //坐标线-->
<!--              show: true ,-->
<!--              symbol: "pin",-->
<!--              lineStyle: {-->
<!--                color: this.Colors[0],-->
<!--                width: '0.5',-->
<!--              }-->
<!--            },-->
<!--            axisTick: { // 坐标刻度-->
<!--              show: true,-->
<!--            },-->
<!--            axisLabel: {-->
<!--              show: true,-->
<!--              fontFamily: "Arial",-->
<!--              verticalAlign: "middle",-->
<!--              formatter: '{value}'-->
<!--            },-->
<!--            splitLine: {-->
<!--              show: false-->
<!--            },-->
<!--            axisPointer: {-->
<!--              show: true,-->
<!--              type: 'line',-->
<!--              lineStyle: {-->
<!--                color: '#16a951'-->
<!--              }-->
<!--            }-->
<!--          },-->
<!--          {-->
<!--            //供电电压 第一个y轴位置在左侧-->
<!--            position:'left',-->
<!--            offset: 45,-->
<!--            type : 'value',-->
<!--            // 名字-->
<!--            name: '供电电压(V)',-->
<!--            nameLocation: 'middle',-->
<!--            nameGap: 22,-->
<!--            min: -10, // 坐标Y轴刻度最小值-->
<!--            max: 30, // 坐标Y轴刻度最大值-->
<!--            interval: 2, // 坐标Y轴刻度间隔-->
<!--            axisLine: { //坐标线-->
<!--              show: true ,-->
<!--              symbol: "pin",-->
<!--              lineStyle: {-->
<!--                color: this.Colors[1],-->
<!--                width: '0.5',-->
<!--              }-->
<!--            },-->
<!--            axisTick: { // 坐标刻度-->
<!--              show: true,-->
<!--            },-->
<!--            axisLabel: {-->
<!--              show: true,-->
<!--              fontFamily: "Arial",-->
<!--              verticalAlign: "middle",-->
<!--              formatter: '{value}'-->
<!--            },-->
<!--            splitLine: {-->
<!--              show: false-->
<!--            },-->
<!--            axisPointer: {-->
<!--              show: true,-->
<!--              type: 'line',-->
<!--              lineStyle: {-->
<!--                color: '#16a951'-->
<!--              }-->
<!--            }-->
<!--          },-->
<!--          {-->
<!--            //计数率 第一个y轴位置在左侧-->
<!--            position:'left',-->
<!--            offset: 90,-->
<!--            type : 'value',-->
<!--            // 名字-->
<!--            name: '采集盒温度(℃)',-->
<!--            nameLocation: 'middle',-->
<!--            nameGap: 23,-->
<!--            min: 0, // 坐标Y轴刻度最小值-->
<!--            max: 100, // 坐标Y轴刻度最大值-->
<!--            interval: 5, // 坐标Y轴刻度间隔-->
<!--            axisLine: { //坐标线-->
<!--              show: true ,-->
<!--              symbol: "pin",-->
<!--              lineStyle: {-->
<!--                color: this.Colors[2],-->
<!--                width: '0.5',-->
<!--              }-->
<!--            },-->
<!--            axisTick: { // 坐标刻度-->
<!--              show: true,-->
<!--            },-->
<!--            axisLabel: {-->
<!--              show: true,-->
<!--              fontFamily: "Arial",-->
<!--              verticalAlign: "middle",-->
<!--              formatter: '{value}'-->
<!--            },-->
<!--            splitLine: {-->
<!--              show: false-->
<!--            },-->
<!--            axisPointer: {-->
<!--              show: true,-->
<!--              type: 'line',-->
<!--              lineStyle: {-->
<!--                color: '#16a951'-->
<!--              }-->
<!--            }-->
<!--          },-->
<!--          {-->
<!--            //计数 第一个y轴位置在左侧-->
<!--            position:'right',-->
<!--            offset: 0,-->
<!--            type : 'value',-->
<!--            // 名字-->
<!--            name: '计数(cnt)',-->
<!--            nameLocation: 'middle',-->
<!--            nameGap: 28,-->
<!--            min: 0.9, // 坐标Y轴刻度最小值-->
<!--            max: 1.1, // 坐标Y轴刻度最大值-->
<!--            interval: 0.02, // 坐标Y轴刻度间隔-->
<!--            axisLine: { //坐标线-->
<!--              show: true ,-->
<!--              symbol: "pin",-->
<!--              lineStyle: {-->
<!--                color: this.Colors[3],-->
<!--                width: '0.5',-->
<!--              }-->
<!--            },-->
<!--            axisTick: { // 坐标刻度-->
<!--              show: true,-->
<!--            },-->
<!--            axisLabel: {-->
<!--              show: true,-->
<!--              fontFamily: "Arial",-->
<!--              verticalAlign: "middle",-->
<!--              formatter: '{value}'-->
<!--            },-->
<!--            splitLine: {-->
<!--              show: false-->
<!--            },-->
<!--            axisPointer: {-->
<!--              show: true,-->
<!--              type: 'line',-->
<!--              lineStyle: {-->
<!--                color: '#16a951'-->
<!--              }-->
<!--            }-->
<!--          },-->
<!--          {-->
<!--            //计数率 第一个y轴位置在左侧-->
<!--            position:'right',-->
<!--            offset: 45,-->
<!--            type : 'value',-->
<!--            // 名字-->
<!--            name: '计数率(cnts/s)',-->
<!--            nameLocation: 'middle',-->
<!--            nameGap: 28,-->
<!--            min: 0, // 坐标Y轴刻度最小值-->
<!--            max: 1, // 坐标Y轴刻度最大值-->
<!--            interval: 0.05, // 坐标Y轴刻度间隔-->
<!--            axisLine: { //坐标线-->
<!--              show: true ,-->
<!--              symbol: "pin",-->
<!--              lineStyle: {-->
<!--                color: this.Colors[4],-->
<!--                width: '0.5',-->
<!--              }-->
<!--            },-->
<!--            axisTick: { // 坐标刻度-->
<!--              show: true,-->
<!--            },-->
<!--            axisLabel: {-->
<!--              show: true,-->
<!--              fontFamily: "Arial",-->
<!--              verticalAlign: "middle",-->
<!--              formatter: '{value}'-->
<!--            },-->
<!--            splitLine: {-->
<!--              show: false-->
<!--            },-->
<!--            axisPointer: {-->
<!--              show: true,-->
<!--              type: 'line',-->
<!--              lineStyle: {-->
<!--                color: '#16a951'-->
<!--              }-->
<!--            }-->
<!--          },-->
<!--          {-->
<!--            //计数率 第一个y轴位置在左侧-->
<!--            position:'right',-->
<!--            offset: 90,-->
<!--            type : 'value',-->
<!--            name: '供电开关',-->
<!--            nameLocation: 'middle',-->
<!--            nameGap: 32,-->
<!--            min: -2, // 坐标Y轴刻度最小值-->
<!--            max: 2, // 坐标Y轴刻度最大值-->
<!--            interval: 0.2, // 坐标Y轴刻度间隔-->
<!--            axisLine: { //坐标线-->
<!--              show: true ,-->
<!--              symbol: "pin",-->
<!--              lineStyle: {-->
<!--                color: this.Colors[5],-->
<!--                width: '0.5',-->
<!--              }-->
<!--            },-->
<!--            axisTick: { // 坐标刻度-->
<!--              show: true,-->
<!--            },-->
<!--            axisLabel: {-->
<!--              show: true,-->
<!--              fontFamily: "Arial",-->
<!--              verticalAlign: "middle",-->
<!--              formatter: '{value}'-->
<!--            },-->
<!--            splitLine: {-->
<!--              show: false-->
<!--            },-->
<!--            axisPointer: {-->
<!--              show: true,-->
<!--              type: 'line',-->
<!--              lineStyle: {-->
<!--                color: '#16a951'-->
<!--              }-->
<!--            }-->
<!--          },-->
<!--        ],-->
<!--        series: [-->
<!--          {-->
<!--            name: "供电电流",-->
<!--            type: "line",-->
<!--            symbol: "circle",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[0],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[0],-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 0,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.CurrentData[0],-->
<!--          },-->
<!--          {-->
<!--            name: "供电电压",-->
<!--            type: "line",-->
<!--            symbol: "rect",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[1],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[1]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 1,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.VoltageData[0],-->
<!--          },-->
<!--          {-->
<!--            name: "采集盒温度",-->
<!--            type: "line",-->
<!--            symbol: "roundRect",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[2],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[2]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 2,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.Temperature,-->
<!--          },-->
<!--          {-->
<!--            name: "计数",-->
<!--            type: "line",-->
<!--            symbol: "diamond",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[3],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[3]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 3,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.Count[0],-->
<!--          },-->
<!--          {-->
<!--            name: "计数率",-->
<!--            type: "line",-->
<!--            symbol: "triangle",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[4],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[4]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 4,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.CountRate[0],-->
<!--          },-->
<!--          {-->
<!--            name: "电源开关",-->
<!--            type: "line",-->
<!--            symbol: "emptyCircle", // 拐点形状-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[5],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[5]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 5,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.PowerSwitch[0],-->
<!--          },-->
<!--        ],-->
<!--      };-->

<!--      this.myChart.setOption(option, true);-->

<!--      window.addEventListener('resize',function () {-->
<!--        this.myChart.resize();-->
<!--      });-->

<!--    },-->
<!--    updateEcharts(es_data) {-->
<!--      const option = {-->
<!--        series: [-->
<!--          {-->
<!--            name: "供电电流",-->
<!--            type: "line",-->
<!--            symbol: "circle",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[0],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[0],-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 0,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.CurrentData[0],-->
<!--          },-->
<!--          {-->
<!--            name: "供电电压",-->
<!--            type: "line",-->
<!--            symbol: "rect",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[1],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[1]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 1,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.VoltageData[0],-->
<!--          },-->
<!--          {-->
<!--            name: "采集盒温度",-->
<!--            type: "line",-->
<!--            symbol: "roundRect",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[2],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[2]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 2,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.Temperature,-->
<!--          },-->
<!--          {-->
<!--            name: "计数",-->
<!--            type: "line",-->
<!--            symbol: "diamond",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[3],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[3]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 3,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.Count[0],-->
<!--          },-->
<!--          {-->
<!--            name: "计数率",-->
<!--            type: "line",-->
<!--            symbol: "triangle",-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[4],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[4]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 4,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.CountRate[0],-->
<!--          },-->
<!--          {-->
<!--            name: "电源开关",-->
<!--            type: "line",-->
<!--            symbol: "emptyCircle", // 拐点形状-->
<!--            symbolSize: 10, //拐点大小-->
<!--            itemStyle: {-->
<!--              normal: {-->
<!--                color: this.Colors[5],-->
<!--                lineStyle: {-->
<!--                  width: 2, // 设置线宽-->
<!--                  color: this.Colors[5]-->
<!--                }-->
<!--              }-->
<!--            },-->
<!--            yAxisIndex: 5,//使用第一个y轴，序号从0开始-->
<!--            data: es_data.PowerSwitch[0],-->
<!--          },-->
<!--        ],-->
<!--      };-->

<!--      this.myChart.setOption(option, true);-->

<!--    },-->

<!--    // ========== Collection Footer Function ========== //-->

<!--    // ========== Collection Right Aside Function ========== //-->
<!--    checkChange () {-->
<!--      this.typeValue = []-->
<!--      // 将tree选择的id赋值给select-->
<!--      this.$refs['tree'].getCheckedNodes().forEach(value => {-->
<!--        // 父级在select中不展示-->
<!--        if (value.id.indexOf('-') > 0) {-->
<!--          this.typeValue.push(value.id)-->
<!--        }-->
<!--      })-->
<!--    },-->
<!--    // 清空select-->
<!--    clearTag () {-->
<!--      // 清空tree选择-->
<!--      this.$refs['tree'].setCheckedKeys([])-->
<!--    },-->
<!--    // 从select中单个移除时，保持tree选择值同步移除-->
<!--    removeTag (data) {-->
<!--      // 获取tree目前选择的值-->
<!--      const chooseData = this.$refs['tree'].getCheckedKeys()-->
<!--      let deleteIndex = ''-->
<!--      // 找到chooseData中与清除的data相同的值-->
<!--      chooseData.forEach((value, index) => {-->
<!--        if (value === data) {-->
<!--          deleteIndex = index-->
<!--        }-->
<!--      })-->
<!--      // 从tree目前选择值中去掉-->
<!--      chooseData.splice(deleteIndex, 1)-->
<!--      // 若有全选情况，tree的选择值中有父级id，而select中无父级id，需用children的id找到父级id并去掉-->
<!--      // 查找其父级id是否在chooseData中（即原来此父级是否全选），若在则去掉-->
<!--      const findFatherData = chooseData.find(element => element === (data.split('-')[0]))-->
<!--      if (findFatherData) {-->
<!--        chooseData.splice(chooseData.indexOf(findFatherData), 1)-->
<!--      }-->
<!--      // 将修改后的值再赋给tree-->
<!--      this.$refs['tree'].setCheckedKeys(chooseData)-->
<!--    },-->
<!--  },-->

<!--  created() {-->
<!--    window.addEventListener("resize", this.getWindowSize);-->
<!--    this.getWindowSize();-->


<!--  },-->

<!--}-->
<!--</script>-->

<!--<style scoped>-->
<!--/*========== Collection Header ==========*/-->
<!--.collection-index-header{-->
<!--  padding: 0;-->
<!--  background: #2B2D30;-->
<!--  height: 60px;-->
<!--}-->

<!--.collection-header-button{-->
<!--  width: 350px;-->
<!--  margin-left: 30px;-->
<!--  padding: 15px 15px 15px 15px;-->
<!--}-->
<!--.collection-header-text{-->
<!--  padding: 20px 20px 20px 20px;-->
<!--}-->

<!--.el-button{-->
<!--  /*border-color: #73767a;*/-->
<!--}-->

<!--/*========== Collection Main ==========*/-->
<!--.collection-main{-->
<!--  padding: 1px 1px 1px 1px;-->
<!--  background: #1E1F22;-->
<!--}-->

<!--#collection-main-echarts {-->
<!--  font-family: Avenir, Helvetica, Arial, sans-serif;-->
<!--  -webkit-font-smoothing: antialiased;-->
<!--  -moz-osx-font-smoothing: grayscale;-->
<!--  text-align: center;-->
<!--  color: #eeeeee;-->
<!--  margin-top: 0;-->
<!--}-->
<!--::v-deep .el-form-item__label {-->
<!--  font-size: 14px;-->
<!--  color: #FFFFFF;-->
<!--}-->
<!--.el-input{-->
<!--  &#45;&#45;el-input-text-color: var(&#45;&#45;el-fill-color-blank);-->
<!--  &#45;&#45;el-input-bg-color: var(&#45;&#45;el-text-color-primary);-->
<!--  &#45;&#45;el-input-border-color: var(&#45;&#45;el-color-info);-->
<!--}-->


<!--/*==========Collection Right Aside ==========*/-->
<!--.collection-right-aside{-->
<!--  width: 180px;-->
<!--  margin-top: 2px;-->
<!--  margin-right: 2px;-->
<!--  background: #2B2D30;-->

<!--}-->
<!--.el-tree {-->
<!--  &#45;&#45;el-tree-node-hover-bg-color: var(&#45;&#45;el-text-color-regular);-->
<!--  &#45;&#45;el-tree-text-color: var(&#45;&#45;el-text-color-regular);-->
<!--  &#45;&#45;el-tree-expand-icon-color: var(&#45;&#45;el-text-color-placeholder);-->
<!--}-->
<!--/*========== Collection Footer ==========*/-->
<!--.collection-footer{-->
<!--  padding: 0;-->
<!--  background: #2B2D30;-->
<!--  height: 250px;-->
<!--}-->
<!--.collection-footer-card{-->
<!--  width: 1100px;-->
<!--  height: 230px;-->
<!--  margin-top: 10px;-->
<!--  margin-left: 10px;-->
<!--  padding: 0;-->
<!--  color: #FFFFFF;-->
<!--  background: #2B2D30;-->
<!--  border-color: #73767a;-->
<!--}-->


<!--/*========== Collection Footer el-table ==========*/-->
<!--.collection-footer-table{-->
<!--  background: #1E1F22;-->
<!--  color: #c8c9cc;-->
<!--}-->

<!--.el-table{-->
<!--  height: 200px;-->
<!--  margin: 0;-->
<!--  border-radius: 4px;-->
<!--  color: #c8c9cc;-->
<!--  background: #1E1F22;-->
<!--  &#45;&#45;el-table-border-color: var(&#45;&#45;el-text-color-regular);-->
<!--  &#45;&#45;el-table-border: 1px solid var(&#45;&#45;el-table-border-color);-->
<!--  &#45;&#45;el-table-text-color: var(&#45;&#45;el-text-color-regular);-->
<!--  &#45;&#45;el-table-header-text-color: var(&#45;&#45;el-fill-color-light);-->
<!--}-->

<!--/*table 背景透明*/-->
<!--::v-deep .el-table__expanded-cell {-->
<!--  background-color: transparent !important;-->
<!--}-->
<!--::v-deep .el-table th,-->
<!--::v-deep .el-table tr,-->
<!--::v-deep .el-table td {-->
<!--  background-color: transparent;-->
<!--  border-color: #73767a;-->
<!--}-->

<!--/*鼠标悬停修改背景色*/-->
<!--.el-table ::v-deep tbody tr:hover > td{-->
<!--  background-color: #545454 !important;-->
<!--}-->


<!--</style>-->