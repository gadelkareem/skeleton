<template>
  <v-container>
    <alert
      :errors="errors"
      :success="success"
      :success-txt="successTxt"
    />
    <v-row>
      <div>
        <payment-methods :show-add-payment-method="true" @updateAlert="updateAlert" />
      </div>
    </v-row>
  </v-container>
</template>
<script>
import PaymentMethods from '@@/components/payment/paymentMethods'
import Alert from '@@/components/base/Alert'
export default {
  components: { PaymentMethods, Alert },
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
