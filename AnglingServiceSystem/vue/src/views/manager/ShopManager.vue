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
            @click="AddCommodityVis()"
            >发布商品</el-button
          >
        </el-row>
        <el-dialog
          title="发布商品"
          :visible.sync="AddCommodityVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">商品名</el-tag>
            <el-input
              v-model="commodity.name"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">类型</el-tag>
            <el-select
              v-model="commodity.type"
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
            <el-tag type="info">价格</el-tag>
            <el-input
              type="number"
              v-model="commodity.price"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">数量</el-tag>
            <el-input
              type="number"
              v-model="commodity.amount"
              style="width: 200px; margin-left: 20px"
            ></el-input>
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
            <el-tag type="info">简介</el-tag>
          </el-row>
          <el-row>
            <QuillEditor v-model="commodity.introduction"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="AddCommodity()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="编辑商品"
          :visible.sync="EditCommodityVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">商品名</el-tag>
            <el-input
              v-model="editCommodity.name"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">价格</el-tag>
            <el-input
              type="number"
              v-model="editCommodity.price"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">数量</el-tag>
            <el-input
              type="number"
              v-model="editCommodity.amount"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">类型</el-tag>
            <el-select
              v-model="editCommodity.type"
              :placeholder="editCommodity.type"
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
            <el-tag type="info">状态</el-tag>
            <el-select
              v-model="editCommodity.status"
              :placeholder="editCommodity.status"
              style="margin-left: 20px"
            >
              <el-option
                v-for="item in statusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-row>
          <el-row>
            <el-tag type="info">简介</el-tag>
          </el-row>
          <el-row>
            <QuillEditor v-model="editCommodity.introduction"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="EditCommodity()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="提示"
          :visible.sync="DeleteCommodityVisible"
          width="30%"
          :before-close="handleClose"
        >
          <span>是否确认删除</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="DeleteCommodityVisible = false">取 消</el-button>
            <el-button type="primary" @click="DeleteCommodity()"
              >确 定</el-button
            >
          </span>
        </el-dialog>
        <el-dialog
          title="商品信息"
          :visible.sync="DetailedCommodityVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">名称</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedCommodity.name
            }}</el-tag>
          </el-row>
          <el-row>
            <el-tag type="info">价格</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedCommodity.price
            }}</el-tag>
          </el-row>
          <el-row>
            <el-tag type="info">数量</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedCommodity.amount
            }}</el-tag>
          </el-row>
          <el-row> <el-tag type="info">简介</el-tag> </el-row>
          <el-row>
            <p
              v-html="this.detailedCommodity.introduction"
              style="margin-left: 60px; margin-top: 20px; width: 600px"
            >
              {{ this.detailedCommodity.introduction }}
            </p>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
          </span>
        </el-dialog>

        <el-table
          :data="
            commodity_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )
          "
        >
          <el-table-column prop="name" label="名称" width="150">
          </el-table-column>
          <el-table-column prop="price" label="价格" width="150">
          </el-table-column>
          <el-table-column prop="image" label="主图" width="300">
            <template slot-scope="scope">
              <img :src="GetImangePath(scope.row.id)" style="width: 50px" />
            </template>
          </el-table-column>
          <el-table-column prop="type" label="类型" width="100">
          </el-table-column>
          <el-table-column prop="create_time" label="发布时间" width="200">
          </el-table-column>
          <el-table-column prop="amount" label="数量" width="150">
          </el-table-column>
          <el-table-column prop="status" label="状态" width="100">
          </el-table-column>
          <el-table-column label="操作" width="300">
            <template slot-scope="scope"
              ><el-button
                type="success"
                @click="DetailedCommodityVis(scope.row)"
                >查看</el-button
              >
              <el-button type="warning" @click="EditCommodityVis(scope.row)"
                >编辑</el-button
              >
              <el-button type="danger" @click="DeleteCommodityVis(scope.row)"
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
      commodity_list_vis: [],
      searchPlhText: "",
      commodity_list: [],
      commodity: {
        name: "",
        introduction: "",
        amount: 0,
        price: 0.0,
        type: "",
        status: "",
      },
      editCommodity: {
        id: "",
        name: "",
        introduction: "",
        amount: 0,
        price: 0.0,
        type: "",
        status: "",
      },
      detailedCommodity: {
        id: "",
        name: "",
        introduction: "",
        amount: 0,
        price: 0.0,
        type: "",
        status: "",
      },
      deleteCommodity: {
        id: "",
      },
      AddCommodityVisible: false,
      EditCommodityVisible: false,
      DeleteCommodityVisible: false,
      DetailedCommodityVisible: false,
      options: [
        {
          value: "钓竿",
          label: "钓竿",
        },
        {
          value: "鱼钩",
          label: "鱼钩",
        },
        {
          value: "鱼饵",
          label: "鱼饵",
        },
        {
          value: "鱼线",
          label: "鱼线",
        },
        {
          value: "鱼漂",
          label: "鱼漂",
        },
        {
          value: "配件",
          label: "配件",
        },
      ],
      statusOptions: [
        {
          value: "上架",
          label: "上架",
        },
        {
          value: "下架",
          label: "下架",
        },
      ],
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllCommodity",
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
          this.commodity_list = res.data.commodity_list;
          for (let i = 0; i < this.commodity_list.length; i++) {
            this.commodity_list[i].create_time = this.commodity_list[
              i
            ].create_time.slice(0, 10);
          }
          this.commodity_list_vis = this.commodity_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    AddCommodityVis() {
      this.AddCommodityVisible = true;
    },
    Cancel() {
      this.AddCommodityVisible = false;
      this.EditCommodityVisible = false;
      this.DetailedCommodityVisible = false;
      (this.commodity.title = ""), (this.commodity.content = "");
    },
    AddCommodity() {
      if (this.commodity.title == "" || this.commodity.content == "") {
        alert("标题或内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/addCommodity",
          data: {
            name: this.commodity.name,
            introduction: this.commodity.introduction,
            amount: parseInt(this.commodity.amount),
            price: parseFloat(this.commodity.price),
            type: this.commodity.type,
            status: "",
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
              alert("发布成功!");
              (this.AddCommodityVisible = false),
                (this.commodity.title = ""),
                (this.commodity.content = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
    EditCommodity() {
      if (this.editCommodity.title == "" || this.editCommodity.content == "") {
        alert("标题或内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/updateCommodity",
          data: {
            id: this.editCommodity.id,
            name: this.editCommodity.name,
            introduction: this.editCommodity.introduction,
            amount: parseInt(this.editCommodity.amount),
            price: parseFloat(this.editCommodity.price),
            type: this.editCommodity.type,
            status: this.editCommodity.status,
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
              alert("修改成功!");
              (this.EditCommodityVisible = false),
                (this.editCommodity.name = ""),
                (this.editCommodity.price = ""),
                (this.editCommodity.status = ""),
                (this.editCommodity.id = ""),
                (this.editCommodity.type = ""),
                (this.editCommodity.amount = ""),
                (this.editCommodity.introduction = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
    DeleteCommodity() {
      this.axios({
        url: this.global.apiUrl + "/deleteCommodity",
        data: this.deleteCommodity,
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
            (this.DeleteCommodityVisible = false),
              (this.deleteCommodity.id = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditCommodityVis(row) {
      this.EditCommodityVisible = true;
      this.editCommodity.name = row.name;
      this.editCommodity.id = row.id;
      this.editCommodity.introduction = row.introduction;
      this.editCommodity.type = row.type;
      this.editCommodity.price = row.price;
      this.editCommodity.amount = row.amount;
      this.editCommodity.status = row.status;
    },
    DetailedCommodityVis(row) {
      this.DetailedCommodityVisible = true;
      this.detailedCommodity.name = row.name;
      this.detailedCommodity.id = row.id;
      this.detailedCommodity.introduction = row.introduction;
      this.detailedCommodity.type = row.type;
      this.detailedCommodity.price = row.price;
      this.detailedCommodity.amount = row.amount;
      this.detailedCommodity.status = row.status;
    },
    DeleteCommodityVis(row) {
      this.deleteCommodity.id = row.id;
      this.DeleteCommodityVisible = true;
    },
    toDetail(row) {
      this.$router.push({
        name: "detailedcommodity",
        query: { param: row },
      });
    },
    search() {
      if (this.searchPlhText == "") {
        this.commodity_list_vis = this.commodity_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.commodity_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.name.toLowerCase().indexOf(_search) !== -1) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.commodity_list_vis = newListData;
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
        "/getImage?imageName=./web/static/images/commodity" +
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
