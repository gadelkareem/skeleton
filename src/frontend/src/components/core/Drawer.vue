<template>
  <v-navigation-drawer
    id="core-navigation-drawer"
    v-model="drawer"
    :dark="barColor !== 'rgba(228, 226, 226, 1), rgba(255, 255, 255, 0.7)'"
    :right="$vuetify.rtl"
    :src="barImage"
    mobile-break-point="960"
    app
    width="auto"
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
          <a class="logo" href="/">
            <v-img
              :src="require('@@/static/logo-dark.svg')"
              contain
              position="left"
            />
          </a>
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
            to: '/dashboard/admin/users/'
          },
          {
            title: 'Audit Logs',
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
