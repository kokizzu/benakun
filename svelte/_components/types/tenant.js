/**
  * @typedef {Object} Tenant
  * @property {string} id
  * @property {string} tenantCode
  * @property {number} createdAt
  * @property {string} createdBy
  * @property {number} updatedAt
  * @property {string} updatedBy
  * @property {number} deletedAt
  * @property {number} productsCoaId
  * @property {number} suppliersCoaId
  * @property {number} customersCoaId
  * @property {number} customerReceivablesCoaId
  * @property {number} staffsCoaId
  * @property {number} banksCoaId
  * @property {number} fundersCoaId
  */
module.exports = {};

/**
  * @typedef {Object} BankAccount
  * @property {string} id
  * @property {string} tenantCode
  * @property {number} createdAt
  * @property {string} createdBy
  * @property {number} updatedAt
  * @property {string} updatedBy
  * @property {number} deletedAt
  * @property {string} name
  * @property {string} parentBankAccountId
  * @property {string} childBankAccountId
  * @property {number} accountNumber
  * @property {string} bankName
  * @property {string} accountName
  * @property {boolean} isProfitCenter
  * @property {boolean} isCostCenter
  * @property {boolean} isFunder 
  * @property {string} staffId
  * @property {string|number} coaIdProfitCenter
  * @property {string|number} coaIdCostCenter
  */
module.exports = {};