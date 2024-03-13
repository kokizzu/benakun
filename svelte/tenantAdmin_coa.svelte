<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import { onMount } from 'svelte';
  import { TenantAdminCreateCoaChild } from './jsApi.GEN.js';
  import { notifier } from './_components/notifier.js';
  import CoaTree from './_components/CoaTree.svelte';
  import PopUpCoaChild from './_components/PopUpCoaChild.svelte';
  import MainLayout from './_layouts/mainLayout.svelte';

  /**
   * @type {any}
   */
  let segments = {/* segments */};
  let user = {/* user */};
  let coas = [/* coas */];

  /**
     * @typedef {Object} CoA
     * @property {string} id
     * @property {string} name
     * @property {number} level
     * @property {string} parentId
     * @property {Array<CoA>} children
    */
  /**
   * @type {Array<CoA>}
   */
  let REFORMAT_COAS = [];

  function coaMaker(id) {
    let coaFormatted = {
      id: '',
      name: '',
      level: 0,
      parentId: '',
      children: []
    }
    for (let i in coas) {
      if (coas[i].id == String(id)) {
        const children = coas[i].children;
        if (children && children.length) {
          for (let j in children) {
            const childId = children[j]
            const child = coaMaker(childId)
            coaFormatted.children = [...coaFormatted.children, child];
          }
        } else {
          coaFormatted.children = [];
        }

        coaFormatted.name = coas[i].name;
        coaFormatted.level = coas[i].level;
        coaFormatted.id = coas[i].id;
        coaFormatted.parentId = coas[i].parentId;
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

  onMount(() => {REFORMAT_COAS = reformatCoas(); console.log(REFORMAT_COAS)});

  let popUpCoaChild;

  let childName = '', childParentId = '';
  let isSubmitted = false;

  async function submitAddCoaChild() {
    isSubmitted = true;
    if (childName === '') {
      isSubmitted = false;
      notifier.showWarning('coa name cannot be empty');
      return;
    }
    await TenantAdminCreateCoaChild(
      {
        name: childName,
        parentId: Number(childParentId)
      },
      // @ts-ignore
      function (o) {
        // @ts-ignore
        if (o.error) {
          popUpCoaChild.hide();
          isSubmitted = false;
          // @ts-ignore
          notifier.showError(o.error);
          // @ts-ignore
          console.log(o.error);
          return;
        }
        popUpCoaChild.hide();
        isSubmitted = false;
        // @ts-ignore
        coas = o.coa;

        REFORMAT_COAS = [];
        REFORMAT_COAS = reformatCoas();
        notifier.showSuccess('Coa child created');
      }
    );
    childName = '';
    childParentId = '';
  }

  function updateEventHandler(e) {
    coas = e.detail.coas;
    
    REFORMAT_COAS = [];
    REFORMAT_COAS = reformatCoas();
  }
</script>

<PopUpCoaChild
  bind:this={popUpCoaChild}
  bind:isSubmitted={isSubmitted}
  bind:childName={childName}
  onSubmit={submitAddCoaChild}
/>

<MainLayout>
  <div class="coa_levels shadow">
    {#if REFORMAT_COAS && REFORMAT_COAS.length}
      {#each REFORMAT_COAS as c, _ (c.id)}
        <div class="coa">
          <div class="parent">
            <h5>{c.level}.&nbsp;{c.name}</h5>
            <div class="options">
              <button class="btn" title="Add child" on:click={() => {
                childParentId = c.id;
                popUpCoaChild.show();
              }}>
                <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
              </button>
            </div>
          </div>
          {#if c.children && c.children.length}
            <div class="children">
              {#each c.children as cc, idx (cc.id)}
                <CoaTree
                  coa={cc}
                  rootNum={idx+1}
                  indent={1}
                  on:update={updateEventHandler}
                />
              {/each}
            </div>
          {/if}
        </div>
      {/each}
    {/if}
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
    /* justify-content: space-between; */
    gap: 20px;
    align-items: center;
    font-size: 22px;
    font-weight: 700;
    color: var(--blue-006);
    padding: 10px 10px 10px 15px;
  }

  .coa_levels .coa .parent h5 {
    font-size: 18px;
    line-height: 2.2rem;
    margin: 0;
  }

  .coa_levels .coa .parent:hover .options {
    display: flex;
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
    gap: 5px;
    overflow: hidden;
    padding: 0 20px 0 10px;
  }
</style>