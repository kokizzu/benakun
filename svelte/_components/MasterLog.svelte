<script>
  /** @typedef {import('./types/admin.js').AccessLog} AccessLog */
  import { utcDatetime } from './formatter';

  export let logs = /** @type {AccessLog[]} */ ([]);
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
          <th class="request_id">Request ID</th>
          <th>IPv4</th>
          <th>IPv6</th>
          <th>Ref ID</th>
          <th>Latency</th>
          <th>Tenant Code</th>
          <th class="user_agent">User Agent</th>
        </tr>
      </thead>
      <tbody>
        {#each (logs || []) as log, i (log.refId)}
          <tr>
            <td class="num_row">{i+1}</td>
            <td>{utcDatetime(log.createdAt)}</td>
            <td>{log.action}</td>
            <td>{log.actorId}</td>
            <td>{log.requestId}</td>
            <td>{log.ipAddr4}</td>
            <td>{log.ipAddr6}</td>
            <td>{log.refId}</td>
            <td>{log.latency}</td>
            <td>{log.tenantCode}</td>
            <td>{log.userAgent}</td>
          </tr>
        {/each}
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
    border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table_root .table_container table thead tr th.datetime {
		min-width: 140px !important;
	}

  .table_root .table_container table thead tr th.user_agent {
    min-width: 240px !important;
  }

  .table_root .table_container table thead tr th.request_id {
    min-width: 150px !important;
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