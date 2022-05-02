import store from "@/store";

function getAllowList(field) {
  let list = [];
  switch (field) {
    case "high":
    case "del": // 删除
      list = [1];
      break;
    case "mid":
    case "cus_edit":
    case "cus_audit": // 用户管理-审核
      list = [1, 2];
      break;
    case "cus_add": // 用户管理-审核
      list = [1, 2, 3];
      break;
    case "cus_detial_sale":
    case "cus_detial_visit":
    case "cus_query":
      list = [1, 2, 3, 4];
      break;
    default:
      list = [1, 2, 3, 4, 5];
  }
  return list;
}

/**
 * 控制某些按钮显示
 */
const has = {
  mounted(el, binding) {
    const roleId = store.getters.userInfo.roleId;
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
  const roleId = store.getters.userInfo.roleId;
  const list = getAllowList(field);
  return list.includes(roleId);
};

export { has, hasFun };

// ADMIN(1, "管理员"),
// SALES_DIRECTOR(2, "销售总监"),
// SALESMAN(3, "销售"),
// CUSTOMER_SERVICE(4, "客户服务"),
// ADMIN_OFFICER(5, "行政");
