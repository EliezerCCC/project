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
          <el-button type="info" style="margin-left: 20px" @click="AddInfoVis()"
            >发布资讯</el-button
          >
        </el-row>
        <el-dialog
          title="发布资讯"
          :visible.sync="AddInfoVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">标题</el-tag>
            <el-input
              v-model="info.title"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">类型</el-tag>
            <el-select
              v-model="info.type"
              placeholder="请选择"
              style="margin-left: 20px"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-row>
          <el-row>
            <el-tag type="info">主图</el-tag>
            <el-upload
              class="avatar-uploader"
              :action="AudioAndVideoPath()"
              :show-file-list="false"
              :on-success="handleAvatarSuccess"
              :before-upload="beforeAvatarUpload"
              style="margin-left: 20px"
            >
              <img v-if="imageUrl" :src="imageUrl" class="avatar" />
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
          </el-row>
          <el-row>
            <el-tag type="info">内容</el-tag>
          </el-row>
          <el-row>
            <QuillEditor v-model="info.content"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="AddInfo()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="编辑资讯"
          :visible.sync="EditInfoVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">标题</el-tag>
            <el-input
              v-model="editInfo.title"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">类型</el-tag>
            <el-select
              v-model="editInfo.type"
              :placeholder="editInfo.type"
              style="margin-left: 20px"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-row>
          <el-row>
            <el-tag type="info">内容</el-tag>
          </el-row>
          <el-row>
            <QuillEditor v-model="editInfo.content"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="EditInfo()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="提示"
          :visible.sync="DeleteInfoVisible"
          width="30%"
          :before-close="handleClose"
        >
          <span>是否确认删除</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="DeleteInfoVisible = false">取 消</el-button>
            <el-button type="primary" @click="DeleteInfo()">确 定</el-button>
          </span>
        </el-dialog>

        <el-table
          :data="
            info_list_vis.slice(
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
          <el-table-column prop="type" label="类型" width="150">
          </el-table-column>
          <el-table-column prop="create_time" label="发布时间" width="300">
          </el-table-column>
          <el-table-column label="操作" width="300">
            <template slot-scope="scope"
              ><el-button type="success" @click="toDetail(scope.row)"
                >查看</el-button
              >
              <el-button type="warning" @click="EditInfoVis(scope.row)"
                >编辑</el-button
              >
              <el-button type="danger" @click="DeleteInfoVis(scope.row)"
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
  data() {
    return {
      imageUrl: "",
      currentPage: 1,
      pagesize: 10,
      info_list_vis: [],
      searchPlhText: "",
      info_list: [],
      info: {
        title: "",
        content: "",
        type: "",
      },
      editInfo: {
        id: "",
        title: "",
        content: "",
        type: "",
        create_time: "",
      },
      deleteInfo: {
        id: "",
      },
      AddInfoVisible: false,
      EditInfoVisible: false,
      DeleteInfoVisible: false,
      options: [
        {
          value: "钓鱼新闻",
          label: "钓鱼新闻",
        },
        {
          value: "钓鱼菜谱",
          label: "钓鱼菜谱",
        },
        {
          value: "钓鱼比赛",
          label: "钓鱼比赛",
        },
        {
          value: "钓鱼技巧",
          label: "钓鱼技巧",
        },
      ],
    };
  },
  created() {
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
    AddInfoVis() {
      this.AddInfoVisible = true;
    },
    Cancel() {
      this.AddInfoVisible = false;
      this.EditInfoVisible = false;
      (this.info.title = ""), (this.info.content = "");
    },
    AddInfo() {
      if (this.info.title == "" || this.info.content == "") {
        alert("标题或内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/addInfo",
          data: this.info,
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
              alert("发布成功!");
              (this.AddInfoVisible = false),
                (this.info.title = ""),
                (this.info.content = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
    EditInfo() {
      if (this.editInfo.title == "" || this.editInfo.content == "") {
        alert("标题或内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/updateInfo",
          data: this.editInfo,
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
              alert("修改成功!");
              (this.EditInfoVisible = false),
                (this.editInfo.title = ""),
                (this.editInfo.content = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
    DeleteInfo() {
      this.axios({
        url: this.global.apiUrl + "/deleteInfo",
        data: this.deleteInfo,
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
            alert("删除成功!");
            (this.DeleteInfoVisible = false), (this.deleteInfo.id = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditInfoVis(row) {
      this.EditInfoVisible = true;
      this.editInfo.title = row.title;
      this.editInfo.id = row.id;
      this.editInfo.content = row.content;
      this.editInfo.type = row.type;
    },
    DeleteInfoVis(row) {
      this.deleteInfo.id = row.id;
      this.DeleteInfoVisible = true;
    },
    toDetail(row) {
      this.$router.push({
        name: "detailedinfo",
        query: { param: row },
      });
    },
    search() {
      if (this.searchPlhText == "") {
        this.info_list_vis = this.info_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.info_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.title.toLowerCase().indexOf(_search) !== -1) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.info_list_vis = newListData;
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
        "/getImage?imageName=./web/static/images/info" +
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
