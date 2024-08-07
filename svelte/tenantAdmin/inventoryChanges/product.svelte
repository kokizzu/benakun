<script>
  import MainLayout from '../../_layouts/mainLayout.svelte';
  import MasterTable from '../../_components/MasterTable.svelte';
  import { TenantAdminInventoryChangesProduct } from '../../jsApi.GEN';
  import { notifier } from '../../_components/notifier';
  
  /** @typedef {import('../../_components/types/master.js').Field} Field */
	/** @typedef {import('../../_components/types/access.js').Access} Access */
  /** @typedef {import('../../_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('../../_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('../../_components/types/user.js').User} User */
  /** @typedef {import('../../_components/types/inventory').InventoryChanges} InventoryChanges */
  /** @typedef {import('../../_components/types/product').Product} Product */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let product = /** @type Product */ ({/* product */});
  let inventoryChanges = /** @type any[][] */ ([/* inventoryChanges */]);
  let products = /** @type Object */ ({/* products */});

  let cogsIDR = new Intl.NumberFormat('id', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0, // to ensure no decimal places
    maximumFractionDigits: 0  // to ensure no decimal places
  }).format(Number(product.cogsIDR) || 0);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await TenantAdminInventoryChangesProduct( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminInventoryChangesCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
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
    await TenantAdminInventoryChangesProduct( // @ts-ignore
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
    await TenantAdminInventoryChangesProduct( // @ts-ignore
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
    await TenantAdminInventoryChangesProduct( // @ts-ignore
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
</script>

<MainLayout>
  <div class="product_inventory">
    <div class="product_container">
      <h4>Product detail</h4>
      <div class="product_detail">
        <div class="detail">
          <span class="label">Name</span>
          <span>:</span>
          <p>{product.name}</p>
        </div>
        <div class="detail">
          <span class="label">Detail</span>
          <span>:</span>
          <p>{product.detail}</p>
        </div>
        <div class="detail">
          <span class="label">Rule</span>
          <span>:</span>
          <p>{product.rule}</p>
        </div>
        <div class="detail">
          <span class="label">Kind</span>
          <span>:</span>
          <p>{product.kind}</p>
        </div>
        <div class="detail">
          <span class="label">Cogs IDR</span>
          <span>:</span>
          <p>{cogsIDR}</p>
        </div>
      </div>
    </div>
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

      {OnDelete}
      {OnEdit}
      {OnRefresh}
      {OnRestore}
    />
  </div>
</MainLayout>

<style>
  .product_inventory {
    display: flex;
    flex-direction: column;
    gap: 20px
  }
  
  .product_container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    background-color: var(--gray-002);
    padding: 20px;
    border-radius: 10px;
    width: 100%;
  }

  .product_container h4 {
    margin: 0;
    font-size: var(--font-lg);
  }

  .product_container .product_detail {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .product_container .product_detail .detail {
    display: grid;
    grid-template-columns: 150px 2px auto;
    gap: 5px;
  }

  .product_container .product_detail .detail .label {
    font-weight: 600;
  }

  .product_container .product_detail .detail p {
    word-break: break-all;
    margin: 0;
  }
</style>

