<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PoUpForms from './_components/PoUpForms.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminProducts } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/product.js').Product} Product */

  let segments = /** @type Access */ ({/* segments */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let user = /** @type User */ ({/* user */});
  let product = {/* product */}
  let products = /** @type any[][] */ ([/* products */]);

  let isPopUpFormsReady = false;
  let popUpForms = null;
  onMount(() => {
    isPopUpFormsReady = true;
  })

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await TenantAdminProducts( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminProductsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddProduct = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        pager = o.pager;
        products = o.products;

        console.log('Products:', products)
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = {
      pager,
      product: {
        id: row[0]
      },
      cmd: 'restore'
    };
    await TenantAdminProducts( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        products = o.products;
        notifier.showSuccess(row[1]+' restored')

        OnRefresh(pager);
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = {
      pager,
      product: {
        id: row[0]
      },
      cmd: 'delete'
    };
    await TenantAdminProducts( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        products = o.products;
        notifier.showSuccess(row[1]+' deleted')

        OnRefresh(pager);
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    /** @type Product */ // @ts-ignore
    const product = {
      id: payloads[0],
      name: payloads[1],
      detail: payloads[2],
      rule: payloads[3],
      kind: payloads[4],
      cogsIDR: payloads[5],
      profitPercentage: payloads[6],
    }
    const i = {
      pager, product,
      cmd: 'upsert'
    };
    await TenantAdminProducts( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBankAccountsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        products = o.products;
        notifier.showSuccess(product.name+' edited')

        OnRefresh(pager);
      }
    );
  }

  let isSubmitAddProduct = false;
  async function OnAddProduct(/** @type any[] */ payloads) {
    isSubmitAddProduct = true;
    /** @type Product */ // @ts-ignore
    const product = {
      name: payloads[1],
      detail: payloads[2],
      rule: payloads[3],
      kind: payloads[4],
      cogsIDR: payloads[5],
      profitPercentage: payloads[6],
    }
    const i = {
      pager, product,
      cmd: 'upsert'
    }
    await TenantAdminProducts( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminProductsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddProduct = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        products = o.products;

        notifier.showSuccess('product created')
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
    heading="Add product"
    FIELDS={fields}
    bind:isSubmitted={isSubmitAddProduct}
    OnSubmit={OnAddProduct}
  />
{/if}

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={products}
      
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
