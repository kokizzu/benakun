<script>
  /** @typedef {import('./types/coa').SelectCoa} SelectCoa */

  import { onMount } from 'svelte';
  import Select from '../node_modules/svelte-select';

  export let label = /** @type {string} */ ('');
  export let values = /** Record<string|number, string> */ ({});
  export let value = /** @type {SelectCoa} */ ({
    value: 'cake', label: 'Cake'
  });

  let items = /** @type {SelectCoa[]} */ ([]);

  /** @returns {number[]} */
  function extractNumParts(/** @type {string} */ s) {
    const parts = /** @type {string[]} */ (s.split(` `));
    const numberStr = /** @type {string} */ (parts[0]);
    const numberParts = /** @type {string[]} */ (numberStr.split(`.`));

    let result = /** @type {number[]} */ ([]);
    for (const np of numberParts) {
      const num = Number(np);
      if (!isNaN(num)) result = [...result, num];
    }

    return result;
  }

  /** @returns {boolean} */
  function compareCoaNums(
    /** @type {number[]} */ va, /** @type {number[]} */ vb
  ) {
    for (let i = 0; i < va.length && i < vb.length; i++) {
      if (va[i] != vb[i]) {
        return va[i] < vb[i];
      } 
    }

    return va.length < vb.length;
  }

  onMount(() => {
    Object.keys(values).forEach(key => {
      items = [...items, { value: key, label: values[key] }];
    });

    if (items.length > 0) {
      if (!value) value = items[0];
      items.sort((a, b) => {
        const va = extractNumParts(a.label);
        const vb = extractNumParts(b.label);

        return compareCoaNums(va, vb) ? -1 : 1;
      })
    }
  })
</script>

<div class="input_coa">
  <span class="label">{label}</span>
  <Select
    class="select_coa"
    items={items}
    bind:value
  />
</div>

<style>
  .input_coa {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input_coa .label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  :global(.input_coa .select_coa) {
    width: 100%;
  }

  :global(.input_coa .select_coa .selected-item) {
    overflow: hidden !important;
    display: -webkit-box !important;
    -webkit-box-orient: vertical !important;
    -webkit-line-clamp: 1 !important;
    line-clamp: 1 !important;
  }
</style>