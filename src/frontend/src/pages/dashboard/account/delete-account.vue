<template>
  <v-container
    fluid
    tag="section"
  >
    <v-row justify="center">
      <material-card
        class="pa-6"
        min-width="50%"
      >
        <template v-slot:heading>
          <div class="display-2 font-weight-light">
            Delete Account
          </div>
          <div class="subtitle-1 font-weight-light">
            Once the account is delete it cannot be restored
          </div>
        </template>
        <alert
          :errors="errors"
          :success="success"
          success-txt="Your account has been deleted successfully"
        />
        <v-container>
          <v-col
            cols="12"
            class="text-center"
          >
            <v-btn
              color="error"
              class="mr-0"
              type="submit"
              :loading="$store.state.loading.status"
              @click="deleteUser"
            >
              Delete Account
            </v-btn>
          </v-col>
        </v-container>
      </material-card>
    </v-row>
    <confirm-modal ref="confirm" />
  </v-container>
</template>

<script>
import UserAPI from '@@/api/user'
import MaterialCard from '@@/components/base/MaterialCard'
import Alert from '@@/components/base/Alert'
import ConfirmModal from '@@/components/base/ConfirmModal'

export default {
  components: { ConfirmModal, Alert, MaterialCard },
  layout: 'Dashboard',
  data: () => ({
    errors: [],
    success: false
  }),
  computed: {
    user: {
      get () {
        return this.$store.getters['user/user']
      }
    }
  },
  mounted () {
    this.$store.dispatch('page/title', '2-step verification')
  },
  methods: {
    deleteUser () {
      this.$refs.confirm.open(
        'Delete Account',
        'Are you sure you want to delete your account?'
      ).then((confirm) => {
        if (confirm) {
          this.clearAlert()
          UserAPI.delete(this.user.id)
            .then(() => {
              this.success = true
              this.$store.dispatch('auth/logout')
            })
            .catch((err) => {
              this.errors = this.parseError(err)
            })
            .finally(() => {
              this.$store.dispatch('loading/finish')
            })
        }
      })
    },
    clearAlert () {
      this.success = false
      this.errors = []
      this.$store.dispatch('loading/start')
    }
  }
}
</script>
