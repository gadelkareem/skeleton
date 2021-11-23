<template>
  <v-container>
    <v-row>
      <alert
        :errors="errors"
        :success="success"
        :success-txt="successTxt"
      />
      <div>
        <payment-methods @updateAlert="updateAlert" />
      </div>
      <div class="float-left ma-2">
        <add-payment-method />
      </div>
    </v-row>
  </v-container>
</template>
<script>
import PaymentMethods from '@@/components/payment/paymentMethods'
import Alert from '@@/components/base/Alert'
import AddPaymentMethod from '@@/components/payment/AddPaymentMethod'
export default {
  components: { PaymentMethods, Alert, AddPaymentMethod },
  layout: 'Dashboard',
  data () {
    return {
      errors: [],
      success: false,
      successTxt: '',
      paymentMethods: []
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Payment Methods')
  },
  methods: {
    updateAlert (v) {
      if (!v) { v = {} }
      this.success = v.success || false
      this.successTxt = v.successTxt || ''
      this.errors = v.errors || null
    }
  }
}
</script>
