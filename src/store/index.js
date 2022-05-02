import { createStore } from "vuex";
import getters from "./getters";
import user from "./modules/user";
import permission from "./modules/permission";
import crm from "./modules/crm";
import common from "./modules/common";

export default createStore({
  getters,
  modules: {
    user,
    permission,
    crm,
    common
  }
});
