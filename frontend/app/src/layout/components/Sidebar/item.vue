<script>
import { h, defineComponent } from 'vue';

export default defineComponent({
  name: 'MenuItem',
  props: {
    title: {
      type: String,
      default: ''
    },
    color: {
      type: String,
      default: ''
    }
  },
  setup(props) {
    return () => {
      const { title, color } = props;
      const vnodes = [];

      if (title) {
        vnodes.push(h('span', { slot: 'title', style: { color: color }}, title));
      }
      return vnodes;
    };
  }
});
</script>

<style lang="scss" scoped></style>
