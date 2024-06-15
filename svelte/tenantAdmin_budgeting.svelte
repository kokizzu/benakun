<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import BudgetPlanTree from './_components/BudgetPlanTree.svelte';
  import { localeDatetime } from './_components/formatter';

  /** @typedef {import('./_components/types/organization').Org} Org */
  /** @typedef {import('./_components/types/budget').BudgetPlan} BudgetPlan */

  let orgs = /** @type {Org[]} */ [/* orgs */];

  /**  @type {Org[]} */
  let REFORMAT_ORGS = [];

  function orgMaker(id) {
    /** @type {Org} */
    let orgFormatted = {
      id: '',
      name: '',
      orgType: 0,
      parentId: '',
      tenantCode: '',
      createdAt: 0,
      createdBy: '',
      updatedAt: 0,
      updatedBy: '',
      deletedAt: 0,
      children: [],
      headTitle: ''
    }
    for (let i in orgs) {
      if (orgs[i].id == String(id)) {
        const children = orgs[i].children;
        if (children && children.length) {
          for (let j in children) {
            const childId = children[j]
            const child = orgMaker(childId)
            orgFormatted.children = [...orgFormatted.children, child];
          }
        }
        orgFormatted.id = orgs[i].id
        orgFormatted.name = orgs[i].name
        orgFormatted.headTitle = orgs[i].headTitle
        orgFormatted.orgType = orgs[i].orgType
        orgFormatted.tenantCode = orgs[i].tenantCode
        orgFormatted.parentId = orgs[i].parentId
        orgFormatted.createdAt = orgs[i].createdAt
        orgFormatted.createdBy = orgs[i].createdBy
        orgFormatted.updatedAt = orgs[i].updatedAt
        orgFormatted.updatedBy = orgs[i].updatedBy
        orgFormatted.deletedAt = orgs[i].deletedAt
      }
    }
    return orgFormatted;
  }

  function reformatorgs() {
    let toorgs = [];

    if (orgs && orgs.length) {
      for (let i in orgs) {
        const id = orgs[i].id;
        const coa = orgMaker(id);
        toorgs = [...toorgs, coa];
      }
    }

    if (toorgs && toorgs.length) {
      toorgs = toorgs.filter(obj => obj.parentId <= 0);
      toorgs.sort((a, b) => a.level - b.level);
    }

    return toorgs;
  }

  // TODO:HABIBI show all budget from beginning, this one still empty everytime
  // eg. refresh the budget page, previously added program still empty

  // TODO:HABIBI find out why "unit" not saved in database or not showing after editing

  onMount(() => {
	  REFORMAT_ORGS = reformatorgs();
	  console.log('REFORMAT_ORGS',REFORMAT_ORGS);
	  console.log('orgs',orgs);
  });

  /** @type {BudgetPlan} */
  let planDetail = {
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
  let isShowPlanDetail = false;

  const onPlanDetails = (e) => {
    isShowPlanDetail = true;
    if (planDetail.id == (e.detail).id) return;
    planDetail = e.detail; 
    planDetail.budgetIDR = new Intl.NumberFormat('id', {
      style: 'currency',
      currency: 'IDR',
      maximumSignificantDigits: 3,
    }).format(Number(planDetail.budgetIDR) || 0);
  }
</script>

<MainLayout>
  <div class="orgs_container">
    {#if REFORMAT_ORGS && REFORMAT_ORGS.length}
      <div class="orgs">
        {#each REFORMAT_ORGS as org, _ (org.id)}
          <BudgetPlanTree
            org={org}
            on:details={onPlanDetails}
          />
        {/each}
      </div>
    {/if}
    <article class="plan_detail {isShowPlanDetail ? 'show' : ''}">
      <button class="close" on:click={() => isShowPlanDetail = false}>close</button>
      <main>
        <h2>{planDetail.title || '--'}</h2>
        <div class="detail">
          <span>Description</span>
          <p>{planDetail.description || '--'}</p>
        </div>
        <div class="detail">
          <span>Type</span>
          <p>{planDetail.planType || '--'}</p>
        </div>
        <div class="detail">
          <span>Year</span>
          <p>{planDetail.yearOf || '0'}</p>
        </div>
        <div class="detail">
          <span>Budget IDR</span>
          <p>{planDetail.budgetIDR || '0'}</p>
        </div>
        <div class="detail">
          <span>Qty</span>
          <p>{planDetail.quantity || '0'} {planDetail.unit || ''}</p>
        </div>
        <div class="detail">
          <span>Last modified</span>
          <p>{localeDatetime(planDetail.updatedAt || 0)}</p>
        </div>
      </main>
    </article>
  </div>
</MainLayout>

<style>
  .orgs_container {
    position: relative;
    display: flex;
    flex-direction: row;
    gap: 10px;
    width: 100%;
  }

  .orgs_container .orgs {
    width: 100%;
    max-width: 100%;
    display: flex;
    flex-direction: column;
    display: flex;
    height: fit-content;
    background-color: #FFF;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    overflow: hidden;
    padding: 15px 20px 25px 20px;
    transition : 0.5s;
  }

  .orgs_container .plan_detail {
    right: -420px;
    position: fixed;
    display: flex;
    flex-direction: column;
    gap: 3px;
    padding: 20px;
    background-color: #FFF;
    width: 0;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    height: fit-content;
    transition : 0.5s;
  }

  .orgs_container .plan_detail.show {
    position: relative;
    right: 0;
    width: 400px;
  }

  .orgs_container .plan_detail .close {
    color: var(--red-006);
    font-size: var(--font-sm);
    position: absolute;
    top: 4px;
    right: 4px;
    width: fit-content;
    height: fit-content;
    display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px 8px;
		border-radius: 3px;
		border: none;
    background-color: #ef444420;
		cursor: pointer;
  }

	.orgs_container .plan_detail .close:hover {
		background-color: #ef444430;
	}

  .orgs_container .plan_detail main {
    display: flex;
    flex-direction: column;
    gap: 15px;
    padding: 0 !important;
    height: fit-content !important;
  }

  .orgs_container .plan_detail main h2 {
    margin: 5px 0;
  }

  .orgs_container .plan_detail main .detail {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .orgs_container .plan_detail main .detail span {
    background-color: var(--sky-transparent);
    padding: 2px 4px;
    border-radius: 3px;
    color: var(--sky-007);
    width: fit-content;
    height: fit-content;
    font-size: var(--font-sm);
  }

  .orgs_container .plan_detail main .detail p {
    margin: 0;
    font-weight: 500;
    font-size: var(--font-md);
    line-height: 1.2em;
  }
</style>