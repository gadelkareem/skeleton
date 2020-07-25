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
            Forgot Password
          </div>

          <div class="subtitle-1 font-weight-light">
            Fill in your email address or user name
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="An email with password reset instructions was sent to your email box!"
        />
        <v-form
          ref="form"
          method="post"
          class="pt-8"
          @submit.prevent="submit"
        >
          <v-container class="py-0">
            <v-row>
              <v-text-field
                v-model.trim="emailOrUsername"
                label="Email or user name*"
                outlined
                type="text"
                :rules="[$validator.required]"
              />
            </v-row>
            <v-spacer />
            <v-col class="text-right">
              <v-btn
                color="blue darken-1"
                text
                :loading="$store.state.loading.status"
                @click="reset()"
              >
                Reset
              </v-btn>
              <v-btn
                color="info"
                class="mr-0"
                type="submit"
                :loading="$store.state.loading.status"
              >
                Send Password
              </v-btn>
            </v-col>
          </v-container>
        </v-form>
      </material-card>
    </v-row>
  </v-container>
</template>

<script>
import UserAPI from '@@/api/user'
import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'

export default {
  name: 'ForgotPassword',
  components: { Alert, MaterialCard },
  layout: 'Home',
  data () {
    return {
      success: false,
      emailOrUsername: '',
      errors: null
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Forgot Password')
  },
  methods: {
    submit () {
      this.clearAlert()
      if (!this.$refs.form.validate()) {
        return
      }
      let username, email
      if (this.emailOrUsername.includes('@')) {
        email = this.emailOrUsername
      } else {
        username = this.emailOrUsername
      }
      this.$store.dispatch('loading/start')
      UserAPI.forgotPassword({ username, email })
        .then(() => {
          this.success = true
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    reset () {
      this.clearAlert()
      this.emailOrUsername = ''
    },

    clearAlert () {
      this.success = false
      this.errors = null
    }
  }
}
</script>
