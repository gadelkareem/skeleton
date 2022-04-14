<template>
  <v-container
    id="user-profile"
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card
        :avatar="user.avatar_url"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Edit Profile
          </div>

          <div class="subtitle-1 font-weight-light">
            Complete your profile
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-txt="Successful!"
        />
        <v-form
          ref="form"
          method="post"
          @submit.prevent="submit"
        >
          <v-container class="py-0">
            <v-row>
              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  v-model.trim="user.username"
                  disabled
                  label="User Name"
                />
              </v-col>
              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  class="purple-input"
                  disabled
                  label="Password"
                  type="password"
                  value="password"
                />
                <v-btn
                  color="blue darken-1"
                  small
                  text
                  to="/dashboard/account/change-password/"
                >
                  Change password
                </v-btn>
              </v-col>
              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  v-model.trim="user.email"
                  :rules="[$validator.required, $validator.email]"
                  class="purple-input"
                  label="Email Address"
                />
              </v-col>
              <v-col
                cols="12"
                md="4"
              >
                <v-text-field
                  ref="mobile"
                  v-model.trim="user.mobile"
                  :rules="[$validator.mobile]"
                  class="purple-input"
                  label="Mobile"
                />
                <v-btn
                  v-if="user.mobile && !user.mobile_verified"
                  color="blue darken-1"
                  small
                  text
                  to="/dashboard/account/verify-mobile/"
                >
                  Verify your mobile number
                </v-btn>
              </v-col>

              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  v-model.trim="user.first_name"
                  :rules="[$validator.name]"
                  class="purple-input"
                  label="First Name"
                />
              </v-col>

              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  v-model.trim="user.last_name"
                  :rules="[$validator.name]"
                  class="purple-input"
                  label="Last Name"
                />
              </v-col>

              <v-col cols="12">
                <v-text-field
                  v-model.trim="user.address.street"
                  class="purple-input"
                  label="Street Address"
                />
              </v-col>

              <v-col
                cols="12"
                md="4"
              >
                <v-text-field
                  v-model.trim="user.address.city"
                  class="purple-input"
                  label="City"
                />
              </v-col>

              <v-col
                cols="12"
                md="4"
              >
                <v-select
                  v-model="user.country"
                  :items="countries"
                  class="purple-input"
                  item-text="name"
                  item-value="code"
                  label="Country"
                />
              </v-col>

              <v-col
                cols="12"
                md="4"
              >
                <v-text-field
                  v-model.trim="user.address.zip_code"
                  class="purple-input"
                  label="Postal Code"
                />
              </v-col>
              <v-col
                cols="12"
                md="4"
              >
                <v-select
                  v-model="user.language"
                  :items="languages"
                  item-text="name"
                  item-value="code"
                  label="Language"
                />
              </v-col>
              <v-col
                class="text-right"
                cols="12"
              >
                <v-btn
                  :loading="$store.state.loading.status"
                  class="mr-0"
                  color="info"
                  type="submit"
                >
                  Update Profile
                </v-btn>
              </v-col>
            </v-row>
          </v-container>
        </v-form>
      </material-card>
    </v-row>
  </v-container>
</template>

<script>
import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'
export default {
  components: { Alert, MaterialCard },
  layout: 'Dashboard',
  data: () => ({
    errors: [],
    success: false,
    languages: [
      { code: '*', name: 'All' },
      { code: 'AF', name: 'Afrikaans' },
      { code: 'AR', name: 'Arabic' },
      { code: 'BG', name: 'Bulgarian' },
      { code: 'CA', name: 'Catalan' },
      { code: 'CS', name: 'Czech' },
      { code: 'DA', name: 'Danish' },
      { code: 'DE', name: 'German' },
      { code: 'EL', name: 'Greek' },
      { code: 'EN', name: 'English' },
      { code: 'ES', name: 'Spanish' },
      { code: 'ET', name: 'Estonian' },
      { code: 'FI', name: 'Finnish' },
      { code: 'FR', name: 'French' },
      { code: 'GL', name: 'Galician' },
      { code: 'HE', name: 'Hebrew' },
      { code: 'HI', name: 'Hindi' },
      { code: 'HR', name: 'Croatian' },
      { code: 'HU', name: 'Hungarian' },
      { code: 'ID', name: 'Indonesian' },
      { code: 'IT', name: 'Italian' },
      { code: 'JA', name: 'Japanese' },
      { code: 'KA', name: 'Georgian' },
      { code: 'KO', name: 'Korean' },
      { code: 'LT', name: 'Lithuanian' },
      { code: 'LV', name: 'Latvian' },
      { code: 'MS', name: 'Malay' },
      { code: 'NL', name: 'Dutch' },
      { code: 'NO', name: 'Norwegian' },
      { code: 'PL', name: 'Polish' },
      { code: 'PT', name: 'Portuguese' },
      { code: 'RO', name: 'Romanian' },
      { code: 'RU', name: 'Russian' },
      { code: 'SK', name: 'Slovak' },
      { code: 'SL', name: 'Slovenian' },
      { code: 'SQ', name: 'Albanian' },
      { code: 'SR', name: 'Serbian' },
      { code: 'SV', name: 'Swedish' },
      { code: 'TH', name: 'Thai' },
      { code: 'TR', name: 'Turkish' },
      { code: 'UK', name: 'Ukrainian' },
      { code: 'ZH', name: 'Chinese' }
    ]
  }),
  computed: {
    user: {
      get () {
        return this.$store.getters['user/user']
      }
    },
    countries () {
      return this.initData().countries
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Profile')
    if (this.$route.query.action === 'add-mobile') {
      this.$refs.mobile.focus()
    }
  },
  methods: {
    // sanitizeMobile () {
    //   this.user.mobile = this.user.mobile.replaceAll(/[^0-9+]/g, '')
    // },
    submit () {
      if (!this.$refs.form.validate()) {
        return
      }
      this.success = false
      this.errors = []
      this.$store.dispatch('loading/start')
      this.$store.dispatch('user/updateUser', this.user)
        .then(() => {
          this.success = true
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        }).finally(() => {
          this.$store.dispatch('loading/finish')
        })
    }
  }
}
</script>
