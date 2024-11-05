const routes = [
  {
    path: "/",
    name: "",
    redirect: {
      name: "Article",
    },
    component: () => import("@/layout/layout.vue"),
    children: [
      // 文章管理
      {
        path: "/article",
        name: "Article",
        component: () => import("@/pages/Article/index.vue"),
      },
      // 用户管理
      {
        path: "/user",
        name: "User",
        component: () => import("@/pages/User/index.vue"),
      },
      // 评论管理
      {
        path: "/comment",
        name: "Comment",
        component: () => import("@/pages/Comment/index.vue"),
      },
      // 文件管理
      {
        path: "/file",
        name: "File",
        component: () => import("@/pages/File/index.vue"),
      },
    ],
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/pages/Login/index.vue"),
  },
];

export default routes;
