<template>
  <v-app>
    <v-app-bar
      app
      text
    >
      <v-container
        mx-auto
        py-0
      >
        <v-layout>
          <router-link to="/">
            <v-img
              :src="require('@@/static/logo.svg')"
              contain
              height="48"
              to="/"
              position="left"
            />
          </router-link>
          <v-spacer />
          <v-btn
            v-for="(link, i) in links"
            :key="i"
            :to="link.to"
            class="ml-0 hidden-sm-and-down"
            text
          >
            {{ link.text }}
          </v-btn>
        </v-layout>
      </v-container>
    </v-app-bar>

    <v-main id="home">
      <material-section>
        <v-row justify="center">
          <material-card
            class="pa-6"
            min-width="40%"
            color="error"
          >
            <template v-slot:heading>
              <div class="display-2 font-weight-light">
                Error!
              </div>

              <div class="subtitle-1 font-weight-light">
                Page not found
              </div>
            </template>
            <v-card-text class="body-1 font-weight-light">
              Looks like you're in uncharted territory. We don't know about this page yet.
              Luckily, we know the way back.
            </v-card-text>
          </material-card>
        </v-row>
      </material-section>
    </v-main>

    <v-footer
      id="dashboard-core-footer"
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
              :href="link.to"
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
  </v-app>
</template>

<script>
import MaterialSection from '@@/components/base/MaterialSection'
import MaterialCard from '@@/components/base/MaterialCard'
export default {
  name: 'Error',
  components: { MaterialSection, MaterialCard },
  data: () => ({
    drawer: true
  }),
  computed: {
    links () {
      if (this.$store.getters['auth/isAuthenticated']) {
        return [
          {
            text: 'Dashboard',
            to: '/dashboard/home/'
          },
          {
            text: 'Logout',
            to: '/auth/logout/'
          }
        ]
      }
      return [
        {
          text: 'Login',
          to: '/auth/login/'
        },
        {
          text: 'Register',
          to: '/auth/register/'
        }
      ]
    }
  },
  head () {
    return {
      title: 'Page not found'
    }
  }
}
</script>
