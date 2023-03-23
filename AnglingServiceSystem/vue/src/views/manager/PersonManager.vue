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
          title="管理用户"
          :visible.sync="EditUserVisible"
          width="60%"
          :before-close="handleClose"
        >
          <el-row>
            <el-tag type="info">账号</el-tag>
              <el-tag>{{editUser.id}}</el-tag>
          </el-row>
          <el-row>
            <el-tag type="info">昵称</el-tag></el-tag>
            <el-tag>{{editUser.name}}</el-tag>
          </el-row>
          <el-row>
            <el-tag type="info">身份</el-tag></el-tag>
            <el-select
              v-model="editUser.identity"
              :placeholder="editUser.identity"
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
              v-model="editUser.status"
              :placeholder="editUser.status"
              style="margin-left: 20px"
            >
              <el-option
                v-for="item in options2"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-row>
          <span slot="footer" class="dialog-footer">
            <el-button @click="Cancel()">取 消</el-button>
            <el-button type="primary" @click="EditUser()">确 定</el-button>
          </span>
        </el-dialog>

        <el-table
          :data="
            user_list_vis.slice(
              (currentPage - 1) * pagesize,
              currentPage * pagesize
            )
          "
        >
                
        <el-table-column prop="id" label="账号" width="150">
        </el-table-column>
          <el-table-column prop="name" label="名称" width="150">
          </el-table-column>
          <el-table-column prop="identity" label="身份" width="150">
          </el-table-column>
          <el-table-column prop="status" label="状态" width="150">
          </el-table-column>

          <el-table-column label="操作" width="300">
            <template slot-scope="scope"
              >
              <el-button type="warning" @click="EditUserVis(scope.row)"
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
      user_list_vis: [],
      searchPlhText: "",
      user_list: [],
      user: {
        name:"",
        id:"",
        identity:"",
        status:"",
        password:"",
        introduction:"",
      },
      editUser: {
        name:"",
        id:"",
        identity:"",
        status:"",
        password:"",
        introduction:"",
      },


      EditUserVisible: false,
      options: [
        {
          value: "administrator",
          label: "administrator",
        },
        {
          value: "user",
          label: "user",
        },
      ],
      options2: [
        {
          value: "正常",
          label: "正常",
        },
        {
          value: "冻结",
          label: "冻结",
        },

      ],
    };
  },
  created() {
    this.axios({
      method: "GET",
      url: this.global.apiUrl + "/getAllUser",
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
          this.user_list = res.data.user_list;
          this.user_list_vis = this.user_list;
        }
      })
      .catch((err) => {});
  },
  watch: {},
  methods: {

    Cancel() {
      this.EditUserVisible = false;
    },

    EditUser() {
      this.axios({
        url: this.global.apiUrl + "/updateUser",
        data: {
          id: this.editUser.id,
          name: this.editUser.name,
          status: this.editUser.status,
          password: this.editUser.password,
          identity: this.editUser.identity,
          introduction: this.editUser.introduction,
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
            (this.EditUserVisible = false),
              (this.editUser.name = ""),
              (this.editUser.status = ""),
              (this.editUser.id = ""),
              (this.editUser.password= ""),
              (this.editUser.introduction = ""),
              (this.editUser.identity = "");
          }
        })
        .catch((Error) => {
          console.log(Error);
        });
    },
    EditUserVis(row) {
      this.EditUserVisible = true;
      this.editUser.name = row.name;
      this.editUser.id= row.id;
      this.editUser.identity = row.identity;
      this.editUser.status = row.status;
      this.editUser.introduction = row.introduction;
      this.editUser.password = row.password;
    },

    search() {
      if (this.searchPlhText == "") {
        this.user_list_vis = this.user_list;
      } else {
        //获取到查询的值，并使用toLowerCase():把字符串转换成小写，让模糊查询更加清晰
        let _search = this.searchPlhText.toLowerCase();
        let newListData = []; // 用于存放搜索出来数据的新数组
        if (_search) {
          //filter 过滤数组
          this.user_list.filter((item) => {
            // newListData中 没有查询的内容，就添加到newListData中
            if (item.name.toLowerCase().indexOf(_search) !== -1||item.id.toLowerCase().indexOf(_search) !== -1||item.identity.toLowerCase().indexOf(_search) !== -1) {
              newListData.push(item);
            }
          });
        }
        //查询后的表格 赋值过滤后的数据
        this.user_list_vis = newListData;
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
