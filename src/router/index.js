import { createRouter, createWebHashHistory } from "vue-router";
import crmRouter from "./modules/crm";
import userRouter from "./modules/user";

const routes = [
  {
    path: "/",
    redirect: "/login"
  },
  {
    path: "/login",
    name: "login",
    component: () => import("@/views/login/index"),
    meta: {
      title: "阳光小学线上教学系统"
    }
  },
  {
    path: "/404",
    component: () => import("@/views/404")
  },
  userRouter,
  crmRouter,
  {
    path: "/:catchAll(.*)",
    redirect: "/404"
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
