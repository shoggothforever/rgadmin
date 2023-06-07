<template>
  <div class="app-wrapper">
    <sidebar class="sidebar-container" />
    <div class="main-container">
      <div>
        <NavBar></NavBar>
      </div>
      <app-main />
    </div>
  </div>
</template>

<script>
import AppMain from './components/AppMain.vue'
import NavBar from './components/NavBar.vue'
import Sidebar from './components/Sidebar/index.vue'
export default {
  components: {
    AppMain,
    NavBar,
    Sidebar
  },
  setup() {
    return {
    };
  }
};
</script>

<style lang="scss" scoped>
  @import "@/styles/mixin.scss";
  @import "@/styles/variables.scss";

  .app-wrapper {
    @include clearfix;
    position: relative;
    height: 100%;
    width: 100%;

    &.mobile.openSidebar {
      position: fixed;
      top: 0;
    }
  }

  .drawer-bg {
    background: #000;
    opacity: 0.3;
    width: 100%;
    top: 0;
    height: 100%;
    position: absolute;
    z-index: 999;
  }

  .fixed-header {
    position: fixed;
    top: 0;
    right: 0;
    z-index: 9;
    width: calc(100% - #{$sideBarWidth});
    transition: width 0.28s;
  }

  .hideSidebar .fixed-header {
    width: calc(100% - 54px)
  }

  .mobile .fixed-header {
    width: 100%;
  }

.header {
  height: 3.5em;
  border-bottom: solid 1px rgb(143, 137, 137);
}
</style>
