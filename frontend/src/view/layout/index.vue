<template>
    <el-container>
      <!--==========Layout Header==========-->
      <el-header class="layout-index-header"  style="--wails-draggable:drag">
        <LayoutHeader />
      </el-header>

      <!--==========Layout Main==========-->
      <el-main class="layout-index-main" :style="defaultMainSize">
        <LayoutMain />
      </el-main>

      <!--==========Layout Footer==========-->
      <el-footer class="layout-index-footer" style="--wails-draggable:drag">
        <LayoutFooter />
      </el-footer>
    </el-container>


</template>

<script>
import LayoutHeader from "./header.vue";
import LayoutFooter from "./footer.vue";
import LayoutMain from "./main.vue";

export default {
  name: "Layout",
  components: { LayoutHeader, LayoutMain, LayoutFooter},
  data() {
    return {
      isCollapse: false,
      defaultMainSize: {
        height: "",
      },
    };
  },
  methods: {
    //定义方法，获取高度减去头尾
    getWinHeight() {
      this.defaultMainSize.height = window.innerHeight - 90 + "px";
    },
  },

  created() {
    this.getWinHeight();
    window.addEventListener("resize", this.getWinHeight);

  },
};
</script>

<style lang="scss" scoped>
.layout-index-header{
  padding: 0;
  background: #BBBBBB;
  height: 50px;
}

.layout-index-footer {
  background: #CCCCCC;
  height: 40px;
  border-bottom: none;
}

.layout-index-main {
  padding: 0;
  background: #FFFFFF;
  color: #c8c9cc;
}

</style>
