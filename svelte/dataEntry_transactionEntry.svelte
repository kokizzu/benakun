<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import InputCustom from './_components/InputCustom.svelte';

  let coa = '';
</script>

<MainLayout>
  <p>Transction entry</p>
  <div>
    <InputCustom
      id="coa"
      label="CoA (Chart of Account)"
      type="text"
      placeholder="CoA"
      bind:value={coa}
    />
  </div>
</MainLayout>

<style>
</style>