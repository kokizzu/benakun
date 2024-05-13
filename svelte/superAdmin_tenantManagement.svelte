<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import CgMathPlus from 'svelte-icons-pack/cg/CgMathPlus';
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
  let tenants   = /** @type any[][] */ ([/* tenants */]);

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
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
        pager = o.pager;
      }
    )
  }

  function AddRow() {
    console.log('AddRow');
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
      /** @returns {void} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          return;
        }

        console.log(o);
        tenants = o.tenants;
        pager = o.pager;
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
      /** @returns {void} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          return;
        }

        console.log(o);
        tenants = o.tenants;
        pager = o.pager;
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

      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnRefresh}
      {OnRestore}
      {OnDelete}
    >
      <button class="btn" on:click={AddRow}>
        <Icon color="var(--gray-007)" size="16" src={CgMathPlus}/>
      </button>
    </MasterTable>
  </div>
</MainLayout>

<style>
  .btn {
    border: none;
    background-color: var(--gray-001);
    border: 1px solid var(--gray-003);
		width: fit-content;
		padding: 10px 15px;
		border-radius: 8px;
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
		gap: 5px;
		cursor: pointer;
  }
</style>
