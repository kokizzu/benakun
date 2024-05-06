<script>
  // @ts-nocheck

  // TODO: fix this to compatible with project

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiDesignBallPenLine from 'svelte-icons-pack/ri/RiDesignBallPenLine';
  import AiOutlineEyeInvisible from 'svelte-icons-pack/ai/AiOutlineEyeInvisible';
  import CgMathPlus from 'svelte-icons-pack/cg/CgMathPlus';
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

	import axios from 'axios';
  import FilterTable from './FilterTable.svelte';
  import InputCustom from './InputCustom.svelte';
  import { onMount } from 'svelte';
  import { notifier } from './xNotifier';
	import { utcDatetime } from './xFormatter';
  
  export const URL = window.location.pathname;
  
  export let PURPOSE = 'station';
  export let CAN_ADD_OR_EDIT = false;
  export let CAN_EXPORT_EXCEL = false;
	export let CAN_SEARCH_ROW = false;
  export let MASTER_ROWS = [];
  export let MASTER_GROUPS = {};
  export let MASTER_FIELDS = [];
  export let MASTER_COLS = [];
  export let ROWS_COUNT = 0;
  export let ROWS_OFFSET = 0;
  export let ROWS_ORDER = '';

  /**
    * @type {Array<string>}
    */
	export let DEFAULT_COLS = [];

  let ShowAddEdit_PopUp = false, SubmitAddEdit = false, AddOrEditState = '';
  let filterTable = null, filterTableReady = false, isAjaxSubmitted = false;
	let sortTableAsc = false, colToSort = '';

  /**
    * @typedef {Object} column
    * @property {string} key
    * @property {string} label
    * @property {string} filterState
    * @property {boolean} isApllied 
    * @property {boolean} isVisible
    */
  /**
    * @type {Array<column>}
    */
  let filterColumns = [], payload = [];
  let roundTotal = 0, paginationTotal = 1, currentRows = 10, paginationsAll = [], paginationShow = [], currentPage = 1;

  function getPaginationShow() {
    roundTotal = Math.ceil(Number(ROWS_COUNT) / currentRows) * currentRows;
		paginationTotal = roundTotal / currentRows, paginationsAll = [];
		if (currentRows > Number(ROWS_COUNT)) paginationTotal = 1;
		for (let i = 0; i < paginationTotal; i++) {paginationsAll = [...paginationsAll, i+1]}
		let start = 0, end = 0;
		if (paginationTotal<5) start = 0, end = paginationTotal;
		else if ((currentPage<5) && (currentPage-3<0) ) start = 0, end = 5;
		else if ((currentPage>paginationTotal-5)&&(currentPage+3<paginationTotal)) start=currentPage-3, end=currentPage+2;
		else if (currentPage+3 >= paginationTotal) start = paginationTotal-5, end = paginationTotal;
		else start = currentPage-3, end = currentPage+2;

		paginationShow = paginationsAll.slice(start, end);
	}

  onMount(() => {
    filterTable = FilterTable, filterTableReady = true;
		if (DEFAULT_COLS && DEFAULT_COLS.length) {
			MASTER_COLS.forEach((col, _) => {
				filterColumns.push({
					key: col.key,
					label: col.label,
					filterState: '',
					isApllied: false,
					isVisible: DEFAULT_COLS.includes(col.key) ? true : false
				})
			});
		} else {
			MASTER_COLS.forEach((col, idx) => {
				filterColumns.push({
					key: col.key,
					label: col.label,
					filterState: '',
					isApllied: false,
					isVisible: idx > 5 ? false : true
				})
			});
		}
		getPaginationShow();
    for (let i in MASTER_FIELDS) {payload = [...payload, '']}

		console.log(MASTER_ROWS)
  })

  async function filterTableOK() {
		filterTable.Hide();
		isAjaxSubmitted = true;
		let bodyFormData = new FormData();
		ROWS_OFFSET = currentRows;
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', `${currentRows}`);
		bodyFormData.append('offset', `${ROWS_OFFSET}`);
		let columnFilter = {};
		for (let i = 0; i < filterColumns.length; i++) {
			if (filterColumns[i].filterState != '' && filterColumns[i].isApllied) {
				columnFilter[filterColumns[i].key] = filterColumns[i].filterState
			}
		}
		bodyFormData.append('filter', JSON.stringify(columnFilter));
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: could not filter rows');
    });
	}

  let searchInput = '';
  async function searchRow() {
		isAjaxSubmitted = true;
		ROWS_OFFSET = 0;
		currentPage = 1;
		let bodyFormData = new FormData();
		bodyFormData.append('a', 'search');
    bodyFormData.append('limit', ''+currentRows);
		bodyFormData.append('offset', ''+ROWS_OFFSET);
		bodyFormData.append('term', searchInput);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			if (response.data.rows.length > 0) {
				MASTER_ROWS = response.data.rows;
				getPaginationShow();
			} else {
				notifier.showWarning(PURPOSE+' not found');
			}
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: could not search rows');
    });
  }

  const enterSearchInput = async (e) => {
		if (e.key == 'Enter') await searchRow();
	}

  // ===============+ Add/Edit/Delete/Restore +==============
	let isAdd = false, isEdit = false, idToEdit = '', isDeleted = false;

  function resetAddEditInput() {
		for (let idx in payload) payload[idx] = '';
  }

	function handleAdd() {
		for (let idx in payload) payload[idx] = '';
		AddOrEditState = 'Creating new '+PURPOSE;
		ShowAddEdit_PopUp = true;
		isAdd = true;
	}

	function handleEdit(id) {
		payload = [];
		let row = {};
		for (let i in MASTER_ROWS) {
			if (MASTER_ROWS[i].id == id) {
				row = MASTER_ROWS[i], isDeleted = MASTER_ROWS[i].is_deleted;
			}
		}
		MASTER_FIELDS.forEach((field, _) => payload = [...payload, row[field.key] || '']);
		AddOrEditState = 'Editing '+PURPOSE+' record #'+id;
		ShowAddEdit_PopUp = true;
		isEdit = true;
		idToEdit = id;
	}

  async function AddNew() {
    if (payload[0] == '') return notifier.showError('Error: field cannot be empty');
    SubmitAddEdit = true;
		let changed = 0;
    let bodyFormData = new FormData();
    bodyFormData.append('a', 'save');
    bodyFormData.append('id', '0');
		for (let i = 0; i < payload.length; i++) {
			if (payload[i]) {
				const key = MASTER_FIELDS[i].key;
				bodyFormData.append(key, payload[i]);
				changed += 1;
			}
		}
		bodyFormData.append('_changed', ''+changed);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
      SubmitAddEdit = false;
			ShowAddEdit_PopUp = false;
      notifier.showSuccess(response.data.info || 'new '+PURPOSE+' created');
			searchInput = '';
			GoToPage(currentPage);
    })
    .catch(function (e) {
      SubmitAddEdit = false;
			ShowAddEdit_PopUp = false;
      notifier.showError('Error: cannot create a new '+PURPOSE);
    });
  }

	async function Edit_N(actionPurpose) {
		if (payload[0] == '')  return notifier.showError('Error: field cannot be empty');
    SubmitAddEdit = true;
		let changed = 0;
    let bodyFormData = new FormData();
    bodyFormData.append('a', actionPurpose);
    bodyFormData.append('id', idToEdit);
		for (let i = 0; i < payload.length; i++) {
			if (payload[i]) {
				bodyFormData.append(MASTER_FIELDS[i].key, payload[i]);
				changed += 1;
			}
		}
		bodyFormData.append('_changed', ''+changed);
		await axios({
			method: 'post',
			url: URL,
			data: bodyFormData,
			headers: { 'Content-Type': 'multipart/form-data' },
		}).then(function (response) {
			SubmitAddEdit = false;
			notifier.showSuccess(response.data.info || PURPOSE+' #'+idToEdit+' updated');
			searchInput = '';
			GoToPage(currentPage);
		})
		.catch(function (e) {
			SubmitAddEdit = false;
			notifier.showError('Cannot '+actionPurpose+' '+PURPOSE+' #'+idToEdit+', with status: '+e.response.status+' '+e.response.statusText);
		});
		idToEdit = '0';
		isEdit = false;
		ShowAddEdit_PopUp = false;
	}
	// ===============+ Add/Edit/Delete/Restore +==============

  // ========== Pagination
	// BUG on backend, cannot get rows 400 and 1000
	// Only show 10, 20, 40, 100, 200 rows
	const rowsToShow = [10, 20, 40, 100, 200, 400, 1000];
	let showRowsNum = false;
	const toggleRowsNum = () => showRowsNum = !showRowsNum;
	
  async function toRow(row) {
		isAjaxSubmitted = true;
		if (row > 200) currentRows = 200; else currentRows = row;
		if (sortTableAsc) {ROWS_ORDER = '["-'+colToSort+'"]'} else ROWS_ORDER = '["+'+colToSort+'"]';
		ROWS_OFFSET = 0;
		let bodyFormData = new FormData();
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', row);
		bodyFormData.append('offset', ''+ROWS_OFFSET);
		bodyFormData.append('order', ROWS_ORDER);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			currentPage = 1;
			isAjaxSubmitted = false;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: ' + e);
			console.log('Error: ', e);
    });
		showRowsNum = false;
		getPaginationShow();
	}

	async function NextPage(page) {
		if (page > paginationTotal) return;
		isAjaxSubmitted = true;
		if (sortTableAsc) {ROWS_ORDER = '["-'+colToSort+'"]'} else ROWS_ORDER = '["+'+colToSort+'"]';
		ROWS_OFFSET = (currentRows*page) - currentRows;
		let bodyFormData = new FormData();
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', `${currentRows}`);
		bodyFormData.append('offset', ''+ROWS_OFFSET);
		bodyFormData.append('order', ROWS_ORDER);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			currentPage = page;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: ' + e);
			console.log('Error: ', e)
    });
		getPaginationShow()
	}

	async function PrevPage(page) {
		isAjaxSubmitted = true;
		if (sortTableAsc) {ROWS_ORDER = '["-'+colToSort+'"]'} else ROWS_ORDER = '["+'+colToSort+'"]';
		ROWS_OFFSET = ROWS_OFFSET - currentRows;
		let bodyFormData = new FormData();
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', `${currentRows}`);
		bodyFormData.append('offset', `${ROWS_OFFSET}`);
		bodyFormData.append('order', ROWS_ORDER);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			currentPage = page;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: ' + e);
			console.log('Error: ', e);
    });
		getPaginationShow()
	}

	async function FirstPage() {
		isAjaxSubmitted = true;
		if (currentPage == 1) return
		if (sortTableAsc) {ROWS_ORDER = '["-'+colToSort+'"]'} else ROWS_ORDER = '["+'+colToSort+'"]';
		ROWS_OFFSET = 0;
		let bodyFormData = new FormData();
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', ''+currentRows);
		bodyFormData.append('offset', ''+ROWS_OFFSET);
		bodyFormData.append('order', ROWS_ORDER);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			currentPage = 1;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: ' + e);
			console.log('Error: ', e);
    });
		getPaginationShow()
	}

	async function LastPage() {
		if (currentPage == paginationTotal) return;
		isAjaxSubmitted = true;
		if (sortTableAsc) {ROWS_ORDER = '["-'+colToSort+'"]'} else ROWS_ORDER = '["+'+colToSort+'"]';
		ROWS_OFFSET = Number(roundTotal) - Number(currentRows);
		let bodyFormData = new FormData();
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', ''+currentRows);
		bodyFormData.append('offset', ''+ROWS_OFFSET);
		bodyFormData.append('order', ROWS_ORDER);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			currentPage = paginationTotal;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: ' + e);
			console.log('Error: ', e);
    });
		getPaginationShow()
	}

	async function GoToPage(page) {
		isAjaxSubmitted = true;
		if (sortTableAsc) {ROWS_ORDER = '["-'+colToSort+'"]'} else ROWS_ORDER = '["+'+colToSort+'"]';
		if (page === 1) { ROWS_OFFSET = 0 } else ROWS_OFFSET = (currentRows*page) - currentRows
		let bodyFormData = new FormData();
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', `${currentRows}`);
		bodyFormData.append('offset', `${ROWS_OFFSET}`);
		bodyFormData.append('order', ROWS_ORDER);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			currentPage = page;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: ' + e);
			console.log('Error: ', e)
    });
		getPaginationShow()
	}
	// ========== Pagination

	async function sortColumn(col) {
		isAjaxSubmitted = true;
		colToSort = col;
		if (sortTableAsc) {ROWS_ORDER = '["+'+col+'"]'} else ROWS_ORDER = '["-'+col+'"]';
		if (currentPage !== 1) ROWS_OFFSET = currentRows*currentPage;
		let bodyFormData = new FormData();
    bodyFormData.append('a', 'search');
    bodyFormData.append('limit', ''+currentRows);
		bodyFormData.append('offset', ''+ROWS_OFFSET);
		bodyFormData.append('order', ROWS_ORDER);
    await axios({
      method: 'post',
      url: URL,
      data: bodyFormData,
      headers: { 'Content-Type': 'multipart/form-data' },
    }).then(function (response) {
			isAjaxSubmitted = false;
			MASTER_ROWS = response.data.rows;
    })
    .catch(function (e) {
			isAjaxSubmitted = false;
			notifier.showError('Error: ' + e);
			console.log('Error: ', e)
    });
		sortTableAsc = !sortTableAsc;
	}

	let table;
	function exportToExcel() {
    // @ts-ignore
    const data = table, excelFile = XLSX.utils.table_to_book(data, {sheet: "sheet1"});
		// @ts-ignore
    XLSX.write(excelFile, { bookType: 'xlsx', bookSST: true, type: 'base64' });
		// @ts-ignore
    XLSX.writeFile(excelFile, PURPOSE+'.xlsx');
	}
</script>

{#if filterTableReady}
  <FilterTable bind:this={filterTable} bind:filterColumns={filterColumns} on:click={filterTableOK}/>
{/if}

{#if CAN_ADD_OR_EDIT}
  {#if ShowAddEdit_PopUp}
	  <div class="popup_container">
		  <div class="popup">
			  <header>
				  <h2>{AddOrEditState}</h2>
				  <button on:click={() => ShowAddEdit_PopUp = false}>
				    <Icon size="22" color="var(--red-005)" src={IoClose}/>
				  </button>
			  </header>
			  <div class="forms">
          {#if MASTER_FIELDS && MASTER_FIELDS.length}
					  {#each MASTER_FIELDS as field, idx}
              {#if field.key === 'group_id'}
                <InputCustom
									id={field.key}
									label={field.label}
									placeholder={field.tooltip || ''}
									bind:value={payload[idx]}
									type={field.type}
									values={MASTER_GROUPS}
									isObject
								/>
              {:else}
                <InputCustom
									id={field.key}
									label={field.label}
									placeholder={field.tooltip || ''}
									bind:value={payload[idx]}
									type={field.type || 'text'}
								/>
              {/if}
            {/each}
          {/if}
			  </div>
			  <div class="foot">
          <div class="left">
            <button class="reset" on:click={resetAddEditInput} title="Reset">Reset</button>
						{#if !isDeleted}
							<button class="delete" on:click={() => Edit_N('delete')} title="Delete">Save and Delete</button>
						{/if}
						{#if isDeleted}
							<button class="restore" on:click={() => Edit_N('restore')} title="Restore">Save and Restore</button>
						{/if}
          </div>
          <div class="right">
						{#if SubmitAddEdit}
							<Icon className="spin" color="var(--sky-006)" size="20" src={FiLoader} />
						{/if}
					  <button class="cancel" on:click={() => {ShowAddEdit_PopUp = false; resetAddEditInput()}}>Cancel</button>
					  <button class="ok" on:click={()=> isAdd?AddNew():Edit_N('save')}>Save</button>
          </div>
			  </div>
		  </div>
	  </div>
  {/if}
{/if}

<div class="table_root">
  <div class="actions_container">
    <div class="left">
      <div class="debug">
        <div class="showing">
          <p>Debug table <span class="text-sky">1</span>/<span class="text-sky">{MASTER_COLS.length}</span> of <span class="text-sky">{MASTER_COLS.length}</span> record(s)</p>
        </div>
        <button class="btn" on:click={() => filterTable.Show()}>
          <Icon color="#FFF" size="16" src={AiOutlineEyeInvisible}/>
        </button>
      </div>
      <div class="actions_btn">
        {#if CAN_ADD_OR_EDIT}
          <button class="btn add" title={'Create new '+PURPOSE} on:click={handleAdd}>
            <Icon color="#FFF" size="18" src={CgMathPlus}/>
          </button>
        {/if}
        {#if CAN_EXPORT_EXCEL}
          <button class="btn export" title="Export to Excel" on:click={exportToExcel}>
            <Icon color="#FFF" size="18" src={AiOutlineFileExcel}/>
          </button>
        {/if}
      </div>
			{#if isAjaxSubmitted}
        <div class="loader">
          <Icon className="spin" color="var(--sky-007)" size="28" src={FiLoader} />
        </div>
      {/if}
    </div>
    <div class="right">
			{#if CAN_SEARCH_ROW}
      	<div class="search_handler">
        	<input
         		on:keydown={enterSearchInput}
          	placeholder="Search..."
          	type="text"
          	name="searchRow"
          	id="searchRow"
          	class="search"
          	bind:value={searchInput}
        	/>
        	<button class="search_btn" on:click={searchRow}>
          	<Icon color="#FFF" size="16" src={IoSearch}/>
        	</button>
      	</div>
			{/if}
    </div>
  </div>

  {#if filterTableReady}
    <div class="table_container">
      <table bind:this={table}>
        <thead>
          <tr>
            <th class="no">No</th>
            {#if CAN_ADD_OR_EDIT}
              <th class="a_row">Actions</th>
            {/if}
            {#each filterColumns as fc}
              {#if fc.isVisible}
                <th class="sort_column" on:click={() => sortColumn(fc.key)}>
									<span>{fc.label}</span>
									{#if sortTableAsc}
										<Icon className="sort_icon" size="15" color="var(--gray-007)" src={IoArrowUpSharp}/>
									{/if}
									{#if !sortTableAsc}
										<Icon size="15" color="var(--gray-007)" src={IoArrowDownSharp}/>
									{/if}
								</th>
              {/if}
            {/each}
          </tr>
        </thead>
        <tbody>
          {#if MASTER_ROWS && MASTER_ROWS.length}
            {#each MASTER_ROWS as row, _ (row.id)}
							<tr class:text-red={row.is_deleted}>
                <th>{ROWS_OFFSET+MASTER_ROWS.indexOf(row) + 1}</th>
                {#if CAN_ADD_OR_EDIT}
                  <td class="a_row">
										<div class="actions">
											<button class="btn edit" title="Edit" on:click={()=>handleEdit(row.id)}>
												<Icon size="13" color="#FFF" src={RiDesignBallPenLine}/>
											</button>
                      {#if PURPOSE == 'station'}
												<a href={'/superadmin/station_log_daily/'+row.id} class="btn info" title="Edit logs (daily)">
													<Icon size="13" color="#FFF" src={RiDeviceDatabase2Line}/>
												</a>
												<a href={'/owner/private_station/'+row.id} class="btn info" title="See graph">
													<Icon size="13" color="#FFF" src={FaChartBar}/>
												</a>
												<a href={'/owner/station_info/'+row.id} class="btn info" title="See/edit station information">
													<Icon size="13" color="#FFF" src={RiSystemInformationLine}/>
												</a>
												<a href={'/superadmin/set_predictions/'+row.id} class="btn info" title="Insert/update prediction">
													<Icon size="13" color="#FFF" src={FaSolidChartLine}/>
												</a>
                      {/if}
										</div>
									</td>
                {/if}
                {#each filterColumns as fc}
                  {#if fc.isVisible}
										{#if fc.key === 'group_id'}
											<td>{MASTER_GROUPS[Number(row[fc.key])]}</td>
										{:else if fc.key === 'updated_at' || fc.key === 'created_at'}
											<td>{utcDatetime(new Date(Number(row[fc.key]) * 1000))}</td>
										{:else}
											<td>{row[fc.key] || '--'}</td>
										{/if}
                  {/if}
                {/each}
              </tr>
            {/each}
          {/if}
        </tbody>
      </table>
    </div>
  {/if}

  <div class="pagination_container">
    <div class="filter">
      <div class="showing">
        <p>Showing <span class="text-sky">{MASTER_ROWS.length}</span>/<span class="text-sky">{currentRows}</span> of <span class="text-sky">{ROWS_COUNT}</span> record(s)</p>
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
          <Icon className={showRowsNum ? 'rotate_right' : 'dropdown'} size="13" src={CgChevronDown}/>
        </button>
      </div>
    </div>
    <div class="pagination">
      <button disabled={currentPage == 1} class="btn to" on:click={FirstPage} title="Go to first page">
        <Icon size="16" src={CgChevronDoubleLeft}/>
      </button>
      <button disabled={currentPage == 1} class="btn to" on:click={() => PrevPage(currentPage - 1)}  title="Go to previous page">
        <Icon size="16" src={CgChevronLeft}/>
      </button>
      {#each paginationShow as i}
        <button disabled={currentPage == i} class={currentPage === i ? 'btn active' : 'btn'} on:click={() => GoToPage(i)}  title={`Go to page ${i}`}>{i}</button>
      {/each}
      <button disabled={currentPage == paginationTotal} class="btn to" on:click={() => NextPage(currentPage + 1)}  title="Go to next page">
        <Icon size="16" src={CgChevronRight}/>
      </button>
      <button disabled={currentPage == paginationTotal} class="btn to" on:click={LastPage} title="Go to last page">
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
    justify-content: space-between;
		gap: 10px;
		align-items: center;
		padding: 15px 20px;
		border-top: 1px solid var(--gray-004);
	}

	.popup_container .popup .foot .left {
		display: flex;
		flex-direction: row;
		gap: 10px;
		align-items: center;
	}

  .popup_container .popup .foot .right {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
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

  .popup_container .popup .foot button.reset {
		background-color: var(--amber-006);
		border: 1px solid var(--amber-006);
	}

  .popup_container .popup .foot button.reset:hover {
    background-color: var(--amber-005);
  }

	.popup_container .popup .foot button.delete {
		background-color: var(--red-006);
		border: 1px solid var(--red-006);
	}

  .popup_container .popup .foot button.delete:hover {
    background-color: var(--red-005);
  }

	.popup_container .popup .foot button.restore {
		background-color: var(--sky-006);
		border: 1px solid var(--sky-006);
	}

  .popup_container .popup .foot button.restore:hover {
    background-color: var(--sky-005);
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

	.popup_container .popup .foot button.cancel {
		background-color: #fbbf2420;
		color: var(--amber-005);
		border: 1px solid var(--amber-005);
	}

  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  :global(.dropdown) {
		transition: all .2s ease-in-out;
	}

	:global(.rotate) {
		transition: all .2s ease-in-out;
		transform: rotate(90deg);
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
    gap: 20px;
  }

  .table_root .actions_container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  .table_root .actions_container .left,
  .table_root .actions_container .right {
		display: flex;
		flex-direction: row;
		align-items: center;
    gap: 10px;
	}

	.table_root .actions_container .left .debug {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 8px;
		padding: 5px 5px 5px 20px;
		border: 1px solid var(--gray-004);
		border-radius: 999px;
	}

	.table_root .actions_container .left .debug .showing .text-sky {
		color: var(--sky-006);
		font-weight: 600;
		padding: 5px;
	}

	.table_root .actions_container .left .debug .btn {
		border: none;
		background-color: var(--sky-006);
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
		background-color: var(--sky-005);
	}

  .table_root .actions_container .right .search_handler {
    display: flex;
    flex-direction: row;
    width: fit-content;
    height: fit-content;
    position: relative;
  }

  .table_root .actions_container .right .search_handler input.search {
    padding: 10px 50px 10px 15px;
    border-radius: 999px;
    border: 1px solid var(--gray-005);
    width: 300px;
  }

  .table_root .actions_container .right .search_handler input.search:focus {
    border-color: var(--sky-005);
    outline: 1px solid var(--sky-005);
  }

  .table_root .actions_container .right .search_handler .search_btn {
    position: absolute;
    right: 5px;
    background-color: var(--sky-006);
    padding: 6px 12px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    border-radius: 999px;
    top: 5px;
    cursor: pointer;
  }

  .table_root .actions_container .right .search_handler .search_btn:hover {
    background-color: var(--sky-005);
  }

  .table_root .actions_container .actions_btn {
    display: flex;
    flex-direction: row;
    gap: 10px;
    align-items: center;
  }

  .table_root .actions_container .actions_btn .btn {
    border: none;
		color: #FFF;
		width: fit-content;
		padding: 7px 15px;
		border-radius: 9999px;
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
		gap: 3px;
		cursor: pointer;
  }

  .table_root .actions_container .actions_btn .btn.add {
    background-color: var(--green-006);
  }

  .table_root .actions_container .actions_btn .btn.add:hover {
    background-color: var(--green-005);
  }

  .table_root .actions_container .actions_btn .btn.export {
    background-color: var(--gray-007);
  }

  .table_root .actions_container .actions_btn .btn.export:hover {
    background-color: var(--gray-006);
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


  .table_root .table_container table thead tr th.a_row {
		max-width: fit-content;
	}

	.table_root .table_container table thead tr th:last-child {
		border-right: none;
	}

	.table_root .table_container table thead tr th.sort_column {
		cursor: pointer;
		text-wrap: nowrap;
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
		gap: 5px;
	}

	.table_root .table_container table tbody tr td .actions .btn {
		border: none;
		padding: 5px 10px;
		border-radius: 999px;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.table_root .table_container table tbody tr td .actions .btn.edit {
		background-color: var(--amber-005);
		color: #FFF;
	}

	.table_root .table_container table tbody tr td .actions .btn.edit:hover {
		background-color: var(--amber-006);
	}

  .table_root .table_container table tbody tr td .actions .btn.info {
		background-color: var(--sky-005);
		color: #FFF;
	}

	.table_root .table_container table tbody tr td .actions .btn.info:hover {
		background-color: var(--sky-006);
	}

  .table_root .pagination_container {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}

	.table_root .pagination_container .filter {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 8px;
		padding: 5px 5px 5px 20px;
		border: 1px solid var(--gray-004);
		border-radius: 999px;
	}

	.table_root .pagination_container .filter .showing .text-sky {
		color: var(--sky-006);
		font-weight: 600;
		padding: 5px;
	}

	.table_root .pagination_container .filter .row_to_show {
		position: relative;
		width: fit-content;
		height: fit-content;
	}

	.table_root .pagination_container .filter .row_to_show .btn {
		border: none;
		background-color: var(--sky-006);
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

	.table_root .pagination_container .filter .row_to_show .btn:hover {
		background-color: var(--sky-005);
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
		background-color: #0ea5e920;
		color: var(--sky-007);
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
		background-color: #FFF;
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
		background-color: #0ea5e920;
		color: var(--sky-006);
		font-weight: 600;
		border: 1px solid var(--sky-004);
	}

	.table_root .pagination_container .pagination .btn.to {
		background-color: var(--sky-006);
		color: #FFF;
		font-weight: 600;
		border: none;
	}

	.table_root .pagination_container .pagination .btn.to:hover {
		background-color: var(--sky-005);
	}

	.table_root .pagination_container .pagination .btn.to:disabled {
		background-color: var(--gray-002);
		color: var(--gray-006);
		font-weight: 600;
		border: 1px solid var(--gray-004);
		cursor: not-allowed;
	}

	@media only screen and (max-width : 768px) {
		.popup_container {
			padding: 15px;
		}
		
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