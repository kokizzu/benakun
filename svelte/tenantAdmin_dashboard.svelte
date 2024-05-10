<script>
  import { onMount } from 'svelte';
  import { notifier } from './_components/notifier.js';
  import ConfirmPopUp from './_components/ConfirmPopUp.svelte';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/user.js').Staff} Staff */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let staffs = /** @type Staff[] */ ([/* staffs */]);
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});

  console.log('segments =', segments);
  console.log('user =', user);
  console.log('staffs =', staffs);
  console.log('fields =', fields);
  console.log('pager =', pager);

  let staffStatus = [];
  let tableReady = false;

  let confirmQuestion = '', confirmVisible = false, confirmAction = function(){};
  let confirm_popup;

  onMount(()=> {
    if (staffs && staffs.length) {
      for (let i = 0; i < staffs.length; i++) {
        const invSts = staffs[i]['invitationState'].split(' ');
        if (invSts.length > 1) {
          invSts.forEach((invSt) => {
            const st = invSt.split(':');
            if (st[1] === user.tenantCode) {
              staffStatus.push(st[2]);
            }
          })
        } else {
          const st = invSts[0].split(':');
          staffStatus.push(st[2]);
        }
      }
      tableReady = true;
    }
    confirm_popup = ConfirmPopUp;
  });

  async function terminateStaff(email) {
    // @ts-ignore
    await TenantAdminTerminateStaff({email}, function (o) {
      // @ts-ignore
      if (o.error) return notifier.showError(o.error);
      notifier.showSuccess(o.message);
      setTimeout(() => window.location.reload(), 1200);
    })
  }

  function confirmTerminateStaff(email) {
    const q = 'Are you sure you want to terminate this staff?';
    if (!confirmVisible || !confirmVisible) {
      confirmVisible = !confirmVisible;
      confirmQuestion = q;
      confirmAction = () => terminateStaff(email);
    } else {
      // reset if there
      confirmQuestion = '', confirmVisible = false, confirmAction = function(){};
      confirmVisible = !confirmVisible;
      confirmQuestion = q;
      confirmAction = () => terminateStaff(email);
    }
  }
</script>

<ConfirmPopUp action={confirmAction} question={confirmQuestion} visible={confirmVisible} />

<MainLayout>
  <div>
    <MasterTable
      URL='/tenantAdmin/dashboard'
      ACCESS={segments}
      FIELDS={fields}
      PAGER={pager}
      MASTER_ROWS={staffs}
      PURPOSE='staff'
    />
  </div>
</MainLayout>

<style>
</style>
