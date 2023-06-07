<template>
  <div>
    <el-upload
      class="upload-demo"
      drag
      ref="uploadRef"
      action="http://rg.sfct.top/admin/upload"
      multiple
      :fileList="fileList"
      :auto-upload="false"
      :on-success="onSuccess"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        拖动文件到这里或 <em>点击这里上传</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          文件大小不能超过200MB!
        </div>
      </template>
    </el-upload>
    <div class="button-container">
      <el-button style="margin-left: 10px;" size="large" type="success" @click="submitUpload" round>上传到服务器</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const fileList = ref([])
const uploadRef = ref(null)

const submitUpload = () => {
  uploadRef.value.submit()
}
const onSuccess = () => {
  ElMessage.success({
    message: '上传成功',
    type: 'success'
  })
}

</script>

<style lang="scss" scoped>
.button-container {
  display: flex;
  justify-content: center;
}
</style>
