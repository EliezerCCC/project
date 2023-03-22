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
      <el-menu-item index="2" @click="ToNoticeManager()">资讯管理</el-menu-item>
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
  },
};
</script>
