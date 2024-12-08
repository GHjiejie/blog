import axios from "axios";
import caches from "@/utils/cache";
const auth = caches.getToken();
const service = axios.create({
  baseURL: "",
  timeout: 5000,
  withCredentials: true,
  headers: {
    Authorization: auth,
  },
});

// Add a request interceptor
service.interceptors.request.use(
  function (config) {
    // 判断是否有token
    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

// Add a response interceptor
service.interceptors.response.use(
  function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data

    return response;
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    // console.log("error", error);
    // 这里面是有问题的,就是token会使用上一次的缓存token,所以会导致token失效,但是我还没有找到解决办法,先这样处理吧
    if (error.response.status === 401) {
      // window.location.reload();
    }

    return Promise.reject(error);
  }
);
export default service;
