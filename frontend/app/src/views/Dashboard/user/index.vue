<template>
    <div class="admin-container">
        <h1>Welcome!</h1>
        <h4>欢迎来到员工系统</h4>
        <p>您的身份是<strong style="color:#489ef7">{{ role }}</strong></p>
        <p>用户姓名:{{ name }}</p>
    </div>
</template>

<script setup>
import { useUserStore } from '../../../stores/user'

const userStore = useUserStore()

const name = userStore.name
const role = userStore.roles[0][0]
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
}
</style>
