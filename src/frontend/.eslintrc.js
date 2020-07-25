module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: '@nuxtjs',
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'vue/singleline-html-element-content-newline': 'off',
  },
  parserOptions: {
    parser: 'babel-eslint',
  },
  ignorePatterns: ['src/dist', 'node_modules/', '.gitignore'],
}
