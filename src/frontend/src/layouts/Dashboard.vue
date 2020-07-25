<template>
  <v-app v-if="userLoaded">
    <dashboard-core-app-bar />

    <dashboard-core-drawer />

    <dashboard-core-view />

    <dashboard-core-settings />
  </v-app>
</template>

<script>
export default {
  name: 'Dashboard',
  middleware: 'auth',
  components: {
    DashboardCoreAppBar: () => import('@@/components/core/AppBar'),
    DashboardCoreDrawer: () => import('@@/components/core/Drawer'),
    DashboardCoreSettings: () => import('@@/components/core/Settings'),
    DashboardCoreView: () => import('@@/components/core/View')
  },
  head () {
    return {
      title: this.$store.state.page.title
    }
  },
  data: () => ({
    userLoaded: false
  }),
  created () {
    this.initUser()
      .then(() => {
        this.userLoaded = true
      })
      .catch(() => {
        this.$router.push('/auth/login/')
      })
  }
}
</script>
