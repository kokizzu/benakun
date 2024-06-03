<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import {
    SuperAdminTenantManagement,
    TenantAdminDashboard
  } from './jsApi.GEN';
  import { notifier } from './_components/notifier';
  import ReadOnlyTable from './_components/ReadOnlyTable.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { IoClose } from './node_modules/svelte-icons-pack/dist/io';

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

  /**
   * @typedef {Object} staff
   * @property {string} id
   * @property {string} email
   * @property {string} fullName
   * @property {string} role
   */
  let staffs        = /** @type staff[] */ ([]);
  let tenantCode    = '';
  let isShowInfo    = false;

  async function OnInfo(/** @type any[] */ row) {
    isShowInfo = true;
    if (tenantCode == row[1] && staffs.length > 0) return;

    tenantCode = row[1];
    const i = {
      tenantCode,
      cmd: 'form'
    }
    await TenantAdminDashboard( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminDashboardCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return;
        }

        staffs = o.staffsForm;
      }
    )
  }

  /**
   * @typedef {Object} staffsField
   * @property {string} key
   * @property {string} label
   */

  /** @type staffsField[] */
  const staffsFields = [
    {
      key: 'email',
      label: 'Email'
    }, {
      key: 'fullName',
      label: 'Full Name'
    }, {
      key: 'role',
      label: 'Role'
    }
  ];
</script>

{#if isShowInfo}
  <div class={`popup_container ${isShowInfo ? 'show' : ''}`}>
    <div class="popup">
      <header class="header">
        <h2>Staffs for {tenantCode}</h2>
        <button on:click={() => isShowInfo = false}>
          <Icon size="22" color="var(--red-005)" src={IoClose}/>
        </button>
      </header>
      <div class="forms">
        {#if staffs && staffs.length > 0}
          <ReadOnlyTable
            FIELDS={staffsFields}
            ROWS={staffs}
          />
        {/if}
      </div>
    </div>
  </div>
{/if}

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
      CAN_SHOW_INFO

      {OnInfo}
      {OnRefresh}
      {OnRestore}
      {OnDelete}
    />
  </div>
</MainLayout>

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .popup_container {
    display: none;
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 2000;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

  .popup_container.show {
    display: flex;
  }

	.popup_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 500px;
		display: flex;
		flex-direction: column;
	}

  .popup_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 10px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup_container .popup header h2 {
		margin: 0;
	}

	.popup_container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup_container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup_container .popup header button:active {
		background-color: #ef444430;
	}

	.popup_container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}
</style>