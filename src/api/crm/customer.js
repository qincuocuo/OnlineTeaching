import axios from "@/utils/request";
export function queryLifecycle(params) {
  const url = "/crm/admin/api/sys/query_lifecycle";
  return axios.post(url, params);
}

export function queryCustomerSource(params) {
  const url = "/crm/admin/api/sys/query_customer_source";
  return axios.post(url, params);
}

export function querySalesman(params) {
  const url = "/crm/admin/api/sys/query_salesman";
  return axios.post(url, params);
}

export function addCustomer(params) {
  const url = "/crm/auth/api/customer/add_customer";
  return axios.post(url, params);
}

export function customerRemove(params) {
  const url = "/crm/auth/api/customer/customer_remove";
  return axios.post(url, params);
}

export function customerAudit(params) {
  const url = "/crm/auth/api/customer/customer_audit";
  return axios.post(url, params);
}

export function updateCustomer(params) {
  const url = "/crm/auth/api/customer/update_customer";
  return axios.post(url, params);
}
//
export function queryCustomerDetail(params) {
  const url = "/crm/auth/api/customer/query_customer";
  return axios.post(url, params);
}

// 新增课程
export function addCourse(params){
  const url = "/api/v1/course";
  return axios.post(url, params);
}
// 编辑课程
export function updateCourse(params){
  const url = "/api/v1/course/update";
  return axios.post(url, params);
}
// 删除课程
export function deleteCourse(params){
  const url = "/api/v1/course";
  return axios({
    method: 'delete',
    url: url,
    params: null,
    data: params
  })
}

// 新增学习内容
export function addContent(params){
  const url = "/api/v1/learning_content";
  return axios.post(url, params);
}

// 查询学习情况
export function getLearningDetail(params){
  let url = "/api/v1/learning_content/result";
  return axios({
    method: 'get',
    url: url,
    params: params,
    data: null
  })
}
// 查询签到情况
export function getQianDaoDetail(params){
  let url = "/api/v1/register/result";
  return axios({
    method: 'get',
    url: url,
    params: params,
    data: null
  })
}

// 新增签到任务
export function addQianDao(params){
  const url = "/api/v1/register";
  return axios.post(url, params);
}
// 创建讨论话题
export function addtalk(params){
  const url = "/api/v1/talk";
  return axios.post(url, params);
}
// 创建课后练习
export function addExercises(params){
  const url = "/api/v1/exercises";
  return axios.post(url, params);
}