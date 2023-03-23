<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-row>
      <h1 style="margin: auto">{{ this.info.title }}</h1>
      <el-button
        type="info"
        round
        style="margin-top: 5px; margin-right: 5px"
        @click="Back()"
        >返回</el-button
      >
    </el-row>
    <el-row>
      <p
        v-html="this.info.content"
        style="margin: auto; margin-top: 20px; width: 1000px"
      >
        {{ this.info.content }}
      </p>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Detail",
  data() {
    return {
      info: {},
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
      url: this.global.apiUrl + "/getOneInfo",
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
          this.info = res.data.info;
        }
      })
      .catch((err) => {});
  },
};
</script>
