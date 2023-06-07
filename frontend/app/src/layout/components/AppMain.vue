<template>
  <section class="app-main">
    <router-view v-slot="{ Component }">
      <transition name="fade-transform" mode="out-in">
        <keep-alive>
          <component :is="Component" :key="Component.name" />
        </keep-alive>
      </transition>
    </router-view>
  </section>
</template>

<script>
import { useRoute } from 'vue-router'
import { computed } from 'vue'
export default {
  name: 'AppMain',
  setup() {
    const route = useRoute()
    const key = computed(() => {
      return route.path
    })
    return {
      key
    }
  }
}
</script>

<style lang="scss" scoped>
.app-main {
  /* 50= navbar  50  */
  min-height: calc(100vh - 50px);
  width: 100%;
  height: 100%;
  position: relative;
  overflow: hidden;
  // display: flex;
  // justify-content: center;
  // align-items: center;
}

.fixed-header+.app-main {
  padding-top: 50px;
}

.hasTagsView {
  .app-main {
    /* 84 = navbar + tags-view = 50 + 34 */
    min-height: calc(100vh - 84px);
  }

  .fixed-header+.app-main {
    padding-top: 84px;
  }
}
</style>

<style lang="scss">
// fix css style bug in open el-dialog
.el-popup-parent--hidden {
    .fixed-header {
        padding-right: 15px;
    }
}
</style>
