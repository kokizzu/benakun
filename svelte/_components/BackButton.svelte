<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { ImArrowLeft2 } from '../node_modules/svelte-icons-pack/dist/im';
</script>

<button class="back_btn" on:click|preventDefault>
  <Icon color="#495057" size={18} src={ImArrowLeft2} />
</button>

<style>
  .back_btn {
    background-color: #FFF;
    padding: 8px;
    border-radius: 9999px;
    border: none;
    display: flex;
    justify-content: center;
    cursor: pointer;
  }

  .back_btn:hover {
    background-color: var(--gray-001);
  }

  .back_btn:active {
    background-color: #7474ec30;
  }

  :global(.back_btn:hover svg path) {
    fill: var(--purple-002);;
  }

  :global(.back_btn svg path) {
    fill: var(--gray-007);
  }
</style>