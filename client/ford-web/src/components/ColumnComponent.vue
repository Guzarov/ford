<template>
  <div class="column">
    <div class="caption">Well went</div>
    <button>+</button>
    <ItemComponent />
    <ItemComponent />
    <ItemComponent />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ItemComponent from "@/components/ItemComponent.vue";

export default defineComponent({
  name: "ColumnComponent",
  components: {
    ItemComponent,
  },
});
</script>

<style scoped>
.column {
  gap: 10px;
  display: grid;
  align-content: start;
  padding: 10px;
}

.caption {
  font-size: large;
}

.column > button {
  height: 30px;
}
</style>
