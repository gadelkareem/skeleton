module.exports = {
  resolve: {
    // for IDE (WebStorm, PyCharm, etc)
    alias: {
      '@': require('path').resolve(__dirname, ''),
      '@@': require('path').resolve(__dirname, 'src')
    }
  }
}
