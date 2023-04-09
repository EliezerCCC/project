<template>
  <div class>
    <el-menu
      default-active="0"
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose"
      background-color="#545c64"
      text-color="#fff"
    >
      <el-menu-item index="1" @click="ToNoticeManager()" style="font-size: 22px"
        >公告管理</el-menu-item
      >
      <el-menu-item index="2" @click="ToInfoManager()" style="font-size: 22px"
        >资讯管理</el-menu-item
      >
      <el-menu-item index="3" @click="ToShopManager()" style="font-size: 22px"
        >商品管理</el-menu-item
      >
      <el-menu-item index="4" @click="ToSiteManager()" style="font-size: 22px"
        >钓场管理</el-menu-item
      >
      <el-menu-item
        index="5"
        @click="ToRecommendManager()"
        style="font-size: 22px"
        >推荐管理</el-menu-item
      >
      <el-menu-item index="6" @click="ToForumManager()" style="font-size: 22px"
        >论坛管理</el-menu-item
      >
      <el-menu-item index="7" @click="ToPersonManager()" style="font-size: 22px"
        >用户管理</el-menu-item
      >
      <el-menu-item index="8" @click="ToOrderManager()" style="font-size: 22px"
        >订单管理</el-menu-item
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
    ToNoticeManager() {
      this.$router.push("/noticemanager");
    },
    ToInfoManager() {
      this.$router.push("/infomanager");
    },
    ToShopManager() {
      this.$router.push("/shopmanager");
    },
    ToSiteManager() {
      this.$router.push("/sitemanager");
    },
    ToRecommendManager() {
      this.$router.push("/recommendmanager");
    },
    ToPersonManager() {
      this.$router.push("/personmanager");
    },
    ToForumManager() {
      this.$router.push("/forummanager");
    },
    ToOrderManager() {
      this.$router.push("/ordermanager");
    },
  },
};
</script>
