/**
 * @typedef {Object} User
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} deletedAt
 * @property {string} email
 * @property {string} fullName
 * @property {string} id
 * @property {string} invitationState
 * @property {number} lastLoginAt
 * @property {string} password
 * @property {number} passwordSetAt
 * @property {string} role
 * @property {string} secretCode
 * @property {number} secretCodeAt
 * @property {string} tenantCode
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} verificationSentAt
 * @property {number} verifiedAt
 * @property {number} supportExpiredAt
 */
module.exports = {};

/**
 * @typedef {Object} Session
 * @property {string} sessionToken
 * @property {number} userId
 * @property {number} expiredAt
 * @property {string} device
 * @property {number} loginAt
 * @property {string} loginIPs
 * @property {string} tenantCode
 */
module.exports = {};