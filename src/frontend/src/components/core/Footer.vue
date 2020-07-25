<template>
  <v-footer
    id="dashboard-core-footer"
    absolute
  >
    <v-container>
      <v-row
        align="center"
        no-gutters
      >
        <v-col
          v-for="(link, i) in links"
          :key="i"
          class="text-center mb-sm-0 mb-5"
          cols="auto"
        >
          <a
            :href="link.href"
            class="mr-0 grey--text text--darken-3"
            rel="noopener"
            v-text="link.text"
          />
        </v-col>

        <v-spacer class="hidden-sm-and-down" />

        <v-col
          cols="12"
          md="auto"
        >
          <div class="body-1 font-weight-light pt-6 pt-md-0 text-center">
            &copy; 2020, <a href="https://gitlab.com/gadelkareem/skeleton">Skeleton</a>.
          </div>
        </v-col>
      </v-row>
    </v-container>
  </v-footer>
</template>

<script>
export default {
  name: 'DashboardCoreFooter',

  data: () => ({
    links: [
      {
        href: '/',
        text: 'Home'
      },
      {
        href: '/dashboard/home/',
        text: 'Dashboard'
      }
    ]
  })
}
</script>

<style lang="sass">
  #dashboard-core-footer
      font-size: .825rem
      font-weight: 500
      text-decoration: none
      text-transform: uppercase
      /*position: absolute*/
      /*bottom: 0*/
      /*width: 100%*/
</style>
