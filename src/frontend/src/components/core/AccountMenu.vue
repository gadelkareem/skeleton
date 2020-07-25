<template>
  <v-menu offset-y open-on-click>
    <template v-slot:activator="{ on }">
      <v-btn
        icon
        tile
        v-on="on"
      >
        <v-avatar
          size="40"
          class="elevation-3"
          color="#013163"
          tile
        >
          <v-img
            v-if="user.avatar_url && !brokenAvatar"
            :src="avatarURL"
            @error="avatarError"
          />
          <span v-else class="white--text headline font-weight-black">{{ firstLetterUpper(user.username) }}</span>
        </v-avatar>
      </v-btn>
    </template>
    <v-list>
      <v-list-item to="/dashboard/account/update-profile/">
        <v-list-item-icon>
          <v-icon>mdi-account</v-icon>
        </v-list-item-icon>
        Profile
      </v-list-item>
      <v-list-item href="/auth/logout/">
        <v-list-item-icon>
          <v-icon>mdi-logout</v-icon>
        </v-list-item-icon>
        Logout
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script>

export default {
  name: 'AccountMenu',
  data: () => ({
    brokenAvatar: false
  }),
  computed: {
    user: {
      get () {
        return this.$store.getters['user/user']
      }
    },
    avatarURL () {
      if (this.user.avatar_url.includes('www.gravatar.com')) {
        return this.user.avatar_url + '?s=150&d=404'
      }
      return this.user.avatar_url
    }
  },
  methods: {
    avatarError () {
      this.brokenAvatar = true
      this.$emit('error', false)
    }
  }
}
</script>
