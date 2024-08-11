<script>
  /** @typedef {import('./types/transaction').TransactionJournal} TransactionJournal */
  /**
   * @typedef {Object} pySales
   * @property {string|number} creditIDR
   * @property {string|number} debitIDR
   * @property {string} description
   * @property {string|Date} date
   * @property {number} salesCount
   * @property {string|number} salesPriceIDR
   * @property {string|number} coaId
	 * @property {boolean} isChildOnly
   */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
	import { RiSystemDeleteBinLine } from '../node_modules/svelte-icons-pack/dist/ri';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputCustom from './InputCustom.svelte';

  export let heading      = 'Add journal';
  export let isSubmitted  = false;
  let isShow              = false;

  export let coas = {}
	export let coaId = '0';

  export let debitIDR = 0;
  export let creditIDR = 0;
  export let description = '';
  export let date = '';

  export let startDate = new Date();
  export let endDate = new Date();
  let payloadsSales = /** @type pySales[] */ ([]);

  export let OnSubmitWithSales = async function(payloadsSales) {}
  export const Show = () => {
    payloadsSales = [{
			creditIDR: 0,
			debitIDR: 0,
			description: '',
			date: new Date(),
			salesCount: 0,
			salesPriceIDR: 0,
			coaId: '',
			isChildOnly: false
		}];
    isShow = true;
  }
  export const ShowWithSales = () => {
    payloadsSales = [{
      creditIDR: 0,
      debitIDR: 0,
      description: '',
      date: new Date(),
      salesCount: 0,
      salesPriceIDR: 0,
      coaId: '',
			isChildOnly: false
    }]
		isShow = true;
  }

  export const Hide = () => isShow = false;

  export const Reset = () => {
    payloadsSales = [];

    debitIDR = 0;
    creditIDR = 0;
    description = '';
    date = '';

    startDate = new Date();
    endDate = new Date();
  }
  
  const cancel = () => {
    isShow = false;
  }
</script>

<div class={`popup_container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>{heading}</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <div class="form_rows">
				<InputCustom
					className="input_custom"
					id="coaId"
					label="CoA "
					bind:value={coaId}
					type="combobox"
					values={coas}
					isObject
				/>
        <InputCustom
          className="input_custom"
          id="startDate"
          label="Start Date"
          bind:value={startDate}
          type="date"
        />
        <InputCustom
          className="input_custom"
          id="endDate"
          label="End Date"
          bind:value={endDate}
          type="date"
        />
      </div>
        <div class="sales_forms_container">
          <div class="actions">
            <button class="btn" on:click={() => {
              payloadsSales = [
                ...payloadsSales,
                {
                  creditIDR: 0,
                  debitIDR: 0,
                  description: '',
                  date: new Date(),
                  salesCount: 0,
                  salesPriceIDR: 0,
                  coaId: '',
									isChildOnly: false
                }
              ]
            }}>
              Add row
            </button>
          </div>
          <div class="table_container">
            <table>
              <thead>
								{#if payloadsSales && payloadsSales.length > 0}
									<tr>
										<th>#</th>
										<th>Debit (IDR)</th>
										<th>Credit (IDR)</th>
										<th>Description</th>
										<th>Date</th>
										<th>Sales Count</th>
										<th>Sales Price (IDR)</th>
									</tr>
								{/if}
								{#if !payloadsSales || payloadsSales.length === 0}
									<tr>
										<th>Add row to see inputs</th>
									</tr>
								{/if}
              </thead>
              <tbody>
                {#each (payloadsSales || []) as pySales, i (pySales)}
                  <tr>
                    <td>
											<button on:click={() => payloadsSales = payloadsSales.filter((item) => item !== pySales)}>
												<Icon
													src={RiSystemDeleteBinLine}
													size="16"
												/>
											</button>
										</td>
										<td>
											<input type="number" min=0 bind:value={pySales.debitIDR}/>
										</td>
										<td>
											<input type="number" min=0 bind:value={pySales.creditIDR}/>
										</td>
                    <td>
                      <textarea
                        placeholder="Description"
                        bind:value={pySales.description}
                        rows="1"
                      />
                    </td>
                    <td>
                      <input type="date" bind:value={pySales.date}/>
                    </td>
                    <td>
                      <input type="number" min=0 bind:value={pySales.salesCount}/>
                    </td>
                    <td>
                      <input type="number" min=0 bind:value={pySales.salesPriceIDR}/>
                    </td>
                  </tr>
                {/each}
              </tbody>
            </table>
          </div>
        </div>
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
        <button
          class="ok"
          on:click|preventDefault={() => {
						OnSubmitWithSales(payloadsSales)
          }}
          disabled={isSubmitted}
        >
          {#if !isSubmitted}
            <span>Save</span>
          {/if}
          {#if isSubmitted}
            <Icon className="spin" color="#FFF" size="14" src={FiLoader} />
          {/if}
        </button>
      </div>
    </div>
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

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .popup_container {
    display: none;
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
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

  .popup_container.show {
    display: flex;
  }

	.popup_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 100%;
    margin: 80px;
		display: flex;
		flex-direction: column;
	}

  .popup_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 10px 20px;
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

	.popup_container .popup header button:active {
		background-color: #ef444430;
	}

	.popup_container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

  .popup_container .popup .forms .form_rows {
    display: flex;
    flex-direction: row;
    gap: 10px;
  }

  :global(.popup_container .popup .forms .form_rows .input_custom) {
    width: 100%;
  }

  .popup_container .popup .forms .sales_forms_container {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .popup_container .popup .forms .sales_forms_container .table_container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .popup_container .popup .forms .sales_forms_container .table_container table thead tr th {
    font-weight: normal;
    text-align: left;
		white-space: nowrap;
  }

  .popup_container .popup .forms .sales_forms_container .table_container input,
  .popup_container .popup .forms .sales_forms_container .table_container textarea {
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid var(--gray-003);
    width: 100%;
  }

  .popup_container .popup .forms .sales_forms_container .table_container textarea {
    resize: vertical;
    min-height: 40px;
    max-height: 100px;
  }

  .popup_container .popup .forms .sales_forms_container .actions {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    gap: 10px;
  }

  .popup_container .popup .forms .sales_forms_container .actions .btn {
    padding: 8px 13px;
    border-radius: 9999px;
    border: none;
    color: #FFF;
    cursor: pointer;
    font-weight: 600;
    background-color: var(--blue-006);
  }

  .popup_container .popup .forms .sales_forms_container .actions .btn:hover {
    background-color: var(--blue-005);
  }

	.popup_container .popup .foot {
		display: flex;
		flex-direction: row;
    justify-content: space-between;
		gap: 10px;
		align-items: center;
		padding: 10px 20px;
		border-top: 1px solid var(--gray-004);
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
	}

	.popup_container .popup .foot button.ok {
		background-color: var(--green-006);
    display: flex;
    justify-content: center;
    align-items: center;
	}

	.popup_container .popup .foot button.ok:hover {
		background-color: var(--green-005);
	}

	.popup_container .popup .foot button.ok:disabled {
		cursor: not-allowed;
		background-color: var(--gray-003);
		color: var(--gray-007);
	}

	.popup_container .popup .foot button.cancel {
		background-color: #fbbf2430;
		color: var(--amber-005);
	}

	.popup_container .popup .foot button.cancel:hover {
		background-color: #fbbf2450;
	}
	
	/* The switch - the box around the slider */
  .switcher {
    position: relative;
    display: inline-block;
    width: 50px;
    height: 24px;
    margin-left: 0 !important;
  }
  
  .switcher input {
    opacity: 0 !important;
    width: 0 !important;
    height: 0 !important;
  }
  
  .switcher .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: var(--gray-005);
    -webkit-transition: .2s;
    transition: .2s;
  }

  .switcher .slider:before {
    position: absolute;
    content: "";
    height: 16px;
    width: 16px;
    left: 4px;
    bottom: 4px;
    background-color: white;
    -webkit-transition: .2s;
    transition: .2s;
  }

  .switcher input:checked + .slider {
    background-color: var(--sky-006) !important;
  }

  .switcher input:focus + .slider {
    box-shadow: 0 0 1px var(--sky-006) !important;
  }

  .switcher input:checked + .slider:before {
    -webkit-transform: translateX(26px);
    -ms-transform: translateX(26px);
    transform: translateX(26px);
  }
  
  .switcher .slider {
    border-radius: 34px;
  }

  .switcher .slider:before {
    border-radius: 50%;
  }
</style>