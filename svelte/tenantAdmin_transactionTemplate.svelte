<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PoUpForms from './_components/PoUpForms.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminTransactionTemplate } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/transaction.js').TransactionTemplate} TransactionTemplate */

  let segments = /** @type Access */ ({/* segments */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let user = /** @type User */ ({/* user */});
  let transactionTemplates = /** @type any[][] */ ([/* transactionTemplates */]);
  let trxTemplate = /** @type TransactionTemplate */ ({/* transactionTemplate */});

  let isPopUpFormsReady = false;
  let popUpForms = null;
  onMount(() => {
    isPopUpFormsReady = true;
  })

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await TenantAdminTransactionTemplate( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminTransactionTemplateCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddTrxTemplate = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        transactionTemplates = o.transactionTemplates;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = {
      pager,
      transactionTemplate: {
        id: row[0]
      },
      cmd: 'restore'
    };
    await TenantAdminTransactionTemplate( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        transactionTemplates = o.transactionTemplates;
        notifier.showSuccess(row[1]+' restored')

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = {
      pager,
      transactionTemplate: {
        id: row[0]
      },
      cmd: 'delete'
    };
    await TenantAdminTransactionTemplate( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        transactionTemplates = o.transactionTemplates;
        notifier.showSuccess(row[1]+' deleted')

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    /** @type TransactionTemplate */ // @ts-ignore
    const transactionTemplate = {
      id: payloads[0],
      name: payloads[1],
    }
    const i = {
      pager, transactionTemplate,
      cmd: 'upsert'
    };
    await TenantAdminTransactionTemplate( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        transactionTemplates = o.transactionTemplates;
        notifier.showSuccess(transactionTemplate.name+' edited')

        OnRefresh(pager);
      }
    );
  }

  let isSubmitAddTrxTemplate = false;
  async function OnAddTransactionTemplate(/** @type any[] */ payloads) {
    isSubmitAddTrxTemplate = true;
    /** @type TransactionTemplate */ // @ts-ignore
    const transactionTemplate = {
      id: 0,
      name: payloads[1],
      color: payloads[2],
    }
    const i = {
      pager, transactionTemplate,
      cmd: 'upsert'
    }
    await TenantAdminTransactionTemplate( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminTransactionTemplateCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddTrxTemplate = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        transactionTemplates = o.transactionTemplates;

        notifier.showSuccess('transaction template created')
        popUpForms.Reset();

        OnRefresh(pager);
      }
    );
    popUpForms.Hide();
  }
</script>


{#if isPopUpFormsReady}
  <PoUpForms
    bind:this={popUpForms}
    heading="Add transaction template"
    FIELDS={fields}
    bind:isSubmitted={isSubmitAddTrxTemplate}
    OnSubmit={OnAddTransactionTemplate}
  />
{/if}

<MainLayout>
  <div>
    <div>Change UI to tree 2 level</div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={transactionTemplates}

      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnDelete}
      {OnRestore}
      {OnRefresh}
      {OnEdit}
    >
      {#if user.tenantCode !== ''}
        <button
          class="action_btn"
          on:click={() => popUpForms.Show()}
          title="add account"
        >
          <Icon
            color="var(--gray-007)"
            size="16"
            src={RiSystemAddBoxLine}
          />
        </button>
      {/if}
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>
