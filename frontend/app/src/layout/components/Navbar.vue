<template>
  <div class="navbar">
    <div class="right-menu">
      <el-dropdown class="avatar-container right-menu-item hover-effect" trigger="click">
        <span class="el-dropdown-link">
          下拉菜单 <el-icon><ArrowDown /></el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <router-link to="/profile/index">
              <el-dropdown-item>个人中心</el-dropdown-item>
            </router-link>
            <router-link to="/">
              <el-dropdown-item>首页</el-dropdown-item>
            </router-link>
            <a target="_blank" href="https://github.com/shoggothforever/rgadmin">
              <el-dropdown-item>项目地址</el-dropdown-item>
            </a>
            <el-dropdown-item>
              <span style="display: block" @click="logout">登出</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import { useUserStore } from '../../stores/user';
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus';
import { usePermissionStore } from '../../stores/permission';

const userStore = useUserStore()
const permissionStore = usePermissionStore()
const router = useRouter()

const logout = async() => {
  await userStore.logout()
  permissionStore.RESET_ROUTES()
  ElMessage.success('登出成功')
  router.push('/login')
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;

        &:hover {
          background: rgba(0, 0, 0, 0.025);
        }
      }
    }

    .avatar-container {
      margin-right: 30px;

      .avatar-wrapper {
        margin-top: 5px;
        position: relative;

        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
    }
  }
}
  .el-dropdown-link {
    line-height: 50px !important;
    cursor: pointer;
    color: #409EFF !important;
  }
  .el-icon-arrow-down {
    font-size: 12px !important;
  }
</style>
