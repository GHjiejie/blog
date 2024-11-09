/*
需要注意的是，我们个工具库的作用是将服务端生成的token进行加密存储到客户端的sessionStrorage中,如果我们不加密直接存储的话，别人可能根据根据token进行枚举解密，导致我们的token不安全，所以我们需要对token进行加密存储，这里我们使用了crypto-js这个库进行加密存储，这样我们就可以保证token的安全性。
然后在发起请求的时候，外面需要用同样的加密算法对存储在客户端的token进行解密，然后再进行请求，这样我们就可以保证token的安全性。

 */

import Crypto from "crypto-js";

const TOKEN_KEY = "authorization";
const SECKET_KEY = "CONSOLE";
const API_KEY = "apiKey";

const SessionDecode = (str) => {
  if (!str) {
    return "";
  }
  const ret = Crypto.AES.decrypt(str, SECKET_KEY).toString(Crypto.enc.Utf8);
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
      return getRealVal(JSON.parse(deCodeObj));
    } catch (error) {
      return "";
    }
  },

  sessionSet(key, value) {
    const encodeObj = SessionEncode(JSON.stringify(formValueObj(value)));
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
