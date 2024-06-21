<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiBuildingsCommunityLine,
    RiBuildingsBuilding2Line,
    RiUserFacesGroup3Line,
    RiBusinessBriefcaseLine,
    RiSystemAddBoxLine,
    RiDesignPencilLine,
    RiArrowsArrowRightSLine
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { TrOutlineRefresh } from '../node_modules/svelte-icons-pack/dist/tr';
  import PlanProgramTree from './PlanProgramTree.svelte';
  import { TenantAdminBudgeting } from '../jsApi.GEN.js';
  import { notifier } from './notifier.js';
  import PopUpBudgetPlan from './PopUpBudgetPlan.svelte';
  import { createEventDispatcher } from 'svelte';

  /** @typedef {import('./types/budget.js').BudgetPlan} BudgetPlan */

  const dispatch = createEventDispatcher();

  export let org = /** @type {import('./types/organization.js').Org} */ ({});

  const PlanTypeVision    = 'vision';
  const PlanTypeMission   = 'mission';
  const PlanTypeProgram   = 'program';
  const PlanTypeActivity  = 'activity';

  let orgType = 'company';
  let orgIcon = RiBuildingsCommunityLine;

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
      orgType = 'division', orgIcon = RiUserFacesGroup3Line;
      break;
    }
    case OrgTypeJob: {
      orgType = 'job', orgIcon = RiBusinessBriefcaseLine;
      break;
    }
  }

  // Budget plans to render
  let budgetPlans = /** @type {BudgetPlan[]} */ ([]);

  // only show description for vision and mission
  let visionDesc = '', missionDesc = '';
  // define vision and mission id for update
  let visionId = 0, missionId = 0;

  // programs and activities to render
  let programsActivity = /** @type {BudgetPlan[]} */ ([]);

  let headingPopUp = 'Add budget plan';

  // make it as payload
  let id = 0, planType = PlanTypeVision, title = '', description = '',
    yearOf = 0, budgetIDR = '', unit = '', quantity = 0;

  // reset payload
  const resetPayload = () => {
    planType = '', title = '', description = '';
    yearOf = 0, budgetIDR = '', unit = '', quantity = 0;
  }

  // state render budget plans
  let isSearching = false, isShowPlans = false;

  // +==================== CONTROL FLOW, REFORMAT DATA ====================+

  function activityMaker(/** @type {string} */ programId) {
    let activities = /** @type {BudgetPlan[]} */ ([]);
    for (let i in budgetPlans) {
      if (budgetPlans[i].planType === PlanTypeActivity && budgetPlans[i].parentId === programId) {
        activities = [...activities, budgetPlans[i]];
      }
    }

    return activities
  }

  function reformatPrograms() {
    if (!budgetPlans || budgetPlans.length === 0) return;
    programsActivity = [];

    for (let i in budgetPlans) {
      if (budgetPlans[i].planType === PlanTypeProgram) {
        let program = /** @type {BudgetPlan} */ (budgetPlans[i]);
        program.children = activityMaker(program.id);

        programsActivity = [...programsActivity, program];
      }
    }
  }

  // reformat budget plans data
  function reformatPlans() {
    if (budgetPlans && budgetPlans.length > 0) {
      for (let i in budgetPlans) {
        if (budgetPlans[i].planType === PlanTypeMission) {
          missionDesc = budgetPlans[i].description;
          missionId = Number(budgetPlans[i].id);
        }
        if (budgetPlans[i].planType === PlanTypeVision) {
          visionDesc = budgetPlans[i].description;
          visionId = Number(budgetPlans[i].id);
        }
      }
    }

    reformatPrograms();
    console.log('REFORMATTED PLANS:', programsActivity);
  }

  // +================================================================+

  async function getBugetPlans() {
    isSearching = true;
    await TenantAdminBudgeting( //@ts-ignore
      {
        cmd: 'form',
        orgId: Number(org.id)
      },
      /** @type {import('../jsApi.GEN').TenantAdminBudgetingCallback}*/
      function (/** @type {any} */ o) {
        isSearching = false;
        if (o.error) {
          notifier.showError(o.error);
          return;
        }

        console.log(o);
        budgetPlans = o.plans;
        reformatPlans();
      }
    )
  }

  const toggleShowPlans = async () => {
    if (isShowPlans) {
      isShowPlans = false;
      return;
    } else {
      isShowPlans = true;
      if (!budgetPlans || budgetPlans.length === 0) {
        await getBugetPlans();
      }
    }
  }

  // popup budget plan bindings
  let popUpBudgetPlan = /** @type {PopUpBudgetPlan} */ ({});

  // state submit
  const submitStateAdd = 'add', submitStateEdit = 'edit';

  let isSubmitPlan = false;
  let submitState = submitStateAdd;


  async function submitUpsertPlan() {
    isSubmitPlan = true;
    const i = { //@ts-ignore
      cmd: 'upsert',
      plan: {
        id: submitState == 'add' ? 0 : id,
        planType, title, description,
        yearOf: Number(yearOf),
        orgId: Number(org.id),
        budgetIDR: ''+(budgetIDR || 0),
        unit,
        quantity: +quantity
      }
    }
    await TenantAdminBudgeting( //@ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminBudgetingCallback} */
      function (/** @type {any} */ o) {
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }

        notifier.showSuccess(planType + ' updated');
        const out = /** @type {import('../jsApi.GEN').TenantAdminBudgetingOut}*/ (o);
        budgetPlans = out.plans;

        resetPayload();
        reformatPlans();
        popUpBudgetPlan.hide();
      }
    )
  }

  async function submitPlan() {
    switch (submitState) {
      case submitStateAdd || submitStateEdit: {
        await submitUpsertPlan();
        break;
      }
      default: {
        console.log('invalid submit state, use add, or edit');
        return;
      }
    }
  }

  /**
   * @param {string} state
   * @param {string} [pType]
   * @param {number} [toId]
   */
  function togglePopUp(state, pType, toId) {
    if (!state || !pType) {
      console.log('invalid toggle state, use add, or edit');
      return;
    }

    resetPayload();

    switch (state) {
      case submitStateAdd: {
        submitState = submitStateAdd;
        break;
      }
      case submitStateEdit: {
        submitState = submitStateEdit;
        id = Number(toId);
        break;
      }
      default: {
        console.log('invalid submit state, use add, or edit');
        return;
      }
    }
    
    switch (pType) {
      case PlanTypeVision: {
        planType = PlanTypeVision;
        description = visionDesc;
        break;
      }
      case PlanTypeMission: {
        planType = PlanTypeMission;
        description = missionDesc;
        break;
      }
      case PlanTypeProgram: {
        planType = PlanTypeProgram;
        break;
      }
      default: {
        console.log('invalid plan type, use vision, mission, or program');
        return;
      }
    }

    headingPopUp = state + ' ' + planType;

    popUpBudgetPlan.show();
  }

  function onUpdateProgramActivity(e) {
    const plans = /** @type {BudgetPlan[]}*/ (e.detail);
    budgetPlans = plans;
    reformatPlans();
  }
</script>

<PopUpBudgetPlan
  bind:this={popUpBudgetPlan}
  bind:planType
  bind:title
  bind:description
  bind:yearOf
  bind:budgetIDR
  bind:unit
  bind:quantity
  bind:heading={headingPopUp}
  bind:isSubmitted={isSubmitPlan}
  onSubmit={submitPlan}
/>

<div class="org_container {orgType}">
  <div class="org">
    <button class="btn_toggle"  on:click={toggleShowPlans}>
      <Icon
        color="var(--gray-006)"
        className="icon {isShowPlans ? 'rotate' : 'dropdown'}"
        size="20"
        src={RiArrowsArrowRightSLine}
      />
    </button>
    <div class="info">
      <Icon
        className="icon"
        size="17"
        src={orgIcon}
      />
      <p class="title">{org.name}</p>
    </div>
    <button
      class="btn_refresh"
      title="Refresh"
      on:click={getBugetPlans}
      disabled={isSearching}>
      <Icon
        className="icon {isSearching ? 'spin' : ''}"
        size="15"
        color="var(--gray-006)"
        src={TrOutlineRefresh}
      />
    </button>
  </div>
  {#if isShowPlans}
    <div class="org_plans">
      <div class="plan vision">
        <div class="label">
          <span>Vision</span>
          {#if !visionDesc || visionDesc === ''}
            <button
              class="btn"
              on:click={() => togglePopUp(submitStateAdd, PlanTypeVision, 0)}
              title="Add vision"
            >
              <Icon
                className="icon"
                color="var(--gray-007)"
                size="14"
                src={RiSystemAddBoxLine}
              />
            </button>
          {:else}
            <button
              class="btn"
              on:click={() => togglePopUp(submitStateEdit, PlanTypeVision, visionId)}
              title="Edit vision"
            >
              <Icon
                className="icon"
                color="var(--gray-007)"
                size="14"
                src={RiDesignPencilLine}
              />
            </button>
          {/if}
        </div>
        <p>{visionDesc || '--'}</p>
      </div>
      <div class="plan mission">
        <div class="label">
          <span>Mission</span>
          {#if !missionDesc || missionDesc === ''}
            <button
              class="btn"
              on:click={() => togglePopUp(submitStateAdd, PlanTypeMission, 0)}
              title="Add mission"
            >
              <Icon
                className="icon"
                color="var(--gray-007)"
                size="14"
                src={RiSystemAddBoxLine}
              />
            </button>
          {:else}  
            <button
              class="btn"
              on:click={() => togglePopUp(submitStateEdit, PlanTypeMission, missionId)}
              title="Edit mission"
            >
              <Icon
                className="icon"
                color="var(--gray-007)"
                size="14"
                src={RiDesignPencilLine}
              />
            </button>
          {/if}
        </div>
        <p>{missionDesc || '--'}</p>
      </div>
      <div class="plan programs">
        <div class="label">
          <span>Programs</span>
          <button
            class="btn"
            on:click={() => togglePopUp(submitStateAdd, PlanTypeProgram, 0)}
            title="Add program"
          >
            <Icon
              className="icon"
              color="var(--gray-007)"
              size="14"
              src={RiSystemAddBoxLine}
            />
          </button>
        </div>
        <div class="program_activity_list">
          {#if programsActivity && programsActivity.length > 0}
            {#each programsActivity as plan, _ (plan.id)}
              <PlanProgramTree
                {plan}
                on:details={(e) => dispatch('details', e.detail)}
                on:update={onUpdateProgramActivity}
              />
            {/each}
          {/if}
        </div>
      </div>
    </div>
  {/if}
</div>

{#if org.children && org.children.length > 0}
  {#each org.children as child, _ (child.id)}
    {#if child.orgType !== OrgTypeJob}
      <svelte:self org={child} on:details />
    {/if}
  {/each}
{/if}

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(-360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  :global(.dropdown) {
		transition: all .2s ease-in-out;
	}

	:global(.rotate) {
		transition: all .2s ease-in-out;
		transform: rotate(90deg);
	}

  :global(.org_container.company .org .info .icon) {
    fill: var(--blue-005);
  }
  :global(.org_container.department .org .info .icon) {
    fill: var(--orange-005);
  }
  :global(.org_container.division .org .info .icon) {
    fill: var(--green-006);
  }
  :global(.org.job .info .icon) {
    fill: var(--gray-007);
  }

  .org_container {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .org_container.department {
    padding-left: 37px;
  }
  .org_container.division{
    padding-left: 65px;
  }
  .org_container.job {
    padding-left: 93px;
  }

  .org_container .org {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 10px;
    border-radius: 8px;
    width: 100%;
    padding-right: 10px;
  }

  .org_container .btn_toggle {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    padding: 5px;
    background-color: transparent;
    border-radius: 999px;
    cursor: pointer;
  }

  .org_container .btn_toggle:hover {
    background-color: var(--gray-001);
  }

  .org_container .btn_toggle:active {
    background-color: var(--gray-002);
  }
  
  .org_container .org .info {
    margin-left: -7px;
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 10px;
    position: relative;
    flex-grow: 1;
    overflow: hidden;
  }

  :global(.org_container .org .info .icon) {
    flex-shrink : 0;
  }

  .org_container .org .info .title {
    font-size: 16px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin: 0;
  }

  .org_container .org_plans .plan .label span {
    padding: 0;
  }

  .org_container .org .options .btn:hover,
  .org_container .org_plans .plan .label .btn:hover {
    background-color: var(--blue-005);
  }

  .org_container .org_plans {
    width: auto;
    display: flex;
    flex-direction: column;
    margin-left: 35px;
    padding: 5px 5px 5px 10px;
    border-radius: 8px;
    background-color: var(--gray-001);
  }

  .org_container .org .btn_refresh,
  .org_container .org_plans .plan .label .btn {
		border: none;
		padding: 6px;
		border-radius: 8px;
		background-color: transparent;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
	}

  .org_container .org .btn_refresh:hover,
	.org_container .org_plans .plan .label .btn:hover {
		background-color: var(--violet-transparent);
	}

  :global(.org_container .org .btn_refresh:hover svg) {
    stroke: var(--violet-005);
  }
	:global(.org_container .org_plans .plan .label .btn:hover svg) {
		fill: var(--violet-005);
	}

  .org_container .org_plans .plan {
    display: flex;
    flex-direction: column;
    padding: 5px;
    gap: 5px;
  }

  .org_container .org_plans .plan .label {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    gap: 8px;
    font-weight: 600;
  }

  .org_container .org_plans .plan p {
    margin: 0;
  }

  .org_container .org_plans .plan.programs {
    gap: 8px;
  }

  .org_container .org_plans .plan.programs .program_activity_list {
    display: flex;
    flex-direction: column;
  }
</style>