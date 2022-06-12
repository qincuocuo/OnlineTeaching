/*
 * @Author: qiubenyang qiubenyang@mycaiwen.com
 * @Date: 2022-06-05 13:54:24
 * @LastEditors: qiubenyang qiubenyang@mycaiwen.com
 * @LastEditTime: 2022-06-12 21:00:54
 * @FilePath: /OnlineTeaching/src/api/login.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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

// 修改密码
export function password(params) {
  let url = "/api/v1/user/password";
  return axios.post(url, params);
}
