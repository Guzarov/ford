import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import DeskView from "../views/DeskView.vue";
import NewDeskView from "../views/NewDeskView.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "new_desk",
    component: NewDeskView,
  },
  {
    path: "/:date",
    name: "home",
    component: DeskView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
