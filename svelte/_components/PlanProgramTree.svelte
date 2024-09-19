<script>
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FaFolder } from '../node_modules/svelte-icons-pack/dist/fa';
  import {
    RiMediaSpeedMiniFill,
    RiSystemAddBoxLine,
    RiDesignPencilLine,
    RiSystemInformationLine,
		RiSystemDeleteBin5Line,
    RiArrowsArrowGoBackLine,
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { createEventDispatcher } from 'svelte';
  import { TenantAdminBudgeting } from '../jsApi.GEN.js';
  import PopUpBudgetPlan from './PopUpBudgetPlan.svelte';
  import { notifier } from './xNotifier.js';

  /** @typedef {import('./types/budget.js').BudgetPlan} BudgetPlan */

  const dispatch = createEventDispatcher();

  const PlanTypeProgram = 'program', PlanTypeActivity	= 'activity';

  let headingPopUp = 'Add activity';

  let lastYearInput = Number(localStorage.getItem('lastYearInput')) || new Date().getFullYear();

  // make it as payload
  let planType = PlanTypeActivity, title = '', description = '',
  yearOf = lastYearInput, budgetIDR = '', quantity = 0, unit = '';

  // reset payload
  const resetPayload = () => {
    planType = '', title = '', description = '';
    yearOf = lastYearInput, budgetIDR = '', quantity = 0, unit = '';
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
      budgetIDR: (budgetIDR || 0)+'',
      quantity: Number(quantity),
      unit: unit,
      parentId: plan.id,
    }
    await TenantAdminBudgeting( //@ts-ignore
      {
        cmd: 'upsert', //@ts-ignore
        plan: planPayload
      }, /** @type {import('../jsApi.GEN').TenantAdminBudgetingCallback} */
      function (/** @type {any} */ o) {
        isSubmitPlan = false;
        lastYearInput = Number(yearOf);
        localStorage.setItem('lastYearInput', lastYearInput+'');
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }

        notifier.showSuccess(planType + ' updated');
        const out = /** @type {import('../jsApi.GEN').TenantAdminBudgetingOut}*/ (o);
        dispatch('update', out.plans);

        popUpBudgetPlan.hide();
        resetPayload();
      }
    )
  }

  async function deletePlan() {
    isSubmitPlan = true;
    await TenantAdminBudgeting( //@ts-ignore
      {
        cmd: 'delete', //@ts-ignore
        plan: {
          id: Number(plan.id)
        }
      }, /** @type {import('../jsApi.GEN').TenantAdminBudgetingCallback} */
      function (/** @type {any} */ o) {
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }

        notifier.showSuccess(planType + ' deleted');
        const out = /** @type {import('../jsApi.GEN').TenantAdminBudgetingOut}*/ (o);
        dispatch('update', out.plans);

        popUpBudgetPlan.hide();
        resetPayload();
      }
    )
  }

  async function restorePlan() {
    isSubmitPlan = true;
    await TenantAdminBudgeting( //@ts-ignore
      {
        cmd: 'restore', //@ts-ignore
        plan: {
          id: Number(plan.id)
        }
      }, /** @type {import('../jsApi.GEN').TenantAdminBudgetingCallback} */
      function (/** @type {any} */ o) {
        isSubmitPlan = false;
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }

        notifier.showSuccess(planType + ' restore');
        const out = /** @type {import('../jsApi.GEN').TenantAdminBudgetingOut}*/ (o);
        dispatch('update', out.plans);

        popUpBudgetPlan.hide();
        resetPayload();
      }
    )
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

  $: {
    if (lastYearInput) {
      localStorage.setItem('lastYearInput', lastYearInput+'');
    }
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
  onSubmit={submitUpsertPlan}
/>

<div class="item {plan.planType === 'activity' ? 'activity' : ''}">
  <div class="title">
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
  </div>
  <div class="options">
    <button
      disabled={isSubmitPlan}
      on:click={() => togglePopUp(submitStateEdit)}
      class="btn"
      title="Edit {plan.planType}"
    >
      <Icon
        className="icon"
        color="var(--gray-007)"
        size="15"
        src={RiDesignPencilLine}
      />
    </button>
    {#if plan.deletedAt <= 0}
      <button
        disabled={isSubmitPlan}
        on:click={deletePlan}
        class="btn"
        title="Delete {plan.planType}"
      >
        <Icon
          className="icon"
          color="var(--gray-007)"
          size="14"
          src={RiSystemDeleteBin5Line}
        />
      </button>
    {/if}
    {#if plan.deletedAt > 0}
      <button
        disabled={isSubmitPlan}
        class="btn"
        title="Restore {plan.planType}"
      >
        <Icon
          className="icon"
          color="var(--gray-007)"
          size="15"
          src={RiArrowsArrowGoBackLine}
        />
      </button>
    {/if}
    {#if plan.planType === 'program'}
      <button
        disabled={isSubmitPlan}
        on:click={() => togglePopUp(submitStateAdd)}
        class="btn"
        title="Add activity"
      >
        <Icon
          className="icon"
          color="var(--gray-007)"
          size="15"
          src={RiSystemAddBoxLine}
        />
      </button>
    {/if}
    <button
      class="btn info"
      title="Info"
      on:click={showDetails}
    >
      <Icon
        size="15"
        color="var(--gray-007)"
        src={RiSystemInformationLine}
      />
    </button>
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
    overflow: hidden;
  }

  :global(.item .title .icon) {
    flex-shrink : 0;
  }

  .item .title span  {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .item .options {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  
  .item .options .btn {
		border: none;
		padding: 6px;
		border-radius: 8px;
		background-color: transparent;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.item .options .btn:hover {
		background-color: var(--violet-transparent);
	}

	:global(.item .options .btn:hover svg) {
		fill: var(--violet-005);
	}
</style>