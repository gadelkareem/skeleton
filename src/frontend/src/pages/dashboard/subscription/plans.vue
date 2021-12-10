<template>
  <v-container
    id="create-subscription"
    fluid
    tag="section"
  >
    <alert
      :errors="errors"
      :success="success"
      :success-txt="successTxt"
    />
    <select-plan v-show="step === 0 && !displaySubscription" ref="selectPlan" @updatePlan="updatePlan" @updateAlert="updateAlert" />
    <select-payment-method
      v-show="step === 1 && !displaySubscription"
      v-if="plan"
      ref="selectPaymentMethod"
      :plan="plan"
      :subscription="subscription"
      @updateAlert="updateAlert"
      @setPaymentMethod="setPaymentMethod"
      @setPaymentIntent="setPaymentIntent"
    />
    <subscribe
      v-show="step === 2 || displaySubscription"
      v-if="plan && paymentMethod"
      ref="subscribe"
      :plan="plan"
      :payment-method="paymentMethod"
      :checkout="!displaySubscription"
      :new-invoice="newInvoice"
      @setChangeSubscription="setChangeSubscription"
      @cancelSubscription="cancelSubscription"
    />
    <v-spacer style="margin-top:200px" />
    <v-footer
      v-if="!displaySubscription && plan"
      fixed
      elevation="6"
      width="auto"
      app
      inset
    >
      <v-container>
        <v-row
          align="center"
          justify="center"
          dense
        >
          <v-col cols="2" class="mr-2">
            <v-card v-if="newInvoice && newInvoice.total > 0" class="pa-4 mt-0" elevation="2" :loading="$store.state.loading.status">
              <v-card-title class="text--secondary">
                Paid Today:
              </v-card-title>
              <v-card-text>
                <template>
                  <span class=" text-xs-center display-2 checkout-price">
                    € {{ newInvoice.total }}
                  </span>
                </template>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="6">
            <v-card class="pa-4 mt-0" elevation="2">
              <v-card-title class="text--secondary">
                Recurring Payment:
              </v-card-title>
              <v-card-text>
                <template v-if="plan.price && !plan.priceAmount">
                  <span class=" text-xs-center display-2 checkout-price">
                    Free
                  </span>
                </template>
                <template v-else>
                  <span class=" text-xs-center display-2 checkout-price">
                    € {{ plan.priceAmount }}
                  </span>
                  <span class="title interval">
                    / {{ plan.interval }}
                  </span>
                </template>
                <span class="float-right">
                  <v-btn v-if="step> 0" class=" primary" :loading="$store.state.loading.status" @click="previousStep()">
                    Back
                  </v-btn>
                  <v-btn v-else-if="changeSubscription" class=" primary" :loading="$store.state.loading.status" @click="changeSubscription=false">
                    Cancel
                  </v-btn>
                  <v-btn class="primary" :loading="$store.state.loading.status" @click="nextStep()">
                    {{ stepActions[step] }}
                  </v-btn>
                </span>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-footer>
  </v-container>
</template>
<script>

import SelectPlan from '@@/components/payment/SelectPlan'
import SelectPaymentMethod from '@@/components/payment/SelectPaymentMethod'
import Alert from '@@/components/base/Alert'
import Subscribe from '@@/components/payment/Subscribe'
import SubscriptionAPI from '@@/api/subscription'
import CustomerAPI from '@@/api/customer'
import InvoiceAPI from '@@/api/invoice'

export default {
  components: { SelectPaymentMethod, SelectPlan, Alert, Subscribe },
  layout: 'Dashboard',
  data: () => ({
    errors: [],
    success: false,
    successTxt: '',
    plan: null,
    step: -1,
    stepActions: [
      'Proceed to Checkout',
      'Continue',
      'Subscribe'
    ],
    paymentMethod: null,
    subscription: null,
    changeSubscription: false,
    newInvoice: null,
    paymentIntent: null
  }),
  computed: {
    displaySubscription () {
      return this.subscription && !this.changeSubscription
    },
    user () {
      return this.$store.getters['user/user']
    }
  },
  watch: {
    displaySubscription (v) {
      this.$store.dispatch('page/hideFooter', !v)
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Plans')
    this.$store.dispatch('page/hideFooter', true)
    this.$nextTick(async function () {
      await this.getSubscription()
    })
  },
  methods: {
    setPaymentIntent (v) {
      this.paymentIntent = v
    },
    previousStep () {
      const s = this.step - 1
      if (s < 0) { return }
      this.updateAlert()
      this.step--
    },
    async nextStep () {
      const s = this.step + 1
      if (s > 3) { return }
      try {
        switch (s) {
          case 3:
            await this.$refs.selectPaymentMethod.confirmCard()
            if (!this.$refs.selectPaymentMethod.isPaymentMethodConfirmed(this.paymentMethod.id)) {
              throw new Error('Could not verify your Card.')
            }
            await this.subscribe()
            break
        }
      } catch (e) {
        console.log('error', e)
        if (e.message) { this.updateAlert({ errors: [{ title: e.message }] }) }
        return
      }
      if (s < 3) { this.step++ }
    },
    updatePlan (plan) {
      this.plan = plan
      if (this.plan.priceID !== this.$route.query.plan) {
        this.$router.replace({ query: { plan: this.plan.priceID } })
        this.upcomingInvoice()
      }
    },
    updateAlert (v) {
      if (!v) { v = {} } else { window.scrollTo(0, 0) }
      this.success = v.success || false
      this.successTxt = v.successTxt || ''
      this.errors = v.errors || null
    },
    setPaymentMethod (pm) {
      this.paymentMethod = pm
    },
    async subscribe () {
      if (this.subscription) {
        this.updateAlert({ success: true, successTxt: 'Your subscription has been updated successfully' })
        return
      }
      const s = {
        price_id: this.plan.priceID,
        customer_id: this.user.customer_id,
        payment_method_id: this.paymentMethod.id,
        id: this.paymentIntent.id
      }
      const s1 = this.subscription ? {
        price_id: this.plan.priceID,
        customer_id: this.user.customer_id,
        old_price_id: this.subscription.price_id,
        item_id: this.paymentIntent.item_id,
        payment_method_id: this.paymentMethod.id
      } : {}
      await this.$store.dispatch('loading/start')
      await SubscriptionAPI.update({ ...s, ...s1 })
        .then((r) => {
          this.updateAlert({ success: true, successTxt: this.subscription ? 'Your subscription has been updated successfully' : 'You have successfully subscribed' })
          this.changeSubscription = false
          this.paymentIntent = null
          this.newInvoice = null
          this.$refs.selectPaymentMethod.confirmedPaymentMethods = []
        })
        .catch((err) => {
          this.updateAlert({ errors: this.parseError(err) })
        })
      await this.$store.dispatch('loading/finish')
      await this.getSubscription()
    },
    async getSubscription () {
      await this.$store.dispatch('loading/start')
      await CustomerAPI.customerSubscription(this.user.customer_id)
        .then((r) => {
          this.subscription = r.data
          this.$refs.selectPlan.setPlan(this.subscription.price_id)
          this.step = 2
        })
        .catch((err) => {
          err = this.parseError(err)
          if (err[0].status !== '404') { this.errors = this.parseError(err) }
          this.step = 0
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    cancelSubscription () {
      this.updateAlert()
      if (!this.subscription) { return }
      this.$store.dispatch('loading/start')
      SubscriptionAPI.cancel(this.subscription.id)
        .then((r) => {
          this.updateAlert({ success: true, successTxt: 'Your subscription has been canceled successfully' })
          this.subscription = null
          this.paymentIntent = null
          this.$refs.selectPaymentMethod.confirmedPaymentMethods = []
          this.step = 0
          this.newInvoice = null
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    setChangeSubscription () {
      this.changeSubscription = true
      this.step = 0
      this.updateAlert()
    },
    upcomingInvoice () {
      if (this.displaySubscription || !this.subscription || !this.plan.priceID) { return }
      this.$store.dispatch('loading/start')
      InvoiceAPI.upcomingInvoice({
        id: this.subscription.id,
        price_id: this.plan.priceID,
        customer_id: this.user.customer_id,
        old_price_id: this.subscription.price_id,
        item_id: this.subscription.item_id
      })
        .then((r) => {
          this.newInvoice = r.data
          this.newInvoice.total = this.newInvoice.total / 100
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
