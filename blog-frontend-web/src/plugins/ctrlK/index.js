// // 自定义插件,监听ctrl+k事件,作为一个插件,需要在main.js中引入并注册
// // src/keyboardPlugin.js
// export default {
//   install(app) {
//     const handleKeyDown = (event) => {
//       // 判断是否按下 Ctrl + K
//       if (event.ctrlKey && event.key === "k") {
//         event.preventDefault(); // 阻止浏览器默认操作
//         console.log("全局捕获: Ctrl + K 被按下");

//         // 调用全局定义的处理方法
//         if (app.config.globalProperties.$ctrlKPressed) {
//           app.config.globalProperties.$ctrlKPressed(); // 调用全局方法
//         }
//       }
//     };

//     // 在窗口上添加键盘事件监听
//     window.addEventListener("keydown", handleKeyDown);

//     // 确保在应用卸载时移除事件监听避免内存泄漏
//     // app.mixin({
//     //   beforeUnmount() {
//     //     window.removeEventListener("keydown", handleKeyDown);
//     //   },
//     // });
//   },
// };
