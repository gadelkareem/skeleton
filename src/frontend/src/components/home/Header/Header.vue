<template>
  <fragment>
    <v-navigation-drawer
      v-show="isTablet"
      v-model="openNavMobile"
      fixed
      temporary
      class="mobile-nav"
    >
      <mobile-menu :is-logged-in="isLoggedIn" />
    </v-navigation-drawer>
    <v-app-bar
      v-scroll="handleScroll"
      :class="{ fixed: fixed }"
      class="header"
      fixed
      app
    >
      <v-container>
        <div class="header-content">
          <nav class="nav-menu">
            <v-btn
              v-if="isTablet"
              text
              icon
              @click.stop="handleToggleOpen"
            >
              <v-icon>mdi-menu</v-icon>
            </v-btn>
            <a class="logo" href="/">
              <v-img
                :src="logo"
                contain
                position="left"
              />
            </a>
            <div v-if="loaded">
              <scrollactive
                v-if="isDesktop"
                :offset="navOffset"
                active-class="active"
                tag="div"
              >
                <v-btn
                  v-for="(item, index) in menuList"
                  :key="index"
                  :href="item.url"
                  class="anchor-link scrollactive-item"
                  text
                  @click="setOffset(item.offset)"
                >
                  {{ item.name }}
                </v-btn>
              </scrollactive>
            </div>
          </nav>
          <nav v-if="isDesktop" class="user-menu">
            <template v-if="isLoggedIn">
              <v-btn text to="/dashboard/home/">Dashboard</v-btn>
              <v-btn color="primary" to="/auth/logout/">Logout</v-btn>
            </template>
            <template v-else>
              <v-btn text to="/auth/login/">Login</v-btn>
              <v-btn color="primary" to="/auth/register/">Register</v-btn>
            </template>
            <v-spacer
              class="vertical-divider"
            />
            <settings />
          </nav>
        </div>
      </v-container>
    </v-app-bar>
  </fragment>
</template>

<style lang="sass" scoped>
@import './header-style'
</style>

<script>
import navMenu from './menu'
import MobileMenu from './MobileMenu'
import Settings from './Settings'

let counter = 0

function createData (name, url, offset) {
  counter += 1
  return {
    id: counter,
    name,
    url,
    offset
  }
}

export default {
  name: 'Header',
  components: {
    MobileMenu,
    Settings
  },
  data () {
    return {
      section: 0,
      fixed: false,
      loaded: false,
      openNavMobile: null,
      navOffset: 20,
      menuList: [
        createData(navMenu[0], '/#' + navMenu[0])
        // createData(navMenu[1], '/#' + navMenu[1]),
        // createData(navMenu[2], '/#' + navMenu[2]),
        // createData(navMenu[3], '/#' + navMenu[3], -40),
        // createData(navMenu[4], '/#' + navMenu[4], -40)
      ]
    }
  },
  computed: {
    isTablet () {
      const mdDown = this.$store.state.breakpoints.mdDown
      return mdDown.includes(this.$mq)
    },
    isDesktop () {
      const smUp = this.$store.state.breakpoints.smUp
      return smUp.includes(this.$mq)
    },
    isLoggedIn () {
      return this.$store.getters['auth/isAuthenticated']
    },
    logo () {
      return require(this.$vuetify.theme.dark ? '@@/static/logo-dark.svg' : '@@/static/logo.svg')
    }
  },
  mounted () {
    this.loaded = true
  },
  methods: {
    handleScroll () {
      this.fixed = (window.scrollY > 100)
      return true
    },
    setOffset (offset) {
      this.navOffset = offset
    },
    handleToggleOpen () {
      this.openNavMobile = !this.openNavMobile
    }
  }
}
</script>
