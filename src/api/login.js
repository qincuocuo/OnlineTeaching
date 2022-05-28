import axios from "@/utils/request";

//
export function verifycode(params) {
  let url = "/auth/verifycode";
  return axios.get(url, { params: params });
}

//登录
export function login(params) {
  let url = "/auth/login";
  return axios.post(url, params);
}

//注册
export function register(params) {
  let url = "/auth/register";
  return axios.post(url, params);
}

//登出
export function logout(params) {
  let url = "/auth/logout";
  return axios.post(url, params);
}
