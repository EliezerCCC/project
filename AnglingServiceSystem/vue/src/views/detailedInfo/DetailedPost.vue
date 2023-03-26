<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-row gutter="20">
      <el-col :span="10" :offset="9">
        <h1 style="margin: auto">{{ this.post.title }}</h1>
      </el-col>
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
      <el-col :span="2" :offset="18">作者:{{ this.post.user_name }}</el-col>
    </el-row>
    <el-row>
      <p
        v-html="this.post.content"
        style="margin: auto; margin-top: 20px; width: 1000px"
      >
        {{ this.post.content }}
      </p>
    </el-row>
    <el-row>
      <h3 style="margin: auto; margin-top: 20px; width: 1000px">评论</h3>
    </el-row>
    <el-row>
      <el-table
        :data="comment_list"
        :show-header="false"
        ref="table"
        style="margin: auto; margin-top: 20px; width: 1000px"
      >
        <el-table-column>
          <template slot-scope="scope">
            <span style="font-size: 15px">{{ scope.row.user_name }}</span>
            <i class="el-icon-time" style="margin-left: 10px"></i>
            <span>{{ scope.row.create_time }}</span>
            <p style="margin-left: 2px">{{ scope.row.content }}</p>
          </template>
        </el-table-column>
      </el-table>
    </el-row>
    <el-row style="margin: auto; margin-top: 20px; width: 1000px">
      <el-input
        placeholder="请输入评论内容"
        v-model="comment"
        style="width: 800px"
      >
      </el-input>
      <el-button type="info" style="margin-left: 20px" @click="AddComment()"
        >发表评论</el-button
      >
    </el-row>
  </div>
</template>

<script>
export default {
  inject: ["reload"],
  name: "Detail",
  data() {
    return {
      post: {},
      comment: "",
      comment_list: [],
    };
  },
  methods: {
    Back() {
      this.$router.back();
    },
    AddComment() {
      if (this.comment == "") {
        alert("内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/addComment",
          data: {
            content: this.comment,
            post_id: this.$route.query.param.id,
          },
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: "Bearer " + sessionStorage.getItem("token"),
          },
        })
          .then((res) => {
            console.log(res.data);
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
              this.comment = "";
              alert("发布成功!");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
  },
  created() {
    this.axios({
      method: "POST",
      url: this.global.apiUrl + "/getOnePost",
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
          this.post = res.data.post;
        }
      })
      .catch((err) => {});

    this.axios({
      method: "POST",
      url: this.global.apiUrl + "/getComment",
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
          this.comment_list = res.data.comment_list;
          for (let i = 0; i < this.comment_list.length; i++) {
            this.comment_list[i].create_time = this.comment_list[
              i
            ].create_time.slice(0, 10);
          }
        }
      })
      .catch((err) => {});
  },
};
</script>
