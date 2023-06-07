<template>
  <div>
    <el-scrollbar wrap-class="scrollbar-wrapper">
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :background-color="variables.menuBg"
        :text-color="variables.menuText"
        :unique-opened="false"
        :active-text-color="variables.menuActiveText"
        :collapse-transition="false"
        mode="vertical"
      >
        <sidebar-item v-for="route in permission_routes" :key="route.path" :item="route" :base-path="route.path" />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script setup>
import SidebarItem from './SidebarItem.vue'
import Variables from '@/styles/variables.scss?inline'
import { computed } from 'vue'
import { usePermissionStore } from '@/stores/permission.js'
import { useAppStore } from '@/stores/app.js'
import { useRoute } from 'vue-router'

const permissionStore = usePermissionStore()
const appStore = useAppStore()
const permission_routes = permissionStore.routes
const sidebar = appStore.sidebar

const activeMenu = computed(() => {
  const route = useRoute()
  const { meta, path } = route
  if (meta.activeMenu) {
    return meta.activeMenu
  }
  return path
})

const variables = computed(() => {
  return Variables
})

const isCollapse = computed(() => {
  return !sidebar.opened
})
</script>

<style lang="scss" scoped></style>
