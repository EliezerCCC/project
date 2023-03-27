<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-container style="width: 1904px; height: 892px; border: 1px solid #eee">
      <el-main>
        <el-row>
          <el-button type="info" style="margin-left: 5px" @click="AddPostVis()"
            >发布帖子</el-button
          >
          <el-input
            placeholder="请输入搜索内容"
            prefix-icon="el-icon-search"
            v-model="searchPlhText"
            style="width: 500px; margin-left: 55px"
          >
          </el-input>
          <el-button type="info" style="margin-left: 20px" @click="search()"
            >搜索</el-button
          >
          <el-button
            type="info"
            style="margin-left: 20px"
            @click="MyPostVisible = true"
            >我的帖子</el-button
          >
        </el-row>

        <el-row style="margin-left: 40px">
          <el-col
            :span="4.5"
            v-for="(o, index) in post_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )"
            :key="o"
          >
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
                    :src="GetPostImangePath(o.id)"
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
                      @click="toDetailPost(o)"
                    >
                      {{ o.title }}
                    </h4>
                  </el-row>
                  <el-row style="margin-top: 20px; margin-left: 150px">
                    <label>{{ o.user_name }}</label>
                  </el-row>
                </el-col>
              </el-row>
            </el-card>
          </el-col>
        </el-row>

        <el-row style="margin-top: 15px; margin-left: 50px">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-size="pagesize"
            background
            layout="prev, pager, next"
            :total="1000"
          >
          </el-pagination
        ></el-row>

        <el-dialog
          title="发布帖子"
          :visible.sync="AddPostVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <el-tag type="info">标题</el-tag>
            <el-input
              v-model="post.title"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px"> </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">主图</el-tag>
            <el-upload
              class="avatar-uploader"
              :action="AudioAndVideoPath()"
              :show-file-list="false"
              :on-success="handleAvatarSuccess"
              :before-upload="beforeAvatarUpload"
              style="margin-left: 20px; margin-top: 15px"
            >
              <img v-if="imageUrl" :src="imageUrl" class="avatar" />
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">内容</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <QuillEditor v-model="post.content"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="AddPost()">确 定</el-button>
          </span>
        </el-dialog>

        <el-dialog
          title="我的帖子"
          :visible.sync="MyPostVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-table :data="my_post_list">
            <el-table-column prop="title" label="标题" width="150">
            </el-table-column>
            <el-table-column prop="image" label="主图" width="200">
              <template slot-scope="scope">
                <img
                  :src="GetPostImangePath(scope.row.id)"
                  style="width: 50px"
                />
              </template>
            </el-table-column>
            <el-table-column prop="create_time" label="发布时间" width="300">
            </el-table-column>
            <el-table-column label="操作" width="300">
              <template slot-scope="scope"
                ><el-button type="success" @click="toDetail(scope.row)"
                  >查看</el-button
                >
                <el-button type="danger" @click="DeletePostVis(scope.row)"
                  >删除</el-button
                >
              </template>
            </el-table-column>
          </el-table>
          <span slot="footer" class="dialog-footer">
            <el-button @click="MyPostVisible = false">取 消</el-button>
          </span>
        </el-dialog>

        <el-dialog
          title="提示"
          :visible.sync="DeletePostVisible"
          width="30%"
          :before-close="handleClose"
        >
          <span>是否确认删除</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="DeletePostVisible = false">取 消</el-button>
            <el-button type="primary" @click="DeletePost()">确 定</el-button>
          </span>
        </el-dialog>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  inject: ["reload"],
  data() {
    return {
      MyPostVisible: false,
      currentPage: 1,
      pagesize: 16,
      searchPlhText: "",
      post_list: [],
      post_list_vis: [],
      my_post_list: [],
      post: {
        title: "",
        content: "",
      },
      deletePost: {
        id: "",
      },
      AddPostVisible: false,
      DeletePostVisible: false,
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllPost",
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
          this.post_list = res.data.post_list;
          for (let i = 0; i < this.post_list.length; i++) {
            this.post_list[i].create_time = this.post_list[i].create_time.slice(
              0,
              10
            );
            if (
              this.post_list[i].user_id == sessionStorage.getItem("user_id")
            ) {
              this.my_post_list.push(this.post_list[i]);
            }
          }
          this.post_list_vis = this.post_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    DeletePost() {
      this.axios({
        url: this.global.apiUrl + "/deletePost",
        data: this.deletePost,
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
            alert("删除成功!");
            (this.DeletePostVisible = false), (this.deletePost.id = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    Cancel() {
      this.AddPostVisible = false;
    },
    AddPostVis() {
      this.AddPostVisible = true;
    },
    AddPost() {
      if (this.post.title == "" || this.post.content == "") {
        alert("标题或内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/addPost",
          data: this.post,
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
              alert("发布成功!");
              (this.AddPostVisible = false),
                (this.post.title = ""),
                (this.post.content = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
    search() {
      if (this.searchPlhText == "") {
        this.post_list_vis = this.post_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.post_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (
              item.title.toLowerCase().indexOf(_search) !== -1 ||
              item.user_name.toLowerCase().indexOf(_search) !== -1
            ) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.post_list_vis = newListData;
      }
    },
    handleSizeChange: function (size) {
      this.pagesize = size;
      console.log(this.pagesize); //每页下拉显示数据
    },
    handleCurrentChange: function (currentPage) {
      this.currentPage = currentPage;
      console.log(this.currentPage); //点击第几页
    },
    handleAvatarSuccess(res, file) {
      this.imageUrl = URL.createObjectURL(file.raw);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG) {
        this.$message.error("上传图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传图片大小不能超过 2MB!");
      }
      return isJPG && isLt2M;
    },
    AudioAndVideoPath: function () {
      return this.global.apiUrl + "/upload";
    },
    toDetail(row) {
      this.$router.push({
        name: "detailedpost",
        query: { param: row },
      });
    },
    toDetailPost(row) {
      this.$router.push({
        name: "detailedpost",
        query: { param: row },
      });
    },
    DeletePostVis(row) {
      this.deletePost.id = row.id;
      this.DeletePostVisible = true;
    },
    GetPostImangePath: function (id) {
      return (
        this.global.apiUrl +
        "/getImage?imageName=./web/static/images/post" +
        id +
        ".jpg"
      );
    },
    handleSizeChange: function (size) {
      this.pagesize = size;
      console.log(this.pagesize); //每页下拉显示数据
    },
    handleCurrentChange: function (currentPage) {
      this.currentPage = currentPage;
      console.log(this.currentPage); //点击第几页
    },
  },
};
</script>

<style></style>
