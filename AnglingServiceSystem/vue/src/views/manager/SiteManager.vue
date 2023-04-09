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
            @click="AddAnglingSiteVis()"
            >发布钓场</el-button
          >
        </el-row>
        <el-dialog
          title="发布钓场"
          :visible.sync="AddAnglingSiteVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <el-tag type="info" style="font-size: 20px">钓场名</el-tag>
            <el-input
              v-model="anglingSite.name"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info" style="font-size: 20px">地址</el-tag>
            <el-input
              v-model="anglingSite.address"
              style="width: 500px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info" style="font-size: 20px">主图</el-tag>
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
            <el-tag type="info" style="font-size: 20px">简介</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <QuillEditor v-model="anglingSite.introduction"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()" style="font-size: 20px"
              >取 消</el-button
            >
            <el-button
              type="primary"
              @click="AddAnglingSite()"
              style="font-size: 20px"
              >确 定</el-button
            >
          </span>
        </el-dialog>
        <el-dialog
          title="编辑钓场"
          :visible.sync="EditAnglingSiteVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <el-tag type="info" style="font-size: 20px">钓场名</el-tag>
            <el-input
              v-model="editAnglingSite.name"
              style="width: 200px; margin-left: 20px; font-size: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info" style="font-size: 20px">地址</el-tag>
            <el-input
              v-model="editAnglingSite.address"
              style="width: 500px; margin-left: 20px; font-size: 20px"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info" style="font-size: 20px">简介</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <QuillEditor v-model="editAnglingSite.introduction"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()" style="font-size: 20px"
              >取 消</el-button
            >
            <el-button
              type="primary"
              @click="EditAnglingSite()"
              style="font-size: 20px"
              >确 定</el-button
            >
          </span>
        </el-dialog>
        <el-dialog
          title="提示"
          :visible.sync="DeleteAnglingSiteVisible"
          width="30%"
          :before-close="handleClose"
        >
          <span>是否确认删除</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="DeleteAnglingSiteVisible = false"
              >取 消</el-button
            >
            <el-button type="primary" @click="DeleteAnglingSite()"
              >确 定</el-button
            >
          </span>
        </el-dialog>
        <el-dialog
          title="钓场信息"
          :visible.sync="DetailedAnglingSiteVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <el-tag type="info">名称</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedAnglingSite.name
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">地址</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedAnglingSite.address
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">简介</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <p
              v-html="this.detailedAnglingSite.introduction"
              style="margin-left: 60px; margin-top: 20px; width: 600px"
            >
              {{ this.detailedAnglingSite.introduction }}
            </p>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
          </span>
        </el-dialog>

        <el-table
          :data="
            anglingSite_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )
          "
        >
          <el-table-column prop="name" label="名称" width="150">
          </el-table-column>
          <el-table-column prop="image" label="主图" width="150">
            <template slot-scope="scope">
              <img :src="GetImangePath(scope.row.id)" style="width: 50px" />
            </template>
          </el-table-column>
          <el-table-column prop="address" label="地址" width="500">
          </el-table-column>
          <el-table-column prop="create_time" label="发布时间" width="200">
          </el-table-column>
          <el-table-column label="操作" width="300">
            <template slot-scope="scope"
              ><el-button
                type="success"
                @click="DetailedAnglingSiteVis(scope.row)"
                >查看</el-button
              >
              <el-button type="warning" @click="EditAnglingSiteVis(scope.row)"
                >编辑</el-button
              >
              <el-button type="danger" @click="DeleteAnglingSiteVis(scope.row)"
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
          :total="pagetotal"
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
      pagetotal: 1000,
      imageUrl: "",
      currentPage: 1,
      pagesize: 10,
      anglingSite_list_vis: [],
      searchPlhText: "",
      anglingSite_list: [],
      anglingSite: {
        name: "",
        introduction: "",
        address: "",
      },
      editAnglingSite: {
        id: "",
        name: "",
        introduction: "",
        address: "",
      },
      detailedAnglingSite: {
        id: "",
        name: "",
        introduction: "",
        address: "",
      },
      deleteAnglingSite: {
        id: "",
      },
      AddAnglingSiteVisible: false,
      EditAnglingSiteVisible: false,
      DeleteAnglingSiteVisible: false,
      DetailedAnglingSiteVisible: false,
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllAnglingSite",
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
          this.anglingSite_list = res.data.anglingSite_list;
          for (let i = 0; i < this.anglingSite_list.length; i++) {
            this.anglingSite_list[i].create_time = this.anglingSite_list[
              i
            ].create_time.slice(0, 10);
          }
          this.anglingSite_list_vis = this.anglingSite_list;
          this.pagetotal = this.anglingSite_list_vis.length;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    AddAnglingSiteVis() {
      this.AddAnglingSiteVisible = true;
    },
    Cancel() {
      this.AddAnglingSiteVisible = false;
      this.EditAnglingSiteVisible = false;
      this.DetailedAnglingSiteVisible = false;
      (this.anglingSite.title = ""), (this.anglingSite.content = "");
    },
    AddAnglingSite() {
      this.axios({
        url: this.global.apiUrl + "/addAnglingSite",
        data: {
          name: this.anglingSite.name,
          introduction: this.anglingSite.introduction,
          address: this.anglingSite.address,
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
            alert("发布成功!");
            (this.AddAnglingSiteVisible = false),
              (this.anglingSite.name = ""),
              (this.anglingSite.address = ""),
              (this.anglingSite.introduction = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditAnglingSite() {
      this.axios({
        url: this.global.apiUrl + "/updateAnglingSite",
        data: {
          id: this.editAnglingSite.id,
          name: this.editAnglingSite.name,
          introduction: this.editAnglingSite.introduction,
          address: this.editAnglingSite.address,
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
            (this.EditAnglingSiteVisible = false),
              (this.editAnglingSite.name = ""),
              (this.editAnglingSite.address = ""),
              (this.editAnglingSitey.introduction = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    DeleteAnglingSite() {
      this.axios({
        url: this.global.apiUrl + "/deleteAnglingSite",
        data: this.deleteAnglingSite,
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
            (this.DeleteAnglingSiteVisible = false),
              (this.deleteAnglingSite.id = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditAnglingSiteVis(row) {
      this.EditAnglingSiteVisible = true;
      this.editAnglingSite.name = row.name;
      this.editAnglingSite.id = row.id;
      this.editAnglingSite.introduction = row.introduction;
      this.editAnglingSite.address = row.address;
    },
    DetailedAnglingSiteVis(row) {
      this.DetailedAnglingSiteVisible = true;
      this.detailedAnglingSite.name = row.name;
      this.detailedAnglingSite.id = row.id;
      this.detailedAnglingSite.introduction = row.introduction;
      this.detailedAnglingSite.address = row.address;
    },
    DeleteAnglingSiteVis(row) {
      this.deleteAnglingSite.id = row.id;
      this.DeleteAnglingSiteVisible = true;
    },
    search() {
      if (this.searchPlhText == "") {
        this.anglingSite_list_vis = this.anglingSite_list;
        this.pagetotal = this.anglingSite_list_vis.length;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.anglingSite_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (
              item.name.toLowerCase().indexOf(_search) !== -1 ||
              item.address.toLowerCase().indexOf(_search) !== -1
            ) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.anglingSite_list_vis = newListData;
        this.pagetotal = this.anglingSite_list_vis.length;
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
        "/getImage?imageName=./web/static/images/anglingSite" +
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
