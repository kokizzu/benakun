<script>
  /**
     * @typedef {Object} CoA
     * @property {string} id
     * @property {string} name
     * @property {number} level
     * @property {string} parentId
     * @property {Array<CoA>} children
    */
  /**
   * @type {CoA}
   */
  export let coa = {
    id: '',
    name: '',
    level: 0,
    parentId: '',
    children: []
  };
</script>

<div>
  {#if coa.children && coa.children.length}
    {#each coa.children as ch}
      <svelte:self {...ch} />
    {/each}
  {:else}
    <div>
      <!-- render -->
    </div>
  {/if}
</div>

