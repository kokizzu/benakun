<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminOrganization } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  let segments = {/* segments */};
  let user = {/* user */};
  let orgs = [/* orgs */];

  let lastSent = {};
  let elementIsDrag;

  function renderTopLevelOrganization() {
    // empty content element
    document.getElementById('content').innerHTML = '';
    const topLevelOrg = orgs.filter(org => org.parentId === '0');

    for (let org of topLevelOrg) {
      // get the childs
      let parent = getOrginazationItemWithChilds(org.id);
      document.getElementById('content').appendChild(parent);
    }
  }

  /**
   * @param {string} id
   * @return {string}
   */
  function getOrginazationItemWithChilds(id) {
    let orgEntry = orgs.find(e => e.id === id),
      parentBoxTag = document.createElement('div'),
      titleTag = document.createElement('h4');

    titleTag.innerText = `- ${orgEntry.name}`;

    // title attribute setting
    titleTag.className = 'title';

    // style title parent element
    titleTag.style.margin = '0';

    // parent element attibrute setting
    parentBoxTag.draggable = true;
    parentBoxTag.id = id;
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

    for (let childId of orgEntry.children) {
      let childrenBoxTag = getOrginazationItemWithChilds(childId.toString());

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
    let dragedId = parseInt(elementIsDrag.getAttribute('id')),
      dragedEntry = orgs.find(e => e.id == dragedId),
      oldParentId = parseInt(dragedEntry.parentId),
      newParentElement = event.srcElement,
      newParentId = parseInt(newParentElement.getAttribute('id'));

    if (isNaN(newParentId)) newParentId = 0;

    if (lastSent.dragedId == dragedId && lastSent.oldParentId == oldParentId && lastSent.newParentId == newParentId) {
      console.warn('update has been applied', lastSent);
      renderTopLevelOrganization();
      return;
    }

    saveChanges(dragedId, newParentId, oldParentId);
    lastSent = {
      dragedId,
      newParentId,
      oldParentId,
    };

    // make new parent back to white background
    newParentElement.style.backgroundColor = 'white';
    
  }

  /**
   * Save new change after droping element
   * @param {string} id
   * @param {string} newParentId
   * @param {string} oldParentId
   */
  async function saveChanges(id, newParentId, oldParentId) {
    await TenantAdminOrganization(
      {
        cmd: 'upsert',
        id,
        newParentId,
        oldParentId,
      },
      res => {
        if (res.error) return notifier.showError(res.error);
        orgs = res.Orgs;
        renderTopLevelOrganization();
      }
    );
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
    console.log(event);

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
    console.log(segments, user, orgs);
    renderTopLevelOrganization();

    // asing ondragover for content
    document.getElementById('content').ondragover = dragOver;
  });
</script>

<MainLayout>
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div id="content" on:dragleave={dragLeave} on:drop={drop}>
    <!-- Content here -->
  </div>
</MainLayout>

<style>
  #content {
    padding: 12px;
    border: 1px black solid;
    border-radius: 8px;
  }
</style>