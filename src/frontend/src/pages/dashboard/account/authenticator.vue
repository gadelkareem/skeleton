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
            2-step verification
          </div>
          <div class="subtitle-1 font-weight-light">
            <span v-if="user.authenticator_enabled">Disable</span><span v-else>Enable</span> Authenticator
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          :info="info"
          :success-text="successText"
        />
        <v-row v-if="!user.recovery_questions_set" align="center" justify="center" class="mt-5">
          <p class="display-1 ma-10">
            Please <router-link to="/dashboard/account/recovery-questions/">
              set your recovery questions
            </router-link> to be able to use this feature.
          </p>
        </v-row>
        <v-row v-else-if="!showForm" align="center" justify="center">
          <v-btn
            color="info"
            class="mr-0 mt-4"
            type="submit"
            @click="showForms"
          >
            <span v-if="user.authenticator_enabled">Remove</span><span v-else>Add</span>&nbsp;Authenticator
          </v-btn>
        </v-row>
        <template v-else>
          <v-container v-if="!user.authenticator_enabled && img">
            <p class="py-4">
              Install an authenticator app on your mobile device if you don't already have one.<br>
              Scan QR code with the authenticator (or tap it in mobile browser).<br><br>
            </p>
            <v-row align="center" justify="center">
              <a :href="url">
                <v-img :src="img" max-width="200" max-height="200" />
              </a>
            </v-row>
            <v-row align="center" justify="center">
              <v-btn icon class="icon_btn" @click="generateAuthenticator(true)">
                <v-icon>mdi-refresh</v-icon>
              </v-btn>
              <v-btn icon class="icon_btn" @click="showSeed = !showSeed">
                <v-icon>mdi-eye</v-icon>
              </v-btn>
            </v-row>
            <v-container v-if="showSeed">
              <v-row align="center" justify="center">
                <v-text-field v-model="seed" disabled label="Seed" />
              </v-row>
            </v-container>
          </v-container>
          <v-container v-else-if="user.authenticator_enabled">
            <p class="py-4">
              Enter the authenticator code to disable 2-step verification on your account.<br>
            </p>
          </v-container>
          <v-form
            ref="form"
            method="post"
            @submit.prevent="submit"
          >
            <v-container class="py-0">
              <v-row align="center" justify="center">
                <v-text-field
                  v-model="code"
                  label="Code*"
                  type="number"
                  class="purple-input shrink"
                  :rules="[$validator.required,$validator.mfa]"
                  maxlength="6"
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
                  <span v-if="user.authenticator_enabled">Disable</span><span v-else>Enable</span>&nbsp;Authenticator
                </v-btn>
              </v-col>
            </v-container>
          </v-form>
        </template>
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
    successText: '',
    info: '',
    code: null,
    img: '',
    seed: null,
    password: null,
    url: null,
    showSeed: false,
    showForm: false

  }),
  computed: {
    user: {
      get () {
        return this.$store.getters['user/user']
      }
    }
  },
  mounted () {
    this.$store.dispatch('page/title', '2-step verification')
  },
  methods: {
    submit () {
      if (!this.$refs.form.validate()) {
        return
      }
      this.clearAlert()
      const enable = !this.user.authenticator_enabled
      UserAPI.authenticator(this.user.id, { enable, code: this.code })
        .then(() => {
          this.success = true
          this.successText = '2-step verification ' + (enable ? 'enabled' : 'disabled') + ' successfully!'
          if (enable) {
            this.successText += '|Your 2-step seed is: ' + this.seed
          }
          this.code = null
          this.$store.dispatch('user/fetchUser', this.user.id)
            .then(() => {
              this.showForm = false
            })
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    showForms () {
      this.showForm = true
      if (!this.user.authenticator_enabled) {
        this.generateAuthenticator()
      }
    },
    generateAuthenticator (refresh) {
      if (!this.user.id || this.user.authenticator_enabled) {
        return
      }
      UserAPI.generateAuthenticator(this.user.id, { enable: true, refresh: !!refresh })
        .then((r) => {
          this.img = 'data:image/png;base64,' + r.data.image
          this.seed = r.data.seed
          this.password = r.data.password
          this.url = r.data.url
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    clearAlert () {
      this.success = false
      this.errors = []
      this.$store.dispatch('loading/start')
    }
  }
}
</script>
