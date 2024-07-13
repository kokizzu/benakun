<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterLog from './_components/MasterLog.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/admin.js').AccessLog} AccessLog */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let logs = /** @type AccessLog[] */ ([/* logs */]);

  console.log('Logs: ', logs);
</script>

<MainLayout>
  <div class="superadmin_accesslog">
    <MasterLog {logs} />
  </div>
</MainLayout>

<style>
</style>