import axios from "@/utils/request";

export function queryList(type = "post", url, params) {
  return axios({
    method: type,
    url: url,
    params: type === "get" ? params : null,
    data: type === "post" ? params : null
  });
}
