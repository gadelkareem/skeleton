<template>
  <v-container
    fluid
    tag="section"
  >
    <v-row justify="center">
      <paginator
        ref="paginator"
        :headers="headers"
        :model="`customers/${user.customer_id}/invoices`"
        :show-search="false"
        icon="mdi-file-document-outline"
        title="Invoices"
      >
        <template #item.created="{item}">
          {{ (new Date(item.created*1000)).toLocaleDateString() }}
        </template>
        <template #item.total="{item}">
          {{ formatMoney(Math.abs(item.total/100)) }}
        </template>
        <template #item.invoice_pdf="{item}">
          <a v-if="item.status !== 'draft'" :href="item.invoice_pdf">
            Download
          </a>
        </template>
      </paginator>
    </v-row>
  </v-container>
</template>

<script>
import Paginator from '@@/components/helpers/Paginator'

export default {
  components: { Paginator },
  layout: 'Dashboard',
  data () {
    return {
      headers: [
        { text: 'Date', value: 'created', sortable: false },
        // { text: 'ID', align: 'start', value: 'id', sortable: false },
        { text: 'Status', value: 'status', sortable: false },
        { text: 'Amount', value: 'total', sortable: false },
        { text: 'Download', value: 'invoice_pdf', sortable: false }
      ]
    }
  },

  computed: {
    user () {
      return this.$store.getters['user/user']
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'Invoices')
    // this.listInvoices()
  }
}
</script>
