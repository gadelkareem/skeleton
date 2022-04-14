<template>
  <v-container>
    <h3 class="display-2 text-center mb-4">
      Pricing and Plan <strong>( Demo )</strong>
    </h3>
    <p class="body-1 text-center mb-4" />
    <div class="pricing-wrap">
      <v-row justify="center">
        <v-col md="9">
          <v-row align="end">
            <v-col
              v-for="(tier, index) in tiers"
              :key="index"
              class="px-5"
              md="4"
              sm="tier.title === 'enterprise' ? 12 : 4"
            >
              <pricing-card
                :card="tier"
              />
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<style lang="sass" scoped>
@import './pricing-style'
</style>

<script>
import ProductAPI from '@@/api/product'
import PricingCard from '../Cards/PricingCard'

export default {
  components: {
    PricingCard
  },
  data () {
    return {
      tiers: []
    }
  },
  mounted () {
    this.fetchProducts()
  },
  methods: {
    fetchProducts () {
      this.$store.dispatch('loading/start')
      ProductAPI.list()
        .then((r) => {
          console.log(r)
          this.tiers = this.formatTiers(r.data)
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
