<script>
  /** @typedef {import('./types/transaction').TransactionTemplate} TransactionTemplate */
  /** @typedef {import('./types/transaction').TransactionTplDetail} TransactionTplDetail */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FaCircleDot } from '../node_modules/svelte-icons-pack/dist/fa';
  import {
    RiArrowsArrowRightSLine, RiDesignPencilLine,
    RiSystemDeleteBinLine, RiArrowsArrowGoBackLine,
    RiSystemAddBoxLine, RiDesignBallPenLine, RiSystemDeleteBin5Line
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { TrOutlineRefresh } from '../node_modules/svelte-icons-pack/dist/tr';
  import { TenantAdminTransactionTplDetail } from '../jsApi.GEN';
  import { notifier } from './notifier';
  import PopUpTransactionTplDetail from './PopUpTransactionTplDetail.svelte';
  import { onMount } from 'svelte';

  export let transactionTemplate = /** @type {TransactionTemplate} */ ({});
  let transactionTplDetails = /** @type {TransactionTplDetail[]} */ ([]);
  export let coas = {};

  let isPopUpFormsReady = false;
  let popUpTransactionTplDetail = null;
  onMount(() => {
    isPopUpFormsReady = true;
  })

  let isSearching = false;
  let isShowDetails = false;

  async function getTransactionTplDetails() {
    isSearching = true;
    await TenantAdminTransactionTplDetail( //@ts-ignore
      {
        cmd: 'list', //@ts-ignore
        transactionTplDetail: {
          parentId: transactionTemplate.id
        }
      },
      /** @type {import('../jsApi.GEN').TenantAdminTransactionTemplateCallback} */
      function(/** @type any */ o) {
        isSearching = false;
        if (o.error) {
          notifier.showError(o.error);
          return
        }

        transactionTplDetails = o.transactionTplDetails;
        console.log('transactionTplDetails', o);

        return
      }
    ) 
  }

  const toggleShowDetail = async () => {
    if (isShowDetails) {
      isShowDetails = false;
      return
    } else {
      isShowDetails = true;
      if (!transactionTplDetails || transactionTplDetails.length === 0) {
        await getTransactionTplDetails();
      }
    }
  }

  let trxTplDetailId = 0;
  let coaId = '';
  let isDebit = 'debit';
  let isSubmitUpsertTrxTplDetail = false;
  async function SubmitUpsertTrxTplDetail() {
    isSubmitUpsertTrxTplDetail = true;
    await TenantAdminTransactionTplDetail( //@ts-ignore
      {
        cmd: 'upsert', //@ts-ignore
        transactionTplDetail: {
          id: trxTplDetailId,
          parentId: transactionTemplate.id,
          coaId: Number(coaId),
          isDebit: isDebit == 'debit' ? true : false,
        }
      },
      /** @type {import('../jsApi.GEN').TenantAdminTransactionTemplateCallback} */
      function(/** @type any */ o) {
        isSubmitUpsertTrxTplDetail = false;
        if (o.error) {
          notifier.showError(o.error);
          return
        }

        notifier.showSuccess('Transaction Template Detail updated');

        transactionTplDetails = o.transactionTplDetails;
        console.log('transactionTplDetails', o);
        popUpTransactionTplDetail.hide();

        return
      }
    )
  }
</script>

{#if isPopUpFormsReady}
  <PopUpTransactionTplDetail
    bind:this={popUpTransactionTplDetail}
    bind:coaId
    bind:isDebit
    onSubmit={SubmitUpsertTrxTplDetail}
    bind:isSubmitted={isSubmitUpsertTrxTplDetail}
    {coas}
  />
{/if}

<div class="transaction_template">
  <div class="info">
    <div class="left">
      <button class="btn_toggle" on:click={toggleShowDetail}>
        <Icon
          color="var(--gray-006)"
          className="icon {isShowDetails ? 'rotate' : 'dropdown'}"
          size="20"
          src={RiArrowsArrowRightSLine}
        />
      </button>
      <div class="name">
        <Icon
          className="icon"
          size="17"
          src={FaCircleDot}
          color={transactionTemplate.color}
        />
        <p class="title">{transactionTemplate.name}</p>
      </div>
    </div>
    <div class="options">
      <button
        class="btn"
        title="Refresh"
        on:click={getTransactionTplDetails}
        disabled={isSearching}>
        <Icon
          className="icon {isSearching ? 'spin' : ''}"
          size="17"
          color="var(--gray-006)"
          src={TrOutlineRefresh}
        />
      </button>
      {#if transactionTemplate.deletedAt < 0}
        <button class="btn" title="Add template" on:click={() => {
          trxTplDetailId = 0;
          popUpTransactionTplDetail.show()
        }}>
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemAddBoxLine}
          />
        </button>
        <button class="btn" title="Edit template">
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiDesignPencilLine}
          />
        </button>
        <button class="btn" title="Delete template">
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemDeleteBinLine}
          />
        </button>
      {:else}
        <button class="btn" title="Rollback">
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiArrowsArrowGoBackLine}
          />
        </button>
      {/if}
    </div>
  </div>
  {#if isShowDetails}
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
                        class="btn edit"
                        title="Edit"
                        on:click={() => {
                          isSubmitUpsertTrxTplDetail = false;
                          trxTplDetailId = trxTplDetail.id;
                          isDebit = trxTplDetail.isDebit ? 'debit' : 'credit';
                          coaId = trxTplDetail.coaId+'';
                          popUpTransactionTplDetail.show()
                        }}
                      >
                        <Icon
                          size="15"
                          color="var(--gray-007)"
                          src={RiDesignBallPenLine}
                        />
                      </button>
                      <button
                        class="btn delete"
                        title="delete">
                        <Icon
                          size="15"
                          color="var(--gray-007)"
                          src={RiSystemDeleteBin5Line}
                        />
                      </button>
                    </div>
                  </td>
                  <td>{!trxTplDetail.isDebit ? 'Yes' : 'No'}</td>
                  <td>{trxTplDetail.isDebit ? 'Yes' : 'No'}</td>
                  <td>{coas[trxTplDetail.coaId]}</td>
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
  {/if}
</div>

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(-360deg);
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

  .transaction_template {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .transaction_template .info {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    gap: 10px;
  }

  .transaction_template .info .left {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }

  .transaction_template .info .left .btn_toggle {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    padding: 5px;
    background-color: transparent;
    border-radius: 999px;
    cursor: pointer;
  }

  .transaction_template .info .left .btn_toggle:hover {
    background-color: var(--gray-001);
  }

  .transaction_template .info .left .btn_toggle:active {
    background-color: var(--gray-002);
  }
  
  .transaction_template .info .left .name {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 15px;
  }

  .transaction_template .info .left .name .title {
    font-weight: 500;
    font-size: 18px;
    margin: 0;
  }

  .transaction_template .info .options {
    display: flex;
    cursor: pointer;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .transaction_template .info .options .btn {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background-color: transparent;
    border-radius: 5px;
    padding: 5px;
    cursor: pointer;
  }

  .transaction_template .info .options .btn:hover {
    background-color: var(--gray-002);
  }

  .transaction_template .info .options .btn:active {
    background-color: var(--gray-003);
  }

  .transaction_template_detail {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    padding: 0;
    border: 1px solid var(--gray-003);
    overflow: hidden;
    margin-left: 40px;
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
</style>