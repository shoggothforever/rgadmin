<template>
  <div v-if="!item.hidden">
    <template v-if="hasOneShowingChild(item.children, item)&&(!onlyOneChild.children||onlyOneChild.noShowingChildren)&&!item.alwaysShow">
      <app-link v-if="onlyOneChild.meta" :to="resolvePath(onlyOneChild.path)">
        <el-menu-item :index="resolvePath(onlyOneChild.path)">
          <item :title="onlyOneChild.meta.title"></item>
        </el-menu-item>
      </app-link>
    </template>

    <el-sub-menu v-else ref="subMenu" :index="resolvePath(item.path)" teleported @mouseover="handleMouseover" @mouseleave="handleMouseleave">
      <template #title>
        <item v-if="item.meta" :title="item.meta.title" :color="color"/>
      </template>
      <sidebar-item
        v-for="child in item.children"
        :key="child.path"
        :is-nest="true"
        :item="child"
        :base-path="resolvePath(child.path)"
        class="nest-menu"
      />
    </el-sub-menu>
  </div>
</template>

<script>
import { isExternal } from '@/utils/validate'
import { ref } from 'vue'
import path from 'path-browserify'
import AppLink from './Link.vue'
import Item from './item.vue'
export default {
  props: {
    item: {
      type: Object,
      required: true
    },
    isNest: {
      type: Boolean,
      default: false
    },
    basePath: {
      type: String,
      default: ''
    }
  },
  name: 'SidebarItem',
  components: { AppLink, Item },
  setup(props) {
    const onlyOneChild = ref(null)
    const color = ref('#fff')
    const hasOneShowingChild = (children = [], parent) => {
      // 递归查找只要有一个子组件是显示的，那么父组件也显示
      const showingChildren = children.filter((item) => {
        if (item.hidden) {
          return false
        } else {
          onlyOneChild.value = item
          return true
        }
      })
      // 如果只有一个子路由，子路由默认展示
      if (showingChildren.length === 1) {
        return true
      }

      // 如果没有子路由显示父路由
      if (showingChildren.length === 0) {
        onlyOneChild.value = { ...parent, path: '', noShowingChildren: true }
        return true
      }

      return false
    }

    const handleMouseover = () => {
      color.value = '#489ef7'
    }

    const handleMouseleave = () => {
      color.value = '#fff'
    }

    // 递归生成菜单
    const resolvePath = (routePath) => {
      if (isExternal(routePath)) {
        return routePath
      }
      if (isExternal(props.basePath)) {
        return props.basePath
      }
      return path.resolve(props.basePath, routePath)
    }
    return {
      hasOneShowingChild,
      resolvePath,
      onlyOneChild,
      color,
      handleMouseover,
      handleMouseleave
    }
  }
}
</script>

<style lang="scss" scoped>
.el-sub-menu {
  background-color: #304156 !important;
}

:deep(.el-sub-menu__title:hover)  {
background: rgb(234, 235, 237) !important;
}

.el-menu-item.is-active {
    color: var(--el-menu-active-color) !important;
}
</style>
