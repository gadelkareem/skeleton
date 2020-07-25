<template>
  <v-container
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card
        :avatar="user.avatar_url"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Edit User
          </div>

          <div class="subtitle-1 font-weight-light">
            Complete your User
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-text="Successful!"
        />
        <v-form
          v-if="user.id"
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
                  label="User Name"
                  :rules="[$validator.required,$validator.username]"
                />
              </v-col>
              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  v-model.trim="user.password"
                  label="Password"
                  type="password"
                  class="purple-input"
                  :rules="[$validator.password]"
                />
              </v-col>
              <v-col
                cols="12"
                md="8"
              >
                <v-text-field
                  v-model.trim="user.email"
                  :rules="[$validator.required, $validator.email]"
                  label="Email Address"
                  class="purple-input"
                />
              </v-col>

              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  v-model.trim="user.first_name"
                  :rules="[$validator.name]"
                  label="First Name"
                  class="purple-input"
                />
              </v-col>

              <v-col
                cols="12"
                md="6"
              >
                <v-text-field
                  v-model.trim="user.last_name"
                  :rules="[$validator.name]"
                  label="Last Name"
                  class="purple-input"
                />
              </v-col>

              <v-col cols="12">
                <v-text-field
                  v-model.trim="user.address.street"
                  label="Street Address"
                  class="purple-input"
                />
              </v-col>

              <v-col
                cols="12"
                md="4"
              >
                <v-text-field
                  v-model.trim="user.address.city"
                  label="City"
                  class="purple-input"
                />
              </v-col>

              <v-col
                cols="12"
                md="4"
              >
                <v-text-field
                  v-model.trim="user.country"
                  label="City"
                  class="purple-input"
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
                md="8"
              >
                <v-text-field
                  v-model.trim="user.avatar_url"
                  label="Avatar URL"
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
                cols="12"
                md="4"
              >
                <v-checkbox
                  v-model="user.active"
                  label="Active"
                />
              </v-col>
              <v-col
                cols="12"
                md="4"
              >
                <v-checkbox
                  v-model="user.authenticator_enabled"
                  label="Authenticator Enabled"
                />
              </v-col>
              <v-col
                cols="6"
              >
                <v-btn
                  color="error"
                  class="mr-0"
                  :loading="$store.state.loading.status"
                  @click="deleteUser"
                >
                  Delete User
                </v-btn>
              </v-col>
              <v-col
                cols="6"
                class="text-right"
              >
                <v-btn
                  color="info"
                  class="mr-0"
                  type="submit"
                  :loading="$store.state.loading.status"
                >
                  Update User
                </v-btn>
              </v-col>
            </v-row>
          </v-container>
        </v-form>
      </material-card>
    </v-row>
    <confirm-modal ref="confirm" />
  </v-container>
</template>

<script>
import UserAPI from '@@/api/user'
import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'
import ConfirmModal from '@@/components/base/ConfirmModal'

export default {
  components: { ConfirmModal, Alert, MaterialCard },
  layout: 'Dashboard',
  data () {
    return {
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
      ],
      user: this.initData().user
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Edit User')
    this.initUser()
  },
  methods: {
    initUser () {
      if (!this.$route.query.id) {
        this.errors = [{ title: 'Invalid user ID' }]
      }
      UserAPI.get(this.$route.query.id)
        .then((r) => {
          this.user = r.data
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    submit () {
      if (!this.$refs.form.validate()) {
        return
      }
      this.clearAlert()
      this.user.password = this.hash(this.user.password)
      UserAPI.update(this.user)
        .then(() => {
          this.success = true
          this.user.password = ''
        })
        .catch((err) => {
          this.errors = this.parseError(err)
          this.user.password = ''
        }).finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    deleteUser () {
      this.$refs.confirm.open(
        'Delete Account',
        'Are you sure you want to delete ' + this.user.username + '?'
      ).then((confirm) => {
        if (confirm) {
          this.clearAlert()
          UserAPI.delete(this.user.id)
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
