

<template>
  <el-container>
    <el-aside class="layout-index-aside" :style="defaultHeight" style="--wails-draggable:drag">
      <LayoutLeftAside/>
    </el-aside>

    <!--==========Layout Main==========-->
    <el-main class="layout-index-home" :style="defaultHeight">
      <router-view></router-view>
    </el-main>

    <el-aside class="layout-index-aside" :style="defaultHeight" style="--wails-draggable:drag">
      <LayoutRightAside/>
    </el-aside>
  </el-container>
</template>

<script>
import LayoutLeftAside from "./left_aside.vue";
import LayoutRightAside from "./right_aside.vue"
import LayoutHome from "../home/index.vue";

export default {
  name: "main",
  components: { LayoutLeftAside, LayoutHome, LayoutRightAside },
  data() {
    return {
      defaultHeight: {
        height: "",
      },
    };
  },
  methods: {
    getHeight() {
      this.defaultHeight.height = window.innerHeight - 90 + "px";
    },
  },

  created() {
    //页面创建时执行一次getHeight进行赋值，顺道绑定resize事件
    window.addEventListener("resize", this.getHeight);
    this.getHeight();
  },
}
</script>

<style scoped>
.layout-index-aside {
  width: 50px;
  background: #DDDDDD;
}

.layout-index-home {
  padding: 0 0 0 0;
}
</style>