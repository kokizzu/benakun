<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaFolder from 'svelte-icons-pack/fa/FaFolder';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import { createEventDispatcher } from 'svelte';
  import { TenantAdminCreateBudgetPlan, TenantAdminUpdateBudgetPlan } from '../jsApi.GEN.js';
  import PopUpBudgetPlan from './PopUpBudgetPlan.svelte';
  import { notifier } from './notifier.js';

  const dispatch = createEventDispatcher();

  let headingPopUp = 'Add Activity';

  let title = '', description = '', perYear = 0, budgetIDR = 0, budgetUSD = 0, budgetEUR = 0;

  const resetPayload = () => {
    title = '', description = '', perYear = 0, budgetIDR = 0, budgetUSD = 0, budgetEUR = 0;
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

  const showDetails = () => dispatch('details', plan);

  let popUpBudgetPlan = /** @type {PopUpBudgetPlan} */ ({});
  let isSubmitAddPlan = false;

  async function submitAddPlan() {
    isSubmitAddPlan = true;
    /** @type {import('../jsApi.GEN.js').TenantAdminCreateBudgetPlanIn} */
    const i = {
      planType: plan.planType, title, description, parentId: Number(plan.parentId),
      orgId: Number(plan.orgId), perYear, budgetIDR, budgetUSD, budgetEUR
    }
    await TenantAdminCreateBudgetPlan(
      i, /** @type {import('../jsApi.GEN').TenantAdminCreateBudgetPlanCallback} */
      function (/** @type {any} */ o) {
        isSubmitAddPlan = false;
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

  const togglePopUpAdd = () => {
    resetPayload();
    headingPopUp = 'Add ' + (plan.planType === 'activity' ? 'activity' : 'program');
    popUpBudgetPlan.show();
  }

  const togglePopUpEdit = () => {
    resetPayload();
    headingPopUp = 'Edit ' + (plan.planType === 'activity' ? 'activity' : 'program');

    title = plan.title, description = plan.description, perYear = plan.perYear, budgetIDR = plan.budgetIDR, budgetUSD = plan.budgetUSD, budgetEUR = plan.budgetEUR;
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
  bind:isSubmitted={isSubmitAddPlan}
  onSubmit={submitAddPlan}
/>

<button
  class="item {plan.planType === 'activity' ? 'activity' : ''}"
  on:click={showDetails}>
  <div class="title">
    <Icon
      className="icon"
      color="var(--gray-006)"
      size="13"
      src={FaFolder}
    />
    <span>{plan.title}</span>
  </div>
  <div class="options">
    <button
      on:click={togglePopUpAdd}
      class="btn"
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
    <button
      on:click={togglePopUpEdit}
      class="btn"
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
  </div>
</button>

<style>
  .item {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    background-color: transparent;
    border: none;
    border-radius: 8px;
    padding: 8px 0 8px 8px;
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

  .item .title,
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
</style>