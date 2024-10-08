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

/**
 * @typedef {Object} DetailObjectTransaction
 * @property {number} salesCount
 * @property {string|number} salesPriceIDR
 */
module.exports = {};

/**
  * @typedef {Object} TransactionJournal
  * @property {number} id
  * @property {string} tenantCode
  * @property {number} coaId
  * @property {string|number} debitIDR
  * @property {string|number} creditIDR
  * @property {string} descriptions
  * @property {string|Date} date
  * @property {string|Object|DetailObjectTransaction} detailObj
  * @property {number} transactionTplId
  * @property {number} createdAt
  * @property {string} createdBy
  * @property {number} updatedAt
  * @property {string} updatedBy
  * @property {number} deletedAt
  * @property {string} deletedBy
  * @property {string} restoredBy
 */
module.exports = {};

/**
 * @typedef {Object} BusinessTransaction
 * @property {number} id
 * @property {string} tenantCode
 * @property {string|Date} startDate
 * @property {string|Date} endDate
 * @property {number} transactionTplId
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} deletedBy
 * @property {string} restoredBy
 */
module.exports = {};