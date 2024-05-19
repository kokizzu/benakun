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
  
  // Binding component PopUpAddBankAccount.svelte
  let popUpAddBankAccount = null;
  // For readiness of component PopUpAddBankAccount.svelte, prevent race condition
	let isPopUpAddBankAccountReady = false;

  onMount(() => {
    isPopUpAddBankAccountReady = true;
  })

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
        console.log(o);
        pager = o.pager;
        accounts = o.accounts;
        popUpAddBankAccount.Reset();
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
      
      CAN_EDIT_ROW={false}
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
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
