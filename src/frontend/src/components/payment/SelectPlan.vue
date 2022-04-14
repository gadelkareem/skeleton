<template>
  <span>
    <h3>
      <span class="step__number" />
      Select A Plan
    </h3>
    <v-item-group v-model="selectedTier" mandatory @change="updatePlan">
      <v-container>
        <v-row>
          <v-col
            v-for="(tier,i) in tiers"
            :key="i"
            md="4"
          >
            <plan-card :card="tier" />
          </v-col>
        </v-row>
      </v-container>
    </v-item-group>
    <template v-if="tiers[selectedTier]">
      <h3>
        <span class="step__number" />
        Term
      </h3>
      <v-item-group v-model="selectedTerm" mandatory @change="updatePlan">
        <v-container>
          <v-row>
            <v-col
              v-for="(price, i) in tiers[selectedTier].prices"
              :key="i"
              cols="6"
              md="3"
            >
              <term-card :card="price" :disabled="subscription && price.id === subscription.price_id" />
            </v-col>
          </v-row>
        </v-container>
      </v-item-group>
    </template>
  </span>
</template>
<script>
import PlanCard from '@@/components/payment/form/PlanCard'
import TermCard from '@@/components/payment/form/TermCard'
import ProductAPI from '@@/api/product'
export default {
  name: 'SelectPlan',
  components: { PlanCard, TermCard },
  props: {
    subscription: {
      type: Object,
      required: false,
      default: null
    }
  },
  data: () => ({
    tiers: [],
    selectedTier: 1,
    selectedTerm: 0
  }),
  computed: {
    plan () {
      const tier = this.tiers[this.selectedTier]
      const price = tier ? tier.prices[this.selectedTerm] : { }
      return {
        tier,
        price,
        'interval': price && price.recurring ? price.recurring.interval : '',
        'priceAmount': price ? price.unit_amount / 100 : 0,
        'priceID': price ? price.id : 0
      }
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Plans')
    this.fetchProducts()
  },
  methods: {
    updatePlan () {
      if (this.selectedTerm === undefined) {
        return
      }
      this.$emit('updatePlan', this.plan)
    },
    setPlan (priceID) {
      if (priceID) {
        for (const [i, t] of this.tiers.entries()) {
          for (const [x, p] of t.prices.entries()) {
            if (p.id === priceID) {
              this.selectedTier = i
              this.selectedTerm = x
              break
            }
          }
        }
      }
      this.updatePlan()
    },
    fetchProducts () {
      this.$store.dispatch('loading/start')
      ProductAPI.list()
        .then((r) => {
          console.log(r)
          this.tiers = this.formatTiers(r.data)
          this.$nextTick(() => {
            this.setPlan(this.$route.query.plan)
          })
        })
        .catch((err) => {
          this.updateAlert({ 'errors': this.parseError(err) })
        })
        .finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    updateAlert (v) {
      this.$emit('updateAlert', v)
    }
  }
}
</script>
