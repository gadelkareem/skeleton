<template>
  <v-app-bar
    id="app-bar"
    absolute
    app
    color="transparent"
    flat
    height="75"
  >
    <v-btn
      class="mr-3"
      elevation="1"
      fab
      small
      @click="setDrawer(!drawer)"
    >
      <v-icon v-if="value">
        mdi-view-quilt
      </v-icon>

      <v-icon v-else>
        mdi-dots-vertical
      </v-icon>
    </v-btn>

    <v-toolbar-title
      class="hidden-sm-and-down font-weight-light"
      v-text="$store.state.page.title"
    />

    <v-spacer />

    <!--    <v-text-field-->
    <!--      :label="$t('search')"-->
    <!--      color="secondary"-->
    <!--      hide-details-->
    <!--      style="max-width: 165px;"-->
    <!--    >-->
    <!--      <template-->
    <!--        v-if="$vuetify.breakpoint.mdAndUp"-->
    <!--        v-slot:append-outer-->
    <!--      >-->
    <!--        <v-btn-->
    <!--          class="mt-n2"-->
    <!--          elevation="1"-->
    <!--          fab-->
    <!--          small-->
    <!--        >-->
    <!--          <v-icon>mdi-magnify</v-icon>-->
    <!--        </v-btn>-->
    <!--      </template>-->
    <!--    </v-text-field>-->

    <div class="mx-3" />

    <!--    <v-btn-->
    <!--      class="ml-2"-->
    <!--      min-width="0"-->
    <!--      text-->
    <!--      to="/"-->
    <!--    >-->
    <!--      <v-icon>mdi-view-dashboard</v-icon>-->
    <!--    </v-btn>-->

    <!--    Notifications -->
    <v-menu
      bottom
      left
      offset-y
      origin="top right"
      transition="scale-transition"
    >
      <template v-slot:activator="{ attrs, on }">
        <v-btn
          class="ml-2"
          min-width="0"
          text
          v-bind="attrs"
          v-on="on"
        >
          <v-badge
            color="red"
            overlap
            bordered
            :content="unreadNotifications"
            :value="unreadNotifications"
          >
            <!--            <template v-slot:badge>-->
            <!--              <span>{{ notifications.length }}</span>-->
            <!--            </template>-->

            <v-icon color="gray">mdi-bell</v-icon>
          </v-badge>
        </v-btn>
      </template>

      <v-list
        v-if="notifications.length"
        :tile="false"
        nav
        width="300"
      >
        <div>
          <v-list-item
            v-for="(n, i) in notifications"
            :key="`item-${i}`"
            @click="read(n)"
          >
            <v-list-item-content>
              <v-list-item-title class="text-wrap" v-text="n.message" />
            </v-list-item-content>
            <v-list-item-icon>
              <v-icon :color="n.read_receipt_at ? '#cdcdcd': 'blue'">mdi-circle</v-icon>
            </v-list-item-icon>
          </v-list-item>
        </div>
      </v-list>
    </v-menu>

    <account-menu />
  </v-app-bar>
</template>

<script>

import UserAPI from '@@/api/user'

export default {
  name: 'DashboardCoreAppBar',

  components: {
    AccountMenu: () => import('./AccountMenu')
  },

  props: {
    value: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    drawer () {
      return this.$store.getters['dashboard/drawer']
    },
    user () {
      return this.$store.getters['user/user']
    },
    notifications () {
      let n = this.user.notifications || []
      n = n.sort((a, b) => (a.created_at > b.created_at) ? -1 : ((b.created_at > a.created_at) ? 1 : 0))
      return n || []
    },
    unreadNotifications () {
      return this.notifications.filter(n => !n.read_receipt_at).length
    }
  },
  methods: {
    setDrawer (v) {
      this.$store.commit('dashboard/SET_DRAWER', v)
    },
    open (n) {
      if (n.url.includes('://')) { location.href = n.url } else { this.$router.push(n.url) }
    },
    read (n) {
      if (n.read_receipt_at) {
        return
      }
      UserAPI.readNotification(this.user.id, {
        id: n.id
      })
        .then((r) => {
          this.$store.dispatch('user/fetchUser', this.user.id)
        })
        .catch((err) => {
          console.log(this.parseError(err))
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
          this.open(n)
        })
    }
  }
}
</script>
