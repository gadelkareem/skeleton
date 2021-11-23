<template>
  <v-container
    fluid
    tag="section"
  >
    Your current plan is
  </v-container>
</template>

<script>

export default {
  layout: 'Dashboard',
  data: () => ({
    errors: [],
    success: false,
    successTxt: '',
    info: '',
    code: null,
    img: '',
    seed: null,
    password: null,
    url: null,
    showSeed: false,
    showForm: false

  }),
  computed: {
    user: {
      get () {
        return this.$store.getters['user/user']
      }
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'My Plan')
  },
  methods: {
    clearAlert () {
      this.success = false
      this.errors = []
      this.$store.dispatch('loading/start')
    }
  }
}
</script>
