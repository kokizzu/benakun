<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import PopUpTransactionTemplate from './_components/PopUpTransactionTemplate.svelte';
  import { TenantAdminTransactionTemplate } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier.js';
  import TransactionTemplateTree from './_components/TransactionTemplateTree.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/transaction.js').TransactionTemplate} TransactionTemplate */
  /** @typedef {import('./_components/types/coa.js').CoA} CoA */

  let transactionTemplates = /** @type {TransactionTemplate[]} */ ([/* transactionTemplates */]);
  let coas = /** @type {CoA[]} */ ([/* coas */]);

  let popUpTransactionTemplate;
  let name = '';
  let color = '';
  let isSubmitted = false;

  async function addTransactionTemplate() {
    isSubmitted = true;
    const i = {
      cmd: 'upsert',
      transactionTemplate: { name, color }
    }
    await TenantAdminTransactionTemplate( //@ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminTransactionTemplateCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitted = false;
        name = '';
        color = '';
        popUpTransactionTemplate.hide();
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        transactionTemplates = o.transactionTemplates;

        notifier.showSuccess('transaction template created');
      }
    )
  }

  function updateTrxTemplate(event) {
    const { trxTemplates } = event.detail;
    transactionTemplates = trxTemplates;
  }
</script>

<PopUpTransactionTemplate
  bind:this={popUpTransactionTemplate}
  heading="Add Transaction Template"
  bind:isSubmitted
  bind:name
  bind:color
  onSubmit={addTransactionTemplate}
/>

<MainLayout>
  <div class="container">
    <header class="header_options">
      <button class="add" on:click={() => popUpTransactionTemplate.show()}>
        Add
      </button>
    </header>
    <div class="transaction_templates">
      {#each (transactionTemplates || []) as tt}
        <TransactionTemplateTree
          transactionTemplate={tt}
          coas={coas}
          on:update={updateTrxTemplate}
        />
      {/each}
    </div>
  </div>
</MainLayout>

<style>
  .container {
    display: flex;
    flex-direction: column;
    gap: 15px;
    width: 100%;
    background-color: #FFF;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    overflow: hidden;
    padding: 15px 15px 20px;
  }

  .container .header_options {
    display: flex;
    flex-direction: row-reverse;
    padding-bottom: 10px;
    border-bottom: 1px solid var(--gray-003);
  }

  .container .header_options .add {
    border: none;
    padding: 10px 20px;
    border-radius: 8px;
    background-color: var(--blue-006);
    color: #FFF;
    font-size: var(--font-lg);
    cursor: pointer;
  }

  .container .header_options .add:hover {
    background-color: var(--blue-005);
  }

  .container .transaction_templates {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
</style>
