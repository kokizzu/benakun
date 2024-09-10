<script>
  import MainLayout from './_layouts/mainLayout.svelte';
</script>

<MainLayout>
  <div>
    <div class="table_root">
      <div class="table_container">
        <table>
          <thead>
            <tr>
              <th class="no">No. Acc / No. Rek</th>
              <th>Name / Nama</th>
              <th>Loss / Rugi</th>
              <th>Income / Laba</th>
            </tr>
          </thead>
          <tbody>
              <tr>
                <td class="no">1</td>
                <td>Pendapatan lainnya</td>
                <td>Rp20.000</td>
                <td>Rp5.000</td>
              </tr>
            <tr>
              <td colspan="2">
                <b>Total</b>
              </td>
              <td>
                <b>Rp25.000</b>
              </td>
              <td>
                <b>Rp10.000</b>
              </td>
              <td></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .table_root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    padding: 0;
    border: 1px solid var(--gray-003);
    overflow: hidden;
  }

 .table_root .table_container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .table_root .table_container table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
  }

  .table_root .table_container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table_root .table_container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table_root .table_container table thead tr th.no {
    width: 30px;
  }

  .table_root .table_container table thead tr th:last-child {
    border-right: none;
  }

  .table_root .table_container table tbody tr td {
    padding: 8px 12px;
  }

	.table_root .table_container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table_root .table_container table tbody tr:last-child td,
	.table_root .table_container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table_root .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table_root .table_container table tbody tr:last-child td,
  .table_root .table_container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table_root .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table_root .table_container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table_root .table_container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }
</style>
