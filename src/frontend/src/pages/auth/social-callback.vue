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
            Successfully logged in via <span class="text-capitalize">{{ provider }}</span>
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="You have successfully logged in!"
        />
        <v-alert v-if="!success && !errors" type="info">
          Please wait while logging you into the system...
        </v-alert>
        <v-alert v-if="errors && errors[0].status === '422'" type="error">
          Please <router-link to="/auth/login/">
            login
          </router-link> using your account username and password...
        </v-alert>
      </material-card>
    </v-row>
  </v-container>
</template>

<script>

import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'
export default {
  name: 'VerifyEmail',
  components: { Alert, MaterialCard },
  layout: 'Home',
  data () {
    return {
      provider: this.$route.query.p,
      success: false,
      errors: null
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Social Login')
    this.clearAlert()
    const $this = this
    this.$nextTick(() => {
      $this.$store.dispatch('loading/start')
    })
    this.$store.dispatch('auth/socialCallback', {
      code: this.$route.query.code,
      state: this.$route.query.state
    })
      .then(() => {
        this.success = true
        this.$router.push('/dashboard/home/')
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
