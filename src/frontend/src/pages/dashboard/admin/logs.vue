<template>
  <v-container
    id="regular-tables"
    fluid
    tag="section"
  >
    <paginator
      ref="paginator"
      title="Audit Logs"
      model="audit-logs"
      :headers="headers"
      :sort-by="['created_at']"
      :sort-desc="[true]"
    >
      <template #item.log="{item}">
        <div onclick="this.classList.toggle('json')">{{ JSON.stringify(item.log, undefined, 4) }}</div>
      </template>
      <template #item.created_at="{ item }">
        <span>{{ new Date(item.created_at * 1000).toLocaleString() }}</span>
      </template>
    </paginator>

    <div class="py-3" />
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
        { text: 'Created At', align: 'start', value: 'created_at' },
        { text: 'ID', value: 'id' },
        { text: 'Log', value: 'log', sortable: false }
      ]
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'AuditLogs')
  }
}
</script>
