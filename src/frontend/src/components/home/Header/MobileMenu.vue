<template>
  <fragment>
    <!--    <v-list dense>-->
    <!--      <v-list-item-->
    <!--        v-for="(item, index) in navMenu"-->
    <!--        :key="index"-->
    <!--        :href="'#' + item"-->
    <!--        link-->
    <!--      >-->
    <!--        <v-list-item-content>-->
    <!--          <v-list-item-title class="menu-list">-->
    <!--            {{ item }}-->
    <!--          </v-list-item-title>-->
    <!--        </v-list-item-content>-->
    <!--      </v-list-item>-->
    <!--    </v-list>-->
    <!--    <v-divider />-->
    <v-list dense>
      <template v-if="isLoggedIn">
        <v-list-item
          href="/dashboard/home/"
          link
        >
          <v-list-item-content>
            <v-list-item-title class="menu-list">Dashboard</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
      <v-list-item
        href="/home/contact/"
        link
      >
        <v-list-item-content>
          <v-list-item-title class="menu-list">Contact</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-divider />
      <template v-if="isLoggedIn">
        <v-list-item
          href="/auth/logout/"
          link
        >
          <v-list-item-content>
            <v-list-item-title class="menu-list">Logout</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
      <template v-else>
        <v-list-item
          href="/auth/login/"
          link
        >
          <v-list-item-content>
            <v-list-item-title class="menu-list">Login</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item
          href="/auth/register/"
          link
        >
          <v-list-item-content>
            <v-list-item-title class="menu-list">Register</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
  </fragment>
</template>

<style scoped lang="sass">
@import './header-style'
</style>

<script>
import navMenu from './menu'

export default {
  name: 'MobileMenu',
  data () {
    return {
      navMenu,
      isOpen: false
    }
  },
  computed: {
    isLoggedIn () {
      return this.$store.getters['auth/isAuthenticated']
    }
  }
}
</script>
