import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "collection",
  },
  {
    path: "/collection",
    name: "collection",
    component: () => import("../view/collection/index.vue"),
  },
  {
    path: "/param",
    name: "param",
    component: () => import("../view/param/index.vue"),
  },
  {
    path: "/data",
    name: "data",
    component: () => import("../view/data/index.vue"),
  },
  {
    path: "/inspector",
    name: "inspector",
    component: () => import("../view/inspector/index.vue"),
  },
  {
    path: "/setting",
    name: "setting",
    component: () => import("../view/setting/index.vue"),
  },
  {
    path: "/diagnosis",
    name: "diagnosis",
    component: () => import("../view/diagnosis/index.vue"),
  },
  {
    path: "/users",
    name: "users",
    component: () => import("../view/users/index.vue"),
  },
  {
    path: "/help",
    name: "help",
    component: () => import("../view/help/index.vue"),
  },

];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
