<template>
  <div class="create-view-section">
    <div v-if="title && title.length > 0" class="create-view-section__header">
      <div v-if="!$slots.header" class="create-view-section__mark" />
      <div v-if="!$slots.header" class="create-view-section__title">{{ title }}</div>
      <slot name="header" />
    </div>
    <div class="create-view-section__content">
      <slot />
    </div>
  </div>
</template>
<script type="text/javascript">
export default {
  name: "CreateSections",
  components: {},
  props: {
    title: {
      type: String,
      default: ""
    }
  },
  data() {
    return {};
  },
  computed: {},
  mounted() {},
  methods: {}
};
</script>
<style lang="less" scoped>
.create-view-section {
  position: relative;
  background-color: white;
  margin-top: 8px;

  &:first-child {
    margin-top: 0;
  }

  &__mark {
    border-left-color: #2362fb;
    border-left-width: 4px;
    border-left-style: solid;
    border-radius: 2px;
    height: 14px;
  }

  &__header {
    display: flex;
    align-items: center;
    padding: 5px 15px;
  }
  &__title {
    font-size: 14px;
    color: #333;
    font-weight: 600;
    margin-left: 8px;
    flex-shrink: 0;
  }

  &__content {
    padding: 0 10px;
  }
}
</style>
