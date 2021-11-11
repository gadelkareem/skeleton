<template>
  <v-container>
    <h3 class="display-2 text-center mb-4">
      Pricing and Plan
    </h3>
    <p class="body-1 text-center mb-4" />
    <div class="pricing-wrap">
      <v-row justify="center">
        <v-col md="9">
          <v-row align="end">
            <v-col
              v-for="(tier, index) in tiers"
              :key="index"
              md="4"
              sm="tier.title === 'enterpris' ? 12 : 4"
              class="px-5"
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

<style scoped lang="scss">
@import './pricing-style';
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
