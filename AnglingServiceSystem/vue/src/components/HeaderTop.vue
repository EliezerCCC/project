<template>
  <div class>
    <el-menu
      :default-active="activeIndex2"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
    >
      <el-menu-item
        index="99"
        style="
          font-size: 25px;
          font-weight: bolder;
          letter-spacing: 20px;
          float: right;
        "
        @click="ToLogin()"
        >垂钓服务中心</el-menu-item
      >
      <el-menu-item index="1" @click="ToHome()">首页</el-menu-item>
      <el-menu-item index="2" @click="ToInfo()">资讯</el-menu-item>
      <el-menu-item index="3" @click="ToShop()">商城</el-menu-item>
      <el-menu-item index="4" @click="ToForum()">论坛</el-menu-item>
      <el-menu-item index="5" @click="ToAnglingSite()">钓场</el-menu-item>
      <el-menu-item index="6" @click="ToRecommend()">推荐</el-menu-item>
      <el-menu-item index="7" @click="ToMap()">地图</el-menu-item>
      <el-menu-item index="8" @click="ToChatRoom()">聊天室</el-menu-item></el-menu-item>
      <el-menu-item index="10" @click="ToNoticeManager()" v-if="vis"
        >管理界面</el-menu-item
      >
    </el-menu>
  </div>
</template>

<script>
export default {
  data() {
    return {
      vis: false,
    };
  },
  methods: {
    ToAll() {
      this.$router.push("/");
    },
    ToHome() {
      this.$router.push("/home");
    },
    ToInfo() {
      this.$router.push("/info");
    },
    ToShop() {
      this.$router.push("/shop");
    },
    ToForum() {
      this.$router.push("/forum");
    },
    ToAnglingSite() {
      this.$router.push("/anglingsite");
    },
    ToRecommend() {
      this.$router.push("/recommend");
    },
    ToMap() {
      this.$router.push("/map");
    },
    ToChatRoom() {
      this.$router.push("/chatroom");
    },
    ToPerson() {
      this.$router.push("/person");
    },
    ToNoticeManager() {
      this.$router.push("/noticemanager");
    },
    ToLogin(){
      window.open("http://localhost:8080/");
    }
  },
  created: function () {
    if (sessionStorage.getItem("identity") == "administrator") {
      this.vis = true;
    }
  },
};
</script>

<style></style>
