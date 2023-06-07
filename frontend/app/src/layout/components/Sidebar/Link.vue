<template>
  <router-link v-bind="linkProps(to)">
      <slot />
  </router-link>
</template>

<script>
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    to: {
      type: String,
      required: true
    }
  },
  setup() {
    const linkProps = (to) => {
      return { to: to };
    };
    return {
      linkProps
    };
  }
});
</script>
