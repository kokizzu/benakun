<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { fade } from 'svelte/transition';

  export let visible = false;
  let isSubmitted = false;
  export let question = 'Are you sure?';
  export let action = function(){};

  async function confirmFunc() {
    isSubmitted = true;
    await action();
    isSubmitted = false;
    visible = false;
  }
</script>

{#if visible}
<div class="backdrop">
  <div class="confirm_popup" transition:fade={{ delay: 70, duration: 175 }}>
    <h2>{question}</h2>
    <div class="actions_btn">
      <button class="no" on:click={() => visible = !visible}>No</button>
      <button class="yes" disabled={isSubmitted} on:click={confirmFunc}>
        {#if !isSubmitted}
          <span>Yes</span>
        {/if}
        {#if isSubmitted}
          <Icon className="spin" color="#FFF" size="16" src={FiLoader} />
        {/if}
      </button>
    </div>
  </div>
</div>
{/if}


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
  .backdrop {
    z-index: 8888;
    position: fixed;
    padding: 0;
    border: none;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: rgba(0 0 0 / 40%);
    display: flex;
    justify-content: center;
  }

  .confirm_popup {
    padding: 20px;
    background-color: #FFF;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    gap: 30px;
    width: 400px;
    height: fit-content;
    max-height: 600px;
    margin-top: 60px;
    border: 1px solid var(--gray-003);
  }

  .confirm_popup h2 {
    font-size: var(--font-xl);
    margin: 0;
    text-align: center;
  }

  .confirm_popup .actions_btn {
    width: 100%;
    display: flex;
    flex-direction: row;
    gap: 10px;
    align-items: stretch;
  }

  .confirm_popup .actions_btn button {
    flex-grow: 1;
    padding: 14px 0;
    border-radius: 9999px;
    border: none;
    font-weight: 600;
    color: var(--gray-007);
    cursor: pointer;
  }

  .confirm_popup .actions_btn button.no {
    background-color: transparent;
  }

  .confirm_popup .actions_btn button.no:hover {
    background-color: #0096c740;
    color: var(--blue-006);
  }

  .confirm_popup .actions_btn button.yes {
    background-color: var(--blue-006);
    color: #FFF;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .confirm_popup .actions_btn button.yes:hover,
  .confirm_popup .actions_btn button.yes:disabled {
    background-color: var(--blue-005);
  }
</style>