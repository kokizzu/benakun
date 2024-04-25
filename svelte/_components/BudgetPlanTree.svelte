<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiBuildingsCommunityLine from 'svelte-icons-pack/ri/RiBuildingsCommunityLine';
  import RiBuildingsBuilding2Line from 'svelte-icons-pack/ri/RiBuildingsBuilding2Line';
  import RiUserTeamLine from 'svelte-icons-pack/ri/RiUserTeamLine';
  import RiBusinessBriefcaseLine from 'svelte-icons-pack/ri/RiBusinessBriefcaseLine';
  import RiSystemArrowRightSLine from 'svelte-icons-pack/ri/RiSystemArrowRightSLine';

  /** @typedef {import('./types/organization.js').Org} Org */
  
  /** @type Org */ //@ts-ignore
  export let org = {}

  let orgType = 'company', orgIcon = RiBuildingsCommunityLine
  let OrgTypeCompany = 1, OrgTypeDept = 2, OrgTypeDivision = 3, OrgTypeJob = 4;
  switch (org.orgType) {
    case OrgTypeCompany: {
      orgType = 'company', orgIcon = RiBuildingsCommunityLine;
      break;
    }
    case OrgTypeDept: {
      orgType = 'department', orgIcon = RiBuildingsBuilding2Line;
      break;
    }
    case OrgTypeDivision: {
      orgType = 'division', orgIcon = RiUserTeamLine;
      break;
    }
    case OrgTypeJob: {
      orgType = 'job', orgIcon = RiBusinessBriefcaseLine;
      break;
    }
  }
</script>

<div class="org {orgType}">
  <div class="info">
    <span class="h-line"></span>
    <div class="label">
      <Icon
        className="icon"
        size="13"
        src={orgIcon}
      />
    </div>
    <span class="title">{org.name}</span>
  </div>
  <div class="options">
    <button class="btn arrow show" title="Show plans">
      <Icon
        color="var(--gray-006)"
        className="icon"
        size="17"
        src={RiSystemArrowRightSLine}
      />
    </button>
  </div>
</div>
{#if org.children}
  {#each org.children as child, _ (child.id)}
    <svelte:self org={child} />
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
    justify-content: space-between;
    gap: 40px;
    padding: 10px;
    border-radius: 8px;
    cursor: pointer;
  }

  .org:hover {
    background-color: var(--gray-001);
  }

  .org:active {
    background-color: var(--gray-002);
  }

  .org:hover .h-line {
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

  .org .options {
    display: flex;
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

  .org .options .btn.arrow:hover,
  .org .options .btn.arrow:active {
    background-color: transparent;
  }

  .org .options .btn:active {
    background-color: var(--gray-003);
  }
</style>