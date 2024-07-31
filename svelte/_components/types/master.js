/**
 * @typedef { 'text' | 'textarea' | 'email' | 'password'
 * | 'number' | 'phone' | 'date' | 'bool' | 'checkbox' | 'combobox'
 * | 'select' | 'percentage' | 'float' | 'color' | 'datetime' } InputType
*/
module.exports = {};

/**
 * @typedef {Object} Field
 * @property {string} name
 * @property {string} label
 * @property {string} description
 * @property {string} type
 * @property {InputType} inputType
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

/**
 * @typedef {Object} ExtendedAction 
 * @property {import("svelte-icons-pack").IconType} icon
 * @property {boolean} isTargetBlank - if true, open link in new window
 * @property {(row: any) => void} link
 * @property {string} tooltip
 */
module.exports = {}