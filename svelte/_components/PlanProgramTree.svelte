<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FaFolder } from '../node_modules/svelte-icons-pack/dist/fa';
  import {
    RiMediaSpeedMiniFill,
    RiSystemAddBoxLine,
    RiDesignPencilLine
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { createEventDispatcher } from 'svelte';
  import { TenantAdminUpsertBudgetPlan } from '../jsApi.GEN.js';
  import PopUpBudgetPlan from './PopUpBudgetPlan.svelte';
  import { notifier } from './notifier.js';

  /** @typedef {import('./types/budget.js').BudgetPlan} BudgetPlan */

  const dispatch = createEventDispatcher();

  const PlanTypeProgram = 'program', PlanTypeActivity	= 'activity';

  let headingPopUp = 'Add activity';

  // make it as payload
  let planType = PlanTypeActivity, title = '', description = '',
  yearOf = 0, budgetIDR = '', quantity = 0, unit = '';

  // reset payload
  const resetPayload = () => {
    planType = '', title = '', description = '';
    yearOf = 0, budgetIDR = '', quantity = 0, unit = '';
  }

  /** @type BudgetPlan */
  export let plan = {
    id: '',
    parentId: '',
    title: '',
    description: '',
    orgId: '',
    planType: '',
    yearOf: 0,
    budgetIDR: '',
    quantity: 0,
	  unit: '',
    createdAt: '',
    createdBy: '',
    updatedAt: 0,
    updatedBy: '',
    deletedAt: 0,
    children: []
  };

  console.log('plan:', plan)

  // forward event to the root component, to show program/activity details
  const showDetails = () => dispatch('details', plan);

  let popUpBudgetPlan = /** @type {PopUpBudgetPlan} */ ({});


  // state submit
  const submitStateAdd = 'add', submitStateEdit = 'edit';

  let isSubmitPlan = false;
  let submitState = submitStateAdd;

  async function submitUpsertPlan() {
    isSubmitPlan = true;

    /** @type BudgetPlan */ // @ts-ignore
    let planPayload = { // @ts-ignore
      id: submitState == 'add' ? 0 : Number(plan.id),
      planType,
      title,
      description,
      yearOf: Number(yearOf),
      orgId: plan.orgId,
      budgetIDR: budgetIDR+'',
      quantity: Number(quantity),
      unit: unit,
      parentId: plan.id,
    }
    await TenantAdminUpsertBudgetPlan( //@ts-ignore
      { plan: planPayload }, /** @type {import('../jsApi.GEN').TenantAdminUpsertBudgetPlanCallback} */
      function (/** @type {any} */ o) {
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        notifier.showSuccess(planType + ' updated');
        const out = /** @type {import('../jsApi.GEN').TenantAdminUpsertBudgetPlanOut}*/ (o);
        dispatch('update', out.plans);

        popUpBudgetPlan.hide();
        resetPayload();
      }
    )
  }

  async function submitPlan() {
    switch (submitState) {
      case submitStateAdd:
      case submitStateEdit: {
        await submitUpsertPlan();
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
        planType = PlanTypeActivity;
        break;
      }
      case submitStateEdit: {
        submitState = submitStateEdit;
        headingPopUp = 'Edit ' + plan.planType;
        title = plan.title, description = plan.description, yearOf = plan.yearOf;
        budgetIDR = plan.budgetIDR, quantity = plan.quantity, unit = plan.unit;
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
  bind:yearOf
  bind:budgetIDR
  bind:quantity
  bind:unit
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

{#each (plan.children ||[]) as child, _ (child.id)}
  {#if child.planType == 'activity'}
    <svelte:self
      plan={child}
      on:update
      on:details
    />
  {/if}
{/each}

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
    text-align: left;
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