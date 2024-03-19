<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import PopUpOrgChild from './_components/PopUpOrgChild.svelte';
  import { TenantAdminCreateOrganizationChild } from './jsApi.GEN';
  import { notifier } from './_components/notifier';
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

  onMount(() => {
    console.log('Segments:', segments);
    console.log('User:', user);
    console.log('Orgs:', orgs);
    REFORMAT_ORGS = reformatorgs();

    console.log('Reformatted ORGS:', REFORMAT_ORGS)
  })

  let popUpOrgChild, isSubmitAddOrgChild = false;
  let childName = '', headTitle = '', parentId = Number(orgs[0].id);
  async function submitAddOrgChild() {
    isSubmitAddOrgChild = true;
    if (childName === '') {
      isSubmitAddOrgChild = false;
      notifier.showWarning('coa name cannot be empty');
      return;
    }
    await TenantAdminCreateOrganizationChild(
      {
        name: childName,
        headTitle: headTitle,
        parentId: parentId,
      },
      // @ts-ignore
      function (o) {
        // @ts-ignore
        if (o.error) {
          popUpOrgChild.hide();
          isSubmitAddOrgChild = false;
          // @ts-ignore
          notifier.showError(o.error);
          // @ts-ignore
          console.log(o.error);
          return;
        }
        popUpOrgChild.hide();
        isSubmitAddOrgChild = false;
        console.log(o)
        notifier.showSuccess('Department child created');
      }
    );
  }
</script>

<PopUpOrgChild
  bind:this={popUpOrgChild}
  bind:isSubmitted={isSubmitAddOrgChild}
  bind:childName={childName}
  bind:headTitle={headTitle}
  onSubmit={submitAddOrgChild}
/>

<MainLayout>
  <div>
    {#if REFORMAT_ORGS && REFORMAT_ORGS.length}
      <div class="orgs">
        {#each REFORMAT_ORGS as org}
          <OrgTree {org} />
        {/each}
      </div>
    {/if}
  </div>
</MainLayout>

<style>
  .orgs {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
</style>