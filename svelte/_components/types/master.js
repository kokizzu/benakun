/**
 * @typedef {Object} Field
 * @property {string} name
 * @property {string} label
 * @property {string} description
 * @property {string} type
 * @property {string} inputType
 * @property {boolean} readOnly
 * @property {string[]} validations
 * @property {string[]} ref
 * @property {string} refEndpoint
 */
module.exports = {};

/**
 * @typedef {Object} PagerIn
 * @property {number} page
 * @property {number} perPage
 * @property {Object} filters
 * @property {string[]} order
 */
module.exports = {};

/**
 * @typedef {Object} PagerOut
 * @property {number} page
 * @property {number} perPage
 * @property {number} pages
 * @property {number} countResult
 * @property {Object} filters
 * @property {string[]} order
 */
module.exports = {};

// TODO: add function for return row
/**
 * @typedef {Object} ExtendedAction 
 * @property {import("svelte-icons-pack").IconType} icon 
 * @property {string} name 
 * @property {boolean} isTargetBlank - if true, open link in new window
 * @property {string} path - default window.location.pathname
 */
module.exports = {}