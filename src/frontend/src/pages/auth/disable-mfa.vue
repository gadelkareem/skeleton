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
            Disable 2-step verification
          </div>

          <div class="subtitle-1 font-weight-light">
            <template v-if="!showQuestions">
              Fill in your current username and password
            </template>
            <template v-else>
              Answer your security questions
            </template>
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-txt="2-step verification has been disabled successfully!"
        />
        <v-form
          v-if="!showQuestions"
          ref="form"
          method="post"
          class="pt-8"
          @submit.prevent="login"
        >
          <v-container class="py-0">
            <template>
              <v-row>
                <v-text-field
                  v-model.trim="username"
                  label="Username*"
                  type="text"
                  outlined
                  :rules="[$validator.required]"
                />
              </v-row>
              <v-row>
                <v-text-field
                  v-model="password"
                  label="Your Password*"
                  type="password"
                  outlined
                  :rules="[$validator.required,$validator.password]"
                />
              </v-row>
              <v-spacer />
              <v-col class="text-right">
                <v-btn
                  color="info"
                  class="mr-0"
                  type="submit"
                  :loading="$store.state.loading.status"
                >
                  Login
                </v-btn>
              </v-col>
            </template>
          </v-container>
        </v-form>
        <v-form
          v-else
          ref="form"
          method="post"
          class="pt-8"
          @submit.prevent="disableMFA"
        >
          <v-container class="py-0">
            <template>
              <recovery-questions :sets="recoveryQuestions" predefined />
              <v-spacer />
              <v-col class="text-right">
                <v-btn
                  color="info"
                  class="mr-0"
                  text
                  @click="showQuestions = false"
                >
                  Back
                </v-btn>
                <v-btn
                  color="info"
                  class="mr-0"
                  type="submit"
                  :loading="$store.state.loading.status"
                >
                  Disable 2-step verification
                </v-btn>
              </v-col>
            </template>
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
  layout: 'Home',
  components: {
    Alert,
    MaterialCard,
    RecoveryQuestions: () => import('@@/components/helpers/RecoveryQuestions')
  },
  data () {
    return {
      success: false,
      showQuestions: false,
      username: '',
      password: '',
      errors: null,
      recoveryQuestions: [{}, {}, {}]
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Disable 2-step verification')
  },
  methods: {
    login () {
      this.clearAlert()
      if (!this.$refs.form.validate()) {
        return
      }
      this.$store.dispatch('loading/start')
      UserAPI.getRecoveryQuestions({
        password: this.hash(this.password),
        username: this.username
      })
        .then((r) => {
          this.recoveryQuestions = r.data.questions.data
          console.log(this.recoveryQuestions)
          this.showQuestions = true
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    disableMFA () {
      this.clearAlert()
      if (!this.$refs.form.validate()) {
        return
      }
      const questions = this.recoveryQuestions
      for (const r in questions) {
        questions[r].answer = this.hash(questions[r].answer.toLowerCase())
      }
      this.$store.dispatch('loading/start')
      UserAPI.disableMFA({
        password: this.hash(this.password),
        username: this.username,
        questions
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
    clearAlert () {
      this.success = false
      this.errors = null
    }
  }
}
</script>
