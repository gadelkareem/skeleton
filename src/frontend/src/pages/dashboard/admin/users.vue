<template>
  <v-container
    id="regular-tables"
    fluid
    tag="section"
  >
    <paginator
      ref="paginator"
      title="Users"
      model="users"
      :headers="headers"
    >
      <template #item.username="{item}">
        <router-link :to="{path:'/dashboard/admin/user/', query:{id:item.id}}">
          {{ item.username }}
        </router-link>
      </template>
      <template #item.first_name="{item}">
        {{ item.first_name }} {{ item.last_name }}
      </template>
      <template #item.id="{item}">
        <router-link :to="{path:'/dashboard/admin/user/', query:{id:item.id}}">
          {{ item.id }}
        </router-link>
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
        { text: 'ID', align: 'start', value: 'id' },
        { text: 'Username', value: 'username' },
        { text: 'Name', value: 'first_name' },
        { text: 'Email', value: 'email' },
        { text: 'country', value: 'country', sortable: false },
        { text: 'City', value: 'city', sortable: false }
      ]
    }
  },
  mounted () {
    this.$store.dispatch('page/title', 'AuditLogs')
  }
}
</script>
