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
            @click="AddNoticeVis()"
            >发布公告</el-button
          >
        </el-row>
        <el-dialog
          title="发布公告"
          :visible.sync="AddNoticeVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">标题</el-tag>
            <el-input
              v-model="notice.title"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">内容</el-tag>
          </el-row>
          <el-row>
            <QuillEditor v-model="notice.content"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="AddNotice()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="编辑公告"
          :visible.sync="EditNoticeVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">标题</el-tag>
            <el-input
              v-model="editNotice.title"
              style="width: 200px; margin-left: 20px"
            ></el-input>
          </el-row>
          <el-row>
            <el-tag type="info">内容</el-tag>
          </el-row>
          <el-row>
            <QuillEditor v-model="editNotice.content"></QuillEditor>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="EditNotice()">确 定</el-button>
          </span>
        </el-dialog>
        <el-dialog
          title="提示"
          :visible.sync="DeleteNoticeVisible"
          width="30%"
          :before-close="handleClose"
        >
          <span>是否确认删除</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="DeleteNoticeVisible = false">取 消</el-button>
            <el-button type="primary" @click="DeleteNotice()">确 定</el-button>
          </span>
        </el-dialog>

        <el-table
          :data="
            notice_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )
          "
        >
          <el-table-column prop="title" label="标题" width="400">
          </el-table-column>
          <el-table-column prop="create_time" label="发布时间" width="300">
          </el-table-column>
          <el-table-column label="操作" width="300">
            <template slot-scope="scope"
              ><el-button type="success" @click="toDetail(scope.row)"
                >查看</el-button
              >
              <el-button type="warning" @click="EditNoticeVis(scope.row)"
                >编辑</el-button
              >
              <el-button type="danger" @click="DeleteNoticeVis(scope.row)"
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
      currentPage: 1,
      pagesize: 10,
      notice_list_vis: [],
      searchPlhText: "",
      notice_list: [],
      notice: {
        title: "",
        content: "",
      },
      editNotice: {
        title: "",
        content: "",
      },
      deleteNotice: {
        id: "",
      },
      AddNoticeVisible: false,
      EditNoticeVisible: false,
      DeleteNoticeVisible: false,
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllNotice",
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
          this.notice_list = res.data.notice_list;
          for (let i = 0; i < this.notice_list.length; i++) {
            this.notice_list[i].create_time = this.notice_list[
              i
            ].create_time.slice(0, 10);
          }
          this.notice_list_vis = this.notice_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {
    AddNoticeVis() {
      this.AddNoticeVisible = true;
    },
    Cancel() {
      this.AddNoticeVisible = false;
      this.EditNoticeVisible = false;
      (this.notice.title = ""), (this.notice.content = "");
    },
    AddNotice() {
      if (this.notice.title == "" || this.notice.content == "") {
        alert("标题或内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/addNotice",
          data: this.notice,
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
              (this.AddNoticeVisible = false),
                (this.notice.title = ""),
                (this.notice.content = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
    EditNotice() {
      if (this.editNotice.title == "" || this.editNotice.content == "") {
        alert("标题或内容不能为空!");
      } else {
        this.axios({
          url: this.global.apiUrl + "/updateNotice",
          data: this.editNotice,
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
              (this.EditNoticeVisible = false),
                (this.editNotice.title = ""),
                (this.editNotice.content = "");
            }
          })
          .catch((Error) => {
            console.log(Error);
          });
      }
    },
    DeleteNotice() {
      this.axios({
        url: this.global.apiUrl + "/deleteNotice",
        data: this.deleteNotice,
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
            (this.DeleteNoticeVisible = false), (this.deleteNotice.id = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditNoticeVis(row) {
      this.EditNoticeVisible = true;
      this.editNotice.title = row.title;
      this.editNotice.id = row.id;
      this.editNotice.content = row.content;
    },
    DeleteNoticeVis(row) {
      this.deleteNotice.id = row.id;
      this.DeleteNoticeVisible = true;
    },
    toDetail(row) {
      this.$router.push({
        name: "detailednotice",
        query: { param: row },
      });
    },
    search() {
      if (this.searchPlhText == "") {
        this.notice_list_vis = this.notice_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.notice_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.title.toLowerCase().indexOf(_search) !== -1) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.notice_list_vis = newListData;
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
