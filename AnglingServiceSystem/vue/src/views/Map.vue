<template>
  <div class="map">
    <HeaderTop></HeaderTop>
    <h1 class="header" style="width: 300px; margin: 0 auto">鱼类分布信息</h1>
    <el-button style="font-size: 20px; margin-left: 1300px" @click="toT()"
      >切换</el-button
    >
    <div id="main" class="charts"></div>
  </div>
</template>

<script>
import * as echarts from "echarts";
import dataJson from "../assets/data";
export default {
  name: "home",
  data() {
    return {
      dataList: [
        {
          name: "北京",
          value: "黄颡鱼、白鲢、花鲢、小口鱼、细鳞青鱼、鲤鱼等。",
        },
        {
          name: "天津",
          value:
            "鳜鱼、大头鲢、红嘴鲌、天鹅鱼、天津鲫鱼、鲤鱼、青鱼、草鱼、鲢鱼、鲶鱼、鲈鱼等。",
        },
        {
          name: "上海",
          value:
            "青鱼、鲤鱼、鲢鱼、鳙鱼、银鱼、草鱼、鲫鱼、鲈鱼、黄颡鱼、白鲑等。",
        },
        {
          name: "重庆",
          value: "鲫鱼、鲢鱼、青鱼、草鱼、鲤鲫杂交鱼、鲤鱼、鳙鱼、罗非鱼等。",
        },
        {
          name: "河北",
          value:
            "塘鳢、鲤鱼、鳊鱼、鲢鱼、草鱼、鳙鱼、鲈鱼、鲮鱼、鲂鱼、鲫鱼、青鱼等。",
        },
        {
          name: "河南",
          value:
            "鲤鱼、鳊鱼、草鱼、鲢鱼、鲫鱼、鲤鱼、罗非鱼、青鱼、泥鳅、鳙鱼、鲈鱼、鳝鱼等。",
        },
        {
          name: "云南",
          value:
            "罗非鱼、草鱼、鲫鱼、青鱼、鲤鱼、鳊鱼、鳅鱼、鲢鱼、白鲑、鲑鱼、鲈鱼、黑鱼、鳝鱼等。",
        },
        {
          name: "辽宁",
          value:
            "鲅鱼、鳕鱼、鲽鱼、三文鱼、黄鱼、鲈鱼、鲫鱼、鲤鱼、草鱼、鳊鱼、青鱼、鳙鱼、鳜鱼等。",
        },
        {
          name: "黑龙江",
          value:
            "鲤鱼、草鱼、鲢鱼、鲟鱼、马哈鱼、扬子鳢、鳜鱼、鲈鱼、鳊鱼、鲑鱼等。",
        },
        {
          name: "湖南",
          value:
            "草鱼、鲢鱼、鲤鱼、鳙鱼、青鱼、白鲟、鳙鱼、花鱼、罗非鱼、黑鱼等。",
        },
        {
          name: "安徽",
          value:
            "黄鳝、鲢鱼、鳙鱼、丽鱼、团头鱼、鳊鱼、太湖银鱼、鲤鱼、鲫鱼、鲢鱼、鳊鱼、草鱼、鲈鱼、青鱼、鲶鱼、鳙鱼、鳝鱼",
        },
        {
          name: "山东",
          value:
            "黄鱼、花鲢、青头鱼、带鱼、鲥鱼、海鳗、鲭鱼、马鲛鱼、鲫鱼、鲤鱼、鳊鱼、草鱼、鲢鱼、鲇鱼、鳙鱼、鲳鱼、黄鳝、泥鳅、鲈鱼",
        },
        {
          name: "新疆",
          value:
            "刺鲤、狗鱼、红目鲌、锦鲤、草鱼、鲢鱼、罗非鱼、苏尼特右旗鱼、野生马哈鱼等。",
        },
        {
          name: "江苏",
          value: "鲢鱼、鳙鱼、鲈鱼、青鱼、草鱼、鲫鱼、鲤鱼、鳊鱼等。",
        },
        {
          name: "浙江",
          value: "鲈鱼、黄鳝、东方鲈鱼、乌鳢鱼、花鲢鱼、鲤鱼、鲢鱼、草鱼等。",
        },
        {
          name: "江西",
          value:
            "草鱼、鲢鱼、鲤鱼、鳙鱼、鳊鱼、鲮鱼、鳜鱼、鲈鱼、青鱼、鲳鱼、罗非鱼、金鱼等。",
        },
        {
          name: "湖北",
          value:
            "草鱼、鲢鱼、青鱼、鲤鱼、鳙鱼、罗非鱼、鲈鱼、鲶鱼、鲫鱼、黑鱼等。",
        },
        {
          name: "广西",
          value:
            "鳙鱼、黄颡鱼、草鱼、鲤鱼、黑鱼、鲈鱼、鲮鱼、红鲌、鲮鲤、蒲鱼、沙鱼、乌鱼、罗非鱼等。",
        },
        { name: "甘肃", value: "马面鲃、银鲢、黑鲢、黄颡鱼、驴头鱼等。" },
        {
          name: "山西",
          value:
            "黄河鲤鱼、山西草鱼、太行八仙鱼、沁河鲢鱼、草鱼、鲫鱼、鲤鱼、青鱼、鳊鱼、鲢鱼等",
        },
        {
          name: "内蒙古",
          value: "鲤鱼、鲫鱼、草鱼、鳙鱼、鲇鱼、鲶鱼、鳊鱼、青鱼、白鲢、石首鱼",
        },
        { name: "陕西", value: "鳜鱼、鲤鱼、青鱼、草鱼、黑鱼、鲫鱼等。" },
        {
          name: "吉林",
          value:
            "鳟鱼、鲫鱼、鲤鱼、黑鱼、泥鳅、草鱼、鲈鱼、青鱼、鲢鱼、鲑鱼、罗非鱼。",
        },
        {
          name: "福建",
          value: "沙鳅、白鲑、鳜鱼、鲢鱼、鲫鱼、泥鳅、草鱼、罗非鱼等。",
        },
        { name: "贵州", value: "草鱼、鲢鱼、鲫鱼、鲤鱼、青鱼、黄颡鱼等。" },
        {
          name: "广东",
          value:
            "珠江花鱼、大头鱼、红鲂、草鱼、鲢鱼、鳜鱼、鲳鱼、鲈鱼、黑鱼、乌鳢等。",
        },
        {
          name: "青海",
          value:
            "青海湖裸鲤、青海湖大鲵、青海湖银鲤、鲤鱼、鲫鱼、鳜鱼、草鱼等。",
        },
        {
          name: "西藏",
          value:
            "藏鲤、青海湖无鳞鲤、裂腹鱼、高原裂腹鱼、托木尔雅罗鱼、高原哲罗鱼等。",
        },
        {
          name: "四川",
          value:
            "草鱼、青衣、马口鱼、赤尾鱼、鲢鱼、草鱼鲮鱼、鲫鱼、罗非鱼、鳊鱼、黑鱼、鲤鱼",
        },
        {
          name: "宁夏",
          value: "草鱼、铜鱼、裂脊鱼、鲫鱼、鲢鱼、鳙鱼、青鱼等。",
        },
        {
          name: "海南",
          value:
            "草鱼、鲫鱼、鳊鱼、鲈鱼、鲳鱼、鲷鱼、旗鱼、马鲛鱼、石斑鱼、乌鳢、黑鱼、石首鱼、带鱼等。",
        },
        {
          name: "台湾",
          value:
            "鲤鱼、鲫鱼、草鱼、鲈鱼、黑鲷、金线鱼、青斑、竹筴鱼、鲷鱼、鲻鱼、鲢鱼、赤鯮、石斑鱼。",
        },
        {
          name: "香港",
          value:
            "鲈鱼、鲷鱼、鲤鱼、鲫鱼、白鲳、鰆鱼、鲅鱼、石斑、金目鲈、黑鲷等。",
        },
        { name: "澳门", value: "花鲢、乌鳢、青衣、黄鳝、白鲳、鲫鱼等。" },
        {
          name: "南海诸岛",
          value:
            "金枪鱼、鲢鱼、马鲅、鲨鱼、黄鳍金眼鲳、花鮨、剑鱼、带鱼、鲈鱼、鲻鱼、鲳鱼、鳕鱼、鳗鱼、鲨鱼、旗鱼、鲸鱼、海龙、石斑鱼等。",
        },
      ],
    };
  },
  mounted() {
    this.drawLine();
  },
  methods: {
    toT() {
      this.$router.push("/sitemap");
    },
    drawLine() {
      let myChart = echarts.init(document.getElementById("main"));
      let uploadDataUrl = dataJson;
      // 注册地图
      echarts.registerMap("china", uploadDataUrl);
      let option = {
        geo: {
          // 地理坐标系组件,支持在地理坐标系上绘制散点图、线图
          map: "china",
          aspectScale: 0.75,
          zoom: 1.1,
        },
        tooltip: {
          formatter: (params) => {
            let str = params.data.value;
            let newStr = "";
            let max_line = 20;
            let i = 0;
            for (i = 0; i <= str.length; i += max_line) {
              newStr += str.substring(i, i + max_line) + "</br>";
            }

            return newStr;
          },
        },
        series: [
          {
            zoom: 1.1,
            map: "china",
            type: "map",
            itemStyle: {
              normal: {
                borderColor: "rgba(0, 0, 0, 0.2)",
              },
              emphasis: {
                shadowOffsetX: 0,
                shadowOffsetY: 0,
                shadowBlur: 20,
                borderWidth: 0,
                shadowColor: "rgba(0, 0, 0, 0.5)",
              },
            },
            label: {
              normal: {
                show: true,
              },
              emphasis: {
                show: true,
              },
            },
            data: this.dataList,
          },
        ],
      };
      myChart.setOption(option);
      window.onresize = function () {
        myChart.resize();
      };
    },
  },
};
</script>

<style scoped>
.charts {
  min-height: 800px;
  min-width: 1400px;
}
.charts:first-child div {
  position: static !important;
  width: 100% !important;
}
</style>
