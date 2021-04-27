<template>
  <v-app>
    <v-main
      id="main-wrap"
    >
      <main-header />
      <nuxt />
    </v-main>
    <main-footer app absolute />
  </v-app>
</template>

<script>
import Footer from '@@/components/core/Footer/Footer'
import Header from '@@/components/home/Header/Header'

export default {
  name: 'Home',
  loading: false,
  components: {
    'main-footer': Footer,
    'main-header': Header
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

<style lang="sass">
@import @@/assets/sass/home_overrides.sass
@import @@/assets/sass/home/vendors/animate.css
@import @@/assets/sass/home/vendors/animate-extends.css
@import @@/assets/sass/home/vendors/hamburger-menu.css
@import @@/assets/sass/home/vendors/slick-carousel/slick.css
@import @@/assets/sass/home/vendors/slick-carousel/slick-theme.css
</style>
