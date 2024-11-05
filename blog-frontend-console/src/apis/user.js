import httpRequest from "@/request/index";

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
