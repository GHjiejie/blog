import httpRequest from "@/request/index";

// 用户登录接口
export function login(params) {
  return httpRequest({
    url: "/v1/users/login",
    method: "post",
    data: {
      ...params,
      // password: bcryptEncrypt(params.password) // 加密密码(后面在做)
    },
  });
}
// 用户注册接口
export function register(params) {
  return httpRequest({
    url: "/v1/users/register",
    method: "post",
    data: {
      ...params,
      // password: bcryptEncrypt(params.password) // 加密密码(后面在做)
    },
  });
}

// 获取用户列表
export function getUserList(params) {
  return httpRequest({
    url: `/v1/users?page=${params.page}&page_size=${params.pageSize}`,
    method: "get",
  });
}
