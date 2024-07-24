<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiSystemAddBoxLine,
    RiDesignPencilLine,
    RiSystemDeleteBinLine,
    RiArrowsArrowGoBackLine
   } from '../node_modules/svelte-icons-pack/dist/ri';
  import { onMount } from 'svelte';
  import PopUpCoA from './PopUpCoa.svelte';
  import { notifier } from './notifier.js';
  import { TenantAdminCoa } from '../jsApi.GEN';
  import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
  
  /** @typedef {import('./types/coa').CoA} CoA */

  /** @type {CoA} */
  export let coa = {
    id: 0,
    name: '',
    parentId: 0,
    children: [],
    createdAt: 0,
    createdBy: '',
    updatedAt: 0,
    updatedBy: '',
    deletedAt: 0,
    tenantCode: '',
    label: '',
    deletedBy: '',
    restoredBy: ''
  };

  export let num = '1';
  export let parentNum = 0;
  export let indent = 1;
  let indentWidth = '10px';
  const toIndentWidth = (/** @type {number} */ i) => { return `${i * 15 + 10}px` }

  onMount(() => indentWidth = toIndentWidth(indent))

  let popUpCoa;
  let id = 0, parentId = 0;
  let name = '', label = '';
  let isSubmitted = false;
  let heading = 'Add CoA child'

  async function submitUpsertCoa() {
    isSubmitted = true;
    /** @type CoA */ //@ts-ignore
    const coa = { id, parentId, name, label }
    const i = {
      cmd: 'upsert',
      coa
    }
    await TenantAdminCoa( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminCoaCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitted = false;
        popUpCoa.hide();
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        dispatch('update', { coas: o.coas })

        notifier.showSuccess('coa updated');
      }
    )
  }

  function toggleShowPopUpCoa(ids, parentIds, names, labels) {
    if (ids != 0) { heading = 'Edit CoA' }
    id = ids; parentId = parentIds;
    name = names; label = labels;
    popUpCoa.show();
  }

  async function deleteCoaChild() {
    isSubmitted = true;
    const i = {
      cmd: 'delete',
      coa: { id }
    }
    await TenantAdminCoa( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminCoaCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitted = false;
        popUpCoa.hide();
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        dispatch('update', { coas: o.coas })

        notifier.showSuccess('coa deleted');
      }
    )
  }

  async function restoreCoaChild() {
    isSubmitted = true;
    const i = {
      cmd: 'restore',
      coa: { id }
    }
    await TenantAdminCoa( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminCoaCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitted = false;
        popUpCoa.hide();
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        dispatch('update', { coas: o.coas })

        notifier.showSuccess('coa restored');
      }
    )
  }
</script>

<PopUpCoA
  bind:heading
  bind:this={popUpCoa}
  bind:name
  bind:label
  bind:isSubmitted
  onSubmit={submitUpsertCoa}
/>

<div>
  <div class="coa-child" style="--indent-width:{indentWidth};">
    <div class="num-title">
      <span class="h-line"></span>
      <h6 class={`text ${coa.deletedAt > 0 ? 'deleted' : ''}`}>
        <span class="label">{parentNum}.{num}</span>&nbsp;&nbsp;
        <span class="name">{coa.name}</span>
      </h6>
    </div>
    <div class="options">
      {#if coa.label !== ''}
        <span class="label">{coa.label}</span>
      {/if}
      {#if coa.deletedAt <= 0}
        <button class="btn" title="Add child" on:click={() => toggleShowPopUpCoa(0, coa.id, '', '')}>
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemAddBoxLine}
          />
        </button>
        <button class="btn" title="Edit" on:click={() => toggleShowPopUpCoa(coa.id, coa.parentId, coa.name, coa.label)}>
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiDesignPencilLine}
          />
        </button>
        <button class="btn" title="Delete" on:click={deleteCoaChild}>
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemDeleteBinLine}
          />
        </button>
      {:else}
        <button class="btn" title="Rollback" on:click={restoreCoaChild}>
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
  {#if coa.children && coa.children.length}
    {#each coa.children as ch, idx (ch.id)}
      <svelte:self
        coa={ch}
        num={`${num}.${idx+1}`}
        indent={indent + 1}
        on:update
      />
    {/each}
  {/if}
</div>

<style>
  .coa-child {
    width: auto;
    display: flex;
    flex-direction: row;
    gap: 20px;
    justify-content: space-between;
    align-items: center;
    padding: 5px 10px;
    margin-left: var(--indent-width);
    cursor: move;
    position: relative;
    border-radius: 8px;
    user-select: none;
  }

  .coa-child:hover {
    background-color: var(--gray-001);
  }

  .coa-child .num-title {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 30px;
    position: relative;
  }

  .coa-child .num-title .h-line {
    position: absolute;
    left: -14px;
    width: 1px;
    height: 41px;
    background-color: var(--gray-003);
  }

  .coa-child .num-title .text {
    font-size: 14px;
    line-height: 2.2rem;
    font-weight: 500;
    margin: 0;
  }

  .coa-child .num-title .text .label {
    color: var(--blue-006);
    font-weight: 600;
    font-size: 10px;
    background-color: #0077B620;
    padding: 3px 6px;
    border-radius: 9999px;
    text-decoration: none;
    width: fit-content;
    height: fit-content;
    line-height: 2.2rem;
  }

  .coa-child .num-title .text.deleted .name{
    color: var(--red-005);
    text-decoration: line-through;
    line-height: 2.2rem;
  }

  .coa-child .options {
    display: flex;
    cursor: pointer;
    flex-direction: row;
    align-items: center;
  }

  .coa-child .options .label {
    padding: 2px 10px;
    margin-right: 7px;
    border-radius: 999px;
    background-color: var(--violet-transparent);
    color: var(--violet-005);
  }

  .coa-child .options .btn {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background-color: transparent;
    border-radius: 5px;
    padding: 5px;
    cursor: pointer;
  }

  .coa-child .options .btn:hover {
    background-color: var(--violet-transparent);
  }

  :global(.coa-child .options .btn:hover svg) {
		fill: var(--violet-005);
	}

  :global(.coa-child .icon_drag) {
    display: none;
  }

  :global(.coa-child:hover .icon_drag) {
    display: block;
  }

  :global(.coa-child:active .icon_drag) {
    display: block;
  }
</style>

