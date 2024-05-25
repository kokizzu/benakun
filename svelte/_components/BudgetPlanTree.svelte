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
  import PlanProgramTree from './PlanProgramTree.svelte';
  import {
    TenantAdminGetBudgetPlans,
    TenantAdminCreateBudgetPlan,
    TenantAdminUpdateBudgetPlan
  } from '../jsApi.GEN.js';
  import { notifier } from './notifier.js';
  import PopUpBudgetPlan from './PopUpBudgetPlan.svelte';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let org = /** @type {import('./types/organization.js').Org} */ ({});

  const PlanTypeVision = 'vision', PlanTypeMission = 'mission', PlanTypeProgram = 'program', PlanTypeActivity	= 'activity';

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
      orgType = 'division', orgIcon = RiUserFacesGroup3Line;
      break;
    }
    case OrgTypeJob: {
      orgType = 'job', orgIcon = RiBusinessBriefcaseLine;
      break;
    }
  }
  
  /** @typedef {import('./types/budget.js').BudgetPlan} BudgetPlan */

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
    perYear = 0, budgetIDR = 0, budgetUSD = 0, budgetEUR = 0;

  // reset payload
  const resetPayload = () => {
    planType = '', title = '', description = '';
    perYear = 0, budgetIDR = 0, budgetUSD = 0, budgetEUR = 0;
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
  }

  // +================================================================+

  async function getBugetPlans() {
    isSearching = true;
    await TenantAdminGetBudgetPlans(
      { orgId: Number(org.id) },
      /** @type {import('../jsApi.GEN').TenantAdminGetBudgetPlansCallback}*/
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
  const submitStateDelete = 'delete', submitStateRestore = 'restore';

  let isSubmitPlan = false;
  let submitState = submitStateAdd;

  async function submitAddPlan() {
    isSubmitPlan = true;
    /** @type {import('../jsApi.GEN.js').TenantAdminCreateBudgetPlanIn} */
    const i = {
      planType: planType, title, description, parentId: 0, orgId: Number(org.id), perYear: Number(perYear),
      budgetIDR: Number(budgetIDR), budgetUSD: Number(budgetUSD), budgetEUR: Number(budgetEUR)
    }
    await TenantAdminCreateBudgetPlan( i, /** @type {import('../jsApi.GEN').TenantAdminCreateBudgetPlanCallback}*/
      function (/** @type {any} */ o) {
        resetPayload();
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        notifier.showSuccess(title + ' added');
        const out = /** @type {import('../jsApi.GEN').TenantAdminCreateBudgetPlanOut}*/ (o);
        budgetPlans = out.plans;

        reformatPlans();
        popUpBudgetPlan.hide();
      }
    )
  }

  async function submitEditPlan() {
    isSubmitPlan = true;
    /** @type {import('../jsApi.GEN.js').TenantAdminUpdateBudgetPlanIn} */
    const i = {
      id, planType, title, description, perYear: Number(perYear),
      budgetIDR: Number(budgetIDR), budgetUSD: Number(budgetUSD), budgetEUR: Number(budgetEUR)
    }
    await TenantAdminUpdateBudgetPlan(i, /** @type {import('../jsApi.GEN').TenantAdminUpdateBudgetPlanCallback} */
      function (/** @type {any} */ o) {
        resetPayload();
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        notifier.showSuccess(planType + ' edited');
        const out = /** @type {import('../jsApi.GEN').TenantAdminUpdateBudgetPlanOut}*/ (o);
        budgetPlans = out.plans;

        reformatPlans();
        popUpBudgetPlan.hide();
      }
    )
  }

  async function submitPlan() {
    switch (submitState) {
      case submitStateAdd: {
        await submitAddPlan();
        break;
      }
      case submitStateEdit: {
        await submitEditPlan();
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
  bind:perYear
  bind:budgetIDR
  bind:budgetUSD
  bind:budgetEUR
  bind:heading={headingPopUp}
  bind:isSubmitted={isSubmitPlan}
  onSubmit={submitPlan}
/>

<div class="org_container {orgType}">
  <span class="h-line"></span>
  <div class="org_wrapper">
    <button class="org {isShowPlans ? 'active' : ''}" on:click={toggleShowPlans}>
      <div class="info">
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
        <button class="btn arrow" title="Show plans">
          <Icon
            color="var(--gray-006)"
            className="icon {isShowPlans ? 'rotate' : 'dropdown'}"
            size="14"
            src={RiArrowsArrowRightSLine}
          />
        </button>
      </div>
    </button>
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
                  color="#FFF"
                  size="14"
                  src={RiSystemAddBoxLine}
                />
                <span>Add</span>
              </button>
            {:else}
              <button
                class="btn"
                on:click={() => togglePopUp(submitStateEdit, PlanTypeVision, visionId)}
                title="Edit vision"
              >
                <Icon
                  className="icon"
                  color="#FFF"
                  size="14"
                  src={RiDesignPencilLine}
                />
                <span>Edit</span>
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
                  color="#FFF"
                  size="14"
                  src={RiSystemAddBoxLine}
                />
                <span>Add</span>
              </button>
            {:else}  
              <button
                class="btn"
                on:click={() => togglePopUp(submitStateEdit, PlanTypeMission, missionId)}
                title="Edit mission"
              >
                <Icon
                  className="icon"
                  color="#FFF"
                  size="14"
                  src={RiDesignPencilLine}
                />
                <span>Edit</span>
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
                color="#FFF"
                size="14"
                src={RiSystemAddBoxLine}
              />
              <span>Add</span>
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
</div>

{#if org.children && org.children.length > 0}
  {#each org.children as child, _ (child.id)}
    {#if child.deletedAt === 0}
      {#if child.orgType !== OrgTypeJob}
        <svelte:self org={child} on:details />
      {/if}
    {/if}
  {/each}
{/if}

<style>
  :global(.org_container.company .org .info .label .icon) {
    fill: var(--blue-005);
  }

  :global(.org_container.department .org .info .label .icon) {
    fill: var(--orange-005);
  }

  :global(.org_container.division .org .info .label .icon) {
    fill: var(--green-006);
  }

  :global(.org.job .info .label .icon) {
    fill: var(--gray-007);
  }

  .org_container {
    width: 100%;
    display: flex;
    flex-direction: row;
    gap: 5px;
  }

  .org_container .h-line {
    left: -15px;
    width: 1px;
    height: auto;
    background-color: var(--gray-003);
  }

  .org_container.company .h-line {
    display: none !important;
  }

  .org_container.company .org .info .label {
    background-image: var(--blue-gradient);
    color: var(--blue-006);
    border: 1px solid var(--blue-003);
  }

  .org_container.department {
    padding-left: 37px;
  }

  .org_container.department .org .info .label {
    background-image: var(--orange-gradient);
    color: var(--orange-006);
    border: 1px solid var(--orange-003);
  }

  .org_container.division{
    padding-left: 65px;
  }

  .org_container.division .org .info .label {
    background-image: var(--green-gradient);
    color: var(--green-006);
    border: 1px solid var(--green-003);
  }

  .org_container.job {
    padding-left: 93px;
  }

  .org_container.job .org .info .label {
    background-image: var(--gray-gradient);
    color: var(--gray-006);
    border: 1px solid var(--gray-003);
  }

  .org_container .org {
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: space-between;
    gap: 40px;
    padding: 10px;
    border-radius: 8px;
    cursor: pointer;
    border: none;
    background-color: transparent;
    color: var(--gray-007);
  }

  .org_container.company .org:hover,
  .org_container.company .org.active {
    background-color: var(--blue-transparent);
  }
  .org_container.department .org:hover,
  .org_container.department .org.active {
    background-color: var(--orange-transparent);
  }
  .org_container.division .org:hover,
  .org_container.division .org.active {
    background-color: var(--green-transparent);
  }

  .org_container .org:active {
    background-color: var(--gray-002);
  }

  .org_container .org:hover .h-line {
    background-color: transparent !important;
  }
  
  .org_container .org .info {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 10px;
    position: relative;
  }

  .org_container .org .info .label {
    display: flex;
    align-items: center;
    justify-content: center;
    width: fit-content;
    height: fit-content;
    padding: 5px;
    border-radius: 5px;
  }

  .org_container .org .info .title {
    font-size: 16px;
  }

  .org_container .org .options {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .org_container .org .options .btn,
  .org_container .org_wrapper .org_plans .plan .label .btn {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background-color: var(--blue-006);
    border-radius: 5px;
    padding: 4px 8px;
    cursor: pointer;
    color: #FFF;
    gap: 5px;
    font-size: var(--font-sm);
  }

  .org_container .org_wrapper .org_plans .plan .label span {
    padding: 0;
  }

  .org_container .org .options .btn:hover,
  .org_container .org_wrapper .org_plans .plan .label .btn:hover {
    background-color: var(--blue-005);
  }

  :global(.dropdown) {
		transition: all .2s ease-in-out;
	}

	:global(.rotate) {
		transition: all .2s ease-in-out;
		transform: rotate(90deg);
	}

  .org_container .org .options .btn.arrow {
    background-color: transparent !important;
  }
  .org_container .org .options .btn.arrow:hover,
  .org_container .org .options .btn.arrow:active {
    background-color: transparent;
  }

  .org_container .org .options .btn:active {
    background-color: var(--gray-003);
  }

  .org_container .org_wrapper {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .org_container .org_wrapper .org_plans{
    width: auto;
    display: flex;
    flex-direction: column;
    padding: 5px 10px;
    border-radius: 8px;
    background-color: var(--gray-001);
  }

  .org_container .org_wrapper .org_plans .plan {
    display: flex;
    flex-direction: column;
    padding: 5px;
    gap: 5px;
  }

  .org_container .org_wrapper .org_plans .plan .label {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    gap: 8px;
    font-size: var(--font-md);
    font-weight: 600;
  }

  .org_container .org_wrapper .org_plans .plan p {
    margin: 0;
  }

  .org_container .org_wrapper .org_plans .plan.programs {
    gap: 8px;
  }

  .org_container .org_wrapper .org_plans .plan.programs .program_activity_list {
    display: flex;
    flex-direction: column;
  }
</style>