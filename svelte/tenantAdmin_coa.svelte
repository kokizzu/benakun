<script>
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import RiDesignDragMoveLine from 'svelte-icons-pack/ri/RiDesignDragMoveLine';

  /**
   * @type {any}
   */
  let segments = {/* segments */};
  let user = {/* user */};

  /**
     * @typedef {Object} CoA
     * @property {string} id
     * @property {string} name
     * @property {number} level
     * @property {string} parent_id
     * @property {Array<CoA>} children
    */
  /**
   * @type {Array<CoA>}
   */
  let coa = [
    {
      id: 'xxxxxxx',
      name: 'Aktiva',
      level: 1,
      children: [
        {
          id: 'xxxxxxx',
          name: 'Aktiva tetap',
          children: [],
          parent_id: '1',
          level: 1.2
        },
        {
          id: 'xxxxxxx',
          name: 'Aktiva lancar',
          children: [],
          parent_id: '1',
          level: 1.3
        },
        {
          id: 'xxxxxxx',
          name: 'Aktiva tak berwujud',
          children: [],
          parent_id: '1',
          level: 1.4
        },
      ],
      parent_id: null,
    }
  ];

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
        <div class="coa_levels">
          {#if coa && coa.length}
            {#each coa as c, _ (c.id)}
              <div class="coa shadow">
                <div class="parent">
                  <span>{c.level}. {c.name}</span>
                  <div class="options">
                    <button class="btn">
                      <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
                    </button>
                    <button class="btn">
                      <Icon color="var(--gray-006)" className="icon" size="17" src={RiDesignPencilLine} />
                    </button>
                  </div>
                </div>
                {#if c.children && c.children.length}
                  <div class="children" bind:this={parent}>
                    {#each c.children as cc, _ (cc.id)}
                      <!-- svelte-ignore a11y-no-static-element-interactions -->
                      <div
                        class="child"
                        draggable={true}
                        on:dragstart={dragStart}
                        on:dragend={dragEnd}
                        on:dragover={dragOver}
                        >
                        <Icon color="var(--gray-006)" className="icon_drag" size="17" src={RiDesignDragMoveLine} />
                        <span>{cc.level}. {cc.name}</span>
                        <div class="options">
                          <button class="btn">
                            <Icon color="var(--gray-006)" className="icon" size="17" src={RiSystemAddBoxLine}/>
                          </button>
                          <button class="btn">
                            <Icon color="var(--gray-006)" className="icon" size="17" src={RiDesignPencilLine} />
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
    grid-template-columns: 1fr 1fr;
    gap: 10px;
    width: 100%;
  }

  .coa_levels .coa {
    background-color: #FFF;
    display: flex;
    flex-direction: column;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    overflow: hidden;
  }

  .coa_levels .coa .parent {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    font-size: 22px;
    font-weight: 700;
    color: var(--blue-006);
    padding: 10px 10px 10px 15px;
    border-bottom: 1px solid var(--gray-003);
  }

  .coa_levels .coa .options {
    display: flex;
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
    font-size: 16px;
    overflow: hidden;
  }

  .coa_levels .coa .children .child {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 10px 10px 10px 40px;
    cursor: pointer;
    position: relative;
  }

  .coa_levels .coa .children .child:hover {
    background-color: var(--gray-001);
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