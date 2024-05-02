<script>
  /** Input type:
    * text (default)
    * number
    * phone
    * date
    * bool
    * select
    * email
    */

  export let type = 'text';
  export let id;
  export let value;

  /** @type {Array<any>|Object} */
  export let values = [];
  export let label;
  export let placeholder = '';
  export let isObject = false;

  if (isObject) value = value+'';
</script>

<div>
  <div class={type == 'bool' ? 'input_box bool' : 'input_box'}>
    {#if type == 'bool'}
      <label class="label" for={id}>{label}</label>
      <label class="switcher" for={id}>
        <input type="checkbox" id={id} bind:checked={value}>
        <span class="slider"></span>
      </label>
    {:else if type == 'select'}
      {#if isObject}
        <label class="label" for={id}>{label}</label>
        <select name={id} id={id} bind:value={value} {placeholder}>
          <option value="" disabled>-- {placeholder} --</option>
          {#each Object.entries(values) as [k, v]}
            <option value={k} selected={value}>{v}</option>
          {/each}
        </select>
      {:else}
        <label class="label" for={id}>{label}</label>
        <select name={id} id={id} bind:value={value} {placeholder}>
          <option value="" disabled>-- {placeholder} --</option>
          {#each values as v}
            <option value={v}>{v}</option>
          {/each}
        </select>
      {/if}
    {:else if type == 'number'}
      <label class="label" for={id}>{label}</label>
      <input type="number" bind:value={value} {id} {placeholder}/>
    {:else if type == 'float'}
      <label class="label" for={id}>{label}</label>
      <input type="number" bind:value={value} {id} {placeholder}/>
    {:else if type == 'textarea'}
      <label class="label" for={id}>{label}</label>
      <textarea bind:value={value} {id} {placeholder}></textarea>
    {:else if type == 'text'}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder}/>
    {:else if type == 'email'}
      <label class="label" for={id}>{label}</label>
      <input type="email" bind:value={value} {id} {placeholder}/>
    {:else if type == 'date'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder}/>
    {:else}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder}/>
    {/if}
  </div>
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

  .input_box.bool {
    width: fit-content;
  }

  .input_box .label {
    font-size: var(--font-base);
    margin-left: 10px;
  }

  .input_box input,
  .input_box textarea {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 10px 12px;
  }

  .input_box input:focus,
  .input_box textarea:focus {
    border-color: var(--sky-005);
    outline: 1px solid var(--sky-005);
  }

  .input_box textarea {
    resize: vertical;
    height: 90px;
    min-height: 50px;
    max-height: 300px;
  }

  .input_box select {
    cursor: pointer;
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 10px 12px;
  }

  .input_box select:focus {
    border-color: var(--sky-005);
    outline: 1px solid var(--sky-005);
  }
  
  /* The switch - the box around the slider */
  .switcher {
    position: relative;
    display: inline-block;
    width: 50px;
    height: 24px;
    margin-left: 0 !important;
  }
  
  .switcher input {
    opacity: 0 !important;
    width: 0 !important;
    height: 0 !important;
  }
  
  .switcher .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: var(--gray-005);
    -webkit-transition: .2s;
    transition: .2s;
  }

  .switcher .slider:before {
    position: absolute;
    content: "";
    height: 16px;
    width: 16px;
    left: 4px;
    bottom: 4px;
    background-color: white;
    -webkit-transition: .2s;
    transition: .2s;
  }

  .switcher input:checked + .slider {
    background-color: var(--sky-006) !important;
  }

  .switcher input:focus + .slider {
    box-shadow: 0 0 1px var(--sky-006) !important;
  }

  .switcher input:checked + .slider:before {
    -webkit-transform: translateX(26px);
    -ms-transform: translateX(26px);
    transform: translateX(26px);
  }
  
  .switcher .slider {
    border-radius: 34px;
  }

  .switcher .slider:before {
    border-radius: 50%;
  }
</style>
