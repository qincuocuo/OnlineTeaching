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
    params: params,
    data: null
  })
}