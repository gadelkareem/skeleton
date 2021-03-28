export default () => ({
  user: {
    id: 0,
    username: '',
    password: '',
    email: '',
    first_name: '',
    last_name: '',
    mobile: '',
    language: '',
    avatar_url: '',
    address: {
      street: '',
      city: '',
      zip_code: ''
    },
    country: '',
    social_login: '',
    authenticator_enabled: false,
    roles: null,
    active: true,
    mobile_verified: false,
    recovery_questions_set: false,
    notifications: []
  }
})
