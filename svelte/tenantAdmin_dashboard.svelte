<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/user.js').Staff} Staff */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});

  let staffs = /** @type any[][] */ ([/* staffs */]);

  console.log('segments =', segments);
  console.log('user =', user);
  console.log('staffs =', staffs);
  console.log('fields =', fields);
  console.log('pager =', pager);
</script>

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={staffs}
      CAN_EDIT_ROW={false}
      CAN_SEARCH_ROW={true}
    />
  </div>
</MainLayout>

<style>
</style>
