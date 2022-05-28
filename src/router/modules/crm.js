/** CRM 管理路由 */
import CrmLayout from "@/views/layout/CrmLayout";

const crmRouter = {
  path: "/crm",
  component: CrmLayout,
  redirect: "/crm/workspace",
  name: "crm",
  meta: {
    title: "个人中心"
  },
  children: [
    {
      name: "crm-customer",
      path: "customer", // 客户管理
      component: () => import("@/views/crm/customer"),
      meta: {
        title: "课程管理",
        icon: "board"
      }
    }
  ]
};

export default crmRouter;
