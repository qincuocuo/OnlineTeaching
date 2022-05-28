import axios from "@/utils/request";

// 新增课程
export function addCourse(params) {
  const url = "v1/course";
  return axios.post(url, params);
}
// 编辑课程
export function updateCourse(params) {
  const url = "v1/course/update";
  return axios.post(url, params);
}
// 删除课程
export function deleteCourse(params) {
  const url = "v1/course";
  return axios.delete(url, params);
}
