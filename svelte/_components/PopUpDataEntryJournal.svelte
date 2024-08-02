<script>
	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputCustom from './InputCustom.svelte';

  export let heading      = 'Add journal';
  export let isSubmitted  = false;
  let isShow              = false;

  export let isDebit = true;
  export let isSales = false;

  export let debitIDR = 0;
  export let creditIDR = 0;
  export let description = '';
  export let date = '';
  export let salesCount = 0;
  export let salesPriceIDR = 0;

  export let OnSubmit = async function() {}  
  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

  export const Reset = () => {

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
      {#if isDebit}
        <InputCustom
          id="debitIDR"
          label="Debit (IDR)"
          placeholder="0"
          bind:value={debitIDR}
          type="number"
        />
      {/if}
      {#if !isDebit}
        <InputCustom
          id="creditIDR"
          label="Credit (IDR)"
          placeholder="0"
          bind:value={creditIDR}
          type="number"
        />
      {/if}
      {#if isSales}
        <InputCustom
          id="salesCount"
          label="Sales Count"
          placeholder="0"
          bind:value={salesCount}
          type="number"
        />
        <InputCustom
          id="salesPriceIDR"
          label="Sales Price (IDR)"
          placeholder="0"
          bind:value={salesPriceIDR}
          type="number"
        />
      {/if}
      <InputCustom
        id="description"
        label="Description"
        placeholder="Description"
        bind:value={description}
        type="textarea"
      />
      <InputCustom
        id="date"
        label="Date"
        bind:value={date}
        type="date"
      />
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
        <button class="ok" on:click|preventDefault={() => OnSubmit()} disabled={isSubmitted}>
          {#if !isSubmitted}
            <span>Ok</span>
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
		width: 500px;
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
		gap: 10px;
	}

  .popup_container .popup .forms .map_container {
    width: 100%;
    height: 200px;
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