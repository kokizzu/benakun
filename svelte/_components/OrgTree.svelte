<script>
  /** @typedef {import('./types/organization.js').Org} Org */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiBuildingsCommunityLine,
    RiSystemAddBoxLine,
    RiDesignPencilLine,
    RiSystemDeleteBinLine,
    RiSystemInformationLine,
    RiBuildingsBuilding2Line,
    RiUserFacesGroup3Line,
    RiBusinessBriefcaseLine,
    RiArrowsArrowGoBackLine
   } from '../node_modules/svelte-icons-pack/dist/ri';
  import { onMount, createEventDispatcher } from 'svelte';
  import PopUpOrg from './PopUpOrg.svelte';
  import { TenantAdminOrganization } from '../jsApi.GEN';
  import { notifier } from './xNotifier';

  const dispatch = createEventDispatcher();
  
  export let org = /** @type Org */ ({});
  console.log('Org:', org);

  const OrgTypeStrCompany   = 'company';
  const OrgTypeStrDept      = 'department';
  const OrgTypeStrDivision  = 'division';
  const OrgTypeStrJob       = 'job';
  const OrgTypeCompany      = 1;
  const OrgTypeDept         = 2;
  const OrgTypeDivision     = 3;
  const OrgTypeJob          = 4;

  let orgTypeStr  = OrgTypeStrCompany;
  let orgIcon     = RiBuildingsCommunityLine;

  switch (org.orgType) {
    case OrgTypeCompany: {
      orgTypeStr = OrgTypeStrCompany;
      orgIcon = RiBuildingsCommunityLine;
      break;
    }
    case OrgTypeDept: {
      orgTypeStr = OrgTypeStrDept;
      orgIcon = RiBuildingsBuilding2Line;
      break;
    }
    case OrgTypeDivision: {
      orgTypeStr = OrgTypeStrDivision;
      orgIcon = RiUserFacesGroup3Line;
      break;
    }
    case OrgTypeJob: {
      orgTypeStr = OrgTypeStrJob;
      orgIcon = RiBusinessBriefcaseLine;
      break;
    }
  }

  const getOrgTypeStr = (type = 0) => {
    switch (type) {
      case OrgTypeCompany: { return OrgTypeStrCompany }
      case OrgTypeDept: { return OrgTypeStrDept }
      case OrgTypeDivision: { return OrgTypeStrDivision }
      case OrgTypeJob: { return OrgTypeStrJob }
      default: { return '' }
    }
  }

  export let indent = 0;
  let indentWidth = '10px';
  const toIndentWidth = (/** @type {number} */ i) => {
    return `${i * 15 + 15}px`;
  }

  onMount(() => indentWidth = toIndentWidth(indent));

  let popUpOrg;
  let isSubmitted = false;
  let popUpHeading = 'Add organization child';

  // Payload
  let orgId         = 0;
  let orgParentId   = 0;
  let orgName       = '';
  let orgHeadTitle  =  '';
  let isEditOrg     = false;

  const toggleAddOrEdit = (isEdit) => {
    if (isEdit) {
      isEditOrg = true;
      orgId = org.id;
      orgName = org.name;
      orgHeadTitle = org.headTitle;
      popUpHeading = 'Edit '+getOrgTypeStr(org.orgType)+': '+org.name;
    } else {
      isEditOrg = false;
      orgName = '';
      orgHeadTitle = '';
      orgParentId = org.id;
      popUpHeading = 'Add organization child';
    }
    popUpOrg.show();
  }

  async function SubmitUpsertOrg() {
    isSubmitted = true;

    /** @type Org */ //@ts-ignore
    let orgPayload = {
      id: orgId,
      parentId: isEditOrg ? org.parentId : org.id,
      name: orgName,
      headTitle: orgHeadTitle,
    }

    const i = {
      cmd: 'upsert',
      org: orgPayload
    }
    
    await TenantAdminOrganization( //@ts-ignore
      i, /** @returns {Promise<void>}*/
      function (/** @type {any} */ o) {
        isSubmitted = false;
        popUpOrg.hide();
        if (o.error) {
          notifier.showError(o.error);
          console.log(o.error);
          return;
        }

        const orgs = /** @type Org[] */ (o.orgs);
        dispatch('update', { orgs: orgs });

        notifier.showSuccess('organization child updated');
      }
    );
  }

  let isDragOver = false;
  let isDragging = false;

  const updateEventInfo = () => dispatch('info', { org: org });

  const updateOnMoving = () => {
    dispatch('moving', { org: org });
    isDragging = true;
  }

  const updateMovedOrg = () => {
    dispatch('moved', { org: org });
    isDragOver = false;
  };
</script>

<PopUpOrg
  bind:this={popUpOrg}
  bind:isSubmitted={isSubmitted}
  bind:childName={orgName}
  bind:headTitle={orgHeadTitle}
  bind:heading={popUpHeading}
  onSubmit={SubmitUpsertOrg}
/>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
  class="org {orgTypeStr} {isDragOver ? 'drag-over' : ''} {isDragging ? 'dragging' : ''}"
  style="--indent-width:{indentWidth};"
  draggable="true"
  on:dragstart={updateOnMoving}
  on:drop|preventDefault={updateMovedOrg}
  on:dragover|preventDefault={() => (isDragOver = true)}
  on:dragleave|preventDefault={() => (isDragOver = false)}
  on:dragend={() => (isDragging = false)}
  >
  <div class="info">
    <span class="h-line"></span>
    <div class="label">
      <Icon
        className="icon"
        size="13"
        src={orgIcon}
      />
    </div>
    <span class="title {org.deletedAt > 0 ? 'deleted' : ''}">{org.name}</span>
  </div>
  <div class="options">
    {#if org.deletedAt === 0}
      {#if org.orgType !== OrgTypeJob}
        <button
          class="btn"
          title="Add child"
          on:click={() => toggleAddOrEdit(false)}
        >
          <Icon
            color="var(--gray-006)"
            className="icon"
            size="17"
            src={RiSystemAddBoxLine}
          />
        </button>
      {/if}
      <button
        class="btn"
        title="Edit organization"
        on:click={() => toggleAddOrEdit(true)}>
        <Icon
          color="var(--gray-006)"
          className="icon"
          size="17"
          src={RiDesignPencilLine}
        />
      </button>
      <button class="btn" title="Delete organization">
        <Icon
          color="var(--gray-006)"
          className="icon"
          size="17"
          src={RiSystemDeleteBinLine}
        />
      </button>
      <button class="btn" title="Info" on:click={updateEventInfo}>
        <Icon
          color="var(--gray-006)"
          className="icon"
          size="17"
          src={RiSystemInformationLine}
        />
      </button>
    {:else}
      <button class="btn" title="Rollback">
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
{#if org.children}
  {#each org.children as child, _ (child.id)}
    <svelte:self
      draggable="true"
      org={child}
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
  .org.company,
  .org.department,
  .org.division,
  .org.job {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-direction: row;
    gap: 40px;
    padding: 10px;
    border-radius: 8px;
  }

  .org.drag-over {
    background-color: var(--gray-002);
  }

  .org.dragging {
    background-color: #FFF;
  }

  .org:hover {
    background-color: var(--gray-001);
    cursor: move;
  }

  .org:hover .h-line,
  .org.drag-over .h-line {
    background-color: transparent !important;
  }

  .org.company .info .label {
    background-image: var(--blue-gradient);
    color: var(--blue-006);
    border: 1px solid var(--blue-003);
  }

  :global(.org.company .info .label .icon) {
    fill: var(--blue-005);
  }

  .org.department {
    padding-left: 37px;
  }

  .org.department .info .label {
    background-image: var(--orange-gradient);
    color: var(--orange-006);
    border: 1px solid var(--orange-003);
  }

  :global(.org.department .info .label .icon) {
    fill: var(--orange-005);
  }

  .org.division {
    padding-left: 65px;
  }

  .org.division .info .label {
    background-image: var(--green-gradient);
    color: var(--green-006);
    border: 1px solid var(--green-003);
  }

  :global(.org.division .info .label .icon) {
    fill: var(--green-006);
  }

  .org.job {
    padding-left: 93px;
  }

  .org.job .info .label {
    background-image: var(--gray-gradient);
    color: var(--gray-006);
    border: 1px solid var(--gray-003);
  }

  :global(.org.job .info .label .icon) {
    fill: var(--gray-007);
  }

  .org .info {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 10px;
    position: relative;
  }

  .org .info .h-line {
    position: absolute;
    left: -15px;
    width: 1px;
    height: 45px;
    background-color: var(--gray-003);
  }

  .org.company .info .h-line {
    display: none;
  }

  .org .info .label {
    display: flex;
    align-items: center;
    justify-content: center;
    width: fit-content;
    height: fit-content;
    padding: 5px;
    border-radius: 5px;
  }

  .org .info .title {
    font-size: 16px;
  }

  .org .info .title.deleted {
    color: var(--red-005);
    text-decoration: line-through;
    line-height: 2.2rem;
  }

  .org .options {
    display: flex;
    cursor: pointer;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .org .options .btn {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background-color: transparent;
    border-radius: 5px;
    padding: 5px;
    cursor: pointer;
  }

  .org .options .btn:hover {
    background-color: var(--gray-002);
  }

  .org .options .btn:active {
    background-color: var(--gray-003);
  }
</style>