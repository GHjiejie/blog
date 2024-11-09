import { computed } from "vue";
export default function validate(str) {
  // 密码长度8-20
  const checkLength = computed(() => {
    const l = str.value.length;
    return l >= 8 && l <= 20;
  });
  // 用户名长度
  const checkUserNameLength = computed(() => {
    const l = str.value.length;
    return l >= 5 && l <= 20;
  });
  // 只能包含字母 (a-z)、数字 (0-9) 和点 (.)
  const checkUserName = computed(() => {
    // return /^[a-zA-Z0-9][a-zA-Z0-9.]*[a-zA-Z0-9]$/.test(str.value)
    return /^[a-z0-9][a-z0-9.]*[a-z0-9]$/.test(str.value);
  });
  // 包含大写字母
  const checkUpperCase = computed(() => {
    return /[A-Z]+/g.test(str.value);
  });
  // 包含小写字母
  const checkLowerCase = computed(() => {
    return /[a-z]+/g.test(str.value);
  });
  // 包含数字
  const checkNumber = computed(() => {
    return /\d+/g.test(str.value);
  });
  // 包含特殊字符
  const checkSpecialCharacter = computed(() => {
    // return /[`~!@#$%^&*()_\-+=<>?:"{}|,.\/;'\\[\]·~！@#￥%……&*（）——\-+={}|《》？：“”【】、；‘'，。、]/im.test(
    //   str.value
    // )
    return /[~!@#$%^&*()_+]/im.test(str.value);
  });
  return {
    checkLength,
    checkUpperCase,
    checkLowerCase,
    checkNumber,
    checkSpecialCharacter,
    checkUserNameLength,
    checkUserName,
  };
}
