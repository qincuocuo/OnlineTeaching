import axios from "@/utils/request";

//申请口令
export function applyForPassword(params) {
  let url = "/crm/open/api/login/send_email_code";
  return axios.get(url, { params: params });
}

//登录
export function login(params) {
  let url = "/crm/open/api/login/login_email";
  return axios.post(url, params);
}
