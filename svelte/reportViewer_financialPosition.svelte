<script>
  import MainLayout from './_layouts/mainLayout.svelte';
</script>

<MainLayout>
  <div class="table_financial_position">
    <div class="header">
      <span>Benalu DEV</span>
      <h2>Laporan Posisi Keuangan</h2>
      <span>Periode Septermber 2024</span>
    </div>
    <div class="body">
      <div class="part">
        <div class="section">
          <div class="info">
            <b>Aktiva</b>
            <ul>
              <li>iojisdfj</li>
              <li>sdfgsdfg</li>
              <li>48hj8sdjfhgishjdio</li>
            </ul>
          </div>
          <div class="foot">
            Total Aktiva
          </div>
        </div>
      </div>
      <div class="part">
        <div class="section">
          <div class="info">
            <b>Modal</b>
            <ul>
              <li>iojisdfj</li>
              <li>sdfgsdfg</li>
              <li>48hj8sdjfhgishjdio</li>
            </ul>
          </div>
          <div class="foot">
            Total Modal
          </div>
        </div>
        <div class="section">
          <div class="info">
            <b>Hutang</b>
            <ul>
              <li>iojisdfj</li>
              <li>sdfgsdfg</li>
              <li>48hj8sdjfhgishjdio</li>
            </ul>
          </div>
          <div class="foot">
            Total Hutang
          </div>
        </div>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .table_financial_position {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .table_financial_position .header {
    display: flex;
    flex-direction: column;
    gap: 8px;
    justify-content: center;
    align-items: center;
  }

  .table_financial_position .header h2 {
    margin: 0;
  }

  .table_financial_position .body {
    display: grid;
    grid-template-columns: 1fr 1fr;
    border: 2px solid var(--gray-007);
    border-radius: 8px;
  }

  .table_financial_position .body .part {
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 10px;
  }

  .table_financial_position .body .part:nth-child(1) {
    border-right: 2px solid var(--gray-007);
  }

  .table_financial_position .body .part .section {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .table_financial_position .body .part .section .info {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .table_financial_position .body .part .section .info ul {
    padding: 0px 0 0 20px;
    margin: 0;
  }
</style>
