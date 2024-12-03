import axios from "axios";
import caches from "@/utils/cache";
const auth= caches.getToken();
console.log("auth",auth);
const service = axios.create({
  baseURL: "",
  timeout: 5000,
  withCredentials: true,
  headers:{
    "Authorization":auth
  }
});

// Add a request interceptor
service.interceptors.request.use(
  function (config) {
    
    
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
