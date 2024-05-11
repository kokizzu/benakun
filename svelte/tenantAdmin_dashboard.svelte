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
  let staffs = /** @type Staff[] */ ([/* staffs */]);
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});

  console.log('segments =', segments);
  console.log('user =', user);
  console.log('staffs =', staffs);
  console.log('fields =', fields);
  console.log('pager =', pager);

  function GoToPage(page) {
    console.log('gotoPage', page);
  }

  function NextPage(page) {
    console.log('nextPage', page);
  }
  function PreviousPage(page) {
    console.log('previousPage', page);
  }
	function FirstPage() {
    console.log('firstPage');
  }
	function LastPage() {
    console.log('lastPage');
  }
</script>

<MainLayout>
  <div>
    <MasterTable
      URL='/tenantAdmin/dashboard'
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={staffs}
      PURPOSE='staff'
      CAN_EDIT_ROW={false}
      CAN_SEARCH_ROW={true}

      {GoToPage} {NextPage} {PreviousPage} {FirstPage} {LastPage}
    />
  </div>
</MainLayout>

<style>
</style>
