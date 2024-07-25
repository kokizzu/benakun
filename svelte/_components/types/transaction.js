/**
  * @typedef {Object} TransactionTplDetail
  * @property {number} id
  * @property {number} parentId
  * @property {number} coaId
  * @property {string} tenantCode
  * @property {boolean} isDebit
  * @property {string[]} attributes
  * @property {number} createdAt
  * @property {string} createdBy
  * @property {number} updatedAt
  * @property {string} updatedBy
  * @property {number} deletedAt
  * @property {string} deletedBy
  * @property {string} restoredBy
  */

/**
  * @typedef {Object} TransactionTemplate
  * @property {number} id
  * @property {string} tenantCode
  * @property {string} name
  * @property {string} color
  * @property {string} imageURL
  * @property {number} createdAt
  * @property {string} createdBy
  * @property {number} updatedAt
  * @property {string} updatedBy
  * @property {number} deletedAt
  * @property {string} deletedBy
  * @property {string} restoredBy
  * @property {TransactionTplDetail[]} details
 */
module.exports = {};