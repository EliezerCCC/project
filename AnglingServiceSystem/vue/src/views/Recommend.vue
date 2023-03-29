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
            v-for="(o, index) in recommend_list_vis.slice(
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
                height: 400px;
              "
            >
              <img
                :src="GetRecommendImangePath(o.id)"
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
                  {{ o.fish_name }}
                </h4>

                <el-button type="text" @click="DetailedRecommendVis(o)"
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
          title="推荐信息"
          :visible.sync="DetailedRecommendVisible"
          width="40%"
          :before-close="handleClose"
        >
          <el-row style="margin-top: 15px">
            <img
              :src="GetRecommendImangePath(detailedRecommend.id)"
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
          </el-row>
          <el-row style="margin-top: 15px">
            <span style="margin-left: 20px">{{
              detailedRecommend.fish_info
            }}</span>
          </el-row>
          <el-row style="margin-top: 15px">
            <el-tag type="info">推荐信息</el-tag>
          </el-row>
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
      recommend_list: [],
      recommend_list_vis: [],
      detailedRecommend: {
        id: "",
        fish_info: "",
        fish_name: "",
        recommend_info: "",
      },
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
    Cancel() {
      this.DetailedRecommendVisible = false;
    },
    DetailedRecommendVis(row) {
      this.DetailedRecommendVisible = true;
      this.detailedRecommend.fish_name = row.fish_name;
      this.detailedRecommend.id = row.id;
      this.detailedRecommend.fish_info = row.fish_info;
      this.detailedRecommend.recommend_info = row.recommend_info;
    },
    GetRecommendImangePath: function (id) {
      return (
        this.global.apiUrl +
        "/getImage?imageName=./web/static/images/recommend" +
        id +
        ".jpg"
      );
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
  },
};
</script>

<style></style>
