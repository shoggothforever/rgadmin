<template>
  <div class="admin-container">
    <h1>Welcome!</h1>
    <h4>欢迎来到财务管理系统</h4>
    <p>您的身份是<strong style="color:#489ef7">管理员</strong></p>
    <p>用户姓名:{{ name }}</p>
    <br />
    <br />
    <br class="line" />
    <br />
    <br />
    <h4 class="user-number">本月已上报工时人数: {{ userNumber }}</h4>
    <h4 class="all-number">员工总数: {{ allNumber }}</h4>
  </div>
</template>

<script setup>
import { useUserStore } from '../../../stores/user'
import { lookUp } from '../../../api/admin';
import { ref, onMounted } from 'vue';

const userStore = useUserStore()

const name = userStore.name
const userNumber = ref(0)
const allNumber = ref(0)

const getUserNumber = () => {
  lookUp().then((res) => {
    userNumber.value = res.currentCount;
    allNumber.value = res.allCount
  })
}

onMounted(() => {
  getUserNumber()
})
</script>

<style lang="scss" scoped>
.admin-container {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    padding: 5rem 0;

    h1 {
        font-size: 1.5rem;
    }

    h4 {
        font-size: 1.2rem;
    }

    p {
        font-size: 1rem;
    }

    .user-number {
        color:#489ef7;
        border-top: 1px solid lightgray;
        width: 100%;
        display: flex;
        justify-content: center;
        padding: 3rem 0;
    }

    .all-number {
        color: #f76260;
    }

    .line {
      display: block; /* 将<br>标签显示为块级元素 */
      content: ""; /* 添加一个空内容 */
      border-top: 1px solid lightgray; /* 添加上边框，设置为浅灰色 */
      margin: 10px 0; /* 可选：为了给线条留出一些间距，可以调整边距值 */
    }
}
</style>
