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
        redirect: {
          name: "ArticleList",
        },
        component: () => import("@/pages/Article/index.vue"),
        children: [
          {
            path: "articleList",
            name: "ArticleList",
            component: () =>
              import("@/pages/Article/components/articleList.vue"),
          },
          {
            path: "articlePublish",
            name: "ArticlePublish",
            component: () =>
              import("@/pages/Article/components/articlePublish.vue"),
          },
        ],
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
      // 用户中心
      {
        path: "/personal",
        name: "Personal",
        component: () => import("@/pages/PersonalCenter/index.vue"),
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
