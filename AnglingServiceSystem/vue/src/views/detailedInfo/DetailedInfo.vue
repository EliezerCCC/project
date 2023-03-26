<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-row gutter="20">
      <el-col :span="10" :offset="9">
        <h1 style="margin: auto">{{ this.info.title }}</h1>
      </el-col>
      <el-col :span="1">
        <el-button
          v-if="!is_collect"
          round
          style="width: 80px; margin-top: 5px; margin-right: 5px"
          @click="AddCollect()"
          >收藏</el-button
        >
        <el-button
          v-if="is_collect"
          type="warning"
          round
          style="width: 80px; margin-top: 5px; margin-right: 5px"
          @click="DeleteCollect()"
          >已收藏</el-button
        >
      </el-col>
      <el-col :span="1">
        <el-button
          type="info"
          round
          style="margin-left: 30px; margin-top: 5px; margin-right: 5px"
          @click="Back()"
          >返回</el-button
        >
      </el-col>
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
  inject: ["reload"],
  name: "Detail",
  data() {
    return {
      is_collect: false,
      info: {},
    };
  },
  methods: {
    Back() {
      this.$router.back();
    },
    AddCollect() {
      this.axios({
        method: "POST",
        url: this.global.apiUrl + "/addCollect",
        data: {
          info_id: this.$route.query.param.id,
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
            this.reload();
          }
        })
        .catch((err) => {});
    },
    DeleteCollect() {
      this.axios({
        method: "POST",
        url: this.global.apiUrl + "/deleteCollect",
        data: {
          info_id: this.$route.query.param.id,
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
            this.reload();
          }
        })
        .catch((err) => {});
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

    this.axios({
      method: "POST",
      url: this.global.apiUrl + "/isCollect",
      data: {
        info_id: this.$route.query.param.id,
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
          this.is_collect = res.data.is_collect;
        }
      })
      .catch((err) => {});
  },
};
</script>
