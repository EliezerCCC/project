import Vue from "vue";
import VueRouter from "vue-router";
import anglingsite from "../views/AnglingSite.vue";
import chatroom from "../views/ChatRoom.vue";
import detailedinfo from "../views/detailedInfo/DetailedInfo.vue";
import detailednotice from "../views/detailedInfo/DetailedNotice.vue";
import detailedpost from "../views/detailedInfo/DetailedPost.vue";
import forum from "../views/Forum.vue";
import home from "../views/Home.vue";
import info from "../views/Info.vue";
import login from "../views/Login.vue";
import forummanager from "../views/manager/ForumManager.vue";
import infomanager from "../views/manager/InfoManager.vue";
import noticemanager from "../views/manager/NoticeManager.vue";
import personmanager from "../views/manager/PersonManager.vue";
import recommendmanager from "../views/manager/RecommendManager.vue";
import shopmanager from "../views/manager/ShopManager.vue";
import sitemanager from "../views/manager/SiteManager.vue";
import map from "../views/Map.vue";
import person from "../views/Person.vue";
import recommend from "../views/Recommend.vue";
import shop from "../views/Shop.vue";
import test from "../views/test.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/test",
    name: "test",
    component: test,
  },
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
  {
    path: "/noticemanager",
    name: "noticemanager",
    component: noticemanager,
  },
  {
    path: "/infomanager",
    name: "infomanager",
    component: infomanager,
  },
  {
    path: "/shopmanager",
    name: "shopmanager",
    component: shopmanager,
  },
  {
    path: "/sitemanager",
    name: "sitemanager",
    component: sitemanager,
  },
  {
    path: "/personmanager",
    name: "personmanager",
    component: personmanager,
  },
  {
    path: "/recommendmanager",
    name: "recommendmanager",
    component: recommendmanager,
  },
  {
    path: "/forummanager",
    name: "forummanager",
    component: forummanager,
  },
  {
    path: "/detailednotice",
    name: "detailednotice",
    component: detailednotice,
  },
  {
    path: "/detailedinfo",
    name: "detailedinfo",
    component: detailedinfo,
  },
  {
    path: "/detailedpost",
    name: "detailedpost",
    component: detailedpost,
  },
];

const router = new VueRouter({
  routes,
});

export default router;
