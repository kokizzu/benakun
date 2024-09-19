<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { SuperAdminUserManagement } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier.js';
  import { onMount } from 'svelte';
  import PopUpAddUser from './_components/PopUpAddUser.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */

  let segments  = /** @type Access */ ({/* segments */});
  let user      = /** @type User */ ({/* user */});
  let fields    = /** @type Field[] */ ([/* fields */]);
  let pager     = /** @type PagerOut */ ({/* pager */});
  let users   = /** @type any[][] */ ([/* users */]);

  console.log('SEGMENTS', segments);
  console.log('USER', user);
  console.log('FIELDS', fields);
  console.log('PAGER', pager);
  console.log('USERS', users);

  let isPopUpAddUserReady = false;
  let popUpAddUser = null;
  onMount(() => {
    isPopUpAddUserReady = true;
  })

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await SuperAdminUserManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').SuperAdminUserManagementCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        users = o.users;
        user = o.user;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const i = {
      pager,
      user: {
        id: row[0]
      },
      cmd: 'restore'
    };
    await SuperAdminUserManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').SuperAdminUserManagementCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return;
        }

        pager = o.pager;
        users = o.users;
        user = o.user;

        notifier.showSuccess('user '+user.email+' restored');
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = {
      pager,
      user: {
        id: row[0]
      },
      cmd: 'delete'
    };
    await SuperAdminUserManagement( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').SuperAdminUserManagementCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        users = o.users;
        pager = o.pager;
        user = o.user;

        notifier.showSuccess('user '+user.email+' deleted');
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    /** @type User */ //@ts-ignore
    const u = {
      id,
      email: payloads[2],
      fullName: payloads[3],
      role: payloads[4],
    }
    const i = {
      pager, user: u,
      cmd: 'upsert'
    };
    await SuperAdminUserManagement( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').SuperAdminUserManagementCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        users = o.users;
        user = o.user;

        notifier.showSuccess('user '+user.email+' edited')
      }
    );
  }

  async function onAddUser(/** @type Object */ payload) {
    const i = {
      pager,
      tenantAdmin: payload.tenantAdmin,
      user: payload.user,
      cmd: 'upsert'
    };
    await SuperAdminUserManagement( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').SuperAdminUserManagementCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        pager = o.pager;
        users = o.users;
        user = o.user;

        notifier.showSuccess('user '+user.email+' created')
      }
    );
    popUpAddUser.Hide();
  }

  const roles = {
    'user': 'User',
    'superAdmin': 'Super Admin',
    'tenantAdmin': 'Tenant Admin',
    'dataEntry': 'Data Entry',
    'reportViewer': 'Report Viewer',
    'fieldSupervisor': 'Field Supervisor'
  }
</script>

{#if isPopUpAddUserReady}
  <PopUpAddUser
    bind:this={popUpAddUser}
    heading="Add user"
    OnSubmit={onAddUser}
  />
{/if}

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={users}
      REFS={{
        'role': roles
      }}

      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnDelete}
      {OnRestore}
      {OnRefresh}
      {OnEdit}
    >
    <button
      class="action_btn"
      on:click={() => popUpAddUser.Show()}
      title="add account"
    >
      <Icon
        color="var(--gray-007)"
        size="16"
        src={RiSystemAddBoxLine}
      />
    </button>
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>