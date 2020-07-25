import api from './api'

export default {
  register (user) {
    return api.post('user', user)
  },
  forgotPassword (body) {
    return api.request({
      url: '/users/forgot-password',
      method: 'POST',
      body,
      type: 'reset-password-request'
    })
  },
  resetPassword (body) {
    return api.request({
      url: '/users/reset-password',
      method: 'POST',
      body,
      type: 'reset-password-request'
    })
  },
  verifyEmail (body) {
    return api.request({
      url: '/users/verify-email',
      method: 'POST',
      body,
      type: 'email-verify-request'
    })
  },
  get (id) {
    return api.fetch('users/' + id)
  },
  list (params) {
    return api.fetch('user', params)
  },
  update (user) {
    return api.patch('user', user)
  },
  delete (id) {
    return api.remove('user', id)
  },
  changePassword (id, oldPassword, password) {
    return api.request({
      url: '/users/' + id + '/password',
      method: 'PATCH',
      body: {
        id,
        old_pass: oldPassword,
        password
      },
      type: 'update-password-request'
    })
  },
  generateAuthenticator (id, body) {
    body.id = id
    return api.request({
      url: '/users/' + id + '/generate-auth-code',
      method: 'PATCH',
      body,
      type: 'authenticator-request'
    })
  },
  authenticator (id, body) {
    body.id = id
    return api.request({
      url: '/users/' + id + '/authenticator',
      method: 'PATCH',
      body,
      type: 'authenticator-request'
    })
  },
  sendSMS (id) {
    return api.request({
      url: '/users/' + id + '/send-verify-sms',
      method: 'PATCH',
      body: { id },
      type: 'mobile-verify-request'
    })
  },
  verifyMobile (id, body) {
    body.id = id
    return api.request({
      url: '/users/' + id + '/verify-mobile',
      method: 'PATCH',
      body,
      type: 'mobile-verify-request'
    })
  },
  recoveryQuestions (id, questions) {
    return api.request({
      url: '/users/' + id + '/recovery-questions',
      method: 'PATCH',
      body: { id, questions },
      type: 'recovery-questions-request'
    })
  },
  getRecoveryQuestions (body) {
    return api.request({
      url: '/users/recovery-questions',
      method: 'POST',
      body,
      type: 'login-requests'
    })
  },
  disableMFA (body) {
    return api.request({
      url: '/users/disable-mfa',
      method: 'POST',
      body,
      type: 'mfa-disable-request'
    })
  }
}
