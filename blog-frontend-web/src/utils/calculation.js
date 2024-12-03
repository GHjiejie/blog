// 获取文件大小
export const getFileSize = (size) => {
  if (!size) return "0B";
  const formatSize =
    size < 1024
      ? Number(size) + "B"
      : size < 1024 * 1024
      ? (Number(size) / 1024).toFixed(2) + "KB"
      : size < 1024 * 1024 * 1024
      ? (Number(size) / 1024 / 1024).toFixed(2) + "MB"
      : (Number(size) / 1024 / 1024 / 1024).toFixed(2) + "MB";
  return formatSize;
};
