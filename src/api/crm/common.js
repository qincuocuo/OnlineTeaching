import axios from "@/utils/request";

export function queryList(type = "get", url, params) {
  return axios({
    method: type,
    url: url,
    params: type === "get" ? params : null,
    data: type === "post" ? params : null
  });
}
// 查询班级
export function getClass(params){
  let url = "/api/v1/course/class";
  return axios({
    method: 'get',
    url: url,
    params: params,
    data: null
  })
}