<template>
    <div>
        <HeaderTop></HeaderTop>

        <h1 class="header">公告栏</h1>
        <el-row :gutter="20" class="el-row" type="flex" >
                      <el-col :span="8" v-for = "(notice,index) in notices" :key="notice.id" class="el-col" >
                        <el-card class="box-card">
                          <div slot="header" class="clearfix">
                            <span>{{notice.title}}</span>
                          </div>
                          <div class="text ">
                            {{notice.content }}
                          </div>
                          <div style="font-size: small; margin-left: 300px;">
                            发布时间:{{notice.create_time}}
                          </div>
                        </el-card>
                    </el-col>
        </el-row>        
    </div>
</template>
  
  
<script>
    export default {
      data() {
        return {
            notices: [
                    {
                        title:"公告1",
                        content:"公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告公告",
                        create_time:"12345678",
                    },
                    {
                        title:"维护通知",
                        content:"本网站将于本周二凌晨两点开始维护，维护时长2-4小时不等，维护完成即可恢复正常使用。",
                        create_time:"12345678",
                    },
                    {
                        title:"公告3",
                        content:"hahahahha",
                        create_time:"12345678",
                    },
                    {
                        title:"公告4",
                        content:"hahahahha",
                        create_time:"12345678",
                    },
                ],
        }
      },
      created() {
       
      },
      watch: {

      },
      methods: {

      }
    }
  </script>

<style>
  .header{
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
    clear: both
  }

  .box-card {
    width: 480px;
  }
</style>
