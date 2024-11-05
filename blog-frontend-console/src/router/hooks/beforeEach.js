// import { ElMessage } from 'element-plus'
// import store from "@/store";
// import cache from '@/utils/cache'

async function beforeEach(to, from, next) {
  // const token = cache.getToken()

  if (to.name === "Login") {
    console.log("请求前的钩子函数，如果是登录页则直接放行");
    next();
  } else {
    next();
  }
}

export default beforeEach;
