<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import InputBox from './_components/InputBox.svelte';

  let coa = '';
</script>

<MainLayout>
  <p>Transction entry</p>
  <p>
    Berdasarkan transaction tempalte detail, baris2 di dalam notanya, sebagai pilihan untuk di entry
  </p>
  <p>
    Tapi, terbatas hanya ke jenis coa yg ada di transaction template detail
  </p>
  <div>
    <InputBox
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