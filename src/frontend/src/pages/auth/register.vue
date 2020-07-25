<template>
  <v-container
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card
        class="pa-6"
        min-width="40%"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Register
          </div>

          <div class="subtitle-1 font-weight-light">
            Fill in your details
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="Successfully registered!| Check your email for instructions on how to activate your account."
        />
        <v-form
          ref="form"
          method="post"
          class="pt-8"
          @submit.prevent="register"
        >
          <v-container class="py-0">
            <v-row>
              <v-text-field
                v-model="username"
                label="Username*"
                type="text"
                outlined
                :rules="[$validator.required,$validator.username]"
              />
            </v-row>
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
            <v-row>
              <v-text-field
                v-model="email"
                label="Email*"
                type="email"
                outlined
                :rules="[$validator.required, $validator.email]"
              />
            </v-row>
            <v-col
              class="text-right"
            >
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
                :loading="this.$store.state.loading.status"
              >
                Register
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
  components: { Alert, MaterialCard },
  layout: 'Home',
  data () {
    return {
      icons: {
        close: ''
      },
      success: false,
      username: '',
      password: '',
      passwordRepeat: '',
      email: '',
      errors: null
    }
  },
  created () {
    if (this.$store.getters['auth/isAuthenticated']) {
      this.redirect()
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Register')
  },
  methods: {
    goHome (v) {
      if (!v) {
        this.$router.push('/dashboard/home/')
      }
    },
    clearAlert () {
      this.success = false
      this.errors = null
    },
    register () {
      this.clearAlert()
      if (!this.$refs.form.validate()) {
        return
      }
      this.$store.dispatch('loading/start')
      const user = {
        username: this.username,
        password: this.hash(this.password),
        email: this.email
      }
      UserAPI.register(user)
        .then(() => {
          this.success = true
        })
        .catch((errs) => {
          this.errors = this.parseError(errs)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    redirect () {
      const redirect = this.$route.query.redirect
      if (typeof redirect !== 'undefined') {
        this.$router.push({ path: redirect })
      } else {
        this.$router.push('/dashboard/home/')
      }
    },
    reset () {
      this.clearAlert()
      this.username = ''
      this.password = ''
      this.passwordRepeat = ''
      this.email = ''
      this.$refs.form.reset()
    }
  }
}
</script>
