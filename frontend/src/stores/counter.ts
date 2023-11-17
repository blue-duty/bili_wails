import { ref } from "vue";
import { defineStore } from "pinia";

// 地址栏
export const useRouteStore = defineStore("bili_url", () => {
  const path = ref("");
  function setPath(value:string) {
    path.value = value;
  }

  return { path, setPath };
});
