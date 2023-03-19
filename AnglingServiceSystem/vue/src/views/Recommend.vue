<template>
  <div>
    <HeaderTop></HeaderTop>
    <h1 class="header">推荐</h1>
    <el-button style="margin-left: 1750px">添加推荐</el-button>
    <div class="recommend_body">
      <el-table :data="recommend_list" style="width: 100%">
        <el-table-column prop="name" label="名称" width="600">
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="300">
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
      recommend_list: [
        {
          name: "鱼类1",
        },
        {
          name: "鱼类2",
        },
        {
          name: "鱼类3",
        },
        {
          name: "鱼类4",
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
.recommend_body {
  margin: auto;
  width: 900px;
}
</style>
