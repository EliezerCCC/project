<template>
  <div>
    <HeaderTop></HeaderTop>
    <h1 class="header">论坛</h1>
    <el-button style="margin-left: 1750px">发布帖子</el-button>
    <div class="forum_body">
      <el-table :data="post_list" style="width: 100%">
        <el-table-column prop="title" label="主题" width="600">
        </el-table-column>
        <el-table-column prop="create_time" label="发布时间" width="200">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button type="text" size="small">查看</el-button>
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
      post_list: [
        {
          title: "帖子1",
          content:
            "帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子",
          create_time: "12345678",
        },
        {
          title: "帖子2",
          content:
            "帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子帖子",
          create_time: "12345678",
        },
        {
          title: "帖子3",
          content: "hahahahha",
          create_time: "12345678",
        },
        {
          title: "帖子4",
          content: "hahahahha",
          create_time: "12345678",
        },
      ],
    };
  },
  created() {},
  watch: {},
  methods: {},
};
</script>

<style>
.forum_body {
  margin: auto;
  width: 900px;
}
</style>
