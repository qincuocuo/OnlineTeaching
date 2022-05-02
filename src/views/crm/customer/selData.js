/**
 * customer公用下拉选择数据处理
 */

import { useStore } from "vuex";
import { computed, onMounted, getCurrentInstance } from "vue";

export default function selData(fieldList = ["industryId", "area", "lifecycleId"]) {
  const store = useStore();
  const queryIndustry = () => store.dispatch("queryIndustry", {});
  const queryProvinces = () => store.dispatch("queryProvinces");
  const queryLifecycle = () => store.dispatch("queryLifecycle");
  const industryList = computed(() => store.getters.industryList);
  const provincesList = computed(() => store.getters.provincesList);
  const lifecycleList = computed(() => store.getters.lifecycleList);

  const _this = getCurrentInstance();
  onMounted(() => {
    const fApi = _this.data.formOptions.fApi;
    if (!industryList.value.length) {
      queryIndustry().then(() => {
        if (fApi.getRule(fieldList[0])) fApi.getRule(fieldList[0]).options = industryList;
      });
    }
    if (!provincesList.value.length) {
      queryProvinces().then(() => {
        if (fApi.getRule(fieldList[1])) fApi.getRule(fieldList[1]).props.options = provincesList;
      });
    }
    if (!lifecycleList.value.length) {
      queryLifecycle().then(() => {
        if (fApi.getRule(fieldList[2])) fApi.getRule(fieldList[2]).options = lifecycleList;
      });
    }
  });
  return {
    userInfo: computed(() => store.getters.userInfo),
    industryList: computed(() => store.getters.industryList),
    provincesList: computed(() => store.getters.provincesList),
    lifecycleList: computed(() => store.getters.lifecycleList),
    queryIndustry,
    queryProvinces,
    queryLifecycle
  };
}
