
export default function ({ store, route, redirect }) {
  if (!store.getters['auth/isAuthenticated']) {
    return redirect({ path: '/auth/login/', query: { to: route.fullPath } })
  }
}
