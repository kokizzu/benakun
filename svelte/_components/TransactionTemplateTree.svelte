<script>
  /** @typedef {import('./types/transaction').TransactionTemplate} TransactionTemplate */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FaCircleDot } from '../node_modules/svelte-icons-pack/dist/fa';
  import { RiArrowsArrowRightSLine } from '../node_modules/svelte-icons-pack/dist/ri';

  export let transactionTemplate = /** @type {TransactionTemplate} */ ({});

  let isShowDetail = false;
  const toggleShowDetail = () => {
    isShowDetail = !isShowDetail;
  }
</script>

<div class="transaction_template">
  <div class="info">
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
  {#if isShowDetail}
    <div>Show detail</div>
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
    align-items: center;
    gap: 10px;
  }

  .transaction_template .info .btn_toggle {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    padding: 5px;
    background-color: transparent;
    border-radius: 999px;
    cursor: pointer;
  }

  .transaction_template .info .btn_toggle:hover {
    background-color: var(--gray-001);
  }

  .transaction_template .info .btn_toggle:active {
    background-color: var(--gray-002);
  }
  
  .transaction_template .info .name {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }
</style>