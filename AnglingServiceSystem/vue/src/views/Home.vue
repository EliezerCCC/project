<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-row :gutter="20">
      <el-col span="5">
        <div class="block" style="margin-top: 90px; margin-left: 50px">
          <el-carousel style="width: 800px" height="500px">
            <el-carousel-item v-for="item in image_list" :key="item">
              <el-image :src="item" style="width: 100%"></el-image>
            </el-carousel-item>
          </el-carousel>
        </div>
      </el-col>
      <el-col span="11" style="margin-left: 500px">
        <h1>公告栏</h1>
        <el-table :data="notice_list" stripe height="500">
          <el-table-column prop="title" label="标题" width="650">
            <template slot-scope="scope">
              <label @click="toDetailNotice(scope.row)">{{
                scope.row.title
              }}</label>
            </template>
          </el-table-column>

          <el-table-column prop="create_time" label="日期" width="180">
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <el-row gutter="20">
      <el-col :span="3">
        <h1 style="margin-left: 70px">最新资讯</h1>
      </el-col>
      <el-col :span="2" style="margin-top: 25px">
        <el-button type="info" @click="ToInfo()">更多资讯</el-button>
      </el-col>
    </el-row>
    <el-row style="margin-left: 40px">
      <el-col :span="4.5" v-for="(o, index) in info_list.slice(0, 8)" :key="o">
        <el-card
          style="
            width: 400px;
            height: 140px;
            margin-top: 20px;
            margin-left: 30px;
          "
        >
          <el-row>
            <el-col :span="8">
              <img
                :src="GetInfoImangePath(o.id)"
                class="image"
                style="width: 100px; height: 100px"
            /></el-col>
            <el-col :span="16">
              <el-row>
                <h4
                  style="
                    width: 250px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                    -webkit-line-clamp: 2;
                  "
                  @click="toDetailInfo(o)"
                >
                  {{ o.title }}
                </h4>
              </el-row>
              <el-row style="margin-top: 20px; margin-left: 150px">
                <label>{{ o.type }}</label>
              </el-row>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      image_list: [
        require("../assets/1.jpg"),
        require("../assets/2.jpg"),
        require("../assets/3.jpg"),
      ],
      notice_list: [],
      info_list: [],
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllNotice",
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
          this.notice_list = res.data.notice_list;
          for (let i = 0; i < this.notice_list.length; i++) {
            this.notice_list[i].create_time = this.notice_list[
              i
            ].create_time.slice(0, 10);
          }
          this.notice_list_vis = this.notice_list;
        }
      })
      .catch((err) => {});

    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllInfo",
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
          this.info_list = res.data.info_list;
          for (let i = 0; i < this.info_list.length; i++) {
            this.info_list[i].create_time = this.info_list[i].create_time.slice(
              0,
              10
            );
          }
          this.info_list_vis = this.info_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    toDetailNotice(row) {
      this.$router.push({
        name: "detailednotice",
        query: { param: row },
      });
    },
    toDetailInfo(row) {
      this.$router.push({
        name: "detailedinfo",
        query: { param: row },
      });
    },
    GetInfoImangePath: function (id) {
      return (
        this.global.apiUrl +
        "/getImage?imageName=./web/static/images/info" +
        id +
        ".jpg"
      );
    },
    ToInfo() {
      this.$router.push("/info");
    },
  },
};
</script>

<style></style>
