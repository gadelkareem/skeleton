<template>
  <material-card
    icon="mdi-account-group"
    :title="title"
    class="px-5 py-3"
  >
    <v-card-title>
      <v-spacer />
      <v-text-field
        v-model="filter"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
        clearable
        @change="search"
        @click:clear="clear"
      />
    </v-card-title>
    <alert
      :errors="errors"
    />
    <v-data-table
      :headers="headers"
      :items="items"
      :options.sync="options"
      :server-items-length="total"
      :loading="$store.state.loading.status"
      @click:row="$emit('rowClicked', $event)"
    >
      <template v-for="(_, slot) of $scopedSlots" v-slot:[slot]="scope">
        <slot :name="slot" v-bind="scope" />
      </template>
    </v-data-table>
  </material-card>
</template>
<script>
import Api from '@@/api/api'
import MaterialCard from '../base/MaterialCard'
import Alert from '../base/Alert'

export default {
  components: { Alert, MaterialCard },
  layout: 'Dashboard',
  props: {
    title: {
      type: String,
      default: ''
    },
    model: {
      type: String,
      default: ''
    },
    headers: {
      type: Array,
      default: () => [{}]
    },
    sortBy: {
      type: Array,
      default: () => [{}]
    },
    sortDesc: {
      type: Array,
      default: () => [{}]
    }
  },
  data () {
    return {
      errors: [],
      dialog: false,
      items: [],
      options: {},
      filter: '',
      optionsLoaded: false,
      lastQuery: {},
      total: 0
    }
  },
  watch: {
    options: {
      handler (v) {
        if (!this.optionsLoaded) {
          return
        }
        this.fetchItems(v, this.filter)
      },
      deep: true
    }
  },
  mounted () {
    this.$on('click:row', alert)
    this.lastQuery = {}
    this.optionsLoaded = false
    const $this = this
    this.$nextTick(function () {
      this.fetchItems({ sortBy: $this.sortBy, sortDesc: $this.sortDesc })
    })
  },
  methods: {
    fetchItems (options, filter) {
      const query = this.$paginator.queryFromOptions(options, filter, this.$route.query)
      if (this.isEqualObjects(query, this.lastQuery)) {
        return
      }
      this.lastQuery = query
      this.$store.dispatch('loading/start')
      Api.fetch(this.model, query)
        .then((r) => {
          this.items = r.data
          if (r.meta.page) {
            this.total = r.meta.page.total || 0
          }
          this.filter = r.meta.filter || ''
          this.options = this.$paginator.optionsFromAPI(r.meta.page, r.meta.sort, r.meta.filter)
          this.updateRoute(this.options, this.filter)
            .then(() => {
              this.optionsLoaded = true
            })
        })
        .catch((err) => {
          this.errors = this.parseError(err)
        }).finally(() => {
          this.$store.dispatch('loading/finish')
        })
    },
    clear () {
      this.filter = ''
      this.search()
    },
    search () {
      this.lastQuery = {}
      if (!this.filter) {
        this.updateRoute(this.options)
          .then(() => {
            this.fetchItems(this.options)
          })
      } else {
        this.fetchItems(this.options, this.filter)
      }
    },
    updateRoute (options, filter) {
      return this.$router.replace({
        ...this.$router.currentRoute,
        query: {
          ...options,
          filter
        }
      })
        .catch(() => {
        })
    }
  }
}
</script>
