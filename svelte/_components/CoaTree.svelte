<script>
  /** @typedef {import('./types/coa').CoA} CoA */

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiSystemAddBoxLine, RiDesignPencilLine,
    RiSystemDeleteBinLine, RiArrowsArrowGoBackLine
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { onMount } from 'svelte';
  import PopUpCoA from './PopUpCoa.svelte';
  import { notifier } from './notifier.js';
  import { TenantAdminCoa } from '../jsApi.GEN';
  import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

  export let coa = /** @type {CoA} */ ({
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
  });

  export let isRootCoA  /** @type {boolean} */  = false;
  export let num        /** @type {string} */   = '1';
  export let parentNum  /** @type {number} */   = 0;
  export let indent     /** @type {number} */   = 1;

  let   indentWidth   = '10px';
  const toIndentWidth = (/** @type {number} */ i) => { return `${i * 15 + 10}px` }

  onMount(() => indentWidth = toIndentWidth(indent))

  let popUpCoa = /** @type {PopUpCoA} */ ({});
  let id          = 0;
  let parentId    = 0;
  let name        = '';
  let label       = '';
  let isSubmitted = false;
  let heading     = 'Add CoA child'

  async function submitUpsertCoa() {
    isSubmitted = true;
    const coa = /** @type CoA */ ({ id, parentId, name, label });
    const i = {
      cmd: 'upsert', coa
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

  /**
   * @param {number} ids
   * @param {number} parentIds
   * @param {string} names
   * @param {string} labels
   * @returns {void}
   */
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

  let isDragOver = false;
  let isDragging = false;

  const updateOnMoving = () => {
    dispatch('moving', { coa: coa });
    isDragging = true;
  }

  const updateMovedOrg = () => {
    dispatch('moved', { coa: coa });
    isDragOver = false;
  };
</script>

<PopUpCoA
  bind:heading
  bind:this={popUpCoa}
  bind:name
  bind:label
  bind:isSubmitted
  onSubmit={submitUpsertCoa}
/>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="{isDragOver ? 'drag-over' : ''} {isDragging ? 'dragging' : ''}"
  draggable={!isRootCoA || coa.editable}
  on:dragstart={() => {
    if (!isRootCoA) updateOnMoving();
  }}
  on:drop|preventDefault={updateMovedOrg}
  on:dragover|preventDefault={() => (isDragOver = true)}
  on:dragleave|preventDefault={() => (isDragOver = false)}
  on:dragend={() => (isDragging = false)}
  >
  <div class="coa-child" style="--indent-width:{indentWidth};">
    <div class="num-title">
      <span class="h-line"></span>
      <h6 class={`text ${coa.deletedAt > 0 ? 'deleted' : ''}`}>
        {#if isRootCoA}
          <span class="label">{parentNum}</span>&nbsp;&nbsp;
        {:else}
          <span class="label">{num}</span>&nbsp;&nbsp;
        {/if}
        <p class="name">{coa.name}</p>
      </h6>
    </div>
    <div class="options">
      {#if coa.label !== ''}
        <span class="label">{coa.label}</span>
      {/if}
      {#if coa.tenantLabel !== ''}
        <span class="label tenant">{coa.tenantLabel}</span>
      {/if}
      {#if coa.deletedAt <= 0}
        <button class="btn" title="Add child" on:click={() => {
          toggleShowPopUpCoa(0, coa.id, '', '');
        }}>
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemAddBoxLine}
          />
        </button>
        <button class="btn" title="Edit" on:click={() => {
          toggleShowPopUpCoa(coa.id, coa.parentId, coa.name, coa.label);
        }}>
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
</div>

{#if coa.children && coa.children.length}
  {#each coa.children as ch, idx (ch.id)}
    <svelte:self
      coa={ch}
      num={`${num}.${idx+1}`}
      parentNum={parentNum}
      indent={indent + 1}
      draggable="true"
      on:update
      on:info
      on:moving
      on:moved
      on:dragover
      on:dragleave
      on:dragend
    />
  {/each}
{/if}

<style>
  .drag-over {
    background-color: var(--gray-002);
  }

  .dragging {
    background-color: #FFF;
  }
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
    font-weight: 500;
    margin: 0;
    display: flex;
    flex-direction: row;
    gap: 0;
    align-items: center;
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
    flex-shrink: 0;
  }

  .coa-child .num-title .text .name{
    line-height: 1em;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
    text-overflow: ellipsis;
    margin: 0;
  }

  .coa-child .num-title .text.deleted .name{
    text-decoration: line-through;
    color: var(--red-005);
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

  .coa-child .options .label.tenant {
    padding: 2px 10px;
    margin-right: 7px;
    border-radius: 999px;
    background-color: var(--yellow-transparent);
    color: var(--yellow-002);
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

