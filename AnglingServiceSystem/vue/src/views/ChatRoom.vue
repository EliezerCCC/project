<template>
  <div>
    <HeaderTop></HeaderTop>

    <div>
      <el-container>
        <el-aside
          width="200px"
          style="height: 875px; background-color: rgb(223, 223, 223)"
        >
          <h2 style="margin-left: 60px">聊天室</h2>
          <p style="margin-left: 50px">房间人数: {{ userCount }}</p>
        </el-aside>

        <el-main>
          <el-row>
            <el-table
              :data="chatData"
              :show-header="false"
              ref="table"
              style="width: 1688px"
              height="800"
            >
              <el-table-column>
                <template slot-scope="scope">
                  <span style="font-size: 20px">@{{ scope.row.user }}</span>
                  <i class="el-icon-time" style="margin-left: 10px"></i>
                  <span>{{ scope.row.timestamp | formatDate }}</span>
                  <p style="margin-left: 2px">{{ scope.row.text }}</p>
                </template>
              </el-table-column>
            </el-table>
          </el-row>
          <el-row style="margin-top: 10px; margin-left: 20px">
            <el-input
              placeholder="请输入内容"
              v-model="msg"
              ref="input"
              @keyup.enter.native="sendMessage"
              style="width: 500px"
            ></el-input>

            <el-button
              type="primary"
              @click="sendMessage"
              style="margin-left: 10px"
              >发送</el-button
            >
            <el-button type="info" @click="clearInput">清空</el-button>
          </el-row>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: sessionStorage.getItem("user_name"),
      EventTypeMsg: "event-msg",
      EventTypeSystem: "event-system",
      EventTypeJoin: "event-join",
      EventTypeTyping: "event-typing",
      EventTypeLeave: "event-leave",
      EventTypeImage: "event-image",
      socket: null,
      userCount: 0,
      msg: "",
      chatData: [],
    };
  },

  created() {},
  mounted() {
    const socket = this.initSocket(this.name);
    this.setUpSocket(socket);
    this.socket = socket;
  },
  destroyed() {
    this.socket.close();
  },
  watch: {
    socket(val) {
      if (!val) {
        this.socket = this.initSocket(sessionStorage.getItem("user_name"));
        this.setUpSocket(this.socket);
      }
    },
  },
  filters: {
    formatDate(val) {
      const date = new Date(val);
      const y = date.getFullYear();
      const m = date.getMonth() + 1;
      const d = date.getDate();
      const hh = date.getHours();
      const mm = date.getMinutes();
      const ss = date.getSeconds();
      return `${m}-${d} ${hh}:${mm}:${ss}`;
    },
  },
  methods: {
    initSocket(username) {
      let url = `ws://localhost:9090/ws/socket?name=${username}`;
      const socket = new WebSocket(url);
      return socket;
    },
    setUpSocket(socket) {
      socket.onopen = () => {
        this.$message({
          type: "success",
          message: "聊天室连接成功",
        });
      };
      socket.onclose = () => {
        this.$message({
          type: "warning",
          message: "连接断开",
        });
        this.socket = null;
      };
      socket.onmessage = (event) => {
        let dt = JSON.parse(event.data);
        switch (dt.type) {
          case this.EventTypeMsg:
            this.receiveMsg(dt);
            this.userCount = dt.userCount;
            break;
          case this.EventTypeSystem:
            this.userCount = dt.userCount;
            break;
        }
      };
    },
    clearInput() {
      this.msg = "";
      this.$refs.input.focus();
    },
    sendMessage() {
      if (!this.msg) {
        this.$refs.input.focus();
        return;
      }
      const req = JSON.stringify({
        msg: this.msg,
      });
      this.socket &&
        (this.socket.send(req), (this.msg = ""), this.$refs.input.focus());
    },
    receiveMsg(data) {
      this.chatData.push(data);
    },
  },
};
</script>

<style></style>
