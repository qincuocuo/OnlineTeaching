const debounce = require("lodash/debounce");

let fn = null;
export default {
  mounted(el, { value }) {
    const delay = el.getAttribute("delay");
    fn = debounce(value, delay || 300);
    el.addEventListener("click", fn);
  },

  unmounted(el) {
    el.removeEventListener("click", fn);
  }
};
