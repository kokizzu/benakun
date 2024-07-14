<script>
  /** @typedef {import('./types/transaction').TransactionTemplate} TransactionTemplate */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FaCircleDot } from '../node_modules/svelte-icons-pack/dist/fa';
  import {
    RiArrowsArrowRightSLine,
    RiDesignPencilLine,
    RiSystemDeleteBinLine,
    RiArrowsArrowGoBackLine
  } from '../node_modules/svelte-icons-pack/dist/ri';

  export let transactionTemplate = /** @type {TransactionTemplate} */ ({});

  let isShowDetail = false;
  const toggleShowDetail = () => {
    isShowDetail = !isShowDetail;
  }
</script>

<div class="transaction_template">
  <div class="info">
    <div class="left">
      <button class="btn_toggle" on:click={toggleShowDetail}>
        <Icon
          color="var(--gray-006)"
          className="icon {isShowDetail ? 'rotate' : 'dropdown'}"
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
      {#if transactionTemplate.deletedAt < 0}
        <button class="btn" title="Edit organization">
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiDesignPencilLine}
          />
        </button>
        <button class="btn" title="Delete organization">
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
  {#if isShowDetail}
    <div class="transaction_template_detail">
      <div class="info">
        <span class="label">Debit</span>
        <span class="label">Kredit</span>
        <span class="label">COA</span>
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
    gap: 5px;
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

  .transaction_template .transaction_template_detail {
    width: auto;
    display: flex;
    flex-direction: column;
    margin-left: 35px;
    padding: 10px;
    border-radius: 8px;
    background-color: var(--gray-001);
  }

  .transaction_template .transaction_template_detail .info {
    display: flex;
    flex-direction: row;
    gap: 10px;
    align-items: center;
  }

  .transaction_template .transaction_template_detail .info span.label {
    padding: 2px 10px;
    margin-right: 7px;
    border-radius: 999px;
    background-color: var(--violet-transparent);
    color: var(--violet-005);
  }
</style>