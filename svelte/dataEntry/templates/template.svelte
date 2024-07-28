<script>
  /** @typedef {import('../../_components/types/transaction.js').TransactionTemplate} TransactionTemplate */

  import MainLayout from '../../_layouts/mainLayout.svelte';
  import InputCustom from '../../_components/InputCustom.svelte';

  let transactionTemplate = /** @type TransactionTemplate */ ({/* transactiontemplate */});

  let coa = '';
  let debitIDR = 0;

  // jika attribut coa adalah childOnly: maka coa pilihan hanya anaknya

  // Nota bagian atas dan bawah adalah business transaction
  // Nota bagian tengah adalah jurnal
</script>

<MainLayout>
  <div>
    <p>
      Yang ditampilin adalah transaction template detail dari
      transaction template [ini]
    </p>
    <p>
      Jika autoSum, maka autoSum disisi apa dulu ?
      jika kredit, dia autoSum dari debitnya, dan sebaliknya
    </p>
    <p>
      jika ChildOnly, maka yang bisa di entry hanya anak - anak CoA nya
    </p>
    <p>
      Jika sales, maka inputannya = (angka x harga)
    </p>

    <h3>INI MASUK KE TABLE JOURNAL</h3>
  </div>
  <div class="data_entry_template___container">
    <div class="form_data_entry_template">
      <InputCustom
        id="coa"
        label="CoA (Chart of Account)"
        type="text"
        placeholder="CoA"
        bind:value={coa}
      />
      <InputCustom
        id="debitIdr"
        label="Debit IDR"
        type="number"
        placeholder="Debit IDR"
        bind:value={debitIDR}
      />
    </div>
  </div>
</MainLayout>

<style>
  .data_entry_template___container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
  }

  .form_data_entry_template {
    background-color: white;
    padding: 20px;
    border-radius: 10px;
    border: 1px solid var(--gray-300);
    display: flex;
    flex-direction: column;
    gap: 10px;
    
  }
</style>