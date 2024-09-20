<script>
  import MainLayout from './_layouts/mainLayout.svelte';

  /** @typedef {Object} financialPosition
   *  @property {string} coaName
   *  @property {number} amountIDR
  */
  const financialPosition = /** @type Record<string, financialPosition[]> */ ({/* financialPosition */});

  console.log('FINANCIAL POSITION', financialPosition);
</script>

<MainLayout>
  <div class="table_financial_position">
    <div class="header">
      <span>Benalu DEV</span>
      <h2>Laporan Posisi Keuangan</h2>
      <!-- <span>Periode Septermber 2024</span> -->
    </div>
    <div class="body">
      <div class="part">
        <div class="section">
          <div class="info">
            <b>Aktiva</b>
            <ul>
              {#each (financialPosition[`asset`] || []) as asset}
                <li>{asset.coaName}: <b>Rp. {asset.amountIDR}</b></li>
              {/each}
            </ul>
          </div>
          <div class="foot">
            <span>Total Aktiva</span>
            <b>Rp. {(financialPosition[`asset`] || []).map((asset) => asset.amountIDR).reduce((a, b) => a + b, 0) || 0}</b>
          </div>
        </div>
      </div>
      <div class="part">
        <div class="section">
          <div class="info">
            <b>Modal</b>
            <ul>
              {#each (financialPosition[`equity`] || []) as equity}
                <li>{equity.coaName}: <b>Rp. {equity.amountIDR}</b></li>
              {/each}
            </ul>
          </div>
          <div class="foot">
            <span>Total Modal</span>
            <b>Rp. {(financialPosition[`equity`] || []).map((equity) => equity.amountIDR).reduce((a, b) => a + b, 0)}</b>
          </div>
        </div>
        <div class="section">
          <div class="info">
            <b>Hutang</b>
            <ul>
              {#each (financialPosition[`liability`] || []) as liability}
                <li>{liability.coaName}: <b>Rp. {liability.amountIDR}</b></li>
              {/each}
            </ul>
          </div>
          <div class="foot">
            <span>Total Hutang</span>
            <b>Rp. {(financialPosition[`liability`] || []).map((liability) => liability.amountIDR).reduce((a, b) => a + b, 0)}</b>
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
    padding: 20px;
  }

  .table_financial_position .body .part:nth-child(1) {
    border-right: 2px solid var(--gray-007);
  }

  .table_financial_position .body .part .section {
    display: flex;
    flex-direction: column;
    gap: 10px;
    width: 100%;
  }

  .table_financial_position .body .part .section .info {
    display: flex;
    flex-direction: column;
    gap: 5px;
    width: 100%;
  }

  .table_financial_position .body .part .section .info ul {
    padding: 0px 0 0 20px;
    margin: 0;
  }

  .table_financial_position .body .part .section .foot {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>
