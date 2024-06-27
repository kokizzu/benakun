<script>
  import MainLayout from '../../_layouts/mainLayout.svelte';
  
  /** @typedef {import('../../_components/types/inventory').InventoryChanges} InventoryChanges */
  /** @typedef {import('../../_components/types/product').Product} Product */

  let product = /** @type Product */ ({/* product */});
  let inventoryChanges = /** @type InventoryChanges[] */ ([/* inventoryChanges */]);

  console.log('Product:', product);
  console.log('Inventory Changes:', inventoryChanges);
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
          <p>{product.cogsIDR}</p>
        </div>
      </div>
    </div>
    <div class="inventory_changes_container">
      {#each (inventoryChanges || []) as invChange, _ (invChange.id)}
        <div>
          {invChange.stockDelta}
        </div>
      {/each}
    </div>
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

