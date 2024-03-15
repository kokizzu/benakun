<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import RiDesignDragMoveLine from 'svelte-icons-pack/ri/RiDesignDragMoveLine';
  import RiSystemDeleteBinLine from 'svelte-icons-pack/ri/RiSystemDeleteBinLine';
  import { onMount } from 'svelte';
  import PopUpCoaChild from './PopUpCoaChild.svelte';
  import { notifier } from './notifier.js';
  import { TenantAdminCreateCoaChild, TenantAdminUpdateCoaChild } from '../jsApi.GEN';
  import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

  /**
     * @typedef {Object} CoA
     * @property {string} id
     * @property {string} name
     * @property {number} level
     * @property {string} parentId
     * @property {Array<CoA>} children
    */
  /**
   * @type {CoA}
   */
  export let coa = {
    id: '',
    name: '',
    level: 0,
    parentId: '',
    children: []
  };
  export let rootNum = 1;
  export let num = 1;
  export let indent = 1;
  let indentWidth = '10px';
  const toIndentWidth = (/** @type {number} */ i) => { return `${i * 10 + 10}px` }

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
    switch (coaState) {
      case 'edit':
        await TenantAdminUpdateCoaChild(
          {
            name: coaName,
            parentId: Number(coa.parentId),
            id: Number(coa.id)
          },
          // @ts-ignore
          function (o) {
            // @ts-ignore
            if (o.error) {
              isSubmitted = false;
              // @ts-ignore
              notifier.showError(o.error);
              // @ts-ignore
              console.log(o.error);
              return;
            }
            isSubmitted = false;
            popUpCoaChild.hide();
            notifier.showSuccess(coaName + ' edited');
            // @ts-ignore
            dispatch('update', { coas: o.coa })
          }
        );
        break;
      case 'add':
        await TenantAdminCreateCoaChild(
          {
            name: coaName,
            parentId: Number(coa.id)
          },
          // @ts-ignore
          function (o) {
            // @ts-ignore
            if (o.error) {
              isSubmitted = false;
              // @ts-ignore
              notifier.showError(o.error);
              // @ts-ignore
              console.log(o.error);
              return;
            }
            isSubmitted = false;
            popUpCoaChild.hide();
            notifier.showSuccess('Coa child created');
            // @ts-ignore
            dispatch('update', { coas: o.coa })
          }
        );
        break;
    }
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
    <div class="left">
      <h6>{coa.level}.{rootNum}{indent > 1 ? `.${num}` : ``}&nbsp;&nbsp;{coa.name}</h6>
      <div class="options">
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
        <button class="btn" title="Delete">
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemDeleteBinLine}
          />
        </button>
      </div>
    </div>
    <Icon
      color="var(--gray-006)"
      className="icon_drag"
      size="17"
      src={RiDesignDragMoveLine}
    />
  </div>
  {#if coa.children && coa.children.length}
    {#each coa.children as ch, idx (ch.id)}
      <svelte:self
        coa={ch}
        rootNum={rootNum}
        num={idx+1}
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
    cursor: pointer;
    position: relative;
    border-radius: 8px;
    user-select: none;
  }

  .coa-child:hover {
    background-color: var(--gray-001);
  }

  .coa-child .left {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 30px;
  }

  .coa-child .left h6 {
    font-size: 14px;
    line-height: 2.2rem;
    font-weight: 500;
    margin: 0;
  }

  .coa-child:hover .left .options {
    display: flex;
  }

  .coa-child .left .options {
    display: none;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .coa-child .left .options .btn {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background-color: transparent;
    border-radius: 5px;
    padding: 5px;
    cursor: pointer;
  }

  .coa-child .left .options .btn:hover {
    background-color: var(--gray-002);
  }

  .coa-child .left .options .btn:active {
    background-color: var(--gray-003);
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

