<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-row gutter="20">
      <el-col :span="10" :offset="9">
        <h1 style="margin: auto">{{ this.notice.title }}</h1></el-col
      >
      <el-col :span="2">
        <el-button
          type="info"
          round
          style="margin-top: 5px; margin-right: 5px"
          @click="Back()"
          >返回</el-button
        >
      </el-col>
    </el-row>
    <el-row>
      <p
        v-html="this.notice.content"
        style="margin: auto; margin-top: 20px; width: 1000px"
      >
        {{ this.notice.content }}
      </p>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Detail",
  data() {
    return {
      notice: {},
    };
  },
  methods: {
    Back() {
      this.$router.back();
    },
  },
  created() {
    this.axios({
      method: "POST",
      url: this.global.apiUrl + "/getOneNotice",
      data: {
        id: this.$route.query.param.id,
      },
      headers: {
        Authorization: "Bearer " + sessionStorage.getItem("token"),
      },
    })
      .then((res) => {
        console.log(res);
        if (
          res.data.code == 2003 ||
          res.data.code == 2004 ||
          res.data.code == 2005
        ) {
          alert("请先完成登录!");
          this.$router.push("/");
        } else {
          sessionStorage.setItem("token", res.data.token);
          this.notice = res.data.notice;
        }
      })
      .catch((err) => {});
  },
};
</script>
