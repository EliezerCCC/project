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
          <el-button
            type="info"
            style="margin-left: 20px"
            @click="MyAddressVisible = true"
            >我的地址</el-button
          >
          <el-button type="info" style="margin-left: 20px" @click=""
            >我的订单</el-button
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
            <el-button type="primary" @click="Buy()">购买</el-button>
          </span>
        </el-dialog>

        <el-dialog
          title="我的地址"
          :visible.sync="MyAddressVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-table :data="address_list" stripe height="500">
            <el-table-column prop="province" label="省份" width="100">
              <template slot-scope="scope">
                <label>{{ CodeToText[parseInt(scope.row.province)] }}</label>
              </template>
            </el-table-column>
            <el-table-column prop="city" label="城市" width="100">
              <template slot-scope="scope">
                <label>{{ CodeToText[parseInt(scope.row.city)] }}</label>
              </template>
            </el-table-column>
            <el-table-column prop="area" label="区域" width="100">
              <template slot-scope="scope">
                <label>{{ CodeToText[parseInt(scope.row.area)] }}</label>
              </template>
            </el-table-column>
            <el-table-column prop="detail" label="详细地址" width="350">
              <template slot-scope="scope">
                <label>{{ scope.row.detail }}</label>
              </template>
            </el-table-column>

            <el-table-column prop="name" label="姓名" width="180">
            </el-table-column>
            <el-table-column prop="phone" label="电话" width="180">
            </el-table-column>
            <el-table-column label="操作" width="100">
              <template slot-scope="scope">
                <el-button type="danger" @click="DeleteAddressVis(scope.row)"
                  >删除</el-button
                >
              </template>
            </el-table-column>
          </el-table>
          <span slot="footer" class="dialog-footer">
            <el-button type="primary" @click="AddAddressVisible = true"
              >添加地址</el-button
            >
            <el-button @click="MyAddressVisible = false">取 消</el-button>
          </span>
        </el-dialog>

        <el-dialog
          title="提示"
          :visible.sync="DeleteAddressVisible"
          width="30%"
          :before-close="handleClose"
        >
          <span>是否确认删除地址</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="DeleteAddressVisible = false">取 消</el-button>
            <el-button type="primary" @click="DeleteAddress()">确 定</el-button>
          </span>
        </el-dialog>

        <el-dialog
          title="添加地址"
          :visible.sync="AddAddressVisible"
          width="30%"
          :before-close="handleClose"
        >
          <el-row>
            <el-cascader
              size="large"
              :options="options"
              v-model="selectedOptions"
              @change="handleChange"
            >
            </el-cascader>
          </el-row>
          <el-row style="margin-top: 15px">
            <label>详细地址</label>
            <el-input
              style="margin-top: 15px"
              v-model="detailAddress"
              placeholder="请输入内容"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <label>收货人姓名</label>
            <el-input
              style="margin-top: 15px"
              v-model="nameAddress"
              placeholder="请输入内容"
            ></el-input>
          </el-row>
          <el-row style="margin-top: 15px">
            <label>电话</label>
            <el-input
              style="margin-top: 15px"
              v-model="phoneAddress"
              placeholder="请输入内容"
            ></el-input>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button type="primary" @click="AddAddress()">添加</el-button>
            <el-button @click="AddAddressVisible = false">取 消</el-button>
          </span>
        </el-dialog>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { regionData, CodeToText } from "element-china-area-data";
export default {
  inject: ["reload"],
  data() {
    return {
      CodeToText,
      nameAddress: "",
      phoneAddress: "",
      detailAddress: "",
      options: regionData,
      selectedOptions: [],
      currentPage: 1,
      pagesize: 10,
      searchPlhText: "",
      commodity_list: [],
      commodity_list_vis: [],
      address_list: [],
      detailedCommodity: {
        id: "",
        name: "",
        introduction: "",
        amount: 0,
        price: 0.0,
        type: "",
        status: "",
      },
      deleteAddress: { id: "" },
      DeleteAddressVisible: false,
      AddAddressVisible: false,
      MyAddressVisible: false,
      DetailedCommodityVisible: false,
      pay_list: [],
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

          let newListData = []; // 用于存放搜索出来数据的新数组
          //filter 过滤数组
          this.commodity_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.status.toLowerCase().indexOf("上架") !== -1) {
              newListData.push(item);
            }
          });

          this.commodity_list = newListData;
          this.commodity_list_vis = this.commodity_list;
        }
      })
      .catch((err) => {});

    this.axios({
      method: "POST",
      url: this.global.apiUrl + "/getAddress",
      data: {
        id: sessionStorage.getItem("user_id"),
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
          this.address_list = res.data.address_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    Buy() {
      this.$confirm("确认购买此商品", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.pay_list = [];
          this.pay_list.push(this.detailedCommodity);
          console.log(this.pay_list);
          this.axios({
            method: "POST",
            data: { pay_list: this.pay_list },
            url: this.global.apiUrl + "/pay",
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
                console.log(res.data.pay_url);
                window.location.href = res.data.pay_url;
              }
            })
            .catch((err) => {});
        })
        .catch(() => {});
    },

    DeleteAddress() {
      this.axios({
        url: this.global.apiUrl + "/deleteAddress",
        data: this.deleteAddress,
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
            (this.DeleteAddressVisible = false), (this.deleteAddress.id = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },

    DeleteAddressVis(row) {
      this.deleteAddress.id = row.id;
      this.DeleteAddressVisible = true;
    },
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
    AddAddress() {
      if (
        this.detailAddress == "" ||
        this.nameAddress == "" ||
        this.phoneAddress == "" ||
        this.selectedOptions.length != 3
      ) {
        alert("内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/addAddress",
          data: {
            name: this.nameAddress,
            phone: this.phoneAddress,
            detail: this.detailAddress,
            province: this.selectedOptions[0],
            city: this.selectedOptions[1],
            area: this.selectedOptions[2],
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
              (this.AddAddressVisible = false),
                (this.nameAddress = ""),
                (this.detailAddress = ""),
                (this.phoneAddress = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
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
    handleChange(value) {
      console.log(value);
    },
  },
};
</script>

<style></style>
