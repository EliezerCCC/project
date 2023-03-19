import Vue from "vue";
import VueRouter from "vue-router";
import anglingsite from "../views/AnglingSite.vue";
import chatroom from "../views/ChatRoom.vue";
import forum from "../views/Forum.vue";
import home from "../views/Home.vue";
import info from "../views/Info.vue";
import login from "../views/Login.vue";
import map from "../views/Map.vue";
import person from "../views/Person.vue";
import recommend from "../views/Recommend.vue";
import shop from "../views/Shop.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "login",
    component: login,
  },
  {
    path: "/home",
    name: "home",
    component: home,
  },
  {
    path: "/info",
    name: "info",
    component: info,
  },
  {
    path: "/shop",
    name: "shop",
    component: shop,
  },
  {
    path: "/anglingsite",
    name: "anglingsite",
    component: anglingsite,
  },
  {
    path: "/chatroom",
    name: "chatroom",
    component: chatroom,
  },
  {
    path: "/forum",
    name: "forum",
    component: forum,
  },
  {
    path: "/map",
    name: "map",
    component: map,
  },
  {
    path: "/person",
    name: "person",
    component: person,
  },
  {
    path: "/recommend",
    name: "recommend",
    component: recommend,
  },
];

const router = new VueRouter({
  routes,
});

export default router;
