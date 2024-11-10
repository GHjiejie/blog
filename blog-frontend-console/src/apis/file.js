import httpRequest from "@/request/index";

// 获取文件列表
export function getFileList(params) {
  return httpRequest({
    url: `/v1/files/getFileList?page=${params.page}&pageSize=${params.pageSize}`,
    method: "get",
  });
}
