<template>
  <v-app dark>
    <main-header />
    <v-main
      id="main-wrap"
    >
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
      <br>
      <br>
    </v-main>
    <main-footer />
  </v-app>
</template>

<script>
import Footer from '@@/components/core/Footer/Footer'
import Header from '@@/components/home/Header/Header'
import MaterialSection from '@@/components/base/MaterialSection'
import MaterialCard from '@@/components/base/MaterialCard'

export default {
  name: 'Error',
  loading: false,
  components: {
    'main-footer': Footer,
    'main-header': Header,
    MaterialSection,
    MaterialCard
  },
  data () {
    return {
      show: false,
      drawer: true
    }
  },
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
  mounted () {
    // Preloader and Progress bar setup
    this.show = true
    this.$nextTick(() => {
      setTimeout(() => this.$nuxt.$loading.finish(), 500)
      this.$nuxt.$loading.start()
    })
    const preloader = document.getElementById('preloader')
    if (preloader !== null || undefined) {
      preloader.remove()
    }
    // RTL initial
    const rtlURL = document.location.pathname.split('/')[1] === 'ar'
    this.$vuetify.rtl = rtlURL
  },
  methods: {
    changeColor () {
      this.$vuetify.theme.themes = {
        light: {
          primary: '#00af4a',
          secondary: '#ff2020'
        }
      }
    }
  },
  head () {
    return {
      title: this.$store.getters['page/title']
    }
  }
}
</script>
