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
            Reset Password
          </div>

          <div class="subtitle-1 font-weight-light">
            Fill in your new password
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="Your password has been updated successfully!"
        />
        <v-form
          ref="form"
          method="post"
          class="pt-8"
          @submit.prevent="resetPassword"
        >
          <v-container class="py-0">
            <v-row>
              <v-text-field
                v-model="password"
                label="Password*"
                type="password"
                outlined
                :rules="[$validator.required,$validator.password]"
              />
            </v-row>
            <v-row>
              <v-text-field
                v-model="passwordRepeat"
                label="Repeat Password*"
                type="password"
                outlined
                :rules="[$validator.required,$validator.password, $validator.repeatPassword(password)]"
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
                Reset Password
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
  name: 'ResetPassword',
  components: { Alert, MaterialCard },
  layout: 'Home',
  data () {
    return {
      success: false,
      password: '',
      passwordRepeat: '',
      errors: null
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Reset Password')
  },
  methods: {
    resetPassword () {
      this.clearAlert()
      if (!this.$refs.form.validate()) {
        return
      }
      this.$store.dispatch('loading/start')
      UserAPI.resetPassword({
        password: this.hash(this.password),
        token: this.$route.query.t,
        email: this.$route.query.email
      })
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
      this.password = ''
      this.passwordRepeat = ''
    },
    clearAlert () {
      this.success = false
      this.errors = null
    }
  }
}
</script>
