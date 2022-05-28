import store from "@/store";

/**
 * 控制某些按钮显示
 */
const has = {
  mounted(el, binding) {
    const roleId = store.getters.userInfo.roleId;
  }
};

/*
 * 控制某些按钮显示
 * @param String
 * return Boolean
 */
const hasFun = function (field) {
  const roleId = store.getters.userInfo.roleId;
  return true;
};

export { has, hasFun };

// ADMIN(1, "管理员"),
// SALES_DIRECTOR(2, "销售总监"),
// SALESMAN(3, "销售"),
// CUSTOMER_SERVICE(4, "客户服务"),
// ADMIN_OFFICER(5, "行政");
