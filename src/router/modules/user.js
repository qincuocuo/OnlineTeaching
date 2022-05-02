/** 个人中心管理路由 */
import UserLayout from "@/views/layout/UserLayout";

const userRouter = {
  path: "/user",
  component: UserLayout,
  redirect: "/user/index",
  name: "user",
  meta: {
    title: "个人中心"
  },
  children: [
    {
      path: "index",
      component: () => import("@/views/user/index")
    }
  ]
};

export default userRouter;
