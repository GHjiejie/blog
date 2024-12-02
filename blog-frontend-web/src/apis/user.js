import httpRequest from "@/request/index";
// 根据用户id获取用户信息
export function getUserById(params) {
  return httpRequest({
    url: `/v1/users/${params.userId}`,
    method: "get",
  });
}
