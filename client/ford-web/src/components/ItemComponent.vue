<template>
  <div class="item">
    <div class="text">some text</div>
    <div class="actions">vote up vote down</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ItemComponent",
});
</script>

<style scoped>
.well-went > .item {
  background-color: #009886;
}

.to-improve > .item {
  background-color: #e92c64;
}

.action-items > .item {
  background-color: #a63eb9;
}

.item {
  padding: 10px;
  color: white;
  border-radius: 5px;
}

.item textarea {
  width: 100%;
  resize: none;
}

.actions {
  margin-top: 10px;
  text-align: right;
}
</style>
