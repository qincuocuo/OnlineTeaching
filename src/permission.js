import router from "@/router";
import "nprogress/nprogress.css";
import store from "@/store";

// 保存左侧菜单激活状态
function saveNavState(activePath) {
  store.commit("activePathChange", activePath);
}

router.beforeEach((to, from, next) => {
  //页面标题
  window.document.title = to.meta.title === undefined ? "CRM" : to.meta.title;
  saveNavState(to.path);
  next();
});

router.afterEach(() => {});

router.onError(error => {
  console.log("error.message", error.message);
  const pattern = /Loading chunk (\d)+ failed/g;
  const isChunkLoadFailed = error.message.match(pattern);
  console.log("router", router);
  if (isChunkLoadFailed) {
    window.location.reload();
  }
});
