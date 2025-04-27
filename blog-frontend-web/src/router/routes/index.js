const routes = [
  {
    path: "/",
    name: "",
    redirect: {
      name: "Article",
    },
    component: () => import("@/layout/layout.vue"),
    children: [
      {
        path: "/article",
        name: "Article",
        component: () => import("@/pages/Article/index.vue"),
      },
    ],
  },
  {
    path: "/article/:id",
    name: "ArticleDetail",
    component: () => import("@/pages/Article/components/detail.vue"),
  },
  {
    path: "/articleManage",
    name: "ArticleManage",
    redirect: {
      name: "ArticleList",
    },
    component: () => import("@/pages/Articles/index.vue"),
    children: [
      {
        path: "articleList",
        name: "ArticleList",
        component: () => import("@/pages/Articles/components/articleList.vue"),
      },
      {
        path: "articlePublish",
        name: "ArticlePublish",
        component: () =>
          import("@/pages/Articles/components/articlePublish.vue"),
      },
    ],
  },
  // {
  //   path: "articlePublish",
  //   name: "ArticlePublish",
  //   component: () => import("@/pages/Articles/components/articlePublish.vue"),
  // },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/pages/Login/index.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("@/pages/Register/index.vue"),
  },
  // 评论管理
  {
    path: "/comment",
    name: "Comment",
    component: () => import("@/pages/Comment/index.vue"),
  },
  // 标签管理
  {
    path: "/tag",
    name: "Tag",
    component: () => import("@/pages/Tag/index.vue"),
  },
  // 文件管理
  {
    path: "/file",
    name: "File",
    component: () => import("@/pages/File/index.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/pages/Login/index.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("@/pages/Register/index.vue"),
  },
];

export default routes;
