<script>
  import MainLayout from './_layouts/mainLayout.svelte';

  /** @typedef {import('./_components/types/transaction.js').TransactionTemplate} TransactionTemplate */

  let transactionTemplates = /** @type TransactionTemplate[] */ ([/* transactionTemplates */]);

  console.log('Transaction Templates:', transactionTemplates);
</script>

<MainLayout>
  <div class="menu_container">
    {#each (transactionTemplates || []) as ttm, i (ttm.id)}
      <a class="menu_item" style="background-color: {ttm.color};" href="/dataEntry/template/{ttm.id}">
        {ttm.name}
      </a>
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
    text-decoration: none;
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
    transition: 0.3s;
  }

  .menu_container .menu_item:hover {
    font-size: 20px;
  }

  @media only screen and (max-width : 768px) {
    .menu_container {
      display: grid;
      grid-template-columns: 1fr 1fr;
      grid-gap: 10px;
    }

    .menu_container .menu_item {
      width: 100%;
    }
  }
</style>