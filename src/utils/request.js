import axios from "axios";
import { ElMessage } from "element-plus";

const debounce = require("lodash/debounce");

const codeMessage = {
  400: "发出的请求有错误，服务器没有进行操作。",
  401: "用户没有权限。",
  403: "用户得到授权，但是访问是被禁止的。",
  404: "发出的请求不存在，服务器没有进行操作。",
  406: "请求的格式错误。",
  410: "请求的资源被永久删除，且不会再得到的。",
  500: "服务器发生错误，请检查。",
  502: "网关错误。",
  503: "服务不可用，服务器暂时过载或维护。",
  504: "网关超时。"
};

const errorMessage = debounce(({ message, type = "error", onClose }) => {
  ElMessage({
    message,
    type,
    onClose
  });
}, 500);

const errorHandle = (status, other) => {
  const errortext = codeMessage[status] || other;
  errorMessage({
    message: errortext,
    onClose: () => {
      if ([401, 403].includes(status)) {
        window.location.href = "/#/login";
      }
    }
  });
};

// 创建axios实例
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 20000
});
service.defaults.headers.post["Content-Type"] = "application/json;charset=UTF-8";

service.interceptors.request.use(
  config => {
    const TOKEN = localStorage.getItem("crmToken") || "";
    if (TOKEN) config.headers["Authorization"] = "Bearer " + TOKEN;
    return config;
  },
  error => {
    errorMessage({ message: error });
  }
);

service.interceptors.response.use(
  response => {
    // if (400 <= response.data.code && response.data.code < 500) {
    //   window.location.href = "/#/login";
    // }
    return response.data;
  },
  error => {
    if (error.response) {
      errorHandle(error.response.status, error.response.statusText);
    } else {
      ElMessage.error(error.toString());
    }
    return Promise.reject(error);
  }
);

export default service;
