<script>
  // @ts-nocheck
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import { onMount } from 'svelte';
  import { Title } from 'chart.js';

  let segments = {
    /* segments */
  };
  let user = {
    /* user */
  };
  let orgs = [
    {
      ParentId: 2,
      Type: 2,
      Children: [3],
    },
    {
      ParentId: 3,
      Type: 1,
      Children: [4],
    },
    {
      ParentId: 4,
      Type: 1,
      Children: [],
    },
    {
      ParentId: 5,
      Type: 1,
      Children: [],
    },
  ];

  let elementIsDrag;

  function renderTopLevelOrganization() {
    // empty content element
    document.getElementById('content').innerHTML = '';
    const topLevelOrg = orgs.filter(org => {
      let parentId = org.ParentId;
      wasAChild = true;
      for (let org of orgs) {
        wasAChild = org.Children.includes(parentId);
        if (wasAChild) break;
      }
      return !wasAChild;
    });

    for (let org of topLevelOrg) {
      // get the childs
      let parent = getOrginazationItemWithChilds(org.ParentId);
      document.getElementById('content').appendChild(parent);
    }
  }

  /**
   * @param {number} ParentId
   * @return {string}
   */
  function getOrginazationItemWithChilds(ParentId) {
    let parent = orgs.find(e => e.ParentId === ParentId),
      parentBoxTag = document.createElement('div'),
      titleTag = document.createElement('h4');

    titleTag.innerText = '- Parent ' + ParentId;

    // title attribute setting
    titleTag.className = 'title';

    // style title parent element
    titleTag.style.margin = '0';

    // parent element attibrute setting
    parentBoxTag.draggable = true;
    parentBoxTag.id = ParentId;
    parentBoxTag.ondragstart = dragStart;
    parentBoxTag.ondrop = drop;
    parentBoxTag.ondragover = dragOver;
    parentBoxTag.ondragleave = dragLeave;

    // style parent element
    parentBoxTag.style.border = '1px black solid';
    parentBoxTag.style.padding = '8px';
    parentBoxTag.style.borderRadius = '8px';
    parentBoxTag.style.cursor = 'grab';
    parentBoxTag.style.marginBottom = '16px';
    parentBoxTag.style.backgroundColor = 'white';

    parentBoxTag.append(titleTag);

    for (let child of parent.Children) {
      let childrenBoxTag = getOrginazationItemWithChilds(child);

      // style child element
      childrenBoxTag.style.margin = '16px';
      childrenBoxTag.style.marginLeft = '32px';

      parentBoxTag.appendChild(childrenBoxTag);
    }
    return parentBoxTag;
  }

  /**
   *
   * @param {Object} event - event data from drag start
   */
  function dragStart(event) {
    const sourceElement = event.srcElement;

    // change the current element was draged
    elementIsDrag = sourceElement;
  }

  /**
   *
   * @param {Object} event - event data from drag start
   */
  function dragLeave(event) {
    const sourceElement = event.srcElement;

    // if element on the same element
    if (sourceElement === elementIsDrag) return;

    // if element on the title
    if (sourceElement.className.split(' ').find(e => e === 'title')) return;

    sourceElement.style.backgroundColor = 'white';
  }

  /**
   * @param {Object} event - event data from drop
   */
  function drop(event) {
    const dragedParentId = parseInt(elementIsDrag.getAttribute('id')),
      dragedEntry = orgs.find(e => e.ParentId === dragedParentId),
      oldParentElement = elementIsDrag.parentNode,
      oldParentId = parseInt(oldParentElement?.getAttribute('id')),
      oldParentEntry = oldParentId && oldParentId !== 'content' ? orgs.find(e => e.ParentId === oldParentId) : null,
      newParentElement = event.srcElement,
      newParentId = parseInt(newParentElement.getAttribute('id')),
      newParentEntry = orgs.find(e => e.ParentId === newParentId);

    // if orginization append into its child
    if (isChild(dragedParentId, newParentId)) {
      if (oldParentEntry) oldParentEntry.Children = [...oldParentEntry.Children, ...dragedEntry.Children];
      dragedEntry.Children = [];
      console.log(orgs);
    }

    // remove organization from its old parent
    if (oldParentEntry) oldParentEntry.Children = oldParentEntry.Children.filter(e => e != dragedParentId);

    // append into new parrent childs
    const newParentChilds = newParentEntry.Children;
    if (!newParentChilds.includes(dragedParentId)) newParentChilds.push(dragedParentId);

    //rerender orginazation structure
    renderTopLevelOrganization();
  }

  /**
   * Check new parrent is family of current organization
   *
   * @param {number} orgParentId
   * @param {number} newParentId
   * @return {boolean}
   *
   */
  function isChild(orgParentId, newParentId) {
    const org = orgs.find(e => e.ParentId === orgParentId);
    for (let childId of org.Children) {
      if (childId === newParentId || isChild(childId, newParentId)) return true;
    }
    return false;
  }

  /**
   *
   * @param {Object} event - event data from drag over
   */
  function dragOver(event) {
    const sourceElement = event.srcElement;

    // if element on the same element
    if (sourceElement === elementIsDrag) return;

    // if element on the title
    if (sourceElement.className.split(' ').find(e => e === 'title')) return;

    sourceElement.style.backgroundColor = 'red';
    return false;
  }

  /**
   * Drop and make item as top level org
   * @param {object} event - eventdata from drag end
   */
  function dragEnd(event) {
    const dragedParentId = parseInt(elementIsDrag.getAttribute('id')),
      oldParent = elementIsDrag.parentNode,
      oldParentId = oldParent?.getAttribute('id');

    if (oldParentId && oldParentId !== 'content') {
      let oldParentEntry = orgs.find(e => e.ParentId == oldParentId);

      // remove child
      console.log(oldParentEntry.Children);
      oldParentEntry.Children = oldParentEntry.Children.filter(e => e != dragedParentId);
      console.log(oldParentEntry.Children);
    }

    renderTopLevelOrganization();
  }

  onMount(async () => {
    renderTopLevelOrganization();
  });
</script>

<div class="root_layout">
  <div class="root_container">
    <SideMenu access={segments} />
    <div class="root_content">
      <Navbar {user} />
      <div class="content" id="content" on:dragover={dragOver} on:dragleave={dragLeave} on:dragend={dragEnd}></div>
      <Footer />
    </div>
  </div>
</div>

<style>
</style>
