<template>
  <HeaderComponent />
  <div class="container">
    <ColumnComponent :class="'well-went'" />
    <ColumnComponent :class="'to-improve'" />
    <ColumnComponent :class="'action-items'" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ColumnComponent from "@/components/ColumnComponent.vue";
import HeaderComponent from "@/components/HeaderComponent.vue"; // @ is an alias to /src

export default defineComponent({
  name: "DeskView",
  components: {
    HeaderComponent,
    ColumnComponent,
  },
});
</script>

<style scoped>
.container {
  padding: 10px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
}
</style>
