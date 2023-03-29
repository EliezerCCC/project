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
          title="管理订单"
          :visible.sync="EditOrderVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <img
              :src="GetImangePath(editOrder.commodity_id)"
              style="width: 150px"
            />
          </el-row>

          <el-row style="margin-top: 15px">
            <el-tag type="info">订单号</el-tag>
            <el-tag style="margin-left: 20px">{{ editOrder.id }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">商品id</el-tag>
            <el-tag style="margin-left: 20px">{{
              editOrder.commodity_id
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">商品名</el-tag>
            <el-tag style="margin-left: 20px">{{
              editOrder.commodity_name
            }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">价格</el-tag>
            <el-tag style="margin-left: 20px">{{ editOrder.amount }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">收货信息</el-tag>
            <el-tag style="margin-left: 20px">{{ editOrder.address }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">用户id</el-tag>
            <el-tag style="margin-left: 20px">{{ editOrder.user_id }}</el-tag>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">状态</el-tag>
            <el-select
              v-model="editOrder.status"
              :placeholder="editOrder.status"
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
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="EditOrder()">确 定</el-button>
          </span>
        </el-dialog>

        <el-table
          :data="
            order_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )
          "
        >
          <el-table-column prop="id" label="订单id" width="80">
          </el-table-column>
          <el-table-column prop="commodity_name" label="商品名" width="260">
          </el-table-column>
          <el-table-column prop="commodity_id" label="商品id" width="80">
          </el-table-column>
          <el-table-column prop="image" label="图片" width="100">
            <template slot-scope="scope">
              <img
                :src="GetImangePath(scope.row.commodity_id)"
                style="width: 50px"
              />
            </template>
          </el-table-column>
          <el-table-column prop="user_id" label="用户id" width="150">
          </el-table-column>
          <el-table-column prop="address" label="收货信息" width="450">
          </el-table-column>
          <el-table-column prop="status" label="状态" width="150">
          </el-table-column>

          <el-table-column label="操作" width="300">
            <template slot-scope="scope">
              <el-button type="warning" @click="EditOrderVis(scope.row)"
                >编辑</el-button
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
      order_list_vis: [],
      searchPlhText: "",
      order_list: [],
      order: {
        id: "",
        status: "",
      },
      editOrder: {
        id: "",
        status: "",
      },

      EditOrderVisible: false,
      options: [
        {
          value: "已付款",
          label: "已付款",
        },

        {
          value: "已发货",
          label: "已发货",
        },
        {
          value: "已签收",
          label: "已签收",
        },
      ],
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllOrder",
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
          this.order_list = res.data.order_list;
          this.order_list_vis = this.order_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    Cancel() {
      this.EditOrderVisible = false;
    },

    EditOrder() {
      this.axios({
        url: this.global.apiUrl + "/updateOrder",
        data: {
          id: this.editOrder.id,
          commodity_name: this.editOrder.commodity_name,
          commodity_id: this.editOrder.commodity_id,
          amount: this.editOrder.amount,
          address: this.editOrder.address,
          status: this.editOrder.status,
          user_id: this.editOrder.user_id,
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
            (this.EditOrderVisible = false),
              (this.editOrder.commodity_name = ""),
              (this.editOrder.commodity_id = ""),
              (this.editOrder.id = ""),
              (this.editOrder.amount = ""),
              (this.editOrder.address = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditOrderVis(row) {
      this.EditOrderVisible = true;
      this.editOrder.id = row.id;
      this.editOrder.user_id = row.user_id;
      this.editOrder.status = row.status;
      this.editOrder.commodity_id = row.commodity_id;
      this.editOrder.commodity_name = row.commodity_name;
      this.editOrder.amount = row.amount;
      this.editOrder.address = row.address;
    },

    search() {
      if (this.searchPlhText == "") {
        this.order_list_vis = this.order_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.order_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (
              item.user_id.toLowerCase().indexOf(_search) !== -1 ||
              item.commodity_name.toLowerCase().indexOf(_search) !== -1
            ) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.order_list_vis = newListData;
      }
      console.log(newListData);
    },
    handleSizeChange: function (size) {
      this.pagesize = size;
      console.log(this.pagesize); //每页下拉显示数据
    },
    handleCurrentChange: function (currentPage) {
      this.currentPage = currentPage;
      console.log(this.currentPage); //点击第几页
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
