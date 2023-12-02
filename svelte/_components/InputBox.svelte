<script>
  import { onMount } from 'svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import AiOutlineEye from 'svelte-icons-pack/ai/AiOutlineEye';
  import AiOutlineEyeInvisible from 'svelte-icons-pack/ai/AiOutlineEyeInvisible';

  export let id;
  export let value;
  export let label;
  export let type = 'text';
  export let placeholder = '';
  let isShowPassword = false;
  let inputElm;
  
  onMount(() => inputElm.type = type)
  function toggleShowPassword() {
    isShowPassword = !isShowPassword;
    if (isShowPassword) inputElm.type = 'text';
    else inputElm.type = 'password';
  }
</script>

<div class={type === 'password' ? 'input_box with_password' :  'input_box'}>
  <label for={id}>{label}</label>
  <input bind:value={value} {id} bind:this={inputElm} {placeholder}/>
  {#if type === 'password'}
    <button class="eye" on:click={toggleShowPassword}>
      {#if !isShowPassword}
        <Icon color="#495057" size="20" src={AiOutlineEye}/>
      {/if}
      {#if isShowPassword}
        <Icon color="#495057" size="20" src={AiOutlineEyeInvisible}/>
      {/if}
    </button>
  {/if}
</div>

<style>
  .input_box {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input_box.with_password input{
    padding-right: 40px !important;
  }

  .input_box label {
    font-size: var(--font-base);
    margin-left: 10px;
  }

  .input_box input {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 10px 12px;
  }

  .input_box input:focus {
    border-color: var(--blue-005);
    outline: 1px solid var(--blue-005);
  }

  .input_box .eye {
    position: absolute;
    height: fit-content;
    width: fit-content;
    background-color: transparent;
    padding: 0;
    top: 30px;
    bottom: auto;
    right: 10px;
    border: none;
    cursor: pointer;
  }

  :global(.input_box .eye:hover svg) {
    fill: var(--blue-005);
  }
</style>
