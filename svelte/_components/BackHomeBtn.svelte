<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FaSolidArrowRight } from '../node_modules/svelte-icons-pack/dist/fa';

  function backHome() {
    window.location.href = '/';
  }
</script>

<button class="back_home" on:click={backHome}>
  <span>Back to home</span>
  <Icon className="arrow_backhome" color="#FFF" size={18} src={FaSolidArrowRight} />
</button>

<style>
  .back_home {
    position: relative;
    width: fit-content;
    padding: 15px 50px 15px 30px;
    background-color: #3b82f6;
    border: none;
    border-radius: 9999px;
    font-size: 18px;
    cursor: pointer;
  }
  .back_home:hover {
    background-color: #2c74ea;
  }
  .back_home span {
    color: #ffffff;
  }
  :global(.arrow_backhome) {
    position: absolute;
    right: 22px;
    top: auto;
    bottom: auto;
    margin-top: 2px;
    transition-property: all;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    transition-duration: 150ms;
  }
  :global(.back_home:hover .arrow_backhome) {
    right: 16px;
  }
</style>
