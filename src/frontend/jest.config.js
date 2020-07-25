module.exports = {
  moduleFileExtensions: [
    'js',
    'jsx',
    'json',
    'vue'
  ],
  watchman: false,
  moduleNameMapper: {
    '^~/(.*)$': '<rootDir>/$1',
    '^~~/(.*)$': '<rootDir>/src/$1',
    '^@/(.*)$': '<rootDir>/$1',
    '^@@/(.*)$': '<rootDir>/src/$1'
  },
  transform: {
    // process js with `babel-jest`
    '^.+\\.[jt]sx?$': '<rootDir>/node_modules/babel-jest',
    // process `*.vue` files with `vue-jest`
    '.*\\.(vue)$': '<rootDir>/node_modules/vue-jest',
    '.+\\.(css|styl|less|sass|scss|svg|png|jpg|ttf|woff|woff2)$': '<rootDir>/node_modules/jest-transform-stub'
  },
  snapshotSerializers: ['<rootDir>/node_modules/jest-serializer-vue'],
  collectCoverage: false,
  collectCoverageFrom: [
    '<rootDir>/src/components/**/*.vue',
    '<rootDir>/src/pages/*.vue'
  ],
  'verbose': true,
  'modulePaths': [
    '<rootDir>/src',
    '<rootDir>/node_modules'
  ],
  'globals': {
    'NODE_ENV': 'test'
  },
  setupFiles: ['<rootDir>/test/_setup.js'],
  transformIgnorePatterns: [
    '<rootDir>/node_modules/(?!vuetify)'
  ]
}
