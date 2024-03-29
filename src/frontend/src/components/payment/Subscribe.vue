<template>
  <span>
    <v-row justify="center">
      <material-card
        :width="$vuetify.breakpoint.smAndDown ? '65%' : '40%'"
        class="pa-6"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            <template v-if="checkout">Checkout</template>
            <template v-else>Current Plan</template>
          </div>

          <div class="subtitle-1 font-weight-light">
            <template v-if="checkout">Review your order</template>
            <template v-else>Subscription Information</template>
          </div>
        </template>
        <v-card-text>
          <v-container>
            <v-row align="center">
              <v-col>
                <h5>Plan:</h5>
              </v-col>
              <v-col>
                <v-card
                  v-if="plan.tier"
                  class="ma-0"
                  outlined
                  shaped
                  width="300"
                >
                  <v-card-title>{{ plan.tier.title }}</v-card-title>
                  <v-card-text>
                    <ul>
                      <li
                        v-for="(line, index) in plan.tier.description"
                        :key="index"
                        class="subtitle-1"
                      >
                        {{ line }}
                      </li>
                    </ul>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
            <v-row v-if="newInvoice && newInvoice.total > 0" align="center" class="mt-5">
              <v-col>
                <h5>Paid Today:</h5>
              </v-col>
              <v-col>
                <v-card
                  class="ma-0"
                  outlined
                  shaped
                  width="300"
                >
                  <v-card-text>
                    <span class=" text-xs-center display-2 checkout-price">
                      {{ formatMoney(newInvoice.total) }}
                    </span>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
            <v-row align="center" class="mt-5">
              <v-col>
                <h5>Recurring Payment:</h5>
              </v-col>
              <v-col>
                <v-card
                  class="ma-0"
                  outlined
                  shaped
                  width="300"
                >
                  <v-card-text>
                    <template v-if="plan.price && !plan.priceAmount">
                      <span class=" text-xs-center display-2 checkout-price">
                        Free
                      </span>
                    </template>
                    <template v-else>
                      <span class=" text-xs-center display-2 checkout-price">
                        {{ formatMoney(plan.priceAmount) }}
                      </span>
                      <span class="title interval">
                        / {{ plan.interval }}
                      </span>
                    </template>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
            <v-row v-if="paymentMethod" align="center" class="mt-5">
              <v-col>
                <h5>Payment Method:</h5>
              </v-col>
              <v-col>
                <payment-method-card
                  :pm="paymentMethod"
                  :show-delete="false"
                />
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions v-if="!checkout">
          <v-spacer />
          <confirm-modal ref="confirm" />
          <v-btn :loading="$store.state.loading.status" class="error" @click="cancelSubscription">
            Cancel Subscription
          </v-btn>
          <v-btn :loading="$store.state.loading.status" class="primary" @click="$emit('setChangeSubscription')">
            Change Current Plan
          </v-btn>
        </v-card-actions>
      </material-card>
    </v-row></span>
</template>
<script>
import MaterialCard from '@@/components/base/MaterialCard'
import PaymentMethodCard from '@@/components/payment/card/PaymentMethodCard'
import ConfirmModal from '@@/components/base/ConfirmModal'

export default {
  name: 'Subscribe',
  components: { PaymentMethodCard, MaterialCard, ConfirmModal },
  props: {
    plan: {
      type: Object,
      required: true
    },
    paymentMethod: {
      type: Object,
      required: false,
      default: null
    },
    checkout: {
      type: Boolean,
      required: false,
      default: true
    },
    newInvoice: {
      type: Object,
      required: false,
      default: null
    }
  },
  mounted () {
    this.$store.dispatch('page/title', this.checkout ? 'checkout' : 'Current Plan')
  },
  methods: {
    cancelSubscription () {
      this.$refs.confirm.open(
        'Cancel Subscription',
        'Are you sure you want to cancel your subscription?'
      ).then((confirm) => {
        if (confirm) {
          this.$emit('cancelSubscription')
        }
      })
    }

  }
}
</script>
