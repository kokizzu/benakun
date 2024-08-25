<script>
  /** @typedef {import('../../_components/types/transaction.js').TransactionTemplate} TransactionTemplate */
  /** @typedef {import('../../_components/types/transaction.js').TransactionJournal} TransactionJournal */
  /** @typedef {import('../../_components/types/transaction.js').DetailObjectTransaction} DetailObjectTransaction */
  /** @typedef {import('../../_components/types/transaction.js').TransactionTplDetail} TransactionTplDetail */
  /** @typedef {import('../../_components/types/organization.js').Org} Org */
  /** @typedef {import('../../_components/types/user.js').User} User */
  /**
   * @typedef {Object} PayloadTransactionJournals
   * @property {number} creditIDR
   * @property {number} debitIDR
   * @property {string} descriptions
   * @property {string|Date} date
   * @property {number} coaId
   * @property {boolean?} isSales
   * @property {boolean?} isAutoSum
   * @property {number?} salesCount?
   * @property {number?} salesPriceIDR?
   * @property {number?} transactionTplId?
   */

  import MainLayout from '../../_layouts/mainLayout.svelte';
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { RiSystemDeleteBin6Line } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { DataEntryTransactionEntry } from '../../jsApi.GEN';
  import { notifier } from '../../_components/notifier';
  import InputCustom from '../../_components/InputCustom.svelte';
  import { onMount } from 'svelte';

  let transactionTemplate   = /** @type TransactionTemplate */ ({/* transactiontemplate */});
  let transactionTplDetails = /** @type TransactionTplDetail[] */ ([/* transactionTplDetails */]);
  let coas                  = /** @type Record<number|string, string> */ ({/* coas */});
  let org                   = /** @type Org */ ({/* org */});
  let user                  = /** @type User */ ({/* user */});
  let coasWithChildren      = /** @type Record<number|string, Record<number|string, string>> */ ({/* coasWithChildren */});  

  let payloads    = /** @type PayloadTransactionJournals[][] */ ([]);
  let isSubmitted = false;
  let isDataReady = false;

  function dateISOFormat(/** @type number */ dayTo = 0) {
    const dt    = new Date();
    const date  = (dayTo > 0 ? dt.getDate() + dayTo : dt.getDate());
    const month = dt.toLocaleDateString('default', {month: '2-digit'});
    const year  = dt.getFullYear();

    return `${year}-${month}-${date}`;
  }

  onMount(() => {
    if (transactionTplDetails && transactionTplDetails.length > 0) {
      transactionTplDetails.forEach((transactionTplDetail) => {
        payloads.push([{
          coaId: transactionTplDetail.coaId,
          creditIDR: 0,
          debitIDR: 0,
          descriptions: '',
          date: dateISOFormat(0),
          salesCount: 0,
          salesPriceIDR: 0,
          transactionTplId: 0,
          isSales: transactionTplDetail.attributes.includes('sales'),
          // isAutoSum: transactionTplDetail.attributes.includes('autoSum')
          isAutoSum: false
        }])
      })
    }
    isDataReady = true;
  });

  let startDate     = dateISOFormat(0);
  let endDate       = dateISOFormat(4);

  async function Submit() {
    isSubmitted = true;

    let trxJournals = /** @type TransactionJournal[]|any */ ([]);
    let trxTplDetailsId = /** @type number[] */ ([]);
    for (const t of transactionTplDetails) {
      trxTplDetailsId.push(t.id);
    }

    // Payloads is a matrix, loop it twice
    for (const payload of payloads) {
      if (payload.length > 0) {
        for (const trxJournal of payload) {
          let detailObj = '';
          let debit = 0;
          let credit = 0;

          if (trxJournal.isSales) {
            const detail = /** @type DetailObjectTransaction */ ({
              salesCount: trxJournal.salesCount,
              salesPriceIDR: trxJournal.salesPriceIDR+'',
            });
            detailObj = JSON.stringify(detail);

            if (Number(trxJournal.creditIDR) > 0) credit = Number(trxJournal.salesCount) * Number(trxJournal.salesPriceIDR);
            else debit = Number(trxJournal.salesCount) * Number(trxJournal.salesPriceIDR);
          } else {
            if (trxJournal.isAutoSum) {
              // TODO: auto sum
            } else {
              if (Number(trxJournal.creditIDR) > 0) credit = Number(trxJournal.creditIDR);
              else debit = Number(trxJournal.debitIDR);
            }
          }

          trxJournals.push({
            coaId: trxJournal.coaId,
            descriptions: trxJournal.descriptions,
            date: trxJournal.date,
            debit: debit,
            credit: credit,
            detailObj: detailObj
          })
        }
      }
    }

    const i = {
      cmd: 'upsert',
      transactionTplId: transactionTemplate.id,
      transactionTplDetailsId: trxTplDetailsId,
      transactionJournals: trxJournals,
      businessTransaction: {
        startDate: startDate,
        endDate: endDate
      }
    }

    await DataEntryTransactionEntry( // @ts-ignore
      i, /** @returns {Promise<any>} */
      function (/** @type {any} */ o) {
        isSubmitted = false;
        if (o.error) {
          notifier.showError(o.error || 'failed to add journal');
          return
        }
        
        notifier.showSuccess('journal added !');
      }
    )
  }
</script>

<MainLayout>
  <div class="show">
    <header>
      <h1>Entry journal for "{transactionTemplate.name}"</h1>
    </header>
    <div class="data_entry_journal___container">
      <div class="forms_journal">
        <div class="form_date">
          <InputCustom
            type="date"
            className="input_custom"
            id="startDate"
            label="Start Date"
            bind:value={startDate}
          />
          <InputCustom
            id="endDate"
            label="End Date"
            bind:value={endDate}
            type="date"
          />
        </div>
        <div class="company_details">
          <div class="company_name">
            <div class="title">
              <span>Company Name</span>
            </div>
            <h5>{org.name || ''}</h5>
          </div>
          <div class="company_contacts">
            <div class="title">
              <span>Company Contacts</span>
            </div>
            <div class="contacts">
              <p>Head Title: {org.headTitle || ''}</p>
              <p>Email: {user.email}</p>
            </div>
          </div>
        </div>
        {#if isDataReady}
          {#each (transactionTplDetails || []) as transactionTplDetail, idx}
            <div class="form_item">
              <div class="form_header">
                <h3>Coa : {coas[transactionTplDetail.coaId]}</h3>
                {#if (transactionTplDetail.attributes).includes('sales')}
                  <div class="actions">
                    <button class="btn" on:click={() => {
                      payloads[idx] = [...payloads[idx],
                        {
                          coaId: transactionTplDetail.coaId,
                          descriptions: '',
                          creditIDR: 0,
                          debitIDR: 0,
                          date: dateISOFormat(0),
                          salesCount: 0,
                          salesPriceIDR: 0,
                          transactionTplId: transactionTemplate.id,
                          isSales: true,
                          isAutoSum: false
                        }
                      ]
                    }}>
                      Add row
                    </button>
                  </div>
                {/if}
              </div>
              <div class="forms_table">
                <table class="table_transaction_journals">
                  <thead>
                    <tr>
                      {#if (transactionTplDetail.attributes).includes('sales')}
                        <th></th>
                      {/if}
                      <th>Description</th>
                      {#if transactionTplDetail.attributes.includes('childOnly')}
                        <th>CoA</th>
                      {/if}
                      {#if (transactionTplDetail.attributes).includes('sales')}
                        <th>Sales Count</th>
                        <th>Sales Price (IDR)</th>
                      {/if}
                      <th>Date</th>
                      {#if !(transactionTplDetail.attributes).includes('autoSum') || !(transactionTplDetail.attributes).includes('sales')}
                        {#if transactionTplDetail.isDebit}
                          <th>Debit (IDR)</th>
                        {:else}
                          <th>Credit (IDR)</th>
                        {/if}
                      {/if}
                    </tr>
                  </thead>
                  <tbody>
                    {#each (payloads[idx] || []) as py, i}
                      <tr>
                        {#if (transactionTplDetail.attributes).includes('sales')}
                          <td>
                            <button
                              disabled={i === 0}
                              title={i === 0 ? 'Cannot remove first row' : 'Remove row'}
                              class="btn_remove"
                              on:click={() => payloads[idx] = payloads[idx].filter((item) => item !== py)}
                            >
                              <Icon
                                src={RiSystemDeleteBin6Line}
                                size="16"
                              />
                            </button>
                          </td>
                        {/if}
                        <td>
                          <textarea
                            placeholder="Description"
                            rows="1"
                            bind:value={py.descriptions}
                          />
                        </td>
                        {#if transactionTplDetail.attributes.includes('childOnly')}
                          <td>
                            <select name={`coa-${py.coaId}`} bind:value={py.coaId}>
                              {#each Object.entries(coasWithChildren[transactionTplDetail.coaId]) as [k, v], i}
                                <option value={k} selected={i === 0}>{v}</option>
                              {/each}
                            </select>
                          </td>
                        {/if}
                        {#if (transactionTplDetail.attributes).includes('sales')}
                          <td>
                            <input inputmode="numeric" type="number" min=0 bind:value={py.salesCount}/>
                          </td>
                          <td>
                            <input inputmode="numeric" type="number" min=0 bind:value={py.salesPriceIDR} />
                          </td>
                        {/if}
                        <td>
                          <input type="date" bind:value={py.date} />
                        </td>
                        {#if !(transactionTplDetail.attributes).includes('autoSum') || !(transactionTplDetail.attributes).includes('sales')}
                          {#if transactionTplDetail.isDebit}
                            <td>
                              <input inputmode="numeric" type="number" min=0 bind:value={py.debitIDR}/>
                            </td>
                          {:else}
                            <td>
                              <input inputmode="numeric" type="number" min=0 bind:value={py.creditIDR} />
                            </td>
                          {/if}
                        {/if}
                      </tr>
                    {/each}
                  </tbody>
                </table>
              </div>
            </div>
          {/each}
        {/if}
      </div>
      <div class="actions">
        <button
          disabled={isSubmitted}
          class="btn submit"
          on:click={Submit}
        >
          {#if !isSubmitted}
            <span>Submit</span>
          {/if}
          {#if isSubmitted}
            <span>Submitting ...</span>
          {/if}
        </button>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .show {
    display: block;
  }

  header {
    margin-left: 30px;
  }

  .transaction_template_detail .table_container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .transaction_template_detail .table_container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .transaction_template_detail .table_container table thead tr th:last-child {
    border-right: none;
  }

  .transaction_template_detail .table_container table tbody tr td {
    padding: 8px 12px;
  }

	.transaction_template_detail .table_container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.transaction_template_detail .table_container table tbody tr:last-child td,
	.transaction_template_detail .table_container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .transaction_template_detail .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .transaction_template_detail .table_container table tbody tr:last-child td,
  .transaction_template_detail .table_container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .transaction_template_detail .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .transaction_template_detail .table_container table tbody tr td:last-child {
    border-right: none !important;
  }

  .transaction_template_detail .table_container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

  .transaction_template_detail .table_container table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .transaction_template_detail .table_container table tbody tr td .actions .btn:hover {
    background-color: var(--violet-transparent);
  }

  :global(.transaction_template_detail .table_container table tbody tr td .actions .btn:hover svg) {
    fill: var(--violet-005);
  }

  .data_entry_journal___container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 20px;
    border-radius: 8px;
    background-color: #fff;
    border: 1px solid var(--gray-003);
  }

  .data_entry_journal___container .forms_journal {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .data_entry_journal___container .forms_journal .form_date {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
  }

  .data_entry_journal___container .forms_journal .form_item {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 15px;
    background-color: var(--gray-001);
    border-radius: 8px;
    border: 1px solid var(--gray-003);
  }

  .data_entry_journal___container .forms_journal .form_item .form_header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding-bottom: 10px;
    border-bottom: 1px solid var(--gray-003);
  }

  .data_entry_journal___container .forms_journal .form_header h3 {
    margin: 0;
  }

  .data_entry_journal___container .company_details {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: start;
    gap: 100px;
    padding: 30px 0;
  }

  .data_entry_journal___container .company_details .company_name,
  .data_entry_journal___container .company_details .company_contacts {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .data_entry_journal___container .company_details .company_name .title,
  .data_entry_journal___container .company_details .company_contacts .title {
    font-size: var(--font-lg);
    border-bottom: 2px solid var(--blue-005);
    padding-bottom: 8px;
  }

  .data_entry_journal___container .company_details .company_name h5 {
    font-weight: 600;
    font-size: 16px;
    line-height: 24px;
    color: var(--blue-005);
    margin: 0;
  }

  .data_entry_journal___container .company_details .company_contacts .contacts {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .data_entry_journal___container .company_details .company_contacts .contacts p {
    margin: 0;
  }

  .data_entry_journal___container .forms_journal .actions {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    gap: 10px;
  }

  .data_entry_journal___container .forms_journal .actions .btn {
    padding: 8px 13px;
    border-radius: 9999px;
    border: none;
    color: #FFF;
    cursor: pointer;
    font-weight: 600;
    background-color: var(--blue-006);
  }

  .data_entry_journal___container .forms_journal .actions .btn:hover {
    background-color: var(--blue-005);
  }

  .data_entry_journal___container .forms_journal .forms_table {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals thead tr th {
    font-weight: normal;
    text-align: left;
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals input,
  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals textarea,
  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals select {
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid var(--gray-003);
    width: 100%;
    resize: vertical;
    background-color: #FFF;
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals input:focus,
  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals textarea:focus,
  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals select:focus {
    border-color: var(--sky-005);
    outline: 1px solid var(--sky-005);
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals .btn_remove {
    padding: 12px;
    background-color: transparent;
    border: none;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 8px;
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals .btn_remove:hover {
    background-color: var(--red-transparent);
    color: var(--red-005);
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals .btn_remove:disabled {
    cursor: not-allowed;
    opacity: 0.5;
  }

  .data_entry_journal___container .actions {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    gap: 10px;
  }

  .data_entry_journal___container .actions .btn {
    padding: 8px 20px;
    border-radius: 9999px;
    border: none;
    cursor: pointer;
    font-weight: 600;
  }

  .data_entry_journal___container .actions .btn.submit {
    color: #FFF;
    background-color: var(--green-006);
  }

  .data_entry_journal___container .actions .btn.submit:hover {
    background-color: var(--green-005);
  }

  .data_entry_journal___container .actions .btn.submit:disabled {
    cursor: not-allowed;
    background-color: var(--gray-003);
    color: var(--gray-007);
  }
</style>