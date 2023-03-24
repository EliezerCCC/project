<template>
  <div>
    <HeaderTop></HeaderTop>
    <el-container style="height: 892px; border: 1px solid #eee">
      <el-main>
        <el-row>
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
        </el-row>
        <el-row>
          <el-col
            :span="4.5"
            v-for="(o, index) in anglingSite_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )"
            :key="o"
          >
            <el-card
              style="
                margin-top: 20px;
                margin-left: 55px;
                width: 300px;
                height: 430px;
              "
            >
              <img
                :src="GetAnglingSiteImangePath(o.id)"
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
                <h5
                  style="
                    width: 250px;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: nowrap;
                    -webkit-line-clamp: 2;
                  "
                >
                  {{ o.address }}
                </h5>
                <el-button type="text" @click="DetailedAnglingSiteVis(o)"
                  >查看详情</el-button
                >
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
      anglingSite_list: [],
      anglingSite_list_vis: [],
      detailedAnglingSite: {
        id: "",
        name: "",
        introduction: "",
        address: "",
      },
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
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    Cancel() {
      this.DetailedAnglingSiteVisible = false;
    },
    DetailedAnglingSiteVis(row) {
      this.DetailedAnglingSiteVisible = true;
      this.detailedAnglingSite.name = row.name;
      this.detailedAnglingSite.id = row.id;
      this.detailedAnglingSite.introduction = row.introduction;
      this.detailedAnglingSite.address = row.address;
    },
    GetAnglingSiteImangePath: function (id) {
      return (
        this.global.apiUrl +
        "/getImage?imageName=./web/static/images/anglingSite" +
        id +
        ".jpg"
      );
    },
    search() {
      if (this.searchPlhText == "") {
        this.anglingSite_list_vis = this.anglingSite_list;
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
