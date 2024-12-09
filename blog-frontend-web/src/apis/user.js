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

// 用户登出接口
export function logout(params) {
  console.log(params);
  return httpRequest({
    url: "/v1/users/logout",
    method: "post",
    data: {
      ...params,
    },
  });
}

// 用户注册接口
export function registerUser(params) {
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

// 用户删除接口
export function deleteUser(params) {
  return httpRequest({
    url: `/v1/users/${params.userId}`,
    method: "delete",
  });
}

// 用户密码重置接口
export function resetPassword(params) {
  return httpRequest({
    url: `/v1/users/reset_password/${params.userId}`,
    method: "put",
  });
}

// 根据用户id获取用户信息
export function getUserById(params) {
  return httpRequest({
    url: `/v1/users/${params.userId}`,
    method: "get",
  });
}

// 更新用户接口
export function updateUser(data) {
  return httpRequest({
    url: `/v1/users/update`,
    method: "post",
    data,
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
}
