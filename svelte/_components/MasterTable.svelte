<script>
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

	// @typedef {import('.')}

	export let ACCESS = {};

	let paginationShow = [1, 2, 3, 4, 5];
	let currentPage = 1, paginationTotal = 1;
	let sortTableAsc = false;
	
</script>


<div class="table_root">
  <div class="actions_container">
    <div class="left">
      <div class="debug">
        <div class="showing">
          <p>Debug table <span class="text-sky">1</span>/<span class="text-sky">100</span> of <span class="text-sky">100</span> record(s)</p>
        </div>
        <button class="btn">
          <Icon color="#FFF" size="16" src={AiOutlineEyeInvisible}/>
        </button>
      </div>
      <div class="actions_btn">
          <button class="btn add" title="Create new">
            <Icon color="#FFF" size="18" src={CgMathPlus}/>
          </button>
          <button class="btn export" title="Export to Excel">
            <Icon color="#FFF" size="18" src={AiOutlineFileExcel}/>
          </button>
      </div>
        <div class="loader">
          <Icon className="spin" color="var(--sky-007)" size="28" src={FiLoader} />
        </div>
    </div>
    <div class="right">
			<div class="search_handler">
				<input
					placeholder="Search..."
					type="text"
					name="searchRow"
					id="searchRow"
					class="search"
				/>
				<button class="search_btn">
					<Icon color="#FFF" size="16" src={IoSearch}/>
				</button>
			</div>
    </div>
  </div>

	<div class="table_container">
		<table>
			<thead>
				<tr>
					<th class="no">No</th>
						<th class="a_row">Actions</th>
							<th class="sort_column">
								<span>Name</span>
								{#if sortTableAsc}
									<Icon className="sort_icon" size="15" color="var(--gray-007)" src={IoArrowUpSharp}/>
								{/if}
								{#if !sortTableAsc}
									<Icon size="15" color="var(--gray-007)" src={IoArrowDownSharp}/>
								{/if}
							</th>
				</tr>
			</thead>
			<tbody>
				<td>hahaha</td>
			</tbody>
		</table>
	</div>

  <div class="pagination_container">
    <div class="filter">
      <div class="showing">
        <p>Showing <span class="text-sky">10</span>/<span class="text-sky">20</span> of <span class="text-sky">40</span> record(s)</p>
      </div>
      <div class="row_to_show">
        <button class="btn" >
          <span>30</span>
          <Icon size="13" src={CgChevronDown}/>
        </button>
      </div>
    </div>
    <div class="pagination">
      <button disabled={currentPage == 1} class="btn to" title="Go to first page">
        <Icon size="16" src={CgChevronDoubleLeft}/>
      </button>
      <button disabled={currentPage == 1} class="btn to" title="Go to previous page">
        <Icon size="16" src={CgChevronLeft}/>
      </button>
      {#each paginationShow as i}
        <button disabled={currentPage == i} class={currentPage === i ? 'btn active' : 'btn'} title={`Go to page ${i}`}>{i}</button>
      {/each}
      <button disabled={currentPage == paginationTotal} class="btn to" title="Go to next page">
        <Icon size="16" src={CgChevronRight}/>
      </button>
      <button disabled={currentPage == paginationTotal} class="btn to" title="Go to last page">
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