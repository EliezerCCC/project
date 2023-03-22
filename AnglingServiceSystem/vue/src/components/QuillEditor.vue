<!--富文本编辑器-->
<template>
  <div class="RichTextEditor-Wrap" v-loading="loading">
    <quill-editor
      :content="content"
      :options="editorOption"
      class="ql-editor"
      ref="myQuillEditor"
      @change="onEditorChange($event)"
    >
    </quill-editor>

    <!-- 图片上传组件辅助-->
    <el-upload
      v-show="false"
      :show-file-list="false"
      :name="uploadImgConfig.name"
      :multiple="false"
      :action="uploadImgConfig.uploadUrl"
      :before-upload="onBeforeUpload"
      :on-success="onSuccess"
      :on-error="onError"
      :file-list="fileList"
    >
      <!--<i class="el-icon-upload"></i>-->
      <!--<div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>-->
      <!--<div class="el-upload__tip" slot="tip">最多只能上传两个附件</div>-->
      <button ref="myinput">上传文件</button>
    </el-upload>

    <!--视频上传-->
    <div>
      <el-dialog
        :close-on-click-modal="false"
        width="50%"
        style="margin-top: 1px"
        title="视频上传"
        :visible.sync="videoDialog.show"
        append-to-body
      >
        <el-tabs v-model="videoDialog.activeName">
          <el-tab-pane label="添加视频链接" name="first">
            <el-input
              v-model="videoDialog.videoLink"
              placeholder="请输入视频链接"
              clearable
            ></el-input>
            <el-button
              type="primary"
              size="small"
              style="margin: 20px 0px 0px 0px"
              @click="addVideoLink(videoDialog.videoLink)"
              >添加
            </el-button>
          </el-tab-pane>
          <el-tab-pane label="本地视频上传" name="second">
            <el-upload
              v-loading="loading"
              style="text-align: center"
              drag
              :action="uploadVideoConfig.uploadUrl"
              accept="video/*"
              :name="uploadVideoConfig.name"
              :before-upload="onBeforeUploadVideo"
              :on-success="onSuccessVideo"
              :on-error="onErrorVideo"
              :multiple="false"
            >
              <i class="el-icon-upload"></i>
              <div class="el-upload__text">
                将文件拖到此处，或<em>点击上传</em>
              </div>
              <div class="el-upload__tip" slot="tip">
                只能上传MP4文件，且不超过{{ uploadVideoConfig.maxSize }}M
              </div>
            </el-upload>
          </el-tab-pane>
        </el-tabs>
      </el-dialog>
    </div>
  </div>
</template>
<script>
// require styles
import "quill/dist/quill.core.css";
import "quill/dist/quill.snow.css";
import "quill/dist/quill.bubble.css";
import { quillEditor } from "vue-quill-editor";
// 设置title
import { addQuillTitle } from "./quill-title.js";

// 工具栏
const toolbarOptions = [
  ["bold", "italic", "underline", "strike"], // toggled buttons
  ["blockquote", "code-block"],
  [{ header: 1 }, { header: 2 }],
  [{ list: "ordered" }, { list: "bullet" }],
  [{ script: "sub" }, { script: "super" }], // superscript/subscript
  [{ indent: "-1" }, { indent: "+1" }], // outdent/indent
  [{ direction: "rtl" }],
  [{ size: ["small", false, "large", "huge"] }],
  [{ header: [1, 2, 3, 4, 5, 6, false] }],
  [{ font: [] }],
  [{ color: [] }, { background: [] }], // dropdown with defaults from theme
  [{ align: [] }],
  [{ clean: "清除" }], // remove formatting button
  // ['link', 'image', 'video']
  ["image", "video"],
];
export default {
  name: "RichTextEditor",
  model: {
    prop: "content",
    event: "change",
  },
  components: {
    quillEditor,
  },
  props: {
    content: {
      // 返回的html片段
      type: String,
      default: "",
    },
    uploadImgConfig: {
      // 图片上传配置 - 若不配置则使用quillEditor默认方式，即base64方式
      type: Object,
      default() {
        return {
          uploadUrl: "", // 图片上传地址
          maxSize: 2, // 图片上传大小限制，默认不超过2M
          name: "Filedata", // 图片上传字段
        };
      },
    },
    uploadVideoConfig: {
      // 视频上传配置
      type: Object,
      default() {
        return {
          uploadUrl: "", // 上传地址
          maxSize: 10, // 图片上传大小限制，默认不超过2M
          name: "Filedata", // 图片上传字段
        };
      },
    },
  },
  data() {
    let _self = this;
    return {
      loading: false, // 加载loading
      editorOption: {
        placeholder: "",
        theme: "snow", // or 'bubble'
        modules: {
          toolbar: {
            container: toolbarOptions, // 工具栏
            handlers: {
              video: function (value) {
                _self.videoDialog.show = true;
              },
            },
          },
        },
      },

      // 图片上传变量
      fileList: [],

      // 视频上传变量
      videoDialog: {
        show: false,
        videoLink: "",
        activeName: "first",
      },
    };
  },
  mounted() {
    // 初始给编辑器设置title
    addQuillTitle();

    let toolbar = this.$refs["myQuillEditor"].quill.getModule("toolbar");
    // 是否开启图片上传到服务器功能
    if (this.uploadImgConfig.uploadUrl) {
      toolbar.addHandler("image", this.addImageHandler);
    }
  },
  methods: {
    // 文本编辑
    onEditorChange({ quill, html, text }) {
      // console.log('editor change!', quill, html, text)
      // console.log(html.replace(/<[^>]*>|/g, ''), 33333333)
      this.$emit("update:content", html);
      this.$emit("change", html);
    },
    hideLoading() {
      this.loading = false;
    },
    // --------- 图片上传相关 start ---------

    addImageHandler(value) {
      if (value) {
        // 触发input框选择图片文件
        this.$refs["myinput"].click();
      } else {
        this.quill.format("image", false);
      }
    },
    // 把已经上传的图片显示回富文本编辑框中
    uploadSuccess(imgurl) {
      let quill = this.$refs["myQuillEditor"].quill;
      let range = quill.getSelection();
      let index = 0;
      if (range == null) {
        index = 0;
      } else {
        index = range.index; // 获取光标所在位置
      }
      // 插入
      quill.insertEmbed(index, "image", imgurl); // imgurl是服务器返回的图片链接地址
      // 调整光标到最后
      quill.setSelection(index + 1);
    },
    // el-文件上传组件
    onBeforeUpload(file) {
      this.loading = true;
      let acceptArr = ["image/jpg", "image/jpeg", "image/gif", "image/png"];
      const isIMAGE = acceptArr.includes(file.type);
      const isLt1M = file.size / 1024 / 1024 < this.uploadImgConfig.maxSize;
      if (!isIMAGE) {
        this.hideLoading();
        this.$message.error("只能插入图片格式!");
      }
      if (!isLt1M) {
        this.hideLoading();
        this.$message.error(
          `上传文件大小不能超过 ${this.uploadImgConfig.maxSize}MB!`
        );
      }
      return isLt1M && isIMAGE;
    },
    // 文件上传成功时的钩子
    onSuccess(response, file, fileList) {
      // ---- 注意这部分需要改为对应的返回格式
      this.hideLoading();
      if (response.retCode === "00") {
        this.uploadSuccess(response.url);
      } else {
        this.$message.error("上传失败");
      }
    },
    // 文件上传失败时的钩子
    onError(file, fileList) {
      this.hideLoading();
      this.$message.error("上传失败");
    },

    // --------- 图片上传相关 end ---------

    // --------- 视频上传相关 start ---------

    addVideoLink(videoLink) {
      if (!videoLink) return this.$message.error("请输入视频地址");
      this.videoDialog.show = false;
      let quill = this.$refs["myQuillEditor"].quill;
      let range = quill.getSelection();
      let index = 0;
      if (range == null) {
        index = 0;
      } else {
        index = range.index;
      }
      // 插入
      quill.insertEmbed(index, "video", videoLink);
      // 调整光标到最后
      quill.setSelection(index + 1);
    },

    // el-文件上传组件
    onBeforeUploadVideo(file) {
      this.loading = true;
      let acceptArr = ["video/mp4"];
      const isVideo = acceptArr.includes(file.type);
      const isLt1M = file.size / 1024 / 1024 < this.uploadVideoConfig.maxSize;
      if (!isVideo) {
        this.hideLoading();
        this.$message.error("只能上传mp4格式文件!");
      }
      if (!isLt1M) {
        this.hideLoading();
        this.$message.error(
          `上传文件大小不能超过 ${this.uploadVideoConfig.maxSize}MB!`
        );
      }
      return isLt1M && isVideo;
    },
    // 文件上传成功时的钩子
    onSuccessVideo(response, file, fileList) {
      // ---- 注意这部分需要改为对应的返回格式
      this.hideLoading();
      if (response.retCode === "00") {
        this.addVideoLink(response.url);
      } else {
        this.$message.error("上传失败");
      }
    },
    // 文件上传失败时的钩子
    onErrorVideo(file, fileList) {
      this.hideLoading();
      this.$message.error("上传失败");
    },

    // --------- 视频上传相关 end ---------
  },
};
</script>
<style>
.RichTextEditor-Wrap .ql-container {
  height: 300px;
}

.RichTextEditor-Wrap .ql-editor {
  padding: 0;
}

.RichTextEditor-Wrap .ql-tooltip {
  left: 5px !important;
}
</style>
