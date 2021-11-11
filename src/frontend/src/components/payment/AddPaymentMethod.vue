<template>
  <v-dialog
    v-model="dialog"
    persistent
    max-width="600px"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn
        v-bind="attrs"
        :loading="$store.state.loading.status"
        color="white"
        dark
        x-large
        height="150"
        width="300"
        class="mt-7"
        v-on="on"
        @click="init"
      >
        <v-icon
          color="green"
          x-large
        >
          mdi-plus
        </v-icon> Add Credit Card
      </v-btn>
    </template>
    <v-container
      id="user-Card"
      fluid
      tag="section"
    >
      <v-row justify="center">
        <material-card width="100%">
          <template v-slot:heading>
            <div class="display-2 font-weight-light">
              Add Credit Card
            </div>

            <div class="subtitle-1 font-weight-light">
              Complete your Credit Card Information
            </div>
          </template>
          <alert
            :errors="errors"
            :success="success"
            success-txt="Your card has been added successfully!"
          />
          <v-form
            v-if="!success"
            ref="form"
            method="post"
            @submit.prevent="addCard"
          >
            <v-container class="py-0">
              <v-row>
                <v-col
                  cols="12"
                >
                  <v-text-field
                    v-model.trim="name"
                    :rules="[$validator.name]"
                    label="Card Holder Name"
                    class="purple-input"
                  />
                </v-col>
                <v-col cols="12">
                  <label>Card</label>
                  <div>
                    <div id="card-element" />
                  </div>
                </v-col>
                <v-col
                  cols="12"
                  class="text-right mt-6"
                >
                  <v-btn
                    color="info"
                    class="mr-0"
                    :loading="$store.state.loading.status"
                    @click="dialog=false;updateAlert()"
                  >
                    Cancel
                  </v-btn>
                  <v-btn
                    color="info"
                    class="mr-0"
                    type="submit"
                    :loading="$store.state.loading.status"
                  >
                    Add Card
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </v-form>
          <v-col
            v-else
            cols="12"
          >
            <v-row justify="center">
              <v-btn
                color="info"
                class="mr-0"
                @click="resetForm"
              >
                Close
              </v-btn>
            </v-row>
          </v-col>
        </material-card>
      </v-row>
    </v-container>
  </v-dialog>
</template>
<script>

import Alert from '@@/components/base/Alert'
import MaterialCard from '@@/components/base/MaterialCard'
import PaymentMethodsAPI from '@@/api/paymentMethod'

export default {
  components: { Alert, MaterialCard },
  layout: 'Dashboard',
  data: () => ({
    dialog: false,
    errors: [],
    success: false,
    successTxt: '',
    pk: process.env.STRIPE_PK,
    stripe: null,
    elements: null,
    card: null,
    name: ''
  }),
  computed: {
    user () {
      return this.$store.getters['user/user']
    },
    countries () {
      return this.initData().countries
    }
  },
  methods: {
    init () {
      this.name = this.user.first_name || this.user.last_name ? this.user.first_name + ' ' + this.user.last_name : ''
      this.includeStripe('js.stripe.com/v3/', function () {
        this.configureStripe()
      }.bind(this))
    },
    addCard () {
      if (!this.$refs.form.validate()) {
        this.$store.dispatch('loading/finish')
        return
      }
      if (!this.user.mobile) {
        this.updateAlert({ errors: [{ title: 'Please add your mobile number to your Card.' }] })
        this.$store.dispatch('loading/finish')
        return
      }
      this.updateAlert()
      this.$store.dispatch('loading/start')
      this.stripe.createPaymentMethod({
        type: 'card',
        card: this.card,
        billing_details: {
          name: this.name,
          email: this.user.email,
          phone: this.user.mobile
          // address: this.address
        }
      }).then(function (r) {
        if (r.error) {
          this.errors = [{ title: r.error.message }]
          this.$store.dispatch('loading/finish')
          return
        }
        PaymentMethodsAPI.addPaymentMethod({
          id: r.paymentMethod.id,
          card: this.card
        })
          .then((r) => {
            this.success = true
            const id = r.data.id
            this.$emit('setPaymentMethodID', id)
            this.$emit('listPaymentMethods', true)
            this.$refs.form.reset()
            this.card.clear()
          })
          .catch((err) => {
            this.errors = this.parseError(err)
          })
          .finally(() => {
            this.$store.dispatch('loading/finish')
          })
      }.bind(this))
      // this.stripe.confirmCardPayment(
      //   clientSecret, {
      //     payment_method: {
      //       card: this.card,
      //       billing_details: {
      //         name: this.name,
      //         email: this.user.email,
      //         phone: this.user.mobile
      //       }
      //     }
      //   }
      // ).then(function (r) {
      //   if (r.error) {
      //     this.updateAlert({ errors: [{ title: r.error.message }] })
      //   } else {
      //     this.success = true
      //     const id = r.paymentIntent.payment_method
      //     this.$emit('setPaymentMethodID', id)
      //     this.$emit('addConfirmedPaymentMethods', id)
      //     this.$emit('listPaymentMethods', true)
      //     this.$refs.form.reset()
      //     this.card.clear()
      //   }
      //   this.$store.dispatch('loading/finish')
      // }.bind(this))
    },
    includeStripe (URL, callback) {
      const documentTag = document
      const tag = 'script'
      const object = documentTag.createElement(tag)
      const scriptTag = documentTag.getElementsByTagName(tag)[0]
      object.src = '//' + URL
      if (callback) {
        object.addEventListener('load', function (e) {
          callback(null, e)
        }, false)
      }
      scriptTag.parentNode.insertBefore(object, scriptTag)
    },
    configureStripe () {
      // eslint-disable-next-line no-undef
      this.stripe = Stripe(this.pk)

      this.elements = this.stripe.elements()
      this.card = this.elements.create('card', {
        hidePostalCode: false,
        style: {
          base: {
            iconColor: '#666EE8',
            color: '#31325F',
            lineHeight: '40px',
            fontWeight: 300,
            fontFamily: 'Helvetica Neue',
            fontSize: '15px',

            '::placeholder': {
              color: '#CFD7E0'
            }
          }
        }
      })

      this.card.mount('#card-element')
    },
    updateAlert (v) {
      if (!v) { v = {} }
      this.success = v.success || false
      this.successTxt = v.successTxt || ''
      this.errors = v.errors || null
    },
    resetForm () {
      this.updateAlert()
      this.dialog = false
    }
  }
}
</script>
