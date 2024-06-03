<script>
  /**
   * @typedef {Object} field
   * @property {string} key
   * @property {string} label
   */

  export let FIELDS = /** @type field[] */ ([]);
  export let ROWS   = /** @type any[] */ ([]);
</script>

<div class="table_root">
  <div class="table_container">
    <table>
      <thead>
        <tr>
          <th class="no">No</th>
          {#each FIELDS as f}
            <th>{f.label}</th>
          {/each}
        </tr>
      </thead>
      <tbody>
        {#if ROWS && ROWS.length}
          {#each ROWS as row, idx (row.id)}
            <tr>
              <th>{idx + 1}</th>
              {#each FIELDS as f}
                <td>{row[f.key] || '--'}</td>
              {/each}
            </tr>
          {/each}
        {/if}
      </tbody>
    </table>
  </div>
</div>

<style>
  .table_root {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

	.table_root .table_container {
		overflow-x: auto;
		scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
	}

  .table_root .table_container table {
		width: 100%;
    background: #fff;
    border: 1px solid var(--gray-004);
    box-shadow: none;
    border-radius: 8px;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
		overflow: hidden;
  }

	.table_root .table_container table thead {
		box-shadow: none;
	}

  .table_root .table_container table thead tr th{
    padding: 12px;
		background-color: var(--gray-001);
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
		-webkit-user-select: none;
  	-ms-user-select: none;
  	user-select: none;
		text-transform: capitalize;
  }

  .table_root .table_container table thead tr th.no {
    width: 30px;
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

  .table_root .table_container table tbody tr:last-child td:last-child,
  .table_root .table_container table thead tr th:last-child  {
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