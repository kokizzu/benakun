<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaFolder from 'svelte-icons-pack/fa/FaFolder';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

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
</script>

<button class="item {plan.planType === 'activity' ? 'activity' : ''}" on:click={showDetails}>
  <Icon
    className="icon"
    color="var(--gray-006)"
    size="13"
    src={FaFolder}
  />
  <span>{plan.title}</span>
</button>

<style>
  .item {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    background-color: transparent;
    border: none;
    border-radius: 8px;
    padding: 8px;
    color: var(--gray-007);
    cursor: pointer;
  }

  .item:hover {
    background-color: var(--gray-002);
  }

  .item.activity {
    margin-left: 30px !important;
  }
</style>