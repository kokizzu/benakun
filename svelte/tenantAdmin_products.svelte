<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PoUpForms from './_components/PoUpForms.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminProducts } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/product.js').Product} Product */

  let segments = /** @type Access */ ({/* segments */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let product = {/* product */}
  let products = /** @type any[][] */ ([/* products */]);

  let isPopUpFormsReady = false;
  let popUpForms = null;
  onMount(() => {
    isPopUpFormsReady = true;
  })

  console.log('Segments:', segments);
  console.log('Fields:', fields);
  console.log('Pager:', pager);
  console.log('Product:', product);
  console.log('Products:', products);

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
        console.log(o);
        pager = o.pager;
        products = o.products;
        notifier.showSuccess('product created')
        popUpForms.Reset();
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
    bind:isSubmitAddProduct
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
    >
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
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>
