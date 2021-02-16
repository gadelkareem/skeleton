export default function (to, from, savedPosition) {
  return savedPosition && to === from ? savedPosition : { x: 0, y: 0 }
}
