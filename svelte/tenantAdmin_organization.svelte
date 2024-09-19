<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import OrgTree from './_components/OrgTree.svelte';
  import { TenantAdminOrganization } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier.js';

  /** @typedef {import('./_components/types/organization').Org} Org */
  /** @type {Org[]} */
  let orgs = [/* orgs */];

  /**  @type {Org[]} */
  let REFORMAT_ORGS = [];

  function orgMaker(id) {
    /** @type {Org} */
    let orgFormatted = {
      children: [],
      id: 0,
      name: '',
      headTitle: '',
      orgType: 0,
      parentId: 0,
      tenantCode: '',
      createdAt: 0,
      createdBy: 0,
      updatedAt: 0,
      updatedBy: 0,
      deletedAt: 0,
      deletedBy: 0,
      restoredBy: 0
    }
    for (let i in orgs) {
      if (orgs[i].id == id) {
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
        orgFormatted.deletedBy = orgs[i].deletedBy
        orgFormatted.restoredBy = orgs[i].restoredBy
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

  onMount(() => {
	  REFORMAT_ORGS = reformatorgs();
	  console.log('REFORMAT_ORGS',REFORMAT_ORGS);
	  console.log('orgs',orgs);
  });

  let infoOrg;
  if (orgs && orgs.length) infoOrg = orgs[0];

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

  const infoEventHandler = (e) => infoOrg = e.detail.org;

  /** @type {Org|null} */
  let orgMoving = null;
  
  const orgMovingHandler = (e) => orgMoving = e.detail.org;
  
  const orgMovedHandler = async (e) => {
    /** @type {Org} */
    const parentOrg = e.detail.org;
    if (!orgMoving) return;

    if (parentOrg.orgType !== orgMoving.orgType-1) return console.log(`CANNOT MOVE ${orgMoving.id} TO ${parentOrg.id} - WRONG ORG TYPE`);
    if (parentOrg.id == orgMoving.parentId) return;

    const i = {
      cmd: 'move',
      org: {
        id: Number(orgMoving.id),
        parentId: Number(orgMoving.parentId)
      },
      moveToIdx: 0,
      toParentId: Number(parentOrg.id)
    }
    await TenantAdminOrganization( // @ts-ignore
      i, /** @returns {Promise<void>}*/
      function (/** @type {any} */o) {
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        console.log(o);
        orgs = o.orgs;
        REFORMAT_ORGS = [];
        REFORMAT_ORGS = reformatorgs();
      }
    );
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
            on:moving={orgMovingHandler}
            on:moved={orgMovedHandler}
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