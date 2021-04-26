<template>
  <v-snackbar
    v-model="snackbar"
    :timeout="-1"
    class="notification"
  >
    <div class="action">
      {{ msgNotif }}
    </div>
    <v-btn
      outlined
      color="white"
      class="button"
      @click="accept"
    >
      {{ $t('common.accept') }}
    </v-btn>
  </v-snackbar>
</template>

<style lang="sass" scoped>
@import './notification-style'
</style>

<script>
import brand from '@@/static/text/brand'

export default {
  name: 'Notification',
  data () {
    return {
      snackbar: false,
      msgNotif: brand.starter.notifMsg
    }
  },
  mounted () {
    if (!this.accepted()) {
      this.snackbar = true
    }
  },
  methods: {
    accept () {
      this.setCookie('accept_cookie', true)
      this.snackbar = false
    },
    accepted () {
      return this.getCookie('accept_cookie') !== null
    }
  }
}
</script>
