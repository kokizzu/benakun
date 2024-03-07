<script>
// @ts-nocheck

  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import RiDesignDragMoveLine from 'svelte-icons-pack/ri/RiDesignDragMoveLine';
  import RiSystemDeleteBinLine from 'svelte-icons-pack/ri/RiSystemDeleteBinLine';
  import FiLoader from 'svelte-icons-pack/fi/FiLoader';
  import IoClose from 'svelte-icons-pack/io/IoClose';
  import { onMount } from 'svelte';
  import InputBox from './_components/InputBox.svelte';
  import { TenantAdminCreateCoaChild } from './jsApi.GEN.js';
  import { notifier } from './_components/notifier.js';

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

  function reformatCoas() {
    if (coas && coas.length) {
      for (let i in coas) {
        coas[i].children = [];
        if (Number(coas[i].parentId) > 0) {
          for (let j in REFORMAT_COAS) {
            if (REFORMAT_COAS[j].id === coas[i].parentId)
            REFORMAT_COAS[j].children = [...REFORMAT_COAS[j].children, coas[i]];
          }
        } else {
          REFORMAT_COAS = [...REFORMAT_COAS, coas[i]];
        }
      }
      REFORMAT_COAS.sort((a, b) => a.level - b.level);
    }
  }

  onMount(() => {
    reformatCoas();
    console.log('REFORMAT_COAS:', REFORMAT_COAS)
  });

  let parent;

  function dragStart(e) {
    setTimeout(() => e.target.classList.add('dragging'), 0);
  }

  function dragEnd(e) {
    e.target.classList.remove('dragging');
  }

  function dragOver(e) {
    e.preventDefault();
    const draggingItem = document.querySelector(".dragging");
    let siblings = [...parent.querySelectorAll(".child")];
    let nextSibling = siblings.find((sibling) => {
      return e.clientY <= sibling.offsetTop + sibling.offsetHeight / 2;
    });

    parent.insertBefore(draggingItem, nextSibling);
  }

  let parentIdToAddOrEdit = 0, childName = '';
  let isSubmitAddOrEditChild = false;

  let showAddEditChild_popUp = false;
  let addOrEditChild_state = '';
  const toggleAddEditChild_popUp = (parentId, name) => {
    addOrEditChild_state = name ? 'Edit ' + name : 'Add new child';
    childName = name;
    parentIdToAddOrEdit = parentId;
    showAddEditChild_popUp = !showAddEditChild_popUp;
  }
  async function submitAddOrEditChild() {
    isSubmitAddOrEditChild = true;
    if (!childName || !parentIdToAddOrEdit) {
      isSubmitAddOrEditChild = false;
      notifier.showWarning('All fields are required');
      return;
    }
    await TenantAdminCreateCoaChild(
      {
        name: childName,
        parentId: parentIdToAddOrEdit
      },
      function (o) {
        if (o.error) {
          isSubmitAddOrEditChild = false;
          notifier.showError(o.error);
          console.log(o.error);
          return;
        }
        isSubmitAddOrEditChild = false;
        coas = o.coa;
        reformatCoas();
        console.log(o);
        showAddEditChild_popUp = false;
        notifier.showSuccess('Company created successfully');
      }
    );
  }
</script>

{#if showAddEditChild_popUp}
  <div class="popup_container">
    <div class="popup">
      <header class="header">
        <h2>{addOrEditChild_state}</h2>
        <button on:click={() => showAddEditChild_popUp = false}>
          <Icon size="22" color="var(--red-005)" src={IoClose}/>
        </button>
      </header>
      <div class="forms">
        <InputBox id="childName" label="Nama" bind:value={childName} type="text" placeholder="Barang..." />
      </div>
      <div class="foot">
        <div class="left">
        </div>
        <div class="right">
          <button class="cancel" on:click={toggleAddEditChild_popUp}>Cancel</button>
          <button class="ok" on:click={submitAddOrEditChild}>
            {#if !isSubmitAddOrEditChild}
              <span>Ok</span>
            {/if}
            {#if isSubmitAddOrEditChild}
              <Icon className="spin" color="#FFF" size="14" src={FiLoader} />
            {/if}
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

<div class="root_layout">
  <div class="root_container">
    <SideMenu access={segments} />
    <div class="root_content">
      <Navbar {user} />
      <div class="content">
        <div class="coa_levels shadow">
          {#if REFORMAT_COAS && REFORMAT_COAS.length}
            {#each REFORMAT_COAS as c, _ (c.id)}
              <div class="coa">
                <div class="parent">
                  <h5>{c.level}.&nbsp;{c.name}</h5>
                  <div class="options">
                    <button class="btn" title="Add child" on:click={
                      ()=>toggleAddEditChild_popUp(c.id, ``)
                    }>
                      <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
                    </button>
                    <button class="btn" title="Edit" on:click={
                      ()=>toggleAddEditChild_popUp(c.id, c.name)
                    }>
                      <Icon color="var(--gray-006)" className="icon" size="17" src={RiDesignPencilLine} />
                    </button>
                  </div>
                </div>
                {#if c.children && c.children.length}
                  <div class="children" bind:this={parent}>
                    {#each c.children as cc, idx (cc.id)}
                      <!-- svelte-ignore a11y-no-static-element-interactions -->
                      <div
                        class="child"
                        draggable={true}
                        on:dragstart={dragStart}
                        on:dragend={dragEnd}
                        on:dragover={dragOver}
                        >
                        <Icon color="var(--gray-006)" className="icon_drag" size="17" src={RiDesignDragMoveLine} />
                        <h6>{cc.level}.{idx+1}&nbsp;&nbsp;{cc.name}</h6>
                        <div class="options">
                          <button class="btn" title="Add child" on:click={toggleAddEditChild_popUp}>
                            <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
                          </button>
                          <button class="btn" title="Edit" on:click={toggleAddEditChild_popUp}>
                            <Icon color="var(--gray-006)" className="icon" size="17" src={RiDesignPencilLine} />
                          </button>
                          <button class="btn" title="Delete">
                            <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemDeleteBinLine} />
                          </button>
                        </div>
                      </div>
                    {/each}
                  </div>
                {/if}
              </div>
            {/each}
          {/if}
        </div>
      </div>
      <Footer />
    </div>
  </div>
</div>

<style>
  .popup_container {
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 2000;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		display: flex;
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

	.popup_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 500px;
		display: flex;
		flex-direction: column;
	}

  .popup_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 10px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup_container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup_container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup_container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.popup_container .popup .foot {
		display: flex;
		flex-direction: row;
    justify-content: space-between;
		gap: 10px;
		align-items: center;
		padding: 10px 20px;
		border-top: 1px solid var(--gray-004);
	}

  .popup_container .popup .foot .right {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }

	.popup_container .popup .foot button {
		padding: 8px 13px;
		border-radius: 9999px;
		border: none;
		color: #FFF;
		cursor: pointer;
		font-weight: 600;
	}

  .popup_container .popup .foot button.reset {
		background-color: var(--amber-006);
		border: 1px solid var(--amber-006);
	}

  .popup_container .popup .foot button.reset:hover {
    background-color: var(--amber-005);
  }

	.popup_container .popup .foot button.ok {
		background-color: var(--green-006);
		border: 1px solid var(--green-006);
    display: flex;
    justify-content: center;
    align-items: center;
	}

	.popup_container .popup .foot button.ok:hover {
		background-color: var(--green-005);
	}

	.popup_container .popup .foot button.cancel {
		background-color: #fbbf2420;
		color: var(--amber-005);
		border: 1px solid var(--amber-005);
	}

  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }


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
  }

  .coa_levels .coa .children .child {
    width: 100%;
    display: flex;
    flex-direction: row;
    /* justify-content: space-between; */
    gap: 20px;
    align-items: center;
    padding: 5px 5px 5px 40px;
    cursor: pointer;
    position: relative;
  }

  .coa_levels .coa .children .child h6 {
    font-size: 14px;
    line-height: 2.2rem;
    font-weight: 500;
    margin: 0;
  }

  .coa_levels .coa .children .child:hover {
    background-color: var(--gray-001);
  }

  .coa_levels .coa .children .child:hover .options {
    display: flex;
  }

  :global(.coa_levels .coa .children .child .icon_drag) {
    display: none;
    position: absolute;
    left: 10px;
  }

  :global(.coa_levels .coa .children .child:hover .icon_drag) {
    display: block;
  }

  :global(.coa_levels .coa .children .child:active .icon_drag) {
    display: block;
  }
</style>