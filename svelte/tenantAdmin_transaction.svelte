<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { IoSearch } from './node_modules/svelte-icons-pack/dist/io';
  import MainLayout from './_layouts/mainLayout.svelte';

  let search = '';
</script>

<MainLayout>
  <div class="budgeting">
    <div class="search_container">
      <input id="search" placeholder="Search" bind:value={search} />
      <button>
        <Icon src={IoSearch} />
      </button>
    </div>
    <div class="tx_templates">
      <button class="template">Jual Barang</button>
      <button class="template">Pembelian Aset</button>
      <button class="template">Penyusutan</button>
    </div>
  </div>
</MainLayout>

<style>
  .budgeting {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .budgeting .tx_templates {
    display: flex;
    flex-direction: row;
    gap: 10px;
  }

  .budgeting .tx_templates .template {
    padding: 30px;
    cursor: pointer;
  }
</style>