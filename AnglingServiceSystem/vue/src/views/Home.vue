<template>
  <div>
    <HeaderTop></HeaderTop>

    <h1 class="header">公告栏</h1>

    <div class="notice_body">
      <el-table :data="notices" style="width: 100%">
        <el-table-column prop="title" label="主题" width="600">
        </el-table-column>
        <el-table-column prop="create_time" label="发布时间" width="200">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button @click="OneNotice(scope.row)" type="text" size="small"
              >查看</el-button
            >
            <el-button type="text" size="small">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      notices: [
        {
          title: "公告1",
          content:
            "公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告",
          create_time: "12345678",
        },
        {
          title: "维护通知",
          content:
            "本网站将于本周二凌晨两点开始维护，维护时长2-4小时不等，维护完成即可恢复正常使用。",
          create_time: "12345678",
        },
        {
          title: "公告3",
          content: "hahahahha",
          create_time: "12345678",
        },
        {
          title: "公告4",
          content: "hahahahha",
          create_time: "12345678",
        },
      ],
    };
  },
  created() {},
  watch: {},
  methods: {
    OneNotice(notice) {
      this.$alert(notice.content, notice.title, {
        confirmButtonText: "确定",
        callback: (action) => {
          this.$message({
            type: "info",
            message: `action: ${action}`,
          });
        },
      });
    },
  },
};
</script>

<style>
.notice_body {
  margin: auto;
  width: 900px;
}

.header {
  text-align: center;
}

.el-row {
  margin-bottom: 20px;
  display: flex;
  flex-wrap: wrap;
}

.el-col {
  border-radius: 10px;
  align-items: stretch;
  margin-bottom: 40px;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}

.box-card {
  width: 480px;
}
</style>
