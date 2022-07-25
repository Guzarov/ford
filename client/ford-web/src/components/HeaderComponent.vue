<template>
  <div class="header">
    FrontOfficeRetroDesk
    <span>12.12.1212</span>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "HeaderComponent",
  props: {
    msg: String,
  },
});
</script>

<style scoped>
.header {
  padding: 10px;
  font-size: x-large;
  border-bottom: 1px solid #333333;
}

.header > span {
  float: right;
}
</style>
