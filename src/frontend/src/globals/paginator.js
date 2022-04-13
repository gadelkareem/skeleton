export default {
  queryFromOptions (options, filter, query) {
    return {
      page: {
        size: options.itemsPerPage || query.itemsPerPage || 0,
        after: options.page || query.page || 0
      },
      sort: this.getSort(options) || this.getSort(query) || '',
      filter: filter || query.filter || ''
    }
  },
  optionsFromAPI (page, sort) {
    const s = this.parseSort(sort)
    return {
      itemsPerPage: Number.parseInt(page.size),
      page: Number.parseInt(page.before),
      sortBy: s.sortBy,
      sortDesc: s.sortDesc
    }
  },
  parseSort (s) {
    const sort = { sortDesc: [], sortBy: [] }
    if (!s) {
      return sort
    }
    const arr = s.split(',')
    for (s of arr) {
      if (s.startsWith('-')) {
        sort.sortDesc.push(true)
        s = s.substr(1)
      } else {
        sort.sortDesc.push(false)
      }
      sort.sortBy.push(s)
    }
    return sort
  },
  getSort (options) {
    if (!options || !options.sortBy) {
      return ''
    }
    let sort = ''
    options.sortBy = Array.isArray(options.sortBy) ? options.sortBy : [options.sortBy]
    options.sortDesc = Array.isArray(options.sortDesc) ? options.sortDesc : [options.sortDesc]
    for (const i in options.sortBy) {
      if (options.sortBy[i] && Object.prototype.toString.call(options.sortBy[i]) === '[object String]') {
        sort += options.sortDesc[i] ? '-' : ''
        sort += options.sortBy[i]
      }
    }
    return sort
  }
}
