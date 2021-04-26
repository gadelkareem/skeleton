<template>
  <v-app v-if="userLoaded">
    <dashboard-core-drawer />
    <dashboard-core-app-bar />
    <v-main app>
      <router-view />
      <br><br>
    </v-main>
    <dashboard-core-footer />
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
    DashboardCoreFooter: () => import('@@/components/core/Footer')
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

<style lang="sass">
@import @@/assets/sass/dashboard_overrides.sass
</style>
