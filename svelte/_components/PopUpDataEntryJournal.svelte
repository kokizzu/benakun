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
   */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputCustom from './InputCustom.svelte';

  export let heading      = 'Add journal';
  export let isSubmitted  = false;
  let isShow              = false;

  export let isDebit      = true;
  export let isSales      = false;
  export let isChildOnly  = false;
  export let coaChildren = {}

  export let debitIDR = 0;
  export let creditIDR = 0;
  export let description = '';
  export let date = '';

  export let startDate = new Date();
  export let endDate = new Date();
  let payloadsSales = /** @type pySales[] */ ([]);

  export let OnSubmit = async function() {}
  export let OnSubmitWithSales = async function(payloadsSales) {}
  export const Show = () => {
    payloadsSales = [];
    isShow = true;
  }
  export const ShowWithSales = () => {
    isSales = true;
    payloadsSales = [
    {
      creditIDR: 0,
      debitIDR: 0,
      description: '',
      date: new Date(),
      salesCount: 0,
      salesPriceIDR: 0,
      coaId: ''
    }]
    isShow = true;
  }

  export const Hide = () => isShow = false;

  export const Reset = () => {
    payloadsSales = [];
    isDebit      = true;
    isSales      = false;
    isChildOnly  = false;
    coaChildren = {}

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
      {#if isSales}
        <div class="sales_forms_container">
          <div class="table_container">
            <table>
              <thead>
                <tr>
                  <th>#</th>
                  {#if isDebit}
                    <th>Debit (IDR)</th>
                  {:else}
                    <th>Credit (IDR)</th>
                  {/if}
                  
                  <th>Description</th>
                  {#if isChildOnly}
                    <th>COA</th>
                  {/if}
                  <th>Date</th>
                  <th>Sales Count</th>
                  <th>Sales Price (IDR)</th>
                </tr>
              </thead>
              <tbody>
                {#each (payloadsSales || []) as pySales (pySales)}
                  <tr>
                    <td><button on:click={() => payloadsSales = payloadsSales.filter((item) => item !== pySales)}>Remove</button></td>
                    {#if isDebit}
                      <td>
                        <input type="number" min=0 bind:value={pySales.debitIDR}/>
                      </td>
                    {:else}
                      <td>
                        <input type="number" min=0 bind:value={pySales.creditIDR}/>
                      </td>
                    {/if}
                    <td>
                      <textarea
                        placeholder="Description"
                        bind:value={pySales.description}
                        rows="1"
                      />
                    </td>
                    {#if isChildOnly}
                      <td>
                        <select bind:value={pySales.coaId}>
                          {#each Object.entries(coaChildren) as [k, v], idx}
                            <option value={k} selected={pySales.coaId === k}>{`${idx+1}: ${v}`}</option>
                          {/each}
                        </select>
                      </td>
                    {/if}
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
                  coaId: ''
                }
              ]
            }}>
              Add row
            </button>
          </div>
        </div>
      {/if}
      {#if !isSales}
        <div class="sales_forms_container">
          <div class="table_container">
            <table>
              <thead>
                <tr>
                  {#if isDebit}
                    <th>Debit (IDR)</th>
                  {:else}
                    <th>Credit (IDR)</th>
                  {/if}
                  <th>Description</th>
                  <th>Date</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  {#if isDebit}
                    <td>
                      <input type="number" min=0 bind:value={debitIDR}/>
                    </td>
                  {:else}
                    <td>
                      <input type="number" min=0 bind:value={creditIDR}/>
                    </td>
                  {/if}
                  <td>
                    <textarea
                      placeholder="Description"
                      bind:value={description}
                      rows="1"
                    />
                  </td>
                  <td>
                    <input type="date" bind:value={date}/>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      {/if}
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
        <button
          class="ok"
          on:click|preventDefault={() => {
            if (isSales) {
              OnSubmitWithSales(payloadsSales)
            } else {
              OnSubmit()
            }
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
</style>