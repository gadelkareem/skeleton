<template>
  <v-container
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card
        class="pa-6"
        min-width="100%"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Credit Card
          </div>
          <div class="subtitle-1 font-weight-light">
            Add your credit card
          </div>
        </template>

        <section
          id="pricing"
          class="space-top"
        >
          <payment-card />
        </section>
      </material-card>
    </v-row>
  </v-container>
</template>

<script>
// import MaterialCard from '@@/components/base/MaterialCard'
// import PaymentCard from '@@/components/payment/PaymentCard'

export default {
  // components: { PaymentCard, MaterialCard },
  layout: 'Dashboard',

  computed: {

  },
  mounted () {
    this.$store.dispatch('page/title', 'Plans')
  },
  methods: {

  }
}
</script>
