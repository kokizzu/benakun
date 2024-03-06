<script>
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import RiDesignDragMoveLine from 'svelte-icons-pack/ri/RiDesignDragMoveLine';
  import RiSystemDeleteBinLine from "svelte-icons-pack/ri/RiSystemDeleteBinLine";
  import { onMount } from 'svelte';

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
</script>

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
                    <button class="btn" title="Add child">
                      <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
                    </button>
                    <button class="btn" title="Edit">
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
                        <h6>{cc.level}.{idx+2}&nbsp;&nbsp;{cc.name}</h6>
                        <div class="options">
                          <button class="btn" title="Add child">
                            <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
                          </button>
                          <button class="btn" title="Edit">
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