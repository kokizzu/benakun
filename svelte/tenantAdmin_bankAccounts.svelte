<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let accounts = /** @type any[][] */ ([/* accounts */]);

  console.log('Segments:', segments);
  console.log('User:', user);
  console.log('Fields:', fields);
  console.log('Pager:', pager);
  console.log('Accounts:', accounts);
</script>

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={accounts}
      
      CAN_EDIT_ROW={false}
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
    />
  </div>
</MainLayout>

<style>
</style>
