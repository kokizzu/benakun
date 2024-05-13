<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { SuperAdminTenantManagement } from './jsApi.GEN';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */

  let segments  = /** @type Access */ ({/* segments */});
  let user      = /** @type User */ ({/* user */});
  let fields    = /** @type Field[] */ ([/* fields */]);
  let pager     = /** @type PagerOut */ ({/* pager */});

  // Contents
  let tenants   = /** @type any[][] */ ([/* tenants */]);

  console.log('segments =', segments);
  console.log('user =', user);
  console.log('tenants =', tenants);
  console.log('fields =', fields);
  console.log('pager =', pager);

  function GoToPage(/** @type number */ page) {
    console.log('gotoPage', page);
  }

  function NextPage(/** @type number */ page) {
    console.log('nextPage', page);
  }

  function PreviousPage(/** @type number */ page) {
    console.log('previousPage', page);
  }

	function FirstPage() {
    console.log('firstPage');
  }

	function LastPage() {
    console.log('lastPage');
  }

  async function OnRefreshTableView(/** @type PagerIn */ pagerIn) {
    const i = {
      pager: pagerIn,
      cmd: 'list'
    }
    await SuperAdminTenantManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').SuperAdminTenantManagementCallback} */
      /** @returns {void} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          return;
        }

        console.log(o);
        tenants = o.tenants;
      }
    )
  }
</script>

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={tenants}
      CAN_EDIT_ROW={true}
      CAN_SEARCH_ROW={true}

      {GoToPage} {NextPage} {PreviousPage} {FirstPage} {LastPage}
      {OnRefreshTableView}
    />
  </div>
</MainLayout>

<style>
</style>
