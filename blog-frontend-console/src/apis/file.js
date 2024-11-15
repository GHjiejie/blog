import httpRequest from "@/request/index";

// 上传文件
export function uploadFile(data) {
  console.log(data);
  return httpRequest({
    url: "/v1/files/upload",
    method: "post",
    data,
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
}

// 获取文件列表
export function getFileList(params) {
  return httpRequest({
    url: `/v1/files/getFileList?page=${params.page}&pageSize=${params.pageSize}`,
    method: "get",
  });
}

// 删除文件
export function deleteFile(params) {
  return httpRequest({
    url: `/v1/files/deleteFile/${params.fileId}`,
    method: "delete",
  });
}

// 根据文件ID获取文件
export function getFileById(params) {
  return httpRequest({
    url: `/v1/files/queryFileById/${params.fileId}`,
    method: "get",
  });
}

// 下载文件
export function downloadFile(params) {
  return httpRequest({
    url: `/v1/files/downloadFile/${params.fileId}`,
    method: "get",
    responseType: "blob",
  });
}
