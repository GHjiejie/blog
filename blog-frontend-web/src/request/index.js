import axios from "axios";
const service = axios.create({
  baseURL: "",
  timeout: 5000,
  withCredentials: true,
});

// Add a request interceptor
service.interceptors.request.use(
  function (config) {
    // 在发送请求之前我们需要进行token的检查

    // console.log("interceptors request", token);

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
