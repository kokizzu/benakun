<script>
  /** @typedef {import('./_components/types/user').User} User */
  /** @typedef {import('./_components/types/access').Access} Access */
  /** @typedef {import('./_components/types/invoicePayment').InvoicePayment} InvoicePayment */

  import MainLayout from './_layouts/mainLayout.svelte';
  import ProfileSubMenu from './_components/partials/ProfileSubMenu.svelte';
  import { formatPrice, datetime } from './_components/xformatter';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let invoices  = /** @type {InvoicePayment[]} */ ([/* invoices */]);

  console.log('User:', user);
  console.log('Segments:', segments);
  console.log('Invoices:', invoices);
</script>

<MainLayout>
  <ProfileSubMenu />
  <div class="table_root">
    <div class="table_container">
      <table>
        <thead>
          <tr>
            <th class="no">No</th>
            <th>Invoice Number</th>
            <th>Amount</th>
            <th>Currency</th>
            <th>Payment Method</th>
            <th>Status</th>
            <th>Support Start At</th>
            <th>Support End At</th>
          </tr>
        </thead>
        <tbody>
          {#if invoices && invoices.length > 0}
            {#each (invoices || []) as inv, idx (inv.id)}
              <tr>
                <td class="no">{idx + 1}</td>
                <td>{inv.invoiceNumber}</td>
                <td>{formatPrice((inv.amount || 0), inv.currency)}</td>
                <td>{inv.currency}</td>
                <td>{inv.paymentMethod || '--'}</td>
                <td>
                  <span class={inv.status}>
                    {inv.status}
                  </span>
                </td>
                <td>{inv.supportStartAt ? datetime(inv.supportStartAt) : '--'}</td>
                <td>{inv.supportStartAt ? datetime(inv.supportEndAt) : '--'}</td>
              </tr>
            {/each}
          {:else}
            <tr>
              <td class="no">0</td>
              <td colspan="7">No data</td>
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  </div>
</MainLayout>

<style>
  :global(.root-layout .root-container .root-content .content) {
    padding: 10px 20px 20px !important;
  }
  
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

  .table_root .table_container table tbody tr td.no {
    text-align: center;
    font-weight: bold;
  }

  .table_root .table_container table tbody tr td span.pending {
    background-color: var(--yellow-transparent);
    color: var(--yellow-006);
    padding: 4px 8px;
    border-radius: 999px;
    text-transform: capitalize;
    font-size: var(--font-sm);
  }

  .table_root .table_container table tbody tr td span.success {
    background-color: var(--green-transparent);
    color: var(--green-006);
    padding: 4px 8px;
    border-radius: 999px;
    text-transform: capitalize;
    font-size: var(--font-sm);
  }

  .table_root .table_container table tbody tr td span.failed {
    background-color: var(--red-transparent);
    color: var(--red-006);
    padding: 4px 8px;
    border-radius: 999px;
    text-transform: capitalize;
    font-size: var(--font-sm);
  }

  .table_root .table_container table tbody tr td span.cancelled,
  .table_root .table_container table tbody tr td span.expired,
  .table_root .table_container table tbody tr td span.refunded {
    background-color: var(--gray-002);
    color: var(--gray-008);
    padding: 4px 8px;
    border-radius: 999px;
    text-transform: capitalize;
    font-size: var(--font-sm);
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
