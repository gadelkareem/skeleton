export default {
  required: v => !!v || 'This field is required',
  username: v => (!v || (/^(?=.{4,50}$)[A-Za-z0-9]+(?:[-_.][A-Za-z0-9]+)*$/.test(v))) || 'Username should be between 4 and 50 characters and can only contains alphanumeric characters, underscore and dot.',
  password: v => (!v || (v && v.length > 6 && v.length < 100)) || 'Password should be between 6 and 100 characters.',
  repeatPassword: pass => v => (!v || v === pass) || 'Repeat Password does not match Password.',
  email: v => (!v || (/^(?=.{4,100}$)\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,6})+$/.test(v))) || 'Invalid Email address.',
  mobile: v => (!v || (/^[0-9+()\-\s]{10,}$/.test(v))) || 'Invalid Mobile number.',
  name: v => (!v || (v && v.length > 2 && v.length < 100)) || 'Enter a string between 2 and 100 characters.',
  numeric: v => (!v || (v && !Number.isInteger(v))) || 'Enter numbers only.',
  mfa: v => (!v || (v && (v.length === 6))) || 'Enter 6 digits code.'
}
