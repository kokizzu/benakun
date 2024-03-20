<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import OrgTree from './_components/OrgTree.svelte';

  let segments = {/* segments */};
  let user = {/* user */};

  /**
    * @typedef {Object} Org
    * @property {string} id
    * @property {string} name
    * @property {string} headTitle
    * @property {number} orgType
    * @property {string} parentId
    * @property {string} tenantCode
    * @property {number} createdAt
    * @property {string} createdBy
    * @property {number} updatedAt
    * @property {string} updatedBy
    * @property {number} deletedAt
    * @property {Array<Org>} children
    */
  /**
   * @type {Array<Org>}
   */
  let orgs = [/* orgs */];

  /**
   * @type {Array<Org>}
   */
  let REFORMAT_ORGS = [];

  function orgMaker(id) {
    /**
     * @type {Org}
     */
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

  onMount(() => REFORMAT_ORGS = reformatorgs());

  let infoOrg = orgs.length > 0 ? orgs[0] : null;

  function updateEventHandler(e) {
    orgs = e.detail.orgs;
    for (let i in orgs) {
      if (orgs[i].id == infoOrg.id) {
        infoOrg = orgs[i];
        break;
      }
    }
    REFORMAT_ORGS = [];
    REFORMAT_ORGS = reformatorgs();
  }

  function infoEventHandler(e) {
    infoOrg = e.detail.org;
  }
</script>

<MainLayout>
  <div class="orgs_container">
    {#if REFORMAT_ORGS && REFORMAT_ORGS.length}
      <div class="orgs">
        {#each REFORMAT_ORGS as org, _ (org.id)}
          <OrgTree
            org={org}
            on:update={updateEventHandler}
            on:info={infoEventHandler}
          />
        {/each}
      </div>
      <div class="info">
        <h4>{infoOrg.name}</h4>
        <h4>{infoOrg.headTitle}</h4>
        <h4>{infoOrg.updatedAt}</h4>
      </div>
    {/if}
  </div>
</MainLayout>

<style>
  .orgs_container {
    display: grid;
    grid-template-columns: auto 400px;
    gap: 20px;
  }
  .orgs {
    display: flex;
    flex-direction: column;
    user-select: none;
    display: flex;
    height: fit-content;
    background-color: #FFF;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    overflow: hidden;
    padding: 15px 20px 25px 20px;
  }

  .info {
    min-height: 100px;
    height: fit-content;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    overflow: hidden;
    padding: 20px;
    background-color: #FFF;
  }
</style>