<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiSend, FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';

  export let isSubmitted = false;
  export let isFullWidth = false;
  export let label = 'Submit';
</script>

<button class={isFullWidth ? 'submit_btn full_width' : 'submit_btn'} on:click|preventDefault>
  {#if !isSubmitted}
    <Icon color="#FFF" size="14" src={FiSend} />
  {/if}
  {#if isSubmitted}
    <Icon className="spin" color="#FFF" size="14" src={FiLoader} />
  {/if}
  <span>{label}</span>
</button>

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

  .full_width {
    width: 100% !important;
    justify-content: center;
  }
  .submit_btn {
    border: none;
    width: fit-content;
    padding: 10px 15px;
    border-radius: 8px;
    background-color: var(--blue-006);
    color: #fff;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    cursor: pointer;
  }

  .submit_btn:hover {
    background-color: var(--blue-005);
  }
</style>
