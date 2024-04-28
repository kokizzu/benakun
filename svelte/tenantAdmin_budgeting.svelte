<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import BudgetPlanTree from './_components/BudgetPlanTree.svelte';

  /** @typedef {import('./_components/types/organization').Org} Org */
  /** @type {Org[]} */
  let orgs = [/* orgs */];

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

  onMount(() => { REFORMAT_ORGS = reformatorgs(); console.log(orgs); });
</script>

<MainLayout>
  <div class="orgs_container">
    {#if REFORMAT_ORGS && REFORMAT_ORGS.length}
      <div class="orgs">
        {#each REFORMAT_ORGS as org, _ (org.id)}
          <BudgetPlanTree
            org={org}
          />
        {/each}
      </div>
    {/if}
  </div>
</MainLayout>

<style>
  .orgs_container {
    display: flex;
    width: 100%;
  }
  .orgs {
    width: 700px;
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
  }
</style>