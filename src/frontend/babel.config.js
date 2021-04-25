module.exports = {
  presets: [
    ['@vue/app',
      {
        useBuiltIns: 'entry'
      }]
  ],
  plugins: [
    ['@babel/plugin-transform-runtime',
      {
        'regenerator': true
      }
    ]
  ]
}
