<template>
  <div>
    <HeaderTop></HeaderTop>
    <h1 class="header">钓场</h1>
    <el-button style="margin-left: 1750px">发布钓场</el-button>
    <div class="site_body">
      <el-table :data="anglingsite_list" style="width: 100%">
        <el-table-column prop="name" label="名称" width="300">
        </el-table-column>
        <el-table-column prop="place" label="地点" width="300">
        </el-table-column>
        <el-table-column prop="status" label="状态" width="200">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button type="text" size="small">查看</el-button>
            <el-button type="text" size="small">管理</el-button>
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
      anglingsite_list: [
        {
          name: "钓场1",
          place: "地点1",
          status: "开放",
        },
        {
          name: "钓场2",
          place: "地点2",
          status: "开放",
        },
        {
          name: "钓场3",
          place: "地点3",
          status: "开放",
        },
        {
          name: "钓场4",
          place: "地点4",
          status: "关闭",
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
.site_body {
  margin: auto;
  width: 900px;
}
</style>
