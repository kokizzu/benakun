<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */

  let segments = /** @type Access */ ({/* segments */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let user = /** @type User */ ({/* user */});
  let transactionTemplates = /** @type any[][] */ ([/* transactionTemplates */]);


</script>

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={transactionTemplates}
    />
  </div>
</MainLayout>

<style>
</style>
