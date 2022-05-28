import axios from "@/utils/request";

//校验权限是否正常
export function verifyHeat(params) {
  const url = "/crm/admin/api/permission/verify/heat";
  return axios.get(url, { params: params });
}
