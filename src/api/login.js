import axios from "@/utils/request";

//验证码
export function verifycode(params) {
  let url = "/api/auth/verifycode";
  return axios.get(url, { params: params });
}

//登录
export function login(params) {
  let url = "/api/auth/login";
  return axios.post(url, params);
}

//注册
export function register(params) {
  let url = "/api/auth/register";
  return axios.post(url, params);
}

//登出
export function logout(params) {
  let url = "/api/auth/logout";
  return axios.post(url, params);
}

//获取用户信息
export function getUser(params) {
  let url = "/api/v1/user";
  return axios.get(url, { params: params });
}

// 修改密码
export function changePassword(params) {
  let url = "/api/v1/user/change_password";
  return axios.post(url, params);
}
