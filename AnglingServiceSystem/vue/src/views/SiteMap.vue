<template>
  <div class="map">
    <HeaderTop></HeaderTop>
    <h1 class="header" style="width: 300px; margin: 0 auto">
      热门钓场分布信息
    </h1>
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
          value: "渔人码头、青龙湖、密云水库、云蒙山自然风景区、百花潭公园",
        },
        {
          name: "天津",
          value: "大沽河、南河滨河公园、宜兴埠水库、天津青年湖、天津长江道公园",
        },
        {
          name: "上海",
          value:
            "上海钓鱼台、青浦区沪松渔场、奉贤区崇明乡渔场、松江区泖港钓鱼区",
        },
        {
          name: "重庆",
          value: "大竹林水库、洪崖洞、重庆市渔业科学研究所钓场",
        },
        {
          name: "河北",
          value: "白洋淀、德馨园、桃花潭、威海渔场、蔚县滦河水库",
        },
        {
          name: "河南",
          value:
            "伊滨水库、南阳兰考大桥水库、许昌市河南省钓鱼场、鹤壁市金龙水库、安阳市汤阴县三门峡水库",
        },
        {
          name: "云南",
          value: "昆明滇池、大理洱海",
        },
        {
          name: "辽宁",
          value: "沙河口公园、旅顺口水库、卧龙湖、锦州渔场",
        },
        {
          name: "黑龙江",
          value: "呼玛河、五常大米站、双鸭山市水库、太阳岛公园",
        },
        {
          name: "湖南",
          value: "衡山县磨盘山、常德市东湖、益阳市桃江县长沙渔场",
        },
        {
          name: "安徽",
          value: "宿松河、合肥包公山、阜阳菱湖、黄山云谷寺、安庆太湖钓鱼公园",
        },
        {
          name: "山东",
          value: "威海海岛公园、枣庄市城山水库、青岛大学东路钓鱼台",
        },
        {
          name: "新疆",
          value: "天山天池、乌鲁木齐香山、喀纳斯湖",
        },
        {
          name: "江苏",
          value: "太湖、洞庭湖、秦淮河、常州虹桥水库、扬州水文站",
        },
        {
          name: "浙江",
          value: "千岛湖、西湖、钱塘江、富春江、天目湖",
        },
        {
          name: "江西",
          value: "万年县碧阳湖、上饶康山国家地质公园、九江灵山",
        },
        {
          name: "湖北",
          value: "三峡大坝、武汉东湖、洪湖渔场、黄鹤楼、黄鹤楼水库",
        },
        {
          name: "广西",
          value: "桂林漓江、阳朔蝴蝶泉、南宁橘子洲、北海银滩、贺州绿田钓鱼场",
        },
        {
          name: "甘肃",
          value:
            "嘉峪关市丝绸之路国家地质公园、甘南夏河牛场、陇南文县白龙江水库",
        },
        {
          name: "山西",
          value:
            "黄河壶口大峡谷、悬空寺、天龙山国家森林公园、大同水洼、河津五龙潭",
        },
        {
          name: "内蒙古",
          value: "哈巴河、阿尔山、乌素图伦河",
        },
        {
          name: "陕西",
          value: "西安大雁塔、咸阳渭河、渭南市渭南水库、汉中市汉中水库",
        },
        {
          name: "吉林",
          value: "长白山天池、松花江、松原水库、吉林锦江、白山水库",
        },
        {
          name: "福建",
          value: "鼓浪屿、厦门白城沙滩、福州闽江、泉州晋江长泰钓鱼场",
        },
        {
          name: "贵州",
          value:
            "荔波小七孔、贵阳金阳湖公园、安顺市长江水库、铜仁市大方龙宫潭景区",
        },
        {
          name: "广东",
          value:
            "深圳东湖、广州白鹅潭公园、珠海渔女码头、惠州大亚湾湿地公园、清远崖门海滩",
        },
        {
          name: "青海",
          value:
            "青海湖、西宁市黑马河森林公园、海北州祁连山国家级自然保护区、海南藏族自治州共和县黄河湾水库、海西蒙古族藏族自治州德令哈市塔尔寺沟水库",
        },
        {
          name: "西藏",
          value: "昌都市江达县金沙江、日喀则市亚东县",
        },
        {
          name: "四川",
          value: "都江堰渔场、峨眉山景区、九寨沟景区、成都市双流区紫龙塘水库",
        },
        {
          name: "宁夏",
          value:
            "银川市贺兰山国家森林公园、中卫市海原县沙坡头、石嘴山市平罗县天池渔场、固原市原州区塞外湖景区、吴忠市盐池县红旗渠水库",
        },
        {
          name: "海南",
          value:
            "三亚大东海、海口观澜湖、万宁中和农场、琼海嘉积镇双月湾渔场、陵水清水湾渔场",
        },
        {
          name: "台湾",
          value:
            "日月潭、阿里山、台北市士林官邸公园、屏东县琉球乡大武山国家森林游乐区、南投县集集镇兰阳洞瀑布",
        },
        {
          name: "香港",
          value: "大澳渔村、香港仔码头、太平山顶、西贡海滩、沙田城门水塘",
        },
        { name: "澳门", value: "氹仔岛黑沙环公园、路氹连贯公路沿线水库" },
        {
          name: "南海诸岛",
          value: "",
        },
      ],
    };
  },
  mounted() {
    this.drawLine();
  },
  methods: {
    toT() {
      this.$router.push("/map");
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
