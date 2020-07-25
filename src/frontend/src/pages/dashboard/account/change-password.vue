<template>
  <v-container
    id="change-password"
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
            Change Password
          </div>
          <div class="subtitle-1 font-weight-light">
            Fill in your current and new password
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          :info="info"
          success-text="Successful!"
        />
        <v-form
          ref="form"
          method="post"
          @submit.prevent="submit"
        >
          <v-container class="py-0">
            <v-row>
              <v-text-field
                v-model="oldPassword"
                label="Current Password*"
                type="password"
                class="purple-input"
                :rules="[$validator.required,$validator.password]"
              />
            </v-row>
            <v-row>
              <v-text-field
                v-model="password"
                label="New Password*"
                type="password"
                class="purple-input"
                :rules="[$validator.required,$validator.password]"
              />
            </v-row>
            <v-row>
              <v-text-field
                v-model="repeatPassword"
                label="Repeat New Password*"
                type="password"
                class="purple-input"
                :rules="[$validator.required,$validator.password, $validator.repeatPassword(password)]"
              />
            </v-row>
            <v-col
              cols="12"
              class="text-right"
            >
              <v-btn
                color="info"
                class="mr-0"
                type="submit"
                :loading="$store.state.loading.status"
              >
                Change Password
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
  layout: 'Dashboard',
  data: () => ({
    errors: [],
    success: false,
    info: '',
    oldPassword: '',
    password: '',
    repeatPassword: ''

  }),
  mounted () {
    this.$store.dispatch('page/title', 'Change Password')
    const user = this.$store.getters['user/user']
    if (user.id && user.social_login) {
      this.info = 'You are logged in using a social network.' +
        '|Reset your password to get a new password for your account.'
    }
  },
  methods: {
    submit () {
      if (!this.$refs.form.validate()) {
        return
      }
      if (this.password === this.oldPassword) {
        this.errors = [{ title: 'Old and new password cannot be the same' }]
        return
      }
      this.success = false
      this.errors = []
      this.$store.dispatch('loading/start')
      UserAPI.changePassword(this.$store.getters['auth/userId'], this.hash(this.oldPassword), this.hash(this.password))
        .then(() => {
          this.success = true
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    }
  }
}
</script>
