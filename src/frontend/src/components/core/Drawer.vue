<template>
  <v-navigation-drawer
    id="core-navigation-drawer"
    v-model="drawer"
    :dark="barColor !== 'rgba(228, 226, 226, 1), rgba(255, 255, 255, 0.7)'"
    :right="$vuetify.rtl"
    :src="barImage"
    mobile-break-point="960"
    app
    width="260"
  >
    <template v-slot:img="props">
      <v-img
        :gradient="`to bottom, ${barColor}`"
        v-bind="props"
      />
    </template>

    <v-divider class="mb-1" />

    <v-list
      dense
      nav
    >
      <v-list-item>
        <v-list-item-content>
          <router-link to="/">
            <v-img
              :src="require('@@/static/logo-dark.svg')"
              contain
              height="42"
              to="/"
              position="left"
              class="mt-2"
            />
          </router-link>
        </v-list-item-content>
      </v-list-item>
    </v-list>

    <v-divider class="mb-2" />

    <v-list
      expand
      nav
    >
      <template v-for="(item, i) in computedItems">
        <item-group
          v-if="item.children"
          :key="`group-${i}`"
          :item="item"
        >
          <!--  -->
        </item-group>

        <item
          v-else
          :key="`item-${i}`"
          :item="item"
        />
      </template>

      <!-- Style cascading bug  -->
      <!-- https://github.com/vuetifyjs/vuetify/pull/8574 -->
      <div />
    </v-list>

    <template v-slot:append>
      <item
        :item="{
          title: $t('logout'),
          icon: 'mdi-package-up',
          to: '/auth/logout/',
        }"
      />
    </template>
  </v-navigation-drawer>
</template>

<script>
import ItemGroup from '../base/ItemGroup'
import Item from '../base/Item'
export default {
  name: 'DashboardCoreDrawer',
  components: { Item, ItemGroup },
  data: () => ({
    items: [
      {
        icon: 'mdi-view-dashboard',
        title: 'dashboard',
        to: '/dashboard/home/'
      },
      {
        icon: 'mdi-account',
        title: 'Account',
        group: 'account',
        children: [
          {
            title: 'Update Profile',
            to: '/dashboard/account/update-profile/'
          },
          {
            title: 'change password',
            to: '/dashboard/account/change-password/'
          },
          {
            title: 'Recovery Questions',
            to: '/dashboard/account/recovery-questions/'
          },
          {
            title: '2-step Verification',
            to: '/dashboard/account/authenticator/'
          },
          {
            title: 'Mobile Verification',
            to: '/dashboard/account/verify-mobile/'
          },
          {
            title: 'Delete Account',
            to: '/dashboard/account/delete-account/'
          }
        ]
      },
      {
        icon: 'mdi-shield-check',
        title: 'Admin',
        group: 'admin',
        children: [
          {
            title: 'Users',
            icon: 'mdi-account-group',
            to: '/dashboard/admin/users/'
          },
          {
            title: 'Audit Logs',
            icon: 'mdi-account-group',
            to: '/dashboard/admin/logs/'
          }
        ]
      },
      {
        title: 'rtables',
        icon: 'mdi-clipboard-outline',
        to: '/dashboard/tables/regular-tables/'
      },
      {
        title: 'typography',
        icon: 'mdi-format-font',
        to: '/dashboard/component/typography/'
      },
      {
        title: 'icons',
        icon: 'mdi-chart-bubble',
        to: '/dashboard/component/icons/'
      },
      {
        title: 'google',
        icon: 'mdi-map-marker',
        to: '/dashboard/maps/google-maps/'
      },
      {
        title: 'notifications',
        icon: 'mdi-bell',
        to: '/dashboard/component/notifications/'
      }
    ]
  }),
  computed: {
    barColor: {
      get () {
        return this.$store.getters['dashboard/barColor']
      }
    },
    barImage: {
      get () {
        return this.$store.getters['dashboard/barImage']
      }
    },
    drawer: {
      get () {
        return this.$store.getters['dashboard/drawer']
      },
      set (v) {
        this.$store.commit('dashboard/SET_DRAWER', v)
      }
    },
    computedItems () {
      return this.items.map(this.mapItem).filter(this.filter)
    },
    profile () {
      return {
        avatar: true,
        title: this.$t('avatar')
      }
    }
  },
  methods: {
    mapItem (item) {
      if (this.canAccessRoute(item.to)) {
        return {
          ...item,
          children: item.children ? item.children.map(this.mapItem).filter(this.filter) : undefined,
          title: this.$t(item.title)
        }
      }
    },
    filter (item) {
      return item && (item.to || (item.group && item.children.length > 0))
    }
  }
}
</script>

<style lang="sass">
  @import '~vuetify/src/styles/tools/_rtl.sass'

  #core-navigation-drawer
    .v-list-group__header.v-list-item--active:before
      opacity: .24

    .v-list-item
      &__icon--text,
      &__icon:first-child
        justify-content: center
        text-align: center
        width: 20px

        +ltr()
          margin-right: 24px
          margin-left: 12px !important

        +rtl()
          margin-left: 24px
          margin-right: 12px !important

    .v-list--dense
      .v-list-item
        &__icon--text,
        &__icon:first-child
          margin-top: 10px

    .v-list-group--sub-group
      .v-list-item
        +ltr()
          padding-left: 8px

        +rtl()
          padding-right: 8px

      .v-list-group__header
        +ltr()
          padding-right: 0

        +rtl()
          padding-right: 0

        .v-list-item__icon--text
          margin-top: 19px
          order: 0

        .v-list-group__header__prepend-icon
          order: 2

          +ltr()
            margin-right: 8px

          +rtl()
            margin-left: 8px
</style>
