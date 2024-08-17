<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
	import InputCustom from './InputCustom.svelte';

  export let heading = 'Add Transaction Template Detail';

  export let isSubmitted = false;
  export let onSubmit = () => {};

  export let isDebit 	= 'debit';
  const debitOrCredit = {
    'credit': 'Kredit',
    'debit': 'Debit'
  };
  export let coaId = '0';
  export let coas = {};
  export let autoSum = false;
  export let childOnly = false;
  export let sales = false;

  let isShow = false;
  export const show = () => isShow = true;
  export const hide = () => isShow = false;

  const cancel = () => {
    isShow	= false;
    isDebit	= 'debit';
  }

  // TODO: coa nomornya harus ada di combobox
  // combobox bisa di search
</script>

<div class={`popup_container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>{heading}</h2>
      <button on:click={hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputCustom
        bind:value={isDebit}
        id="isDebit"
        label="Debit / Kredit"
        type="select"
        placeholder="Debit / Kredit"
        values={debitOrCredit}
        isObject
			/>
      <InputCustom
        bind:value={coaId}
        id="coa"
        label="CoA (Chart of Accounts)"
        type="select"
        placeholder="CoA (Chart of Accounts)"
        values={coas}
        isObject
        isPlainHTML
			/>
      <InputCustom
        bind:value={autoSum}
        id="autoSum"
        label="Auto Sum ?"
        type="bool"
      />
      <InputCustom
        bind:value={childOnly}
        id="childOnly"
        label="Child Only ?"
        type="bool"
      />
      <InputCustom
        bind:value={sales}
        id="sales"
        label="Sales ?"
        type="bool"
      />
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
        <button class="ok" on:click|preventDefault={onSubmit}>
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

	.popup_container .popup .foot button.cancel {
		background-color: #fbbf2430;
		color: var(--amber-005);
	}

	.popup_container .popup .foot button.cancel:hover {
		background-color: #fbbf2450;
	}
</style>