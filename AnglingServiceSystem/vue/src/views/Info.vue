<template>
  <div>
    <HeaderTop></HeaderTop>
    <h1 class="header">资讯</h1>
    <el-button style="margin-left: 1300px">发布资讯</el-button>
    <div class="info_body">
      <el-table :data="infos" style="width: 100%">
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
      infos: [
        {
          title: "资讯1",
          content:
            "资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯",
          create_time: "12345678",
        },
        {
          title: "资讯2",
          content:
            "资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯资讯",
          create_time: "12345678",
        },
        {
          title: "资讯3",
          content: "hahahahha",
          create_time: "12345678",
        },
        {
          title: "资讯4",
          content: "hahahahha",
          create_time: "12345678",
        },
      ],
    };
  },
  created() {},
  watch: {},
  methods: {
    OneInfo(info) {
      this.$alert(info.content, info.title, {
        confirmButtonText: "确定",
      });
    },
  },
};
</script>

<style>
.info_body {
  margin: auto;
  width: 900px;
}
</style>
