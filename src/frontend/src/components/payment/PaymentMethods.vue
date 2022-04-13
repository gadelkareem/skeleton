<template>
  <v-container>
    <v-item-group v-model="selectedPaymentMethodKey" mandatory>
      <v-container>
        <v-row>
          <div
            v-for="pm in paymentMethods"
            :key="pm.id"
            class="float-left ma-2"
          >
            <payment-method-card
              :pm="pm"
              @deletePaymentMethod="deletePaymentMethod"
              @setPaymentMethodID="setPaymentMethodID"
            />
          </div>
          <div class="float-left ma-2">
            <add-payment-method
              v-if="showAddPaymentMethod"
              @listPaymentMethods="listPaymentMethods"
              @setPaymentMethodID="setPaymentMethodID"
              @updateAlert="updateAlert"
            />
          </div>
        </v-row>
      </v-container>
    </v-item-group>
  </v-container>
</template>
<script>
import CustomerAPI from '@@/api/customer'
import PaymentMethodCard from '@@/components/payment/card/PaymentMethodCard'
import AddPaymentMethod from '@@/components/payment/AddPaymentMethod'
import SubscriptionAPI from '@@/api/subscription'
import PaymentMethodAPI from '@@/api/paymentMethod'

export default {
  name: 'PaymentMethods',
  components: { AddPaymentMethod, PaymentMethodCard },
  props: {
    plan: {
      type: Object,
      default: () => ({ id: '' }),
      required: false
    },
    showAddPaymentMethod: {
      type: Boolean,
      required: false,
      default: false
    },
    subscription: {
      type: Object,
      required: false,
      default: null
    }
  },
  data () {
    return {
      selectedPaymentMethodKey: '',
      selectedPaymentMethodID: '',
      paymentMethods: [],
      pk: process.env.STRIPE_PK,
      stripe: null
    }
  },
  computed: {
    user () {
      return this.$store.getters['user/user']
    }
  },
  mounted () {
    this.listPaymentMethods()
    this.init()
  },
  methods: {
    setPaymentMethodID (id) {
      if (!this.paymentMethods.length) {
        return
      }
      if (!id) {
        id = this.paymentMethods[0].id
      }
      this.selectedPaymentMethodID = id
      const pm = this.paymentMethods.find(o => o.id === id)
      if (pm) { this.$emit('setPaymentMethod', pm) }
    },
    async setupPaymentIntent () {
      if (!this.plan.priceID) {
        const err = 'No plan selected'
        this.updateAlert({ errors: [{ title: err }] })
        return Promise.reject(new Error())
      }
      this.updateAlert()
      let sub = {
        price_id: this.plan.priceID,
        customer_id: this.user.customer_id,
        create_payment_intent: true
      }
      if (this.subscription) {
        sub = { ...sub,
          id: this.subscription.id,
          // old_price_id: this.subscription.price_id,
          item_id: this.subscription.item_id,
          payment_method_id: this.selectedPaymentMethodID
        }
      }
      await this.$store.dispatch('loading/start')
      const cs = await SubscriptionAPI.createOrUpdate(sub)
        .then((r) => {
          this.$emit('setPaymentIntent', r.data)
          console.log('createOrUpdate', r.data)
          console.log('clientSecret', r.data.payment_intent_client_secret)
          return r.data ? r.data.payment_intent_client_secret : null
        })
        .catch((err) => {
          this.updateAlert({ errors: this.parseError(err) })
          return Promise.reject(new Error())
        })
        .finally(() => { this.$store.dispatch('loading/finish') })
      return cs
    },
    listPaymentMethods (resetCache) {
      this.$store.dispatch('loading/start')
      CustomerAPI.customerPaymentMethods(this.user.customer_id, !!resetCache)
        .then((r) => {
          this.paymentMethods = r.data
          for (const pm of this.paymentMethods) {
            if (pm.is_default) {
              this.defaultPaymentMethod = pm.id
            }
            if (pm.id === this.selectedPaymentMethodID) {
              this.selectedPaymentMethodKey = pm.id
            }
          }
          this.setPaymentMethodID(this.selectedPaymentMethodID)
          return r
        })
        .catch((err) => {
          err = this.parseError(err)
          if (err[0].status !== '404') {
            this.updateAlert({ errors: err })
          }
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    deletePaymentMethod (id) {
      this.updateAlert()
      this.$store.dispatch('loading/start')
      PaymentMethodAPI.deletePaymentMethod(id)
        .then((r) => {
          this.updateAlert({ successTxt: 'Payment Method deleted successfully.' })
          if (id === this.selectedPaymentMethodID) {
            this.selectedPaymentMethodID = ''
          }
          this.$emit('setPaymentMethod', null)
          this.listPaymentMethods()
          return r
        })
        .catch((err) => {
          this.updateAlert({ errors: this.parseError(err) })
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    updateAlert (v) {
      this.$emit('updateAlert', v)
    },
    init () {
      this.includeStripe('js.stripe.com/v3/', function () {
        this.configureStripe()
      }.bind(this))
    },
    async confirmCard () {
      this.updateAlert()
      const clientSecret = await this.setupPaymentIntent()
      console.log('clientSecret', clientSecret)
      // no payment intent required
      if (!clientSecret) {
        this.$emit('addConfirmedPaymentMethods', this.selectedPaymentMethodID)
        return 0
      }
      return this.stripeConfirmCard(clientSecret)
    },
    async stripeConfirmCard (clientSecret) {
      if (!clientSecret) {
        const err = 'Internal error (missing client secret)'
        this.updateAlert({ errors: [{ title: err }] })
        return Promise.reject(new Error())
      }
      await this.$store.dispatch('loading/start')
      return this.stripe.confirmCardPayment(
        clientSecret, {
          payment_method: this.selectedPaymentMethodID
        }
      ).then(function (r) {
        this.$store.dispatch('loading/finish')
        if (r.error && r.error.code !== 'payment_intent_unexpected_state') {
          console.log(r)
          this.updateAlert({ errors: [{ title: r.error.message }] })
          return Promise.reject(new Error())
        } else {
          this.$emit('addConfirmedPaymentMethods', this.selectedPaymentMethodID)
        }
      }.bind(this))
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
    }
  }
}
</script>
