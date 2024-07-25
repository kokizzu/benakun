<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import { onMount } from 'svelte';
  import { TenantAdminCoa } from './jsApi.GEN.js';
  import { notifier } from './_components/notifier.js';
  import CoaTree from './_components/CoaTree.svelte';
  import PopUpCoA from './_components/PopUpCoa.svelte';
  import MainLayout from './_layouts/mainLayout.svelte';

  /** @typedef {import('./_components/types/user').User} User */
  /** @typedef {import('./_components/types/access').Access} Access */
  /** @typedef {import('./_components/types/coa').CoA} CoA */

  let segments  = /** @type Access */ ({/* segments */});
  let user      = /** @type User */   ({/* user */});
  let coas      = /** @type CoA[] */  ([/* coas */]);
  
  let REFORMAT_COAS = /** @type CoA[] */ ([]);

  function coaMaker(/** @type any */ id) {
    /** @type CoA */
    let coaFormatted = {
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
    }
    for (let i in coas) {
      if (coas[i].id == Number(id)) { // @ts-ignore
        const children = /** @type number[] */ (coas[i].children);
        if (children && children.length) {
          for (let j in children) {
            const childId = children[j]
            const child = coaMaker(childId); // @ts-ignore
            coaFormatted.children = [...coaFormatted.children, child];
          }
        }
        coaFormatted.id = coas[i].id
        coaFormatted.tenantCode = coas[i].tenantCode
        coaFormatted.label = coas[i].label
        coaFormatted.name = coas[i].name
        coaFormatted.parentId = coas[i].parentId
        coaFormatted.createdAt = coas[i].createdAt
        coaFormatted.createdBy = coas[i].createdBy
        coaFormatted.updatedAt = coas[i].updatedAt
        coaFormatted.updatedBy = coas[i].updatedBy
        coaFormatted.deletedAt = coas[i].deletedAt
      }
    }

    return coaFormatted;
  }

  function reformatCoas() {
    let toCoas = [];

    if (coas && coas.length) {
      for (let i in coas) {
        const id = coas[i].id;
        const coa = coaMaker(id);
        toCoas = [...toCoas, coa];
      }
    }

    if (toCoas && toCoas.length) {
      toCoas = toCoas.filter(obj => obj.parentId <= 0);
      toCoas.sort((a, b) => a.level - b.level);
    }

    return toCoas;
  }

  onMount(() => REFORMAT_COAS = reformatCoas());

  function updateEventHandler(/** @type any */ e) {
    coas = e.detail.coas;
    
    REFORMAT_COAS = [];
    REFORMAT_COAS = reformatCoas();
  }

  let popUpCoa;
  let id = 0, parentId = 0;
  let name = '', label = '';
  let isSubmittedCoA = false;

  async function submitUpsertCoa() {
    isSubmittedCoA = true;
    /** @type CoA */ //@ts-ignore
    const coa = { id, parentId, name, label }
    const i = {
      cmd: 'upsert',
      coa
    }
    await TenantAdminCoa( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminCoaCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmittedCoA = false;
        id = 0, parentId = 0, name = '', label = '';
        popUpCoa.hide();
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        coas = o.coas;
        REFORMAT_COAS = [];
        REFORMAT_COAS = reformatCoas();

        notifier.showSuccess('coa updated');
      }
    )
  }

  function toggleShowPopUpCoa(idz, parentIdz, name, label) {
    id = idz; parentId = parentIdz;
    name = name; label = label;
    popUpCoa.show();
  }

  /** @type {CoA|null} */
  let coaMoving = null;

  const coaMovingHandler = (e) => coaMoving = e.detail.coa;
  
  const coaMovedHandler = async (e) => {
    /** @type {CoA} */
    const parentCoa = e.detail.coa;
    if (!coaMoving) return;
    if (parentCoa.id == coaMoving.parentId) return;

    const i = {
      cmd: 'move',
      coa: {
        id: Number(coaMoving.id),
        parentId: Number(coaMoving.parentId)
      },
      moveToIdx: 0,
      toParentId: Number(parentCoa.id)
    }
    await TenantAdminCoa( // @ts-ignore
      i, /** @returns {Promise<void>}*/
      function (/** @type {any} */o) {
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        coas = o.coas;
        REFORMAT_COAS = [];
        REFORMAT_COAS = reformatCoas();

        notifier.showSuccess('coa moved');
      }
    );
  }
</script>

<PopUpCoA
  bind:this={popUpCoa}
  bind:name
  bind:label
  bind:isSubmitted={isSubmittedCoA}
  onSubmit={submitUpsertCoa}
/>

<MainLayout>
  <div class="coa_levels shadow">
    <div class="header_options">
      <button
        class="add"
        on:click={() => {toggleShowPopUpCoa(0, 0, '', '')}}
      >Add</button>
    </div>
    {#each (REFORMAT_COAS || []) as c, idxParent (c.id)}
      <div class="coa">
        <div class="parent">
          <h5>
            <span class="label">{idxParent+1}</span>
            &nbsp;{c.name}
          </h5>
          <div class="options">
            <button class="btn" title="Add CoA child" on:click={() => {toggleShowPopUpCoa(0, c.id, c.name, c.label)}}>
              <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
            </button>
          </div>
        </div>
        {#if c.children && c.children.length}
          <div class="children">
            {#each c.children as cc, idx (cc.id)}
              <CoaTree
                coa={cc}
                num={`${idx+1}`}
                parentNum={idxParent+1}
                indent={1}
                on:update={updateEventHandler}
                on:moving={coaMovingHandler}
                on:moved={coaMovedHandler}
              />
            {/each}
          </div>
        {/if}
      </div>
    {/each}
  </div>
</MainLayout>

<style>
  .coa_levels {
    display: grid;
    grid-template-columns: 1fr;
    width: 100%;
    background-color: #FFF;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    overflow: hidden;
    padding: 15px 15px 20px;
  }

  .coa_levels .header_options {
    display: flex;
    flex-direction: row-reverse;
    padding-bottom: 10px;
    border-bottom: 1px solid var(--gray-003);
  }

  .coa_levels .header_options .add {
    border: none;
    padding: 10px 20px;
    border-radius: 8px;
    background-color: var(--blue-006);
    color: #FFF;
    font-size: var(--font-lg);
    cursor: pointer;
  }

  .coa_levels .header_options .add:hover {
    background-color: var(--blue-005);
  }

  .coa_levels .coa {
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }

  .coa_levels .coa .parent {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    gap: 20px;
    align-items: center;
    font-size: 22px;
    font-weight: 700;
    color: var(--blue-006);
    padding: 10px 0 10px 15px;
  }

  .coa_levels .coa .parent h5 {
    font-size: 18px;
    line-height: 2.2rem;
    margin: 0;
    user-select: none;
  }

  .coa_levels .coa .parent h5 .label {
    font-size: 14px;
    padding: 3px 7px;
    background-color: var(--blue-006);
    color: #FFF;
    border-radius: 9999px;
  }

  .coa_levels .coa .parent .options {
    display: flex;
    padding-right: 10px;
  }

  .coa_levels .coa .options {
    display: none;
    flex-direction: row;
    align-items: center;
  }

  .coa_levels .coa .options .btn {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background-color: transparent;
    border-radius: 5px;
    padding: 5px;
    cursor: pointer;
  }

  .coa_levels .coa .options .btn:hover {
    background-color: var(--gray-002);
  }

  .coa_levels .coa .options .btn:active {
    background-color: var(--gray-003);
  }

  .coa_levels .coa .children {
    display: flex;
    flex-direction: column;
    overflow: hidden;
    padding: 0 0 0 10px;
  }
</style>