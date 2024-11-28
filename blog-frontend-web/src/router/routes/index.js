const routes = [
  {
    path: "/",
    name: "",
    redirect: {
      name: "Blog",
    },
    component: () => import("@/layout/layout.vue"),
    children: [
      {
        path: "/blog",
        name: "Blog",
        component: () => import("@/pages/Blog/index.vue"),
      },
    ],
  },
];

export default routes;
