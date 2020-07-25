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
            Verify your email address
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="Your email has been successfully verified!"
        />
        <v-alert v-if="!success && !errors" type="info">
          Please wait while verifying your email.
        </v-alert>
      </material-card>
    </v-row>
  </v-container>
</template>

<script>
import UserAPI from '@@/api/user'
import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'

export default {
  name: 'VerifyEmail',
  components: { Alert, MaterialCard },
  layout: 'Home',
  data () {
    return {
      success: false,
      errors: null
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Verify Email')
    this.clearAlert()
    const $this = this
    this.$nextTick(() => {
      $this.$store.dispatch('loading/start')
    })
    UserAPI.verifyEmail({
      token: this.$route.query.t,
      email: this.$route.query.email
    })
      .then(() => {
        this.success = true
      })
      .catch((err) => {
        this.errors = this.parseError(err)
      }).finally(() => {
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
