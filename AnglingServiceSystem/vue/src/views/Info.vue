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
          <el-menu-item
            index="1"
            @click="searchType('')"
            style="font-size: 20px"
            >全部资讯</el-menu-item
          >
          <el-menu-item
            index="2"
            @click="searchType('钓鱼新闻')"
            style="font-size: 20px"
            >钓鱼新闻</el-menu-item
          >
          <el-menu-item
            index="3"
            @click="searchType('钓鱼技巧')"
            style="font-size: 20px"
            >钓鱼技巧</el-menu-item
          >
          <el-menu-item
            index="4"
            @click="searchType('钓鱼比赛')"
            style="font-size: 20px"
            >钓鱼比赛</el-menu-item
          >
          <el-menu-item
            index="4"
            @click="searchType('钓鱼菜谱')"
            style="font-size: 20px"
            >钓鱼菜谱</el-menu-item
          >
          <el-menu-item
            index="5"
            @click="personCollect()"
            style="font-size: 20px"
            >我的收藏</el-menu-item
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
          <el-button
            type="info"
            style="margin-left: 20px; font-size: 18px"
            @click="search()"
            >搜索</el-button
          >
        </el-row>
        <el-row>
          <el-col
            :span="4.5"
            v-for="(o, index) in info_list_vis.slice(
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
                    :src="GetInfoImangePath(o.id)"
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
                      @click="toDetailInfo(o)"
                    >
                      {{ o.title }}
                    </h4>
                  </el-row>
                  <el-row style="margin-top: 20px; margin-left: 150px">
                    <label>{{ o.type }}</label>
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
            :total="pagetotal"
          >
          </el-pagination
        ></el-row>
      </el-main>
    </el-container>
  </div>
</template>
<script>
export default {
  data() {
    return {
      pagetotal: 1000,
      currentPage: 1,
      pagesize: 12,
      searchPlhText: "",
      info_list: [],
      info_list_vis: [],
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
          this.pagetotal = this.info_list_vis.length;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    personCollect() {
      this.axios({
        method: "GET",
        url: this.global.apiUrl + "/personCollect",
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
            this.info_list_vis = res.data.info_list;
          }
        })
        .catch((err) => {});
    },
    toDetailInfo(row) {
      this.$router.push({
        name: "detailedinfo",
        query: { param: row },
      });
    },
    GetInfoImangePath: function (id) {
      return (
        this.global.apiUrl +
        "/getImage?imageName=./web/static/images/info" +
        id +
        ".jpg"
      );
    },
    search() {
      if (this.searchPlhText == "") {
        this.info_list_vis = this.info_list;
        this.pagetotal = this.info_list_vis.length;
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
        this.pagetotal = this.info_list_vis.length;
      }
    },
    searchType(row) {
      if (row == "") {
        this.info_list_vis = this.info_list;
        this.pagetotal = this.info_list_vis.length;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = row.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.info_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.type.toLowerCase().indexOf(_search) !== -1) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据

        this.info_list_vis = newListData;
        this.pagetotal = this.info_list_vis.length;
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
