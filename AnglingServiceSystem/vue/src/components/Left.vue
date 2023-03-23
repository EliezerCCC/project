<template>
  <div class>
    <el-menu
      default-active="2"
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
    >
      <el-menu-item index="1" @click="ToNoticeManager()">公告管理</el-menu-item>
      <el-menu-item index="2" @click="ToInfoManager()">资讯管理</el-menu-item>
      <el-menu-item index="3" @click="ToShopManager()">商品管理</el-menu-item>
      <el-menu-item index="4" @click="ToSiteManager()">钓场管理</el-menu-item>
      <el-menu-item index="5" @click="ToRecommendManager()"
        >推荐管理</el-menu-item
      >
      <el-menu-item index="6" @click="ToForumManager()">论坛管理</el-menu-item>
      <el-menu-item index="7" @click="ToPersonManager()">用户管理</el-menu-item>
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
  },
};
</script>
