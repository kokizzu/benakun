<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import IoSearch from 'svelte-icons-pack/io/IoSearch';

  let search = '';
</script>

<MainLayout>
  <div>
    <div class="search_container">
      <input id="search" placeholder="Search" bind:value={search} />
      <button>
        <Icon src={IoSearch} />
      </button>
    </div>
    <div>
      <button>Jual Barang</button>
      <button></button>
      <button></button>
      <button></button>
    </div>
  </div>
</MainLayout>

<style>
</style>