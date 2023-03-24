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
          <el-button
            type="info"
            style="margin-left: 20px"
            @click="AddRecommendVis()"
            >添加推荐</el-button
          >
        </el-row>
        <el-dialog
          title="添加推荐"
          :visible.sync="AddRecommendVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <el-tag type="info">名称</el-tag>
            <el-input
              v-model="recommend.fish_name"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">信息</el-tag>
            <el-input
              v-model="recommend.fish_info"
              style="width: 500px; margin-left: 20px"
            ></el-input>
          </el-row>
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
            <el-tag type="info">推荐内容</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <QuillEditor v-model="recommend.recommend_info"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="AddRecommend()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="编辑推荐"
          :visible.sync="EditRecommendVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <el-tag type="info">名称</el-tag>
            <el-input
              v-model="editRecommend.fish_name"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">信息</el-tag>
            <el-input
              v-model="editRecommend.fish_info"
              style="width: 500px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">推荐信息</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <QuillEditor v-model="editRecommend.recommend_info"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="EditRecommend()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="提示"
          :visible.sync="DeleteRecommendVisible"
          width="30%"
          :before-close="handleClose"
        >
          <span>是否确认删除</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="DeleteRecommendVisible = false">取 消</el-button>
            <el-button type="primary" @click="DeleteRecommend()"
              >确 定</el-button
            >
          </span>
        </el-dialog>
        <el-dialog
          title="推荐信息"
          :visible.sync="DetailedRecommendVisible"
          width="40%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <img
              :src="GetImangePath(detailedRecommend.id)"
              style="width: 200px; margin-left: 60px"
            />
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">名称</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedRecommend.fish_name
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">信息</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedRecommend.fish_info
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px"> <el-tag type="info">推荐信息</el-tag> </el-row>
          <el-row style="margin-top: 15px">
            <p
              v-html="this.detailedRecommend.recommend_info"
              style="margin-left: 60px; margin-top: 20px; width: 600px"
            >
              {{ this.detailedRecommend.recommend_info }}
            </p>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
          </span>
        </el-dialog>

        <el-table
          :data="
            recommend_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )
          "
        >
          <el-table-column prop="fish_name" label="名称" width="150">
          </el-table-column>
          <el-table-column prop="image" label="主图" width="300">
            <template slot-scope="scope">
              <img :src="GetImangePath(scope.row.id)" style="width: 50px" />
            </template>
          </el-table-column>
          <el-table-column prop="create_time" label="发布时间" width="200">
          </el-table-column>
          <el-table-column label="操作" width="300">
            <template slot-scope="scope"
              ><el-button
                type="success"
                @click="DetailedRecommendVis(scope.row)"
                >查看</el-button
              >
              <el-button type="warning" @click="EditRecommendVis(scope.row)"
                >编辑</el-button
              >
              <el-button type="danger" @click="DeleteRecommendVis(scope.row)"
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
      recommend_list_vis: [],
      searchPlhText: "",
      recommend_list: [],
      recommend: {
        fish_info: "",
        fish_name: "",
        recommend_info: "",
      },
      editRecommend: {
        id: "",
        fish_info: "",
        fish_name: "",
        recommend_info: "",
      },
      detailedRecommend: {
        id: "",
        fish_info: "",
        fish_name: "",
        recommend_info: "",
      },
      deleteRecommend: {
        id: "",
      },
      AddRecommendVisible: false,
      EditRecommendVisible: false,
      DeleteRecommendVisible: false,
      DetailedRecommendVisible: false,
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllRecommend",
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
          this.recommend_list = res.data.recommend_list;
          for (let i = 0; i < this.recommend_list.length; i++) {
            this.recommend_list[i].create_time = this.recommend_list[
              i
            ].create_time.slice(0, 10);
          }
          this.recommend_list_vis = this.recommend_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    AddRecommendVis() {
      this.AddRecommendVisible = true;
    },
    Cancel() {
      this.AddRecommendVisible = false;
      this.EditRecommendVisible = false;
      this.DetailedRecommendVisible = false;
      (this.recommend.fish_name = ""),
        (this.recommend.recommend_info = ""),
        (this.recommend.fish_info = "");
    },
    AddRecommend() {
      this.axios({
        url: this.global.apiUrl + "/addRecommend",
        data: {
          fish_name: this.recommend.fish_name,
          fish_info: this.recommend.fish_info,
          recommend_info: this.recommend.recommend_info,
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
            alert("添加成功!");
            (this.AddRecommendVisible = false),
              (this.recommend.fish_name = ""),
              (this.recommend.fish_info = ""),
              (this.recommend.recommend_info = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditRecommend() {
      this.axios({
        url: this.global.apiUrl + "/updateRecommend",
        data: {
          id: this.editRecommend.id,
          fish_name: this.editRecommend.fish_name,
          fish_info: this.editRecommend.fish_info,
          recommend_info: this.editRecommend.recommend_info,
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
            alert("修改成功!");
            (this.EditRecommendVisible = false),
              (this.editRecommend.name = ""),
              (this.editRecommend.address = ""),
              (this.editRecommendy.introduction = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    DeleteRecommend() {
      this.axios({
        url: this.global.apiUrl + "/deleteRecommend",
        data: this.deleteRecommend,
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
            (this.DeleteRecommendVisible = false),
              (this.deleteRecommend.id = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditRecommendVis(row) {
      this.EditRecommendVisible = true;
      this.editRecommend.fish_name = row.fish_name;
      this.editRecommend.id = row.id;
      this.editRecommend.fish_info = row.fish_info;
      this.editRecommend.recommend_info = row.recommend_info;
    },
    DetailedRecommendVis(row) {
      this.DetailedRecommendVisible = true;
      this.detailedRecommend.fish_name = row.fish_name;
      this.detailedRecommend.id = row.id;
      this.detailedRecommend.fish_info = row.fish_info;
      this.detailedRecommend.recommend_info = row.recommend_info;
    },
    DeleteRecommendVis(row) {
      this.deleteRecommend.id = row.id;
      this.DeleteRecommendVisible = true;
    },
    search() {
      if (this.searchPlhText == "") {
        this.recommend_list_vis = this.recommend_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.recommend_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.fish_name.toLowerCase().indexOf(_search) !== -1) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.recommend_list_vis = newListData;
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
        "/getImage?imageName=./web/static/images/recommend" +
        id +
        ".jpg"
      );
    },
  },
};
</script>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
