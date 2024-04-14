<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiBuildingsCommunityLine from 'svelte-icons-pack/ri/RiBuildingsCommunityLine';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import RiSystemDeleteBinLine from 'svelte-icons-pack/ri/RiSystemDeleteBinLine';
  import RiSystemInformationLine from 'svelte-icons-pack/ri/RiSystemInformationLine';
  import RiBuildingsBuilding2Line from 'svelte-icons-pack/ri/RiBuildingsBuilding2Line';
  import RiUserTeamLine from 'svelte-icons-pack/ri/RiUserTeamLine';
  import RiBusinessBriefcaseLine from 'svelte-icons-pack/ri/RiBusinessBriefcaseLine';
  import { onMount, createEventDispatcher } from 'svelte';
  import PopUpOrgChild from './PopUpOrgChild.svelte';
  import { TenantAdminCreateOrganizationChild, TenantAdminUpdateOrganizationChild } from '../jsApi.GEN';
  import { notifier } from './notifier';

  const dispatch = createEventDispatcher();

  /**
    * @typedef {Object} Org
    * @property {string} id
    * @property {string} name
    * @property {string} headTitle
    * @property {number} orgType
    * @property {string} parentId
    * @property {string} tenantCode
    * @property {number} createdAt
    * @property {string} createdBy
    * @property {number} updatedAt
    * @property {string} updatedBy
    * @property {number} deletedAt
    * @property {Array<Org>} children
    */
  /**
   * @type {Org}
   */
  export let org = {
    id: "",
    name: "",
    headTitle: "",
    orgType: 0,
    parentId: "",
    tenantCode: "",
    createdAt: 0,
    createdBy: "",
    updatedAt: 0,
    updatedBy: "",
    deletedAt: 0,
    children: []
  }

  let orgType = 'company', orgIcon = RiBuildingsCommunityLine
  let OrgTypeCompany = 1, OrgTypeDept = 2, OrgTypeDivision = 3, OrgTypeJob = 4;
  switch (org.orgType) {
    case OrgTypeCompany:
      orgType = 'company', orgIcon = RiBuildingsCommunityLine;
      break;
    case OrgTypeDept:
      orgType = 'department', orgIcon = RiBuildingsBuilding2Line;
      break;
    case OrgTypeDivision:
      orgType = 'division', orgIcon = RiUserTeamLine;
      break;
    case OrgTypeJob:
      orgType = 'job', orgIcon = RiBusinessBriefcaseLine;
      break;
  }

  export let indent = 0;
  let indentWidth = '10px';
  const toIndentWidth = (/** @type {number} */ i) => { return `${i * 15 + 10}px` }

  onMount(() => indentWidth = toIndentWidth(indent))

  let popUpOrgChild, isSubmitted = false, popUpHeading = 'Add organization child';
  let orgName = '', headTitle = '', orgState = 'add';

  const toggleAddOrEdit = (/** @type {string}*/ state) => {
    if (state === 'add') orgState = 'add', orgName = '', headTitle = '', popUpHeading = 'Add coa child';
    else orgState = 'edit', orgName = org.name, headTitle = org.headTitle, popUpHeading = 'Edit: ' + org.name;
    popUpOrgChild.show();
  }

  async function submitAddOrgChild() {
    isSubmitted = true;
    if (orgName === '' || headTitle === '') {
      isSubmitted = false;
      notifier.showWarning('fields cannot be empty');
      return;
    }
    switch (orgState) {
      case 'add':
        await TenantAdminCreateOrganizationChild(
          {
            name: orgName,
            headTitle: headTitle,
            parentId: Number(org.id),
          },
          // @ts-ignore
          function (o) {
            isSubmitted = false;
            popUpOrgChild.hide();
            // @ts-ignore
            if (o.error) {
              // @ts-ignore
              notifier.showError(o.error);
              // @ts-ignore
              console.log(o.error);
              return;
            }
            console.log(o);
            // @ts-ignore
            dispatch('update', { orgs: o.orgs })
            notifier.showSuccess('Organization child created');
          }
        );
        break;
      case 'edit':
        await TenantAdminUpdateOrganizationChild(
          {
            name: orgName,
            id: Number(org.id),
            headTitle: headTitle,
          },
          // @ts-ignore
          function (o) {
            isSubmitted = false;
            popUpOrgChild.hide();
            // @ts-ignore
            if (o.error) {
              // @ts-ignore
              notifier.showError(o.error);
              // @ts-ignore
              console.log(o.error);
              return;
            }
            console.log(o);
            // @ts-ignore
            dispatch('update', { orgs: o.orgs })
            notifier.showSuccess('Organization child updated');
          }
        );
        break;
    }
  }

  function updateEventInfo() {
    // @ts-ignore
    dispatch('info', { org: org })
  }

  function updateOnMoving(/** @type {Org}*/ org) {
    dispatch('moving', { org: org })
  }
</script>

<PopUpOrgChild
  bind:this={popUpOrgChild}
  bind:isSubmitted={isSubmitted}
  bind:childName={orgName}
  bind:headTitle={headTitle}
  bind:heading={popUpHeading}
  onSubmit={submitAddOrgChild}
/>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
  class="org {orgType}"
  style="--indent-width:{indentWidth};"
  draggable="true"
  on:dragstart={() => updateOnMoving(org)}
  >
  <div class="info">
    <span class="h-line"></span>
    <div class="label">
      <Icon
        color="#FFF"
        className="icon"
        size="16"
        src={orgIcon}
      />
    </div>
    <span class="title">{org.name}</span>
  </div>
  <div class="options">
    {#if org.orgType !== OrgTypeJob}
      <button class="btn" title="Add child" on:click={() => toggleAddOrEdit('add')}>
        <Icon
          color="var(--gray-006)"
          className="icon"
          size="17"
          src={RiSystemAddBoxLine}
        />
      </button>
    {/if}
    <button class="btn" title="Edit organization" on:click={() => toggleAddOrEdit('edit')}>
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
  </div>
</div>
{#if org.children}
  {#each org.children as child, _ (child.id)}
    <svelte:self
      org={child}
      on:update
      on:info
      on:moving
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
    flex-direction: row;
    gap: 40px;
    padding: 10px 0;
    border-radius: 8px;
  }

  .org:hover {
    background-color: var(--gray-001);
    cursor: move;
  }

  .org:active {
    background-color: var(--gray-002);
  }

  .org.department {
    padding-left: 30px;
  }

  .org.department .info .label {
    background-color: var(--orange-006);
  }

  .org.division {
    padding-left: 50px;
  }

  .org.division .info .label {
    background-color: var(--green-006);
  }

  .org.job {
    padding-left: 70px;
  }

  .org.job .info .label {
    background-color: var(--gray-006);
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
    height: 50px;
    background-color: var(--gray-003);
  }

  .org.org.company .info .h-line {
    display: none;
  }

  .org .info .label {
    display: flex;
    align-items: center;
    justify-content: center;
    width: fit-content;
    height: fit-content;
    padding: 7px;
    border-radius: 9999px;
    background-color: var(--blue-006);
    color: #FFF;
  }

  .org .info .title {
    font-size: 18px;
    font-weight: 600;
  }

  .org.company:hover .options,
  .org.department:hover .options,
  .org.division:hover .options,
  .org.job:hover .options {
    display: flex;
    cursor: pointer;
  }

  .org .options {
    display: none;
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