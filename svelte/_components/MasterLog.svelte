<script>

</script>

<div class="table_root">
  <div class="actions_container"></div>
  <div class="table_container">
    <table>
      <thead>
        <tr>
          <th class="no">No</th>
          <th class="datetime">Time</th>
          <th>Action</th>
          <th>Actor ID</th>
          <th>Request ID</th>
          <th>IPv4</th>
          <th>IPv6</th>
          <th>Ref ID</th>
          <th>Latency</th>
          <th>Tenant Code</th>
          <th>User Agent</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="num_row">1</td>
          <td>Jul 11, 12:32:12</td>
          <td>/superAdmin/accessLog</td>
          <td>1</td>
          <td>0aYx2B~-----0~0</td>
          <td>0.0.0.0</td>
          <td>0.0.0.0</td>
          <td>0</td>
          <td>0.01</td>
          <td>admin-2642</td>
          <td>Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36</td>
        </tr>
      </tbody>
    </table>
  </div>
</div>

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

	:global(.action_btn:hover svg) {
		fill: var(--violet-005);
	}

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  :global(.dropdown) {
		transition: all .2s ease-in-out;
	}

	:global(.rotate) {
		transition: all .2s ease-in-out;
		transform: rotate(180deg);
	}

	:global(.rotate_right) {
		transition: all .2s ease-in-out;
		transform: rotate(-90deg);
	}

	:global(.sort_icon) {
		margin-bottom: -5px;
	}

  .table_root {
    display: flex;
    flex-direction: column;
		background-color: #FFF;
		border-radius: 10px;
		border: 1px solid var(--gray-003);
		padding: 0;
		overflow: hidden;
  }

  .table_root .actions_container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
		padding: 10px 15px;
		background-color: #FFF;
  }

	.table_root .table_container {
		overflow-x: auto;
		scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
	}

  .table_root .table_container table {
		width: 100%;
    background: #fff;
    border-top: 1px solid var(--gray-003);
    border-bottom: 1px solid var(--gray-003);
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

  .table_root .table_container table thead tr th{
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table_root .table_container table thead tr th.datetime {
		min-width: 140px !important;
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

	.table_root .table_container table tbody tr td.num_row {
		border-right: 1px solid var(--gray-003);
		font-weight: 600;
		text-align: center;
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