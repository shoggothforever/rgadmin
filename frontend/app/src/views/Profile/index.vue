<template>
  <div>
    <div v-if="isAdmin()">
      <admin />
    </div>
    <div v-else>
      <user />
    </div>
  </div>
</template>

<script setup>
import admin from './admin/index.vue'
import user from './user/index.vue'
import { useUserStore } from '../../stores/user';

const userStore = useUserStore()

const roles = userStore.roles

const isAdmin = () => {
  return roles[0].includes('finance staff')
}
</script>

<style lang="scss" scoped>
</style>
