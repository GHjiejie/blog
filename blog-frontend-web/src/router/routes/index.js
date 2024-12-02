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
    component: () => import("@/pages/Article/index.vue"),
  },
];

export default routes;
