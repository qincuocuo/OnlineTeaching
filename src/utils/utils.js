/*
 * @Author: qiubenyang qiubenyang@mycaiwen.com
 * @Date: 2022-06-05 13:54:24
 * @LastEditors: qiubenyang qiubenyang@mycaiwen.com
 * @LastEditTime: 2022-06-11 22:34:26
 * @FilePath: /OnlineTeaching/src/utils/utils.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import JSEncrypt from "jsencrypt";

/**
 * 过滤对象中为空的属性
 * @param obj
 * @returns {*}
 */
export function filterObj(obj) {
  if (!(typeof obj == "object")) {
    return;
  }
  for (let key in obj) {
    if (
      Object.prototype.hasOwnProperty.call(obj, key) &&
      (obj[key] == null || obj[key] == undefined || obj[key] === "")
    ) {
      delete obj[key];
    }
  }
  return obj;
}

/**
 * 获取数组中指定属性值的元素
 * @returns {*}
 */
export function gainAppoint(arr, val, field = "value") {
  if (!(arr instanceof Array)) {
    return;
  }
  const res = arr.find(item => item[field] === val);
  return res || {};
}

/**
 * RSA加密
 * @param val 要加密的内容
 * @returns {*}
 */
export function jsencrypt(val, publicKey) {
  let key = publicKey;
  let jse = new JSEncrypt();
  jse.setPublicKey(key);
  let str = jse.encrypt(val);
  return str;
}

/**
 * RSA解密
 * @param val 加密的内容
 * @returns {*}
 */
export function jsdecrypt(val, publicKey) {
  let key = publicKey;
  let jse = new JSEncrypt();
  jse.setPrivateKey(key);
  let str = jse.decrypt(val);
  return str;
}

/**
 * 格式化默认时间
 * @param dataStr
 * @returns {*}
 */
export function timeStr(dataStr) {
  let date = new Date(dataStr);
  let y = date.getFullYear();
  let m = date.getMonth() + 1;
  m = m < 10 ? "0" + m : m;
  let d = date.getDate();
  d = d < 10 ? "0" + d : d;
  let h = date.getHours();
  h = h < 10 ? "0" + h : h;
  let mm = date.getMinutes();
  mm = mm < 10 ? "0" + mm : mm;
  let ss = date.getSeconds();
  ss = ss < 10 ? "0" + ss : ss;
  return y + "-" + m + "-" + d + " " + h + ":" + mm + ":" + ss;
}

/**
 * 格式化默认时间
 * @param dataStr
 * @returns {*}
 */
export function getTime(dataStr) {
  let date = new Date(dataStr);
  let h = date.getHours();
  h = h < 10 ? "0" + h : h;
  let mm = date.getMinutes();
  mm = mm < 10 ? "0" + mm : mm;

  return h + ":" + mm;
}
