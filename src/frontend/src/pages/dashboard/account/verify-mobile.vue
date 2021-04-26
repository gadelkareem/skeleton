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
            Mobile Verification
          </div>
          <div class="subtitle-1 font-weight-light">
            Verify your mobile
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="Mobile verified successfully"
        />
        <v-row v-if="!user.recovery_questions_set" align="center" justify="center" class="mt-5">
          <p class="display-1 ma-10">
            Please <router-link to="/dashboard/account/recovery-questions/">
              set your recovery questions
            </router-link> to be able to use this feature.
          </p>
        </v-row>
        <v-row v-else-if="!user.mobile" align="center" justify="center" class="mt-5">
          <p class="display-1 ma-10">
            Please <router-link to="/dashboard/account/update-profile/?action=add-mobile">
              add a mobile number to your account
            </router-link> to be able to use this feature.
          </p>
        </v-row>
        <v-row v-else-if="user.mobile_verified" align="center" justify="center" class="mt-5">
          <p class="display-1">
            Your mobile number is verified.
          </p>
        </v-row>
        <v-row v-else-if="!showForm" align="center" justify="center" class="mt-5">
          <v-btn
            color="info"
            class="mr-0"
            type="submit"
            @click="showForms"
          >
            Send SMS code
          </v-btn>
        </v-row>
        <template v-else>
          <v-container>
            <p class="py-4">
              Enter the SMS code you received.<br>
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
                  Verify
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
    code: null,
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
    this.$store.dispatch('page/title', 'Mobile verification')
  },
  methods: {
    submit () {
      if (!this.$refs.form.validate()) {
        return
      }
      this.clearAlert()
      UserAPI.verifyMobile(this.user.id, { code: this.code })
        .then(() => {
          this.success = true
          this.$store.dispatch('user/fetchUser', this.user.id)
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    showForms () {
      this.clearAlert()
      UserAPI.sendSMS(this.user.id)
        .then(() => {
          this.showForm = true
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
