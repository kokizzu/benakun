<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaFolder from 'svelte-icons-pack/fa/FaFolder';
  import RiMediaSpeedMiniFill from "svelte-icons-pack/ri/RiMediaSpeedMiniFill";
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import { createEventDispatcher } from 'svelte';
  import { TenantAdminCreateBudgetPlan, TenantAdminUpdateBudgetPlan } from '../jsApi.GEN.js';
  import PopUpBudgetPlan from './PopUpBudgetPlan.svelte';
  import { notifier } from './notifier.js';

  const dispatch = createEventDispatcher();

  const PlanTypeProgram = 'program', PlanTypeActivity	= 'activity';

  let headingPopUp = 'Add activity';

  // make it as payload
  let planType = PlanTypeActivity, title = '', description = '',
  perYear = 0, budgetIDR = 0, budgetUSD = 0, budgetEUR = 0;

  // reset payload
  const resetPayload = () => {
    planType = '', title = '', description = '';
    perYear = 0, budgetIDR = 0, budgetUSD = 0, budgetEUR = 0;
  }

  /** @type {import('./types/budget.js').BudgetPlan} */
  export let plan = {
    id: '',
    parentId: '',
    title: '',
    description: '',
    orgId: '',
    planType: '',
    perYear: 0,
    budgetIDR: 0,
    budgetUSD: 0,
    budgetEUR: 0,
    createdAt: '',
    createdBy: '',
    updatedAt: 0,
    updatedBy: '',
    deletedAt: 0,
    children: []
  };

  // forward event to the root component, to show program/activity details
  const showDetails = () => dispatch('details', plan);

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
      planType: 'activity', title, description, parentId: Number(plan.id), orgId: Number(plan.orgId),
      perYear: Number(perYear), budgetIDR: Number(budgetIDR), budgetUSD: Number(budgetUSD), budgetEUR: Number(budgetEUR)
    }
    await TenantAdminCreateBudgetPlan(
      i, /** @type {import('../jsApi.GEN').TenantAdminCreateBudgetPlanCallback} */
      function (/** @type {any} */ o) {
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return
        }
        notifier.showSuccess(title + ' added');

        const out = /** @type {import('../jsApi.GEN').TenantAdminGetBudgetPlansOut}*/ (o);
        dispatch('update', out.plans);
        popUpBudgetPlan.hide();
      }
    )
  }

  async function submitEditPlan() {
    isSubmitPlan = true;
    /** @type {import('../jsApi.GEN.js').TenantAdminUpdateBudgetPlanIn} */
    const i = {
      id: Number(plan.id), planType, title, description, perYear: Number(perYear),
      budgetIDR: Number(budgetIDR), budgetUSD: Number(budgetUSD), budgetEUR: Number(budgetEUR)
    }
    await TenantAdminUpdateBudgetPlan(i, /** @type {import('../jsApi.GEN').TenantAdminUpdateBudgetPlanCallback} */
      function (/** @type {any} */ o) {
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        notifier.showSuccess(planType + ' edited');
        const out = /** @type {import('../jsApi.GEN').TenantAdminUpdateBudgetPlanOut}*/ (o);
        dispatch('update', out.plans);

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

  function togglePopUp(state) {
    switch (state) {
      case submitStateAdd: {
        submitState = submitStateAdd;
        headingPopUp = 'Add activity';
        break;
      }
      case submitStateEdit: {
        submitState = submitStateEdit;
        headingPopUp = 'Edit ' + plan.planType;
        title = plan.title, description = plan.description, perYear = plan.perYear;
        budgetIDR = plan.budgetIDR, budgetUSD = plan.budgetUSD, budgetEUR = plan.budgetEUR;
        break;
      }
      default: {
        console.log('invalid submit state, use add, or edit');
        return;
      }
    }

    popUpBudgetPlan.show();
  }
</script>

<PopUpBudgetPlan
  bind:this={popUpBudgetPlan}
  bind:planType={plan.planType}
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

<div class="item {plan.planType === 'activity' ? 'activity' : ''}">
  <button class="title" on:click={showDetails} title="click to show info">
    {#if plan.planType === 'program'}
      <Icon
        className="icon"
        color="var(--gray-006)"
        size="13"
        src={FaFolder}
      />
    {:else}
      <Icon
        className="icon"
        color="var(--gray-006)"
        size="13"
        src={RiMediaSpeedMiniFill}
      />
    {/if}
    <span>{plan.title}</span>
  </button>
  <div class="options">
    <button
      on:click={() => togglePopUp(submitStateEdit)}
      class="btn edit"
      title="Edit {plan.planType}"
    >
      <Icon
        className="icon"
        color="#FFF"
        size="14"
        src={RiDesignPencilLine}
      />
      <span>Edit</span>
    </button>
    {#if plan.planType === 'program'}
      <button
        on:click={() => togglePopUp(submitStateAdd)}
        class="btn"
        title="Add activity"
      >
        <Icon
          className="icon"
          color="#FFF"
          size="14"
          src={RiSystemAddBoxLine}
        />
        <span>Add</span>
      </button>
    {/if}
  </div>
</div>

{#if plan.children && plan.children.length > 0}
  {#each plan.children as child, _ (child.id)}
    {#if child.deletedAt === 0}
      {#if child.planType === 'activity'}
        <svelte:self
          plan={child}
          on:update
          on:details
        />
      {/if}
    {/if}
  {/each}
{/if}

<style>
  .item {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    gap: 10px;
    align-items: center;
    background-color: transparent;
    border: none;
    border-radius: 8px;
    padding: 5px 0 5px 0;
    color: var(--gray-007);
    cursor: pointer;
  }

  .item:hover .title span {
    color: var(--blue-006);
  }

  :global(.item:hover .title .icon) {
    fill: var(--blue-006);
  }

  .item.activity {
    margin-left: 30px !important;
  }

  .item .title {
    width: 100%;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    background-color: transparent;
    border: none;
    border-radius: 5px;
    color: var(--gray-007);
    padding: 4px 8px;
    cursor: pointer;
  }

  .item .title:hover {
    background-color: var(--blue-transparent);
  }

  .item .options {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }

  .item .options .btn {
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

  .item .options .btn:hover {
    background-color: var(--blue-005);
  }

  .item .options .btn.edit {
    background-color: var(--amber-006);
  }

  .item .options .btn.edit:hover {
    background-color: var(--amber-005);
  }
</style>