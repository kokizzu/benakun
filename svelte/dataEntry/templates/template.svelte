<script>
  /** @typedef {import('../../_components/types/transaction.js').TransactionTemplate} TransactionTemplate */
  /** @typedef {import('../../_components/types/transaction.js').TransactionJournal} TransactionJournal */
  /** @typedef {import('../../_components/types/transaction.js').DetailObjectTransaction} DetailObjectTransaction */
  /** @typedef {import('../../_components/types/transaction.js').TransactionTplDetail} TransactionTplDetail */
  /** @typedef {import('../../_components/types/organization.js').Org} Org */
  /** @typedef {import('../../_components/types/user.js').User} User */
  /**
   * @typedef {Object} PayloadTransactionJournals
   * @property {string|number} creditIDR
   * @property {string|number} debitIDR
   * @property {string} descriptions
   * @property {string|Date} date
   * @property {string|number} coaId
   * @property {number?} salesCount?
   * @property {string|number?} salesPriceIDR?
   * @property {number?} transactionTplId?
   */

  import MainLayout from '../../_layouts/mainLayout.svelte';
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine, RiSystemDeleteBin6Line } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { DataEntryTransactionEntry } from '../../jsApi.GEN';
  import { notifier } from '../../_components/notifier';
  import InputCustom from '../../_components/InputCustom.svelte';

  let transactionTemplate   = /** @type TransactionTemplate */ ({/* transactiontemplate */});
  let transactionTplDetails = /** @type TransactionTplDetail[] */ ([/* transactionTplDetails */]);
  let coas                  = /** @type Record<number|string, string> */ ({/* coas */});
  let org                   = /** @type Org */ ({/* org */});
  let user                  = /** @type User */ ({/* user */});

  let isCreatingJournal = false;
  let isSubmitted       = false;

  let isDebit     = true;
  let isSales     = false;
  let isChildOnly = false;
  let isAutoSum   = false;

  function dateISOFormat(/** @type number */ dayTo = 0) {
    const dt    = new Date();
    const date  = (dayTo > 0 ? dt.getDate() + dayTo : dt.getDate());
    const month = dt.toLocaleDateString('default', {month: '2-digit'});
    const year  = dt.getFullYear();

    return `${year}-${month}-${date}`;
  }

  let startDate     = dateISOFormat(0);
  let endDate       = dateISOFormat(4);
  let coaId         = 0;
  let debitIDR      = 0;
  let creditIDR     = 0;
  let descriptions  = '';
  let date          = dateISOFormat(0);

  // If isChildOnly is true, then we will show coa children
  // when input rows in table
  let coaChildren     = /** @type Record<number|string, string> */ ({});
  let payloadsMulti   = /** @type PayloadTransactionJournals[] */ ([]);

  // If autoSum and not sales, count credit/debit from other template details
  let totalCreditIDR  = 0;
  let totalDebitIDR   = 0;

  function reset() {
    isCreatingJournal = false;
    isDebit       = true;
    isSales       = false;
    isChildOnly   = false;
    isAutoSum     = false;
    startDate     = dateISOFormat(0);
    endDate       = dateISOFormat(4);
    coaId         = 0;
    debitIDR      = 0;
    creditIDR     = 0;
    descriptions  = '';
    date          = dateISOFormat(0);
    coaChildren   = {};
    payloadsMulti = [];

    totalCreditIDR = 0;
    totalDebitIDR  = 0;
  }

  async function GetCoaChildren(/** @type number */ coaId) {
    const i = {
      cmd: 'form',
      coaId: coaId
    }
    await DataEntryTransactionEntry( // @ts-ignore
      i, /** @returns {Promise<any>} */
      function (/** @type {any} */ o) {
        isSubmitted = false;
        if (o.error) {
          notifier.showError(o.error || 'failed to get coa children');
          return
        }
        
        coaChildren = o.coaChildren;
      }
    )
  }

  async function Submit() {
    isSubmitted = true;
    let debit = 0;
    let credit = 0;

    if (isDebit) credit = totalCreditIDR;
    else debit = totalDebitIDR;

    const i = {
      cmd: 'upsert',
      coaId: coaId,
      transactionTplId: transactionTemplate.id,
      transactionJournals: [
        {
          debitIDR: debit+'',
          creditIDR: credit+'',
          descriptions: descriptions,
          date: date
        }
      ],
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

        reset();
        notifier.showSuccess('journal added !');
      }
    )
  }

  async function SubmitMulti() {
    isSubmitted = true;
    let trxJournals = /** @type TransactionJournal[]|any */ ([]);

    for (const payload of payloadsMulti) {
      let detailObj = '';
      let debit = 0;
      let credit = 0;

      if (isSales) {
        const detail = /** @type DetailObjectTransaction */ ({
          salesCount: payload.salesCount,
          salesPriceIDR: payload.salesPriceIDR+'',
        });

        detailObj = JSON.stringify(detail);

        if (isDebit) credit = Number(payload.salesCount) * Number(payload.salesPriceIDR);
        else debit = Number(payload.salesCount) * Number(payload.salesPriceIDR);
      }

      if (isAutoSum) {
        if (isDebit) credit = totalCreditIDR;
        else debit = totalDebitIDR;
      }

      trxJournals.push({
        debitIDR: debit+'',
        creditIDR: credit+'',
        descriptions: payload.descriptions,
        date: payload.date,
        coaId: payload.coaId,
        detailObj: detailObj
      });
    }
    const i = {
      cmd: 'upsert',
      coaId: coaId,
      transactionTplId: transactionTemplate.id,
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
        
        reset();
        notifier.showSuccess('journal added !');
      }
    )
  }
</script>

<MainLayout>
  <div class={!isCreatingJournal ? 'show' : 'hidden'}>
    <header>
      <h1>Data entry for {transactionTemplate.name}</h1>
    </header>
    <div class="data_entry_template___container" >
      <div class="transaction_template_detail">
        <div class="table_container">
          <table>
            <thead>
              <tr>
                <th class="no">#</th>
                <th class="a_row">Action</th>
                <th>Kredit</th>
                <th>Debit</th>
                <th>CoA (Chart of Accounts)</th>
                <th>Auto Sum</th>
                <th>Child Only</th>
                <th>Sales</th>
              </tr>
            </thead>
            <tbody>
              {#if transactionTplDetails && transactionTplDetails.length > 0}
                {#each (transactionTplDetails || []) as trxTplDetail, i (trxTplDetail.id)}
                  <tr>
                    <td>{i + 1}</td>
                    <td class="a_row">
                      <div class="actions">
                        <button
                          on:click={ async () => {
                            isDebit = trxTplDetail.isDebit
                            isSales = (trxTplDetail.attributes || []).includes('sales')
                            isChildOnly = (trxTplDetail.attributes || []).includes('childOnly')
                            coaId = trxTplDetail.coaId;

                            let pyCoaId = trxTplDetail.coaId;
                            if (isChildOnly) {
                              await GetCoaChildren(trxTplDetail.coaId);
                              console.log(coaChildren);
                              for (const [id, v] of Object.entries(coaChildren)) {
                                console.log(id, v);
                                pyCoaId = Number(id);
                                break;
                              }
                            };

                            if (isSales) {
                              payloadsMulti = [
                                {
                                  coaId: Number((isChildOnly ? pyCoaId : coaId)),
                                  descriptions: '',
                                  creditIDR: 0,
                                  debitIDR: 0,
                                  date: dateISOFormat(0),
                                  salesCount: 0,
                                  salesPriceIDR: 0,
                                  transactionTplId: transactionTemplate.id
                                }
                              ]
                              console.log(payloadsMulti);
                            } else if (isChildOnly) {
                              payloadsMulti = [
                                {
                                  coaId: pyCoaId,
                                  descriptions: '',
                                  creditIDR: 0,
                                  debitIDR: 0,
                                  date: dateISOFormat(0),
                                  salesCount: 0,
                                  salesPriceIDR: 0,
                                  transactionTplId: transactionTemplate.id
                                }
                              ]
                            }

                            isCreatingJournal = true;
                          }}
                          class="btn"
                          title="Add journal"
                        >
                          <Icon
                            size="15"
                            color="var(--gray-007)"
                            src={RiSystemAddBoxLine}
                          />
                        </button>
                      </div>
                    </td>
                    <td>{!trxTplDetail.isDebit ? 'Yes' : 'No'}</td>
                    <td>{trxTplDetail.isDebit ? 'Yes' : 'No'}</td>
                    <td>{coas[trxTplDetail.coaId]}</td>
                    <td>{(trxTplDetail.attributes || []).includes('autoSum') ? 'Yes' : 'No'}</td>
                    <td>{(trxTplDetail.attributes || []).includes('childOnly') ? 'Yes' : 'No'}</td>
                    <td>{(trxTplDetail.attributes || []).includes('sales') ? 'Yes' : 'No'}</td>
                  </tr>
                {/each}
              {:else}
                <tr>
                  <td>0</td>
                  <td>no-data</td>
                </tr>
              {/if}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  <div class={isCreatingJournal ? 'show' : 'hidden'}>
    <header>
      <h1>Entry journal for "{coas[coaId]}"</h1>
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
        {#if isSales || isChildOnly}
          <div class="actions">
            <button class="btn" on:click={() => {
              payloadsMulti = [
                ...payloadsMulti,
                {
                  coaId: coaId,
                  descriptions: '',
                  creditIDR: 0,
                  debitIDR: 0,
                  date: dateISOFormat(0),
                  salesCount: 0,
                  salesPriceIDR: 0,
                  transactionTplId: transactionTemplate.id
                }
              ]
            }}>
              Add row
            </button>
          </div>
        {/if}
        <div class="forms_table">
          <table class="table_transaction_journals">
            {#if !isSales && !isChildOnly}
              <thead>
                <tr>
                  {#if !isAutoSum}
                    {#if isDebit}
                      <th>Debit (IDR)</th>
                    {:else}
                      <th>Credit (IDR)</th>
                    {/if}
                  {/if}
                  <th>Description</th>
                  <th>Date</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  {#if !isAutoSum}
                    {#if isDebit}
                      <td>
                        <input type="number" min=0 bind:value={debitIDR}/>
                      </td>
                    {:else}
                      <td>
                        <input type="number" min=0 bind:value={creditIDR}/>
                      </td>
                    {/if}
                  {/if}
                  <td>
                    <textarea
                      placeholder="Description"
                      bind:value={descriptions}
                      rows="1"
                    />
                  </td>
                  <td>
                    <input type="date" bind:value={date}/>
                  </td>
                </tr>
              </tbody>
            {/if}
            {#if isSales || isChildOnly}
              <thead>
                <tr>
                  <th></th>
                  {#if !isAutoSum}
                    {#if isDebit}
                      <th>Debit (IDR)</th>
                    {:else}
                      <th>Credit (IDR)</th>
                    {/if}
                  {/if}
                  <th>Description</th>
                  <th>Date</th>
                  {#if isChildOnly}
                    <th>CoA</th>
                  {/if}
                  {#if isSales}
                    <th>Sales Count</th>
                    <th>Sales Price (IDR)</th>
                  {/if}
                </tr>
              </thead>
              <tbody>
                {#each (payloadsMulti || []) as py, i}
                  <tr>
                    <td>
                      <button
                        disabled={i === 0}
                        title={i === 0 ? 'Cannot remove first row' : 'Remove row'}
                        class="btn_remove"
                        on:click={() => payloadsMulti = payloadsMulti.filter((item) => item !== py)}
                      >
                        <Icon
                          src={RiSystemDeleteBin6Line}
                          size="16"
                        />
                      </button>
                      </td>
                    {#if !isAutoSum}
                      {#if isDebit}
                        <td>
                          <input type="number" min=0 bind:value={py.debitIDR}/>
                        </td>
                      {:else}
                        <td>
                          <input type="number" min=0 bind:value={py.creditIDR}/>
                        </td>
                      {/if}
                    {/if}
                    <td>
                      <textarea
                        placeholder="Description"
                        bind:value={py.descriptions}
                        rows="1"
                      />
                    </td>
                    <td>
                      <input type="date" bind:value={py.date}/>
                    </td>
                    {#if isChildOnly}
                      <td>
                        <select bind:value={py.coaId}>
                          {#each Object.entries(coaChildren) as [k, v], idx}
                            <option value={k}>{v}</option>
                          {/each}
                        </select>
                      </td>
                    {/if}
                    {#if isSales}
                      <td>
                        <input type="number" min=0 bind:value={py.salesCount}/>
                      </td>
                      <td>
                        <input type="number" min=0 bind:value={py.salesPriceIDR}/>
                      </td>
                    {/if}
                  </tr>
                {/each}
              </tbody>
            {/if}
          </table>
        </div>
      </div>
      <div class="actions">
        <button
          class="btn cancel"
          on:click={() => isCreatingJournal = false}
        >
          Cancel
        </button>
        <button
          disabled={isSubmitted}
          class="btn submit"
          on:click={() => {
            if (!isSales && !isChildOnly) Submit();
            else SubmitMulti();
          }}
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

  .hidden {
    display: none;
  }

  header {
    margin-left: 30px;
  }

  .data_entry_template___container {
    display: flex;
  }

  .transaction_template_detail {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    padding: 0;
    border: 1px solid var(--gray-003);
    overflow: hidden;
    width: 100%;
  }

  .transaction_template_detail .table_container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .transaction_template_detail .table_container table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
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

  .transaction_template_detail .table_container table thead tr th.no {
    width: 30px;
  }

  .transaction_template_detail .table_container table thead tr th.a_row {
    max-width: 100px;
    width: 100px;
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

  .transaction_template_detail .table_container table tbody tr td .actions {
    display: flex;
    flex-direction: row;
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

  .data_entry_journal___container .actions .btn.cancel {
    background-color: transparent;
  }

  .data_entry_journal___container .actions .btn.cancel:hover {
    background-color: var(--gray-003);
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