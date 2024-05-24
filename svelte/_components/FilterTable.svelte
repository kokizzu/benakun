<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
	import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';

  /**
    * @typedef {Object} FilterTable
    * @property {string} key
    * @property {string} label
    */

  /** @type {FilterTable[]} */
  export let filterColumns = [];
	export let filtersMap = {};

	let isVisible = false;

  export function Show() { isVisible = true }
  export function Hide() { isVisible = false }
</script>

{#if isVisible}
	<div class="popup_filter_container">
		<div class="popup">
			<header>
				<h2>Visibility and filter for DebugColumns</h2>
				<button on:click={() => isVisible = false}>
					<Icon size="22" color="var(--red-005)" src={IoClose}/>
				</button>
			</header>
			<div class="filters">
				{#if filterColumns && filterColumns.length}
					{#each filterColumns as col}
						<div class="row">
							<div class="inputs">
								<label for={col.key}>{col.label}</label>
								<input
									type="text"
									name={col.key}
									id={col.key}
									bind:value={filtersMap[col.key]}
								/>
							</div>
							<div class="actions">
							</div>
						</div>
					{/each}
				{/if}
			</div>
			<div class="foot">
				<button class="cancel" on:click={() => isVisible = false}>Cancel</button>
				<button class="ok" on:click|preventDefault>Ok</button>
			</div>
		</div>
	</div>
{/if}

<style>
	.popup_filter_container {
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 99999999;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		display: flex;
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

	.popup_filter_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 600px;
		display: flex;
		flex-direction: column;
	}

	.popup_filter_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 15px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup_filter_container .popup header h2 {
		margin: 0;
	}

	.popup_filter_container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup_filter_container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup_filter_container .popup .filters {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.popup_filter_container .popup .filters .row {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
	}

	.popup_filter_container .popup .filters .row .inputs {
		display: flex;
		width: 100%;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
	}

	.popup_filter_container .popup .filters .row .inputs label {
		padding: 8px 12px;
		background-color: #0ea5e920;
		color: var(--sky-007);
		font-weight: 600;
		border-radius: 9999px;
		text-transform: capitalize;
	}

	.popup_filter_container .popup .filters .row .inputs input {
		width: 320px;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 10px 12px;
	}

	.popup_filter_container .popup .filters .row .inputs input:focus {
    border-color: var(--sky-005);
    outline: 1px solid var(--sky-005);
  }

	.popup_filter_container .popup .filters .row .actions {
		display: flex;
		flex-direction: row;
		gap: 8px;
	}

	.popup_filter_container .popup .foot {
		display: flex;
		flex-direction: row;
		justify-content: flex-end;
		gap: 10px;
		align-items: center;
		padding: 15px 20px;
		border-top: 1px solid var(--gray-004);
	}

	.popup_filter_container .popup .foot button {
		padding: 8px 13px;
		border-radius: 9999px;
		border: none;
		color: #FFF;
		cursor: pointer;
		font-weight: 600;
	}

	.popup_filter_container .popup .foot button.ok {
		background-color: var(--green-006);
		border: 1px solid var(--green-006);
	}

	.popup_filter_container .popup .foot button.ok:hover {
		background-color: var(--green-005);
	}

	.popup_filter_container .popup .foot button.cancel {
		background-color: #fbbf2420;
		color: var(--amber-005);
		border: 1px solid var(--amber-005);
	}

	@media only screen and (max-width : 768px) {
		.popup_filter_container {
			padding: 15px;
		}
	}
</style>