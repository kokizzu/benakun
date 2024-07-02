<script>
  import MainLayout from './_layouts/mainLayout.svelte';

  /** @typedef {import('./_components/types/transaction.js').TransactionTemplate} TransactionTemplate */

  let transactionTemplates = /** @type TransactionTemplate[] */ ([/* transactionTemplates */]);

  console.log('Transaction Templates:', transactionTemplates);
</script>

<MainLayout>
  <div class="menu_container">
    <p>TODO: fix query</p>
    {#each (transactionTemplates || []) as ttm, i (ttm.id)}
      <button class="menu_item" style="background-color: {ttm.color};">
        {ttm.name}
      </button>
    {/each}
  </div>
</MainLayout>

<style>
  .menu_container {
    display: flex;
    flex-direction: row;
    gap: 15px;
    flex-wrap: wrap;
  }

  .menu_container .menu_item {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 250px;
    height: 120px;
    border-radius: 10px;
    color: #FFF;
    border: none;
    font-size: var(--font-lg);
    font-weight: 600;
    cursor: pointer;
  }
</style>