import store from "@/store";

function getAllowList(field) {
  let list = [];
  switch (field) {
    case "teach": // 删除
      list = [1];
      break;
    case "student":
      list = [2];
      break;
    default:
      list = [1, 2];
  }
  return list;
}

/**
 * 控制某些按钮显示
 */
const has = {
  mounted(el, binding) {
    const roleId = store.getters.userInfo.role;
    const list = getAllowList(binding.value);
    if (!list.includes(roleId)) {
      el.parentNode.removeChild(el);
    }
  }
};

/*
 * 控制某些按钮显示
 * @param String
 * return Boolean
 */
const hasFun = function (field) {
  const roleId = store.getters.userInfo.role;
  const list = getAllowList(field);
  return list.includes(roleId);
};

export { has, hasFun };

// ADMIN(1, "管理员"),
// SALES_DIRECTOR(2, "销售总监"),
// SALESMAN(3, "销售"),
// CUSTOMER_SERVICE(4, "客户服务"),
// ADMIN_OFFICER(5, "行政");
