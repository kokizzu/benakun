<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiDesignBallPenLine from 'svelte-icons-pack/ri/RiDesignBallPenLine';
  import AiOutlineEyeInvisible from 'svelte-icons-pack/ai/AiOutlineEyeInvisible';
  import AiOutlineFileExcel from 'svelte-icons-pack/ai/AiOutlineFileExcel';
  import IoSearch from 'svelte-icons-pack/io/IoSearch';
  import IoClose from 'svelte-icons-pack/io/IoClose';
  import FiLoader from 'svelte-icons-pack/fi/FiLoader';
  import CgChevronDown from 'svelte-icons-pack/cg/CgChevronDown';
	import CgChevronLeft from 'svelte-icons-pack/cg/CgChevronLeft';
  import CgChevronRight from 'svelte-icons-pack/cg/CgChevronRight';
  import CgChevronDoubleRight from 'svelte-icons-pack/cg/CgChevronDoubleRight';
  import CgChevronDoubleLeft from 'svelte-icons-pack/cg/CgChevronDoubleLeft';
	import IoArrowUpSharp from 'svelte-icons-pack/io/IoArrowUpSharp';
	import IoArrowDownSharp from 'svelte-icons-pack/io/IoArrowDownSharp';
	import RiDeviceDatabase2Line from 'svelte-icons-pack/ri/RiDeviceDatabase2Line';
	import FaChartBar from 'svelte-icons-pack/fa/FaChartBar';
	import RiSystemInformationLine from 'svelte-icons-pack/ri/RiSystemInformationLine';
	import FaSolidChartLine from 'svelte-icons-pack/fa/FaSolidChartLine';
	import RiSystemDeleteBin5Line from 'svelte-icons-pack/ri/RiSystemDeleteBin5Line';
	import RiSystemFilterLine from 'svelte-icons-pack/ri/RiSystemFilterLine';
	import RiSystemArrowGoBackLine from 'svelte-icons-pack/ri/RiSystemArrowGoBackLine';
	import InputCustom from './InputCustom.svelte';
	import { onMount } from 'svelte';
	import FilterTable from './FilterTable.svelte';
	import { datetime } from './formatter';

	/** @typedef {import('./types/master.js').Field} Field */
	/** @typedef {import('./types/access.js').Access} Access */
	/** @typedef {import('./types/master.js').PagerOut} PagerOut */ // @ts-ignore
	/** @typedef {import('./types/master.js').PagerIn} PagerIn */

	export let FIELDS					= /** @type Field[] */  ([]);	// bind
	export let PAGER					= /** @type PagerOut */ ({}); // bind
	export let MASTER_ROWS		= /** @type any[][] */	([]); // bind
	
	export let ACCESS						= /** @type Access */ ({});
	export let ARRAY_OF_ARRAY		= true;
	export let CAN_SEARCH_ROW 	= true;
	export let CAN_EDIT_ROW 		= true;
	export let CAN_DELETE_ROW 	= false;
	export let CAN_RESTORE_ROW	= false;

	// State for loading if hit ajax
	let isAjaxSubmitted = false;

	// Binding component FilterTable.svelte 
  let filterTable = null;
	// For readiness of component FilterTable.svelte, prevent race condition
	let filterTableReady = false;
	// Key and label of column to filter
	let filterColumns = [];
	// Binding value of column, for payloaf
	let filtersMap = {};

	// Index of field 'deletedAt', for marker deleted rows
	let deletedIndex = 0;

	// PopUp for modify rows
	let showPopUp = false;

	// Pagination total, based on total pages
	let paginationTotal = 1;
	// Pagination all, based on total pages
	let paginationsAll = /** @type number[] */ ([]);
	// Pagination to show, based on total pagination
	let paginationShow = /** @type number[] */ ([]);
	// Current page describe which page is currently rendered
	let currentPage = 1;
	// State for sort, wheter is ascending or descending
	let sortTableAsc = false;
	// Total Pages
	let totalPages = 0;
	// Total rows but rounded by current rows
	let totalRound = 0;
	// Rows per page
	let currentRows = PAGER.perPage;
	// Rows per page options
	let rowsToShow = [10, 1, 40, 60, 70, 100, 200];
	// State for show rows options
	let showRowsNum = false;
	// Total rows
	let totalRows = PAGER.countResult;
	// Total rows current
	let totalRowsCurrent = 0;

	// Payloads for modify rows
	let payloads = [];

	// Toggle show rows options
	function toggleRowsNum() { showRowsNum = !showRowsNum }

	// Refresh Pagination
	function getPaginationShow() {
		totalRows = PAGER.countResult;
		totalPages = PAGER.pages;
		currentPage = PAGER.page;
		if (MASTER_ROWS && MASTER_ROWS.length) totalRowsCurrent = MASTER_ROWS.length;
		else totalRowsCurrent = 0;

    totalRound = Math.ceil(totalPages / currentRows) * currentRows;
		paginationTotal = totalRound / currentRows, paginationsAll = [];
		if (currentRows > PAGER.countResult) paginationTotal = 1;
		for (let i = 0; i < paginationTotal; i++) {
			paginationsAll = [...paginationsAll, i+1]
		}
		let start = 0, end = 0;
		if (paginationTotal < 5) {
			start = 0, end = paginationTotal
		} else if ((currentPage < 5) && ((currentPage - 3) < 0) ) {
			start = 0, end = 5;
		} else if ((currentPage > (paginationTotal-5)) && ((currentPage + 3) < paginationTotal)) {
			start = (currentPage - 3), end = (currentPage + 2);
		} else if ((currentPage + 3) >= paginationTotal) {
			start = (paginationTotal - 5), end = paginationTotal;
		} else {
			start = (currentPage - 3), end = (currentPage + 2);
		}
		paginationShow = paginationsAll.slice(start, end);
	}

	/**
	 * @param {any} row
	 * @param {number} i
	 * @param {Field} field
	 */
	function Cell( row, i, field ) {
		if( ARRAY_OF_ARRAY ) return row[ i ] || '';
		return row[ field.name ] || '';
	}

	onMount(() => {
		// FilterTable.svelte component is rendered
		filterTable = FilterTable;
		// FilterTable.svelte component is ready
		filterTableReady = true;
		// Total pages
		totalPages = PAGER.pages;
		// Loop column/fields, fill variable filterColumn for filters
		if (FIELDS && FIELDS.length > 0) {
			filterColumns = [];
			FIELDS.forEach((col, idx) => {
				if (col.name === 'deletedAt') deletedIndex = idx;
				else if (col.name === 'invitationState') deletedIndex = idx;
				filterColumns = [...filterColumns, {
					key: col.name,
					label: col.label
				}]
			});
		}
		// Calculate pagination
		getPaginationShow();
		// Fill initial payloads
		if (FIELDS && FIELDS.length > 0) {
			FIELDS.forEach(() => payloads = [...payloads, '']);
		}
	})

	// Export function, forward parameter to parent
	export let OnRestore = async function(/** @type any[]*/ row) {}
	export let OnDelete = async function(/** @type any[]*/ row) {}
	export let OnEdit = function(/** @type any */ id, /** @type any[]*/ payloads) {}
	export let OnRefresh = async function(/** @type PagerIn */ pagerIn) {}

	function ApplyFilter() {
		// Hide FilterTable.svelte
		filterTable.Hide();
		// Make a 'filters' payload from variable filtersMap
		// Make it with format { key: [value, value] }
		let filters = {};
		for (let key in filtersMap) {
			let value = filtersMap[ key ];
			if (value) filters[ key ] = value.split('|');
		}
		OnRefresh({ ...PAGER, filters });
		// Refresh pagination view
		getPaginationShow();
	}

	// Apply row counts
	async function toRow(/** @type number */ perPage) {
		currentRows = perPage;
		showRowsNum = false;
		await OnRefresh({ ...PAGER, perPage });
		// Refresh pagination view
		getPaginationShow();
	}

	// Go to page, last page, first page
	async function goToPage(/** @type number */ page) {
		currentPage = page;
		await OnRefresh({ ...PAGER, page });
		// Refresh pagination view
		getPaginationShow();
	}

	// Restore row
	async function restoreRow(/** @type any[] */ row) {
		await OnRestore(row);
		// Refresh pagination view
		getPaginationShow();
	}

	// Delete row
	async function deleteRow(/** @type any[] */ row) {
		await OnDelete(row);
		// Refresh pagination view
		getPaginationShow();
	}

	// Row ID to modify
	let idToMod = '';

	function toggleShowPopUp(/** @type any */ id, /** @type any[]*/ row) {
		payloads = [];
		FIELDS.forEach((_, i) => {
			payloads = [...payloads, row[i]]
		});

		showPopUp = true;
		idToMod = id;
	}

	function handleSubmitEdit() {
		showPopUp = false;
		OnEdit(idToMod, payloads);
	}
</script>

{#if filterTableReady}
  <FilterTable
		bind:this={filterTable}
		bind:filterColumns
		bind:filtersMap
		on:click={ApplyFilter}
	/>
{/if}

{#if showPopUp}
	<div class="popup_container">
		<div class="popup">
			<header>
				<h2>Edit row</h2>
				<button on:click={() => showPopUp = false}>
					<Icon
						size="22"
						color="var(--red-005)"
						src={IoClose}
					/>
				</button>
			</header>
			<div class="forms">
					{#each (FIELDS || []) as field, idx}
						{#if field.name !== 'id'}
							{#if !field.readOnly}
								<InputCustom
									id={field.name}
									label={field.label}
									placeholder={field.description}
									bind:value={payloads[idx]}
									type={field.type || 'text'}
								/>
							{/if}
						{/if}
					{/each}
			</div>
			<div class="foot">
				<button class="ok" on:click={handleSubmitEdit}>Save</button>
			</div>
		</div>
	</div>
{/if}

<div class="table_root">
	<div class="actions_container">
    <div class="left">
      <div class="actions_btn">
				<button
					class="action_btn"
					on:click={() => filterTable.Show()}
					title="filter table"
				>
					<Icon
						color="var(--gray-007)"
						size="16"
						src={RiSystemFilterLine}
					/>
				</button>
				<!-- Action buttons -->
				<slot />
      </div>
			{#if isAjaxSubmitted}
        <div class="loader">
          <Icon
						className="spin"
						color="var(--violet-007)"
						size="28"
						src={FiLoader}
					/>
        </div>
      {/if}
    </div>
    <div class="right">
			{#if CAN_SEARCH_ROW}
      	<div class="search_handler">
        	<button class="search_btn" title="Search">
          	<Icon
							color="var(--gray-007)"
							size="16"
							src={IoSearch}
						/>
        	</button>
        	<input
          	placeholder="Search..."
          	type="text"
          	name="searchRow"
          	id="searchRow"
          	class="search"
        	/>
      	</div>
			{/if}
    </div>
  </div>
	<div class="table_container">
		<table>
			<thead>
				<tr>
					<th class="no">No</th>
					{#each (FIELDS || []) as f, _ (f.name)}
						{#if f.name === 'id'}
							<th class="a_row">Actions</th>
						{:else}
							<th>{f.label}</th>
						{/if}
					{/each}
				</tr>
			</thead>
			<tbody>
				{#if MASTER_ROWS && MASTER_ROWS.length > 0}
					{#each MASTER_ROWS as row}
						<tr class={
							(CAN_DELETE_ROW && row[deletedIndex] > 0)
							|| (CAN_DELETE_ROW && row[deletedIndex] === 'terminated')
								? 'deleted'
								: ''
						}>
							<td class="num_row">{MASTER_ROWS.indexOf(row) + 1}</td>
							{#each (FIELDS || []) as f, idx}
								{#if f.name === 'id'}
									<td class="a_row">
										{#if ACCESS.superAdmin
											|| ACCESS.tenantAdmin
											|| ACCESS.entryUser
											|| ACCESS.reportViewer
										}
											<div class="actions">	
												{#if CAN_EDIT_ROW}
													<button
														class="btn edit"
														title="Edit"
														on:click={() => toggleShowPopUp(Cell(row, idx, f), row)}
													>
														<Icon
															size="15"
															color="var(--gray-007)"
															src={RiDesignBallPenLine}
														/>
													</button>
												{/if}
												{#if CAN_DELETE_ROW || CAN_RESTORE_ROW}
													{#if (row[deletedIndex] > 0) || (row[deletedIndex] === 'terminated')}
														<button
															class="btn info"
															title="restore"
															on:click={() => restoreRow(row)}
														>
															<Icon
																size="15"
																color="var(--gray-007)"
																src={RiSystemArrowGoBackLine}
															/>
														</button>
													{:else}
														<button
															class="btn delete"
															title="delete"
															on:click={() => deleteRow(row)}
														>
															<Icon
																size="15"
																color="var(--gray-007)"
																src={RiSystemDeleteBin5Line}
															/>
														</button>
													{/if}
												{/if}
											</div>
										{:else}
											<span>--</span>
										{/if}
									</td>
								{:else if f.inputType === 'datetime'}
									<td>{(row[idx]) ? datetime(row[idx]) : '--'}</td>
								{:else}
									<td>{row[idx] || '--'}</td>
								{/if}
							{/each}
						</tr>
					{/each}
				{:else}
					{#each (Array(5).fill()) as _, i}
						<tr>
							<td class="num_row">{i + 1}</td>
							{#each (FIELDS || []) as _}
								<td>-</td>
							{/each}
						</tr>
						{/each}
				{/if}
			</tbody>
		</table>
	</div>
  <div class="pagination_container">
    <div class="filter">
      <div class="showing">
        <p>Showing <span class="text-violet">{totalRowsCurrent}</span> / </p>
      </div>
      <div class="row_to_show">
				{#if showRowsNum}
          <div class="rows">
            {#each rowsToShow as r}
              <button on:click={() => toRow(r)}>{r}</button>
            {/each}
          </div>
        {/if}
        <button class="btn" on:click={toggleRowsNum}>
          <span>{currentRows}</span>
          <Icon
						className={showRowsNum ? 'rotate_right' : 'dropdown'}
						size="13"
						src={CgChevronRight}
					/>
        </button>
      </div>
			<p>record(s)</p>
    </div>
		<div>
			<p>Total:<span class="text-violet">{totalRows}</span></p>
		</div>
    <div class="pagination">
      <button
				disabled={currentPage == 1}
				class="btn to"
				title="Go to first page"
				on:click={() => goToPage(1)}
			>
        <Icon size="16" src={CgChevronDoubleLeft}/>
      </button>
      <button
				disabled={currentPage == 1}
				class="btn to"
				title="Go to previous page"
				on:click={() => goToPage(currentPage - 1)}
			>
        <Icon size="16" src={CgChevronLeft}/>
      </button>
      {#each paginationShow as i}
        <button
					disabled={currentPage == i}
					class={currentPage === i ? 'btn active' : 'btn'}
					title={`Go to page ${i}`}
					on:click={() => goToPage(i)}
				>{i}</button>
      {/each}
      <button
				disabled={currentPage == paginationTotal}
				class="btn to" title="Go to next page"
				on:click={() => goToPage(currentPage + 1)}
			>
        <Icon size="16" src={CgChevronRight}/>
      </button>
      <button
				disabled={currentPage == paginationTotal}
				class="btn to"
				title="Go to last page"
				on:click={() => goToPage(paginationTotal)}
			>
        <Icon size="16" src={CgChevronDoubleRight}/>
      </button>
    </div>
  </div>
</div>

<style>
	.popup_container {
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 2000;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		display: flex;
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

	.popup_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 500px;
		display: flex;
		flex-direction: column;
	}

	.popup_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 15px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup_container .popup header h2 {
		margin: 0;
	}

	.popup_container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup_container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup_container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.popup_container .popup .foot {
		display: flex;
		flex-direction: row;
    justify-content: flex-end;
		gap: 10px;
		align-items: center;
		padding: 10px 20px;
		border-top: 1px solid var(--gray-004);
	}

	.popup_container .popup .foot button {
		padding: 8px 13px;
		border-radius: 9999px;
		border: none;
		color: #FFF;
		cursor: pointer;
		font-weight: 600;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.popup_container .popup .foot button.ok {
		background-color: var(--green-006);
		border: 1px solid var(--green-006);
    display: flex;
    justify-content: center;
    align-items: center;
	}

	.popup_container .popup .foot button.ok:hover {
		background-color: var(--green-005);
	}

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
		padding: 0 0 20px 0;
		overflow: hidden;
  }

	.table_root .text-violet {
		color: var(--violet-005);
		font-weight: 600;
		padding: 5px;
	}

	.table_root p {
		margin: 0;
	}

  .table_root .actions_container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
		padding: 10px 15px;
		background-color: #FFF;
  }

  .table_root .actions_container .left,
  .table_root .actions_container .right {
		display: flex;
		flex-direction: row;
		align-items: center;
    gap: 10px;
	}

	.table_root .actions_container .left .debug .btn {
		border: none;
		background-color: var(--violet-006);
		color: #FFF;
		width: fit-content;
		padding: 4px 10px;
		border-radius: 9999px;
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
		gap: 3px;
		cursor: pointer;
	}

	.table_root .actions_container .left .debug .btn:hover {
		background-color: var(--violet-005);
	}

  .table_root .actions_container .right .search_handler {
    display: flex;
    flex-direction: row;
    width: fit-content;
    height: fit-content;
		position: relative;
  }

  .table_root .actions_container .right .search_handler input.search {
    padding: 12px 40px 12px 15px;
    border-radius: 8px;
    border: none;
		background-color: var(--gray-001);
    width: 370px;
  }

  .table_root .actions_container .right .search_handler input.search:focus {
    border-color: none;
		outline: 1px solid var(--gray-003);
		box-shadow: var(--shadow-md);
  }

  .table_root .actions_container .right .search_handler .search_btn {
		position: absolute;
    background-color: transparent;
    padding: 8px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
		border-radius: 8px;
    cursor: pointer;
		right: 5px;
		top: 3px;
  }

  .table_root .actions_container .right .search_handler .search_btn:hover {
    background-color: var(--violet-transparent);
  }

	:global(.table_root .actions_container .right .search_handler .search_btn:hover svg) {
		fill: var(--violet-005);
	}

  .table_root .actions_container .actions_btn {
    display: flex;
    flex-direction: row;
    align-items: center;
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
  }

	.table_root .table_container table tbody tr.deleted {
		color: var(--red-005);
	}

  .table_root .table_container table thead tr th.no {
    width: 30px;
  }


  .table_root .table_container table thead tr th.a_row {
		max-width: fit-content;
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

	.table_root .table_container table tbody tr td .actions {
		display: flex;
		flex-direction: row;
	}

	.table_root .table_container table tbody tr td .actions .btn {
		border: none;
		padding: 6px;
		border-radius: 8px;
		background-color: transparent;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.table_root .table_container table tbody tr td .actions .btn:hover {
		background-color: var(--violet-transparent);
	}

	:global(.table_root .table_container table tbody tr td .actions .btn:hover svg) {
		fill: var(--violet-005);
	}

  .table_root .pagination_container {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 15px 15px 0 15px;
	}

	.table_root .pagination_container .filter {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 8px;
	}

	.table_root .pagination_container .filter .row_to_show {
		position: relative;
		width: fit-content;
		height: fit-content;
	}

	.table_root .pagination_container .filter .row_to_show .btn {
		border: none;
		background-color: var(--violet-transparent);
		color: var(--violet-005);
		width: fit-content;
		padding: 3px 3px 3px 6px;
		font-weight: 600;
		border: 1px solid var(--violet-004);
		border-radius: 9999px;
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
		gap: 1px;
		cursor: pointer;
	}

	.table_root .pagination_container .filter .row_to_show .btn:hover {
		background-color: var(--violet-002);
	}

	.table_root .pagination_container .filter .row_to_show .rows {
		display: flex;
		flex-direction: column-reverse;
		position: absolute;
		width: 100%;
		top: -180px;
		border-radius: 5px;
		border: 1px solid var(--gray-004);
		background-color: #FFF;
	}

	.table_root .pagination_container .filter .row_to_show .rows button {
		border: none;
		background-color: transparent;
		padding: 5px;
		cursor: pointer;
		color: var(--gray-007);
	}

	.table_root .pagination_container .filter .row_to_show .rows button:hover {
		background-color: var(--violet-transparent);
		color: var(--violet-007);
	}

	.table_root .pagination_container .pagination {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 5px;
		overflow: hidden;
	}

	.table_root .pagination_container .pagination .btn {
		border: none;
		background-color: transparent;
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		padding: 6px 10px;
		border-radius: 9999px;
		cursor: pointer;
		gap: 5px;
		color: var(--gray-007);
		border: 1px solid transparent;
	}

	.table_root .pagination_container .pagination .btn:hover {
		border: 1px solid var(--gray-004);
	}

	.table_root .pagination_container .pagination .btn.active {
		background-color: var(--violet-transparent);
		color: var(--violet-006);
		font-weight: 600;
		border: 1px solid var(--violet-004);
	}

	.table_root .pagination_container .pagination .btn.to {
		background-color: var(--violet-006);
		color: #FFF;
		font-weight: 600;
		border: none;
	}

	.table_root .pagination_container .pagination .btn.to:hover {
		background-color: var(--violet-005);
	}

	.table_root .pagination_container .pagination .btn.to:disabled {
		background-color: var(--gray-002);
		color: var(--gray-006);
		font-weight: 600;
		border: 1px solid var(--gray-004);
		cursor: not-allowed;
	}

	@media only screen and (max-width : 768px) {
		.table_root .actions_container {
			flex-wrap: wrap;
			gap: 10px;
		}

		.table_root .actions_container .left {
			flex-wrap: wrap;
		}

		.table_root .table_container {
			overflow-x: scroll;
		}

		.table_root .pagination_container {
			flex-wrap: wrap;
			gap: 10px;
		}

		.table_root .pagination_container .pagination {
			gap: 2px;
		}
	}
</style>