<template>
  <v-container
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card
        class="pa-6"
        min-width="50%"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Logging out..
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="You have successfully logged out!"
        />
        <v-alert v-if="!success && !errors" type="info">
          Please wait while logging you out of the system.
        </v-alert>
      </material-card>
    </v-row>
  </v-container>
</template>

<script>

import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'
export default {
  name: 'Logout',
  components: { Alert, MaterialCard },
  layout: 'Home',
  data () {
    return {
      success: false,
      errors: null
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Logout')
  },
  created () {
    this.clearAlert()
    const $this = this
    this.$nextTick(() => {
      $this.$store.dispatch('loading/start')
    })
    this.$store.dispatch('auth/logout')
      .then(() => {
        this.success = true
        this.$store.dispatch('auth/removeSession')
        this.$store.dispatch('reset')
        if (process.browser) {
          window.location.href = '/'
        }
      })
      .catch((err) => {
        this.errors = this.parseError(err)
      })
      .finally(() => {
        this.$store.dispatch('loading/finish')
      })
  },
  methods: {
    clearAlert () {
      this.success = false
      this.errors = null
    }
  }
}
</script>
