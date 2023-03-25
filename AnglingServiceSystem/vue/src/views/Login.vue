<template>
  <div class="loginbody">
    input
    <div class="logindata">
      <div class="logintext">
        <h2>Welcome</h2>
      </div>
      <div class="formdata">
        <el-form ref="form" :model="form" :rules="rules">
          <el-form-item prop="id">
            <el-input
              v-model="form.id"
              clearable
              placeholder="请输入账号"
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="form.password"
              clearable
              placeholder="请输入密码"
              show-password
            ></el-input>
          </el-form-item>
        </el-form>
      </div>
      <div class="butt">
        <el-button type="primary" @click="login('form')">登录</el-button>
        <el-button class="shou" @click="RegisterVisible = true">注册</el-button>
      </div>
    </div>
    <el-dialog
      title="注册"
      :visible.sync="RegisterVisible"
      width="400px
    "
      append-to-body
    >
      <el-form :model="registerForm" :rules="rules2" ref="registerForm">
        <el-form-item label="账号" prop="id" :label-width="formLabelWidth">
          <el-input
            v-model="registerForm.id"
            autocomplete="off"
            style="width: 250px"
          ></el-input>
        </el-form-item>
        <el-form-item label="昵称" prop="name" :label-width="formLabelWidth">
          <el-input
            v-model="registerForm.name"
            autocomplete="off"
            style="width: 250px"
          ></el-input>
        </el-form-item>
        <el-form-item
          label="密码"
          prop="password"
          :label-width="formLabelWidth"
        >
          <el-input
            type="password"
            v-model="registerForm.password"
            autocomplete="off"
            style="width: 250px"
          ></el-input>
        </el-form-item>
        <el-form-item
          label="确认密码"
          prop="secondpassword"
          :label-width="formLabelWidth"
        >
          <el-input
            type="password"
            v-model="registerForm.secondpassword"
            autocomplete="off"
            style="width: 223px"
          ></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="RegisterVisible = false">取 消</el-button>
        <el-button type="primary" @click="Register('registerForm')"
          >确 定</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
import AES from "../util/md5.js";
export default {
  name: "login",
  data() {
    return {
      form: {
        id: "",
        password: "",
      },
      registerForm: {
        id: "",
        password: "",
        secondpassword: "",
        name: "",
      },
      rules: {
        id: [
          { required: true, message: "请输入账号", trigger: "blur" },
          {
            min: 5,
            max: 11,
            message: "长度在 5 到 11 个字符",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
      },
      rules2: {
        id: [
          { required: true, message: "请输入账号", trigger: "blur" },
          {
            min: 5,
            max: 11,
            message: "长度在 5 到 11 个字符",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
        name: [
          { required: true, message: "请输入昵称", trigger: "blur" },
          {
            min: 2,
            max: 10,
            message: "长度在 2 到 10 个字符",
            trigger: "blur",
          },
        ],
      },
      RegisterVisible: false,
    };
  },
  mounted() {},
  methods: {
    login(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.axios({
            url: this.global.apiUrl + "/login",
            data: {
              id: this.form.id,
              password: AES.encrypt(this.form.password, this.global.md5key),
            },
            method: "POST",
            header: {
              "Content-Type": "application/json",
            },
          })
            .then((res) => {
              if (res.data.result == "账号或密码错误") {
                alert("账号或密码错误!");
              } else {
                alert("登录成功!");
                sessionStorage.setItem("token", res.data.token);
                sessionStorage.setItem("identity", res.data.identity);
                sessionStorage.setItem("user_id", res.data.user_id);
                sessionStorage.setItem("user_name", res.data.user_name);
                this.$router.push("/home");
              }
            })
            .catch((Error) => {
              console.log(Error);
            });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    Register(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (
            (console.log(
              this.registerForm.password,
              " ",
              this.registerForm.secondpassword
            ),
            this.registerForm.password != this.registerForm.secondpassword)
          ) {
            alert("密码不一致!");
          } else {
            this.axios({
              url: this.global.apiUrl + "/register",
              data: {
                id: this.registerForm.id,
                name: this.registerForm.name,
                password: AES.encrypt(
                  this.registerForm.password,
                  this.global.md5key
                ),
                identity: "user",
                status: "正常",
              },
              method: "POST",
              header: {
                "Content-Type": "application/json",
              },
            })
              .then((res) => {
                console.log(res.data);
                if (res.data.result == "账号已存在") {
                  alert("账号已存在!");
                } else {
                  alert("注册成功!");
                  this.RegisterVisible = false;
                }
              })
              .catch((Error) => {
                console.log(Error);
              });
          }
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
  },
};
</script>

<style scoped>
.loginbody {
  width: 100%;
  height: 100%;
  min-width: 1000px;
  background-image: url("../assets/background.jpg");
  background-size: 100% 100%;
  background-position: center center;
  overflow: auto;
  background-repeat: no-repeat;
  position: fixed;
  line-height: 100%;
  padding-top: 150px;
  top: 0px;
  left: 0px;
}

.logintext {
  margin-bottom: 20px;
  line-height: 50px;
  text-align: center;
  font-size: 30px;
  font-weight: bolder;
  color: white;
  text-shadow: 2px 2px 4px #000000;
}

.logindata {
  margin-top: 150px;
  width: 400px;
  height: 300px;
  transform: translate(-50%);
  margin-left: 50%;
}

.tool {
  display: flex;
  justify-content: space-between;
  color: #606266;
}

.butt {
  margin-top: 10px;
  text-align: center;
}

.shou {
  cursor: pointer;
  color: #606266;
}
</style>
