<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiBusinessMailAddLine } from './node_modules/svelte-icons-pack/dist/ri';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PopUpInviteUser from './_components/PopUpInviteUser.svelte';
  import { TenantAdminDashboard } from './jsApi.GEN';
  import { notifier } from './_components/notifier';
  import { onMount } from 'svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let staffs = /** @type any[][] */ ([/* staffs */]);

  // Binding component PopUpInviteUser.svelte
  let popUpInviteUser = null;
  // For readiness of component PopUpInviteUser.svelte, prevent race condition
	let isPopUpInviteUserReady = false;

  onMount(() => {
    isPopUpInviteUserReady = true;
  })

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' };
    await TenantAdminDashboard( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminDashboardCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        console.log(o);
        staffs = o.staffs;
        pager = o.pager;
      }
    );
  }

  async function OnRestore(/** @type any[] */ row) {
    const email = row[1];
    const i = {
      pager,
      staffEmail: email,
      cmd: 'restore'
    };
    await TenantAdminDashboard( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminDashboardCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return;
        }

        console.log(o);
        notifier.showSuccess('user ' + email + ' restored');
        staffs = o.staffs;
        pager = o.pager;
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const email = row[1];
    const i = {
      pager,
      staffEmail: email,
      cmd: 'delete'
    };
    await TenantAdminDashboard( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminDashboardCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        console.log(o);
        notifier.showSuccess('user ' + email + ' terminated');
        staffs = o.staffs;
        pager = o.pager;
      }
    );
  }

  // User email to invite
  let email = /** @type string */ ('');
  // User role to invite
  let role = /** @type string */ ('');
  // State for loading if hit ajax
  let isSubmitted = false;

  // Submit invite user
  async function submitInviteUser() {
    isSubmitted = true;
    popUpInviteUser.Hide();
    const i = {
      pager,
      staffEmail: email,
      staffRole: role,
      cmd: 'upsert'
    };
    await TenantAdminDashboard( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminDashboardCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        email = '';
        isSubmitted = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        console.log(o);
        notifier.showSuccess('user ' + email + ' invited');
        staffs = o.staffs;
        pager = o.pager;
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    const i = {
      pager,
      staffEmail: payloads[1],
      role: payloads[3],
      isEdit: true,
      cmd: 'upsert'
    };
    await TenantAdminDashboard( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminDashboardCallback} */
      /** @returns {Promise<any>} */
      function(/** @type any */ o) {
        email = '';
        isSubmitted = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        console.log(o);
        notifier.showSuccess('user ' + email + ' edited');
        staffs = o.staffs;
        pager = o.pager;
      }
    );
  }
</script>

{#if isPopUpInviteUserReady}
  <PopUpInviteUser
    bind:this={popUpInviteUser}
    bind:email
    bind:role
    bind:isSubmitted

    onSubmit={submitInviteUser}
  />
{/if}

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={staffs}
      
      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW

      {OnRefresh}
      {OnRestore}
      {OnDelete}
      {OnEdit}
    >
      {#if user.tenantCode !== ''}
        <button
          class="action_btn"
          on:click={() => popUpInviteUser.Show()}
          title="invite user"
        >
          <Icon
            color="var(--gray-007)"
            size="16"
            src={RiBusinessMailAddLine}
          />
        </button>
      {/if}
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>
