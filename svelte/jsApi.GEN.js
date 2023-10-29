
const axios = require("axios");


// rearrange response to be data first instead of axios error
function wrapErr( cb ) {
  return function( err ) {
    let data = ((err.response || {}).data || {})
    if( !data.error ) data.error = err.code
    data._axios = err
    cb( data )
  }
}

// rearrange response to be data first instead of axios error
function wrapOk( cb ) {
  return function( resp ) {
    let data = resp.data || {}
    data._axios = resp
    cb( data )
  }
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
/**
 * @typedef {Object} GuestAutoLoginIn
 * @property {String} uid
 * @property {String} token
 * @property {String} path
 */
const GuestAutoLoginIn = {
  uid: '', // string
  token: '', // string
  path: '', // string
}
/**
 * @typedef {Object} GuestAutoLoginOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.tenantCode
 * @property {Object} segments
 */
const GuestAutoLoginOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    tenantCode: '', // string
  }, // rqAuth.Users
  segments: { // M.SB
  }, // M.SB
}
/**
 * @callback GuestAutoLoginCallback
 * @param {GuestAutoLoginOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestAutoLoginIn} i
 * @param {GuestAutoLoginCallback} cb
 * @returns {Promise}
 */
exports.GuestAutoLogin = async function GuestAutoLogin( i, cb ) {
  return await axios.post( '/guest/autoLogin', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestDebugIn
 */
const GuestDebugIn = {
}
/**
 * @typedef {Object} GuestDebugOut
 * @property {Object} request
 */
const GuestDebugOut = {
  request: { // RequestCommon
  }, // RequestCommon
}
/**
 * @callback GuestDebugCallback
 * @param {GuestDebugOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestDebugIn} i
 * @param {GuestDebugCallback} cb
 * @returns {Promise}
 */
exports.GuestDebug = async function GuestDebug( i, cb ) {
  return await axios.post( '/guest/debug', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestExternalAuthIn
 * @property {String} provider
 * @property {Object} redirect
 */
const GuestExternalAuthIn = {
  provider: '', // string
  redirect: false, // bool
}
/**
 * @typedef {Object} GuestExternalAuthOut
 * @property {String} link
 * @property {String} clientID
 * @property {String} redirectUrl
 * @property {Array<String>} scopes
 * @property {String} csrfState
 */
const GuestExternalAuthOut = {
  link: '', // string
  clientID: '', // string
  redirectUrl: '', // string
  scopes: [], // []string
  csrfState: '', // string
}
/**
 * @callback GuestExternalAuthCallback
 * @param {GuestExternalAuthOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestExternalAuthIn} i
 * @param {GuestExternalAuthCallback} cb
 * @returns {Promise}
 */
exports.GuestExternalAuth = async function GuestExternalAuth( i, cb ) {
  return await axios.post( '/guest/externalAuth', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestForgotPasswordIn
 * @property {String} email
 */
const GuestForgotPasswordIn = {
  email: '', // string
}
/**
 * @typedef {Object} GuestForgotPasswordOut
 * @property {Object} ok
 * @property {String} resetPassUrl
 */
const GuestForgotPasswordOut = {
  ok: false, // bool
  resetPassUrl: '', // string
}
/**
 * @callback GuestForgotPasswordCallback
 * @param {GuestForgotPasswordOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestForgotPasswordIn} i
 * @param {GuestForgotPasswordCallback} cb
 * @returns {Promise}
 */
exports.GuestForgotPassword = async function GuestForgotPassword( i, cb ) {
  return await axios.post( '/guest/forgotPassword', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestLoginIn
 * @property {String} email
 * @property {String} password
 */
const GuestLoginIn = {
  email: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestLoginOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.tenantCode
 * @property {Object} segments
 */
const GuestLoginOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    tenantCode: '', // string
  }, // rqAuth.Users
  segments: { // M.SB
  }, // M.SB
}
/**
 * @callback GuestLoginCallback
 * @param {GuestLoginOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestLoginIn} i
 * @param {GuestLoginCallback} cb
 * @returns {Promise}
 */
exports.GuestLogin = async function GuestLogin( i, cb ) {
  return await axios.post( '/guest/login', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestOauthCallbackIn
 * @property {String} state
 * @property {String} code
 * @property {String} accessToken
 */
const GuestOauthCallbackIn = {
  state: '', // string
  code: '', // string
  accessToken: '', // string
}
/**
 * @typedef {Object} GuestOauthCallbackOut
 * @property {Object} oauthUser
 * @property {String} email
 * @property {number} currentUser.id
 * @property {String} currentUser.email
 * @property {String} currentUser.password
 * @property {number} currentUser.createdAt
 * @property {number} currentUser.createdBy
 * @property {number} currentUser.updatedAt
 * @property {number} currentUser.updatedBy
 * @property {number} currentUser.deletedAt
 * @property {number} currentUser.passwordSetAt
 * @property {String} currentUser.secretCode
 * @property {number} currentUser.secretCodeAt
 * @property {number} currentUser.verificationSentAt
 * @property {number} currentUser.verifiedAt
 * @property {number} currentUser.lastLoginAt
 * @property {String} currentUser.fullName
 * @property {String} currentUser.tenantCode
 * @property {String} provider
 * @property {Object} segments
 */
const GuestOauthCallbackOut = {
  oauthUser: { // M.SX
  }, // M.SX
  email: '', // string
  currentUser: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    tenantCode: '', // string
  }, // rqAuth.Users
  provider: '', // string
  segments: { // M.SB
  }, // M.SB
}
/**
 * @callback GuestOauthCallbackCallback
 * @param {GuestOauthCallbackOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestOauthCallbackIn} i
 * @param {GuestOauthCallbackCallback} cb
 * @returns {Promise}
 */
exports.GuestOauthCallback = async function GuestOauthCallback( i, cb ) {
  return await axios.post( '/guest/oauthCallback', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestRegisterIn
 * @property {String} email
 * @property {String} password
 */
const GuestRegisterIn = {
  email: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestRegisterOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.tenantCode
 * @property {String} verifyEmailUrl
 */
const GuestRegisterOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    tenantCode: '', // string
  }, // rqAuth.Users
  verifyEmailUrl: '', // string
}
/**
 * @callback GuestRegisterCallback
 * @param {GuestRegisterOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestRegisterIn} i
 * @param {GuestRegisterCallback} cb
 * @returns {Promise}
 */
exports.GuestRegister = async function GuestRegister( i, cb ) {
  return await axios.post( '/guest/register', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestResendVerificationEmailIn
 * @property {String} email
 */
const GuestResendVerificationEmailIn = {
  email: '', // string
}
/**
 * @typedef {Object} GuestResendVerificationEmailOut
 * @property {Object} ok
 * @property {String} verifyEmailUrl
 */
const GuestResendVerificationEmailOut = {
  ok: false, // bool
  verifyEmailUrl: '', // string
}
/**
 * @callback GuestResendVerificationEmailCallback
 * @param {GuestResendVerificationEmailOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestResendVerificationEmailIn} i
 * @param {GuestResendVerificationEmailCallback} cb
 * @returns {Promise}
 */
exports.GuestResendVerificationEmail = async function GuestResendVerificationEmail( i, cb ) {
  return await axios.post( '/guest/resendVerificationEmail', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestResetPasswordIn
 * @property {String} secretCode
 * @property {String} hash
 * @property {String} password
 */
const GuestResetPasswordIn = {
  secretCode: '', // string
  hash: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestResetPasswordOut
 * @property {Object} ok
 */
const GuestResetPasswordOut = {
  ok: false, // bool
}
/**
 * @callback GuestResetPasswordCallback
 * @param {GuestResetPasswordOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestResetPasswordIn} i
 * @param {GuestResetPasswordCallback} cb
 * @returns {Promise}
 */
exports.GuestResetPassword = async function GuestResetPassword( i, cb ) {
  return await axios.post( '/guest/resetPassword', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestVerifyEmailIn
 * @property {String} secretCode
 * @property {String} hash
 */
const GuestVerifyEmailIn = {
  secretCode: '', // string
  hash: '', // string
}
/**
 * @typedef {Object} GuestVerifyEmailOut
 * @property {Object} ok
 * @property {String} email
 */
const GuestVerifyEmailOut = {
  ok: false, // bool
  email: '', // string
}
/**
 * @callback GuestVerifyEmailCallback
 * @param {GuestVerifyEmailOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestVerifyEmailIn} i
 * @param {GuestVerifyEmailCallback} cb
 * @returns {Promise}
 */
exports.GuestVerifyEmail = async function GuestVerifyEmail( i, cb ) {
  return await axios.post( '/guest/verifyEmail', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserAutoLoginLinkIn
 * @property {String} path
 */
const UserAutoLoginLinkIn = {
  path: '', // string
}
/**
 * @typedef {Object} UserAutoLoginLinkOut
 * @property {String} link
 */
const UserAutoLoginLinkOut = {
  link: '', // string
}
/**
 * @callback UserAutoLoginLinkCallback
 * @param {UserAutoLoginLinkOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserAutoLoginLinkIn} i
 * @param {UserAutoLoginLinkCallback} cb
 * @returns {Promise}
 */
exports.UserAutoLoginLink = async function UserAutoLoginLink( i, cb ) {
  return await axios.post( '/user/autoLoginLink', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserChangePasswordIn
 * @property {String} oldPass
 * @property {String} newPass
 */
const UserChangePasswordIn = {
  oldPass: '', // string
  newPass: '', // string
}
/**
 * @typedef {Object} UserChangePasswordOut
 * @property {Object} ok
 */
const UserChangePasswordOut = {
  ok: false, // bool
}
/**
 * @callback UserChangePasswordCallback
 * @param {UserChangePasswordOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserChangePasswordIn} i
 * @param {UserChangePasswordCallback} cb
 * @returns {Promise}
 */
exports.UserChangePassword = async function UserChangePassword( i, cb ) {
  return await axios.post( '/user/changePassword', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserLogoutIn
 */
const UserLogoutIn = {
}
/**
 * @typedef {Object} UserLogoutOut
 * @property {number} logoutAt
 */
const UserLogoutOut = {
  logoutAt: 0, // int64
}
/**
 * @callback UserLogoutCallback
 * @param {UserLogoutOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserLogoutIn} i
 * @param {UserLogoutCallback} cb
 * @returns {Promise}
 */
exports.UserLogout = async function UserLogout( i, cb ) {
  return await axios.post( '/user/logout', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserManagementIn
 */
const UserManagementIn = {
}
/**
 * @typedef {Object} UserManagementOut
 */
const UserManagementOut = {
}
/**
 * @callback UserManagementCallback
 * @param {SuperadminUserManagementOut} o
 * @returns {Promise}
 */
/**
 * @param  {SuperadminUserManagementIn} i
 * @param {UserManagementCallback} cb
 * @returns {Promise}
 */
exports.UserManagement = async function UserManagement( i, cb ) {
  return await axios.post( '/', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserProfileIn
 */
const UserProfileIn = {
}
/**
 * @typedef {Object} UserProfileOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.tenantCode
 * @property {Object} segments
 */
const UserProfileOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    tenantCode: '', // string
  }, // rqAuth.Users
  segments: { // M.SB
  }, // M.SB
}
/**
 * @callback UserProfileCallback
 * @param {UserProfileOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserProfileIn} i
 * @param {UserProfileCallback} cb
 * @returns {Promise}
 */
exports.UserProfile = async function UserProfile( i, cb ) {
  return await axios.post( '/user/profile', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserSessionKillIn
 * @property {String} sessionTokenHash
 */
const UserSessionKillIn = {
  sessionTokenHash: '', // string
}
/**
 * @typedef {Object} UserSessionKillOut
 * @property {number} logoutAt
 * @property {number} sessionTerminated
 */
const UserSessionKillOut = {
  logoutAt: 0, // int64
  sessionTerminated: 0, // int64
}
/**
 * @callback UserSessionKillCallback
 * @param {UserSessionKillOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserSessionKillIn} i
 * @param {UserSessionKillCallback} cb
 * @returns {Promise}
 */
exports.UserSessionKill = async function UserSessionKill( i, cb ) {
  return await axios.post( '/user/sessionKill', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserSessionsActiveIn
 */
const UserSessionsActiveIn = {
}
/**
 * @typedef {Object} UserSessionsActiveOut
 * @property {Object} sessionsActive
 */
const UserSessionsActiveOut = {
  sessionsActive: { // []rqAuth.Sessions
  }, // []rqAuth.Sessions
}
/**
 * @callback UserSessionsActiveCallback
 * @param {UserSessionsActiveOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserSessionsActiveIn} i
 * @param {UserSessionsActiveCallback} cb
 * @returns {Promise}
 */
exports.UserSessionsActive = async function UserSessionsActive( i, cb ) {
  return await axios.post( '/user/sessionsActive', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserUpdateProfileIn
 * @property {String} userName
 * @property {String} fullName
 * @property {String} email
 * @property {String} country
 * @property {String} language
 */
const UserUpdateProfileIn = {
  userName: '', // string
  fullName: '', // string
  email: '', // string
  country: '', // string
  language: '', // string
}
/**
 * @typedef {Object} UserUpdateProfileOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} user.fullName
 * @property {String} user.tenantCode
 * @property {Object} segments
 */
const UserUpdateProfileOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
    fullName: '', // string
    tenantCode: '', // string
  }, // rqAuth.Users
  segments: { // M.SB
  }, // M.SB
}
/**
 * @callback UserUpdateProfileCallback
 * @param {UserProfileOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserUpdateProfileIn} i
 * @param {UserUpdateProfileCallback} cb
 * @returns {Promise}
 */
exports.UserUpdateProfile = async function UserUpdateProfile( i, cb ) {
  return await axios.post( '/user/updateProfile', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}


// Code generated by 1_codegen_test.go DO NOT EDIT.