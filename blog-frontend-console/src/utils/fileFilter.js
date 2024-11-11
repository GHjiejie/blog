// 设置文件类型映射
const typeMap = new Map([
  [
    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
    "docx",
  ],
  ["application/json", "json"],
  ["text/markdown", "md"],
  ["application/pdf", "pdf"],
  ["text/plain", "txt"],

  // 引入图片
  ["image/jpeg", "jpg"],
  ["image/png", "png"],
  ["image/gif", "gif"],
  ["image/bmp", "bmp"],
  ["image/webp", "webp"],
]);

// 获取文件的类型
export function getFileType(fileType) {
  return typeMap.get(fileType);
}
// 将base64解码成为字符串（转换为特定编码）
export function decodeBase64(base64Data) {
  const decodedData = atob(base64Data);
  let decoder = new TextDecoder("utf-8");
  let utf8String = decoder.decode(
    new Uint8Array([...decodedData].map((char) => char.charCodeAt(0)))
  );
  return utf8String;
}
// base64转blob
export function base64ToBlob(base64Data, contentType) {
  const byteCharacters = atob(base64Data);
  const byteNumbers = new Array(byteCharacters.length);
  for (let i = 0; i < byteCharacters.length; i++) {
    byteNumbers[i] = byteCharacters.charCodeAt(i);
  }
  const byteArray = new Uint8Array(byteNumbers);
  const blob = new Blob([byteArray], { type: contentType });
  return blob;
}

// 获取文件地址
export function getFileUrl(data, type) {
  const blob = base64ToBlob(data, type);
  console.log("blob", blob);
  const blobUrl = URL.createObjectURL(blob);
  return blobUrl;
}
