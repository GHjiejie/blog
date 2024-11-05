import Crypto from "crypto-js";

const TOKEN_KEY = "authorization";
const SECKET_KEY = "CONSOLE";
const API_KEY = "apiKey";

const SessionDecode = (str) => {
  if (!str) {
    return "";
  }
  const ret = Crypto.AES.encrypt(str, SECKET_KEY).toString(Crypto.enc.Utf8);
  return ret;
};

const SessionEncode = (str) => {
  if (!str) {
    return "";
  }
  const ret = Crypto.AES.encrypt(str, SECKET_KEY).toString();
  return ret;
};

const getRealVal = (obj) => obj && obj.value;

const formValueObj = (obj) => {
  return {
    value: obj,
  };
};

const cache = {
  sessionGet(key) {
    try {
      const encodeObj = window.sessionStorage.getItem(key);

      const deCodeObj = SessionDecode(encodeObj);

      return GerRealData(JSON.parse(deCodeObj));
    } catch (error) {
      return "";
    }
  },

  sessionSet(key, value) {
    const encodeObj = SessionEncode(JSON.stringify(formValueObj(value)));
    console.log("encodeObj", encodeObj);
    window.sessionStorage.setItem(key, encodeObj);
  },

  // 获取token
  getToken() {
    return this.sessionGet(TOKEN_KEY) || "";
  },
  // 设置token
  setToken(token) {
    this.sessionSet(TOKEN_KEY, token);
  },
  // 移除token
  removeToken() {
    this.sessionRemove(TOKEN_KEY);
  },
};

export default cache;
