<template>
  <v-container
    id="user-profile"
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card>
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Recovery Questions
          </div>

          <div class="subtitle-1 font-weight-light">
            Fill in your answers to be able to recover your account in case you do not have access to your mobile.
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-txt="Successful!"
        />
        <v-row v-if="user.recovery_questions_set" align="center" justify="center" class="mt-5">
          <p class="display-1 ma-10">
            Your recovery questions are set, please let us know if you would like to change them.
          </p>
        </v-row>
        <v-form
          v-else
          ref="form"
          method="post"
          @submit.prevent="submit"
        >
          <v-container class="py-0">
            <recovery-questions :sets="recoveryQuestions" />
            <v-row>
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
                  Save
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
import UserAPI from '@@/api/user'
import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'

export default {
  layout: 'Dashboard',
  components: {
    Alert,
    MaterialCard,
    RecoveryQuestions: () => import('@@/components/helpers/RecoveryQuestions')
  },
  data: () => ({
    errors: [],
    success: false,
    list: [],
    recoveryQuestions: [{}, {}, {}]
  }),
  computed: {
    user: {
      get () {
        return this.$store.getters['user/user']
      }
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Recovery Questions')
  },
  methods: {
    submit () {
      if (!this.$refs.form.validate()) {
        return
      }
      const questions = this.recoveryQuestions
      for (const r in questions) {
        questions[r].answer = this.hash(questions[r].answer.toLowerCase())
      }
      this.success = false
      this.errors = []
      this.$store.dispatch('loading/start')
      UserAPI.recoveryQuestions(this.user.id, questions)
        .then(() => {
          this.success = true
          this.$store.dispatch('user/fetchUser', this.user.id)
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
