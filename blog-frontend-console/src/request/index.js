import axios from "axios";
import router from "@/router";
import cache from "@/utils/cache";

const token = cache.getToken();
console.log("token", token);

const service = axios.create({
  baseURL: "",
  timeout: 5000,
  withCredentials: true,
  headers: {
    Authorization: `${token}`,
  },
});

// Add a request interceptor
service.interceptors.request.use(
  function (config) {
    // 在发送请求之前我们需要进行token的检查
    const token = cache.getToken();
    console.log("interceptors request", token);
    if (!token) {
      router.push("/login");
    } else {
      config.headers.Authorization = token;
    }
    // Do something before request is sent
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
    return Promise.reject(error);
  }
);
export default service;
