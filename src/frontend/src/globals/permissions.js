const routes = {
  admin: [
    '/dashboard/admin/users/',
    '/dashboard/admin/logs/'
  ]
}
const restricted = Object.values(routes).flat()
export default {
  canAccessRoute (route) {
    if (!route) {
      return true
    }
    if (!restricted.includes(route)) {
      return true
    }
    const user = this.$store.getters['user/user']
    if (!user.roles || !user.roles.length) {
      return false
    }
    for (const role of user.roles) {
      if (role in routes && routes[role].includes(route)) {
        return true
      }
    }
    return false
  }

}
