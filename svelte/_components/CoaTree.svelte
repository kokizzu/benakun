<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiSystemAddBoxLine,
    RiDesignPencilLine,
    RiArrowsDragMove2Fill,
    RiSystemDeleteBinLine,
    RiArrowsArrowGoBackLine
   } from '../node_modules/svelte-icons-pack/dist/ri';
  import { onMount } from 'svelte';
  import PopUpCoaChild from './PopUpCoaChild.svelte';
  import { notifier } from './notifier.js';
  import {
    TenantAdminUpsertCoaChild,
    TenantAdminDeleteCoaChild,
    TenantAdminRestoreCoaChild
  } from '../jsApi.GEN';
  import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
  
  /** @typedef {import('./types/coa').CoA} CoA */

  /** @type {CoA} */
  export let coa = {
    id: '',
    name: '',
    level: 0,
    parentId: '',
    children: [],
    createdAt: 0,
    createdBy: '',
    updatedAt: 0,
    updatedBy: '',
    deletedAt: 0
  };

  export let num = '1';
  export let indent = 1;
  let indentWidth = '10px';
  const toIndentWidth = (/** @type {number} */ i) => { return `${i * 15 + 10}px` }

  let popUpCoaChild, heading = 'Add coa child';

  onMount(() => indentWidth = toIndentWidth(indent))

  let coaState = 'add', coaName = '';
  let isSubmitted = false;

  const toggleAddOrEdit = (/** @type {string}*/ state) => {
    if (state === 'add') coaState = 'add', coaName = '', heading = 'Add coa child';
    else coaState = 'edit', coaName = coa.name, heading = 'Edit: ' + coa.name;
    popUpCoaChild.show();
  }

  async function submitAddOrEditChild() {
    isSubmitted = true;
    if (coaName === '') {
      isSubmitted = false;
      notifier.showWarning('Coa name cannot be empty');
      return;
    }

    /** @type CoA */ //@ts-ignore
    let coaPayload = {
      id: coaState === 'edit' ? coa.id : '0',
      name: coaName,
      parentId: coa.id,
    }

    await TenantAdminUpsertCoaChild( //@ts-ignore
      { coa: coaPayload}, /** @type {import('../jsApi.GEN').TenantAdminUpsertCoaChildCallback}*/
      function (/** @type {any} */ o) {
        isSubmitted = false;
        popUpCoaChild.hide();

        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }

        notifier.showSuccess('Coa child updated');
        dispatch('update', { coas: o.coas })
      }
    );
  }

  async function deleteCoaChild() {
    isSubmitted = true;
    await TenantAdminDeleteCoaChild(
      { id: Number(coa.id) },
      /** @type {import('../jsApi.GEN').TenantAdminDeleteCoaChildCallback}*/ function (/** @type {any} */ o) {
        if (o.error) {
          isSubmitted = false;
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        isSubmitted = false;
        notifier.showSuccess(coa.name + ' deleted');
        dispatch('update', { coas: o.coa })
      }
    )
  }

  async function restoreCoaChild() {
    isSubmitted = true;
    await TenantAdminRestoreCoaChild(
      { id: Number(coa.id) },
      /** @type {import('../jsApi.GEN').TenantAdminRestoreCoaChildCallback}*/ function (/** @type {any} */ o) {
        if (o.error) {
          isSubmitted = false;
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        isSubmitted = false;
        notifier.showSuccess(coa.name + ' restored');
        dispatch('update', { coas: o.coa })
      }
    )
  }
</script>

<PopUpCoaChild
  bind:this={popUpCoaChild}
  bind:isSubmitted={isSubmitted}
  bind:childName={coaName}
  bind:heading={heading}
  onSubmit={submitAddOrEditChild}
/>

<div>
  <div class="coa-child" style="--indent-width:{indentWidth};">
    <div class="num-title">
      <span class="h-line"></span>
      <h6 class={`text ${coa.deletedAt > 0 ? 'deleted' : ''}`}>
        <span class="label">{coa.level}.{num}</span>&nbsp;&nbsp;
        <span class="name">{coa.name}</span>
      </h6>
    </div>
    <div class="options">
      {#if coa.deletedAt <= 0}
        <button class="btn" title="Add child" on:click={() => toggleAddOrEdit('add')}>
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemAddBoxLine}
          />
        </button>
        <button class="btn" title="Edit" on:click={() => toggleAddOrEdit('edit')}>
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

