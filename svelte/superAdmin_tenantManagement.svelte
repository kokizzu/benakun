<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { SuperAdminTenantManagement } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/tenant.js').Tenant} Tenant */

  let segments  = /** @type Access */ ({/* segments */});
  let user      = /** @type User */ ({/* user */});
  let fields    = /** @type Field[] */ ([/* fields */]);
  let pager     = /** @type PagerOut */ ({/* pager */});
  let tenants   = /** @type any[][] */ ([/* tenants */]);
  let tenant    = /** @type Tenant */ ({/* tenant */});

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await SuperAdminTenantManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').SuperAdminTenantManagementCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          return;
        }

        tenants = o.tenants;
        pager = o.pager;
        tenant = o.tenant;
      }
    )
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = {
      pager,
      tenant: {
        id: row[0]
      },
      cmd: 'restore'
    };
    await SuperAdminTenantManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').SuperAdminTenantManagementCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return;
        }

        tenants = o.tenants;
        pager = o.pager;
        tenant = o.tenant;

        notifier.showSuccess(tenant.tenantCode + ' restored');
      }
    )
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = {
      pager,
      tenant: {
        id: row[0]
      },
      cmd: 'delete'
    };
    await SuperAdminTenantManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').SuperAdminTenantManagementCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return;
        }

        tenants = o.tenants;
        pager = o.pager;
        tenant = o.tenant;

        notifier.showSuccess(tenant.tenantCode + ' deleted');
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

      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
      CAN_EDIT_ROW={false}

      {OnRefresh}
      {OnRestore}
      {OnDelete}
    />
  </div>
</MainLayout>

<style>
</style>