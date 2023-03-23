<template>
  <div>
    <HeaderTop></HeaderTop>
    <h1 class="header">商城</h1>
    <el-button style="margin-left: 1750px">发布商品</el-button>
    <div class="shop_body">
      <el-row>
        <el-col :span="4.5" v-for="(o, index) in commodity_list" :key="o">
          <el-card :body-style="{ padding: '0px' }">
            <img
              src="http://localhost:9090/getImage?imageName=./web/static/images/tupian2.jpg"
              class="image"
            />
            <div style="padding: 14px">
              <span>{{ o.name }}</span>
              <div class="bottom clearfix">
                <time class="time">{{ currentDate }}</time>
                <el-button type="text" class="button">查看详情</el-button>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      commodity_list: [
        {
          id: 1,
          name: "商品1",
          introduction: "商品1的介绍",
          price: 9.99,
          image: 1,
          status: "sale",
          seller_id: "123",
          seller_name: "卖家1",
          create_time: "123",
        },
        {
          id: 1,
          name: "商品2",
          introduction: "商品1的介绍",
          price: 19.99,
          image: 1,
          status: "sale",
          seller_id: "123",
          seller_name: "卖家1",
          create_time: "123",
        },
        {
          id: 1,
          name: "商品3",
          introduction: "商品1的介绍",
          price: 9.99,
          image: 1,
          status: "sale",
          seller_id: "123",
          seller_name: "卖家1",
          create_time: "123",
        },
        {
          id: 1,
          name: "商品4",
          introduction: "商品1的介绍",
          price: 9.99,
          image: 1,
          status: "sale",
          seller_id: "123",
          seller_name: "卖家1",
          create_time: "123",
        },
        {
          id: 1,
          name: "商品4",
          introduction: "商品1的介绍",
          price: 9.99,
          image: 1,
          status: "sale",
          seller_id: "123",
          seller_name: "卖家1",
          create_time: "123",
        },
        {
          id: 1,
          name: "商品4",
          introduction: "商品1的介绍",
          price: 9.99,
          image: 1,
          status: "sale",
          seller_id: "123",
          seller_name: "卖家1",
          create_time: "123",
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
.shop_body {
  margin: auto;
  margin-left: 80px;
  margin-top: 20px;
}

.el-card {
  width: 300px;
  margin-left: 30px;
}

.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 0;
  float: right;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}
</style>
