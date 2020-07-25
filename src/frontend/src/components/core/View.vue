<template>
  <v-main class="fill-height">
    <router-view />
    <br><br>
    <dashboard-core-footer />
  </v-main>
</template>

<script>
export default {
  name: 'DashboardCoreView',

  components: {
    DashboardCoreFooter: () => import('./Footer')
  }
}
</script>
