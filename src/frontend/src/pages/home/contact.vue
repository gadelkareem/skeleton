<template>
  <v-container
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card
        class="pa-6"
        :min-width="$vuetify.breakpoint.smAndDown ? '95%' : '70%'"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Contact
          </div>

          <div class="subtitle-1 font-weight-light">
            Fill in your message
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-txt="Your Message was successfully sent to us! |We will get back to you ASAP!"
        />
        <template>
          <v-form
            ref="form"
            method="post"
            class="pt-8"
            @submit.prevent="contact"
          >
            <v-container class="py-0">
              <v-row>
                <v-text-field
                  v-model.trim="name"
                  label="Name*"
                  type="text"
                  outlined
                  :rules="[$validator.required]"
                  data-username
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
              <v-row>
                <v-textarea
                  v-model="msg"
                  label="Message*"
                  type="text-area"
                  outlined
                  :rules="[$validator.required]"
                />
              </v-row>
              <v-spacer />
              <v-col
                class="text-right"
              >
                <v-btn
                  color="info"
                  class="mr-0"
                  type="submit"
                  :loading="$store.state.loading.status"
                >
                  Send
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

import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'
import CommonAPI from '@@/api/common'

export default {
  name: 'Contact',
  components: { Alert, MaterialCard },
  layout: 'Home',
  data () {
    return {
      success: false,
      name: '',
      email: '',
      msg: '',
      errors: []
    }
  },
  created () {
  },
  mounted () {
    this.$store.dispatch('page/title', 'Contact')
  },
  methods: {
    contact () {
      if (!this.$refs.form.validate()) {
        return
      }
      this.$store.dispatch('loading/start')
      this.success = false
      this.errors = []

      CommonAPI.contact({
        name: this.name,
        email: this.email,
        message: this.msg
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
    }
  }
}
</script>
