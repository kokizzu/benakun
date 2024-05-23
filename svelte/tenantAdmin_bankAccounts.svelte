<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PopUpAddBankAccount from './_components/PopUpAddBankAccount.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminBankAccounts } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/tenant.js').BankAccount} BankAccount */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let accounts = /** @type any[][] */ ([/* accounts */]);

  console.log('Fields:', fields);
  console.log('Accounts:', accounts);

  // Binding component PopUpAddBankAccount.svelte
  let popUpAddBankAccount = null;
  // For readiness of component PopUpAddBankAccount.svelte, prevent race condition
	let isPopUpAddBankAccountReady = false;

  onMount(() => {
    isPopUpAddBankAccountReady = true;
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await TenantAdminBankAccounts( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        accounts = o.accounts;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = {
      pager,
      account: {
        id: row[0]
      },
      cmd: 'restore'
    };
    await TenantAdminBankAccounts( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return;
        }

        notifier.showSuccess('account '+row[1]+' restored');
        accounts = o.accounts;
        pager = o.pager;
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = {
      pager,
      account: {
        id: row[0]
      },
      cmd: 'delete'
    };
    await TenantAdminBankAccounts( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        notifier.showSuccess('account '+row[1]+' deleted');
        accounts = o.accounts;
        pager = o.pager;
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const i = {
      pager,
      account: {
        id: payloads[0],
        name: payloads[1],
        accountName: payloads[2],
        accountNumber: payloads[3],
        bankName: payloads[4],
        isProfitCenter: payloads[5],
        isCostCenter: payloads[6]
      },
      cmd: 'upsert'
    };
    await TenantAdminBankAccounts( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        accounts = o.accounts;
        notifier.showSuccess(payloads[1]+' edited')
        popUpAddBankAccount.Reset();
      }
    );
  }

  let isSubmitAddBankAccount = false;
  async function addAccount(/** @type BankAccount*/ bankAccount) {
    isSubmitAddBankAccount = true;
    const i = {
      pager: pager,
      account: bankAccount,
      cmd: 'upsert'
    }
    await TenantAdminBankAccounts( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddBankAccount = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        accounts = o.accounts;
        notifier.showSuccess('bank account created')
        popUpAddBankAccount.Reset();
        
        OnRefresh(pager);
      }
    );
    popUpAddBankAccount.Hide();
  }
</script>

{#if isPopUpAddBankAccountReady}
  <PopUpAddBankAccount
    bind:this={popUpAddBankAccount}
    bind:isSubmitAddBankAccount
    OnSubmit={addAccount}
  />
{/if}

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={accounts}
      
      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnDelete}
      {OnRestore}
      {OnRefresh}
      {OnEdit}
    >
    <button
      class="action_btn"
      on:click={() => popUpAddBankAccount.Show()}
      title="add account"
    >
      <Icon
        color="var(--gray-007)"
        size="16"
        src={RiSystemAddBoxLine}
      />
    </button>
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>
