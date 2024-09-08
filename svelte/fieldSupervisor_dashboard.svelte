<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/transaction.js').Transaction} Transaction */

  let segments = /** @type Access */ ({/* segments */});
  // let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let user = /** @type User */ ({/* user */});
  let transaction = {/* transaction */}
  let transactions = /** @type any[][] */ ([/* transactions */]);
</script>

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:PAGER={pager}
      bind:MASTER_ROWS={transactions}
      
      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
    >
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>
