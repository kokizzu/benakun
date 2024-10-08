<script>
  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/master.js').ExtendedAction} ExtendedAction */
  
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import { FaShareFromSquare } from './node_modules/svelte-icons-pack/dist/fa';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PopUpInventoryChanges from './_components/PopUpInventoryChanges.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminInventoryChanges } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier.js';

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let inventoryChanges = /** @type any[][] */ ([/* inventoryChanges */]);
  let products = /** @type Object */ ({/* products */});

  let isPopUpFormsReady = false;
  let popUpInventoryChanges = null;
  onMount(() => {
    isPopUpFormsReady = true;
  })

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await TenantAdminInventoryChanges( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminInventoryChangesCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddInvChange = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        inventoryChanges = o.inventoryChanges;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = {
      pager,
      inventoryChange: {
        id: row[0]
      },
      cmd: 'restore'
    };
    await TenantAdminInventoryChanges( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        inventoryChanges = o.inventoryChanges;
        notifier.showSuccess(row[1]+' restored')

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = {
      pager,
      inventoryChange: {
        id: row[0]
      },
      cmd: 'delete'
    };
    await TenantAdminInventoryChanges( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        inventoryChanges = o.inventoryChanges;
        notifier.showSuccess('inventoryChange deleted')

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const inventoryChange = {
      id: id,
      stockDelta: payloads[1],
    }

    const i = {
      pager, inventoryChange,
      cmd: 'upsert'
    };
    await TenantAdminInventoryChanges( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        inventoryChanges = o.inventoryChanges;
        notifier.showSuccess('inventoryChange edited')

        OnRefresh(pager);
      }
    );
  }

  let isSubmitAddInvChange = false;
  async function OnAddInventoryChange(/** @type any */ payloads) {
    isSubmitAddInvChange = true;
    const i = {
      pager,
      inventoryChange: payloads,
      cmd: 'upsert'
    }
    await TenantAdminInventoryChanges( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminInventoryChangesCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddInvChange = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        inventoryChanges = o.inventoryChanges;

        notifier.showSuccess('inventoryChange created')
        popUpInventoryChanges.Reset();

        OnRefresh(pager);
      }
    );
    popUpInventoryChanges.Hide();
  }

  /** @type {ExtendedAction[]} */
  const LINKS = [
    {
      icon: FaShareFromSquare,
      isTargetBlank: true,
      link: (/** @type any */ row) => `/tenantAdmin/inventoryChanges/${row[2]}`,
      tooltip: 'View',
    }
  ]
</script>

{#if isPopUpFormsReady}
  <PopUpInventoryChanges
    {products}
    bind:this={popUpInventoryChanges}
    bind:isSubmitAddInventoryChanges={isSubmitAddInvChange}
    OnSubmit={OnAddInventoryChange}
  />
{/if}

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={inventoryChanges}
      REFS={{
        'productId': products
      }}

      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
      CAN_OPEN_LINK
      
      LINKS={LINKS}

      {OnDelete}
      {OnEdit}
      {OnRefresh}
      {OnRestore}
    >
      {#if user.tenantCode !== ''}
        <button
          class="action_btn"
          on:click={() => popUpInventoryChanges.Show()}
          title="add inventoryChange"
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