<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-container style="height: 892px; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: #545c64">
        <LeftManager></LeftManager>
      </el-aside>
      <el-main>
        <el-row>
          <el-input
            placeholder="请输入搜索内容"
            prefix-icon="el-icon-search"
            v-model="searchPlhText"
            style="width: 500px"
            @input-changed="searchInputChange"
          >
          </el-input>
          <el-button type="info" style="margin-left: 20px" @click="search()"
            >搜索</el-button
          >
        </el-row>

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

        <el-table
          :data="
            post_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )
          "
        >
          <el-table-column prop="title" label="标题" width="150">
          </el-table-column>
          <el-table-column prop="image" label="主图" width="400">
            <template slot-scope="scope">
              <img :src="GetImangePath(scope.row.id)" style="width: 50px" />
            </template>
          </el-table-column>
          <el-table-column prop="user_name" label="作者名" width="150">
          </el-table-column>
          <el-table-column prop="user_id" label="作者账号" width="150">
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
          </el-table-column> </el-table
        ><el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-size="pagesize"
          background
          layout="prev, pager, next"
          :total="1000"
        >
        </el-pagination>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  inject: ["reload"],
  data() {
    return {
      imageUrl: "",
      currentPage: 1,
      pagesize: 10,
      post_list_vis: [],
      searchPlhText: "",
      post_list: [],

      deletePost: {
        id: "",
      },

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

    DeletePostVis(row) {
      this.deletePost.id = row.id;
      this.DeletePostVisible = true;
    },
    toDetail(row) {
      this.$router.push({
        name: "detailedpost",
        query: { param: row },
      });
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
              item.user_id.toLowerCase().indexOf(_search) !== -1 ||
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
    GetImangePath: function (id) {
      return (
        this.global.apiUrl +
        "/getImage?imageName=./web/static/images/post" +
        id +
        ".jpg"
      );
    },
  },
};
</script>

<style></style>
