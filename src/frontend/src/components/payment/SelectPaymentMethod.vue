<template>
  <span>
    <h3>
      <span class="step__number" />
      Select Payment Method
    </h3>
    <payment-methods
      ref="paymentMethods"
      :plan="plan"
      :show-add-payment-method="Boolean(plan)"
      :subscription="subscription"
      @setPaymentMethod="setPaymentMethod"
      @updateAlert="updateAlert"
      @addConfirmedPaymentMethods="addConfirmedPaymentMethods"
      @setPaymentIntent="setPaymentIntent"
    />
  </span>
</template>
<script>
import PaymentMethods from '@@/components/payment/PaymentMethods'

export default {
  name: 'SelectPaymentMethod',
  components: { PaymentMethods },
  props: {
    plan: {
      type: Object,
      required: true
    },
    subscription: {
      type: Object,
      required: false,
      default: null
    }
  },
  data: () => ({
    selectedPaymentMethod: '',
    confirmedPaymentMethods: []
  }),
  mounted () {
    this.$store.dispatch('page/title', 'Select Payment Method')
  },
  methods: {
    addConfirmedPaymentMethods (id) {
      this.confirmedPaymentMethods.push(id)
      console.log('confirmedPaymentMethods', this.confirmedPaymentMethods)
    },
    async confirmCard () {
      if (!this.isPaymentMethodConfirmed(this.selectedPaymentMethod.id)) {
        await this.$refs.paymentMethods.confirmCard()
      }
    },
    isPaymentMethodConfirmed (id) {
      return this.confirmedPaymentMethods.includes(id)
    },
    updateAlert (v) {
      this.$emit('updateAlert', v)
    },
    setPaymentIntent (v) {
      this.$emit('setPaymentIntent', v)
    },
    setPaymentMethod (pm) {
      this.selectedPaymentMethod = pm
      this.$emit('setPaymentMethod', pm)
      console.log('selectedPaymentMethodID', this.selectedPaymentMethod.id)
    }
  }
}
</script>
