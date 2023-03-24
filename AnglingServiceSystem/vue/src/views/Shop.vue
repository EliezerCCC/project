<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-container style="height: 892px; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: #545c64">
        <el-menu
          default-active="0"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
          background-color="#545c64"
          text-color="#fff"
        >
          <el-menu-item index="1" @click="searchType('')"
            >全部商品</el-menu-item
          >
          <el-menu-item index="2" @click="searchType('钓竿')"
            >钓竿</el-menu-item
          >
          <el-menu-item index="3" @click="searchType('鱼钩')"
            >鱼钩</el-menu-item
          >
          <el-menu-item index="4" @click="searchType('鱼饵')"
            >鱼饵</el-menu-item
          >
          <el-menu-item index="4" @click="searchType('鱼线')"
            >鱼线</el-menu-item
          >
          <el-menu-item index="5" @click="searchType('鱼漂')"
            >鱼漂</el-menu-item
          >
          <el-menu-item index="6" @click="searchType('配件')"
            >配件</el-menu-item
          >
          <el-menu-item index="7" @click="searchType('其他')"
            >其他</el-menu-item
          >
        </el-menu>
      </el-aside>
      <el-main>
        <el-row>
          <el-input
            placeholder="请输入搜索内容"
            prefix-icon="el-icon-search"
            v-model="searchPlhText"
            style="width: 500px"
          >
          </el-input>
          <el-button type="info" style="margin-left: 20px" @click="search()"
            >搜索</el-button
          >
        </el-row>
        <el-row>
          <el-col
            :span="4.5"
            v-for="(o, index) in commodity_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )"
            :key="o"
          >
            <el-card
              style="
                margin-top: 20px;
                margin-left: 20px;
                width: 300px;
                height: 400px;
              "
            >
              <img
                :src="GetCommodityImangePath(o.id)"
                style="width: 260px; height: 250px"
                class="image"
              />
              <div style="padding: 14px">
                <h4
                  style="
                    width: 250px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                    -webkit-line-clamp: 2;
                  "
                >
                  {{ o.name }}
                </h4>

                <el-row>
                  <el-col :span="17">
                    <span style="font-size: larger; color: rgb(115, 167, 226)"
                      >{{ o.price }}元</span
                    ></el-col
                  >
                  <el-col :span="7">
                    <el-button type="text" @click="DetailedCommodityVis(o)"
                      >查看详情</el-button
                    ></el-col
                  >
                </el-row>
              </div></el-card
            >
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
          title="商品信息"
          :visible.sync="DetailedCommodityVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <img
              :src="GetCommodityImangePath(detailedCommodity.id)"
              style="width: 200px; margin-left: 60px"
            />
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">名称</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedCommodity.name
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">价格</el-tag>
            <el-tag style="margin-left: 20px"
              >{{ detailedCommodity.price }}元</el-tag
            >
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">数量</el-tag>
            <el-tag style="margin-left: 20px">{{
              detailedCommodity.amount
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">简介</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
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
      </el-main>
    </el-container>
  </div>
</template>
<script>
export default {
  data() {
    return {
      currentPage: 1,
      pagesize: 10,
      searchPlhText: "",
      commodity_list: [],
      commodity_list_vis: [],
      detailedCommodity: {
        id: "",
        name: "",
        introduction: "",
        amount: 0,
        price: 0.0,
        type: "",
        status: "",
      },
      DetailedCommodityVisible: false,
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
    Cancel() {
      this.DetailedCommodityVisible = false;
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
    GetCommodityImangePath: function (id) {
      return (
        this.global.apiUrl +
        "/getImage?imageName=./web/static/images/commodity" +
        id +
        ".jpg"
      );
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
    searchType(row) {
      if (row == "") {
        this.commodity_list_vis = this.commodity_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = row.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.commodity_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.type.toLowerCase().indexOf(_search) !== -1) {
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
  },
};
</script>

<style></style>
