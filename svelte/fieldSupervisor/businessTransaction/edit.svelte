<script>
	/** @typedef {import('../../_components/types/access.js').Access} Access */
  /** @typedef {import('../../_components/types/user.js').User} User */
  /** @typedef {import('../../_components/types/transaction.js').BusinessTransaction} BusinessTransaction */
  /** @typedef {import('../../_components/types/transaction.js').TransactionJournal} TransactionJournal */

  import MainLayout from '../../_layouts/mainLayout.svelte';

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let transaction = /** @type BusinessTransaction */ ({/* transaction */});
  let transactionJournals = /** @type TransactionJournal[] */ ([/* transactionJournals */]);

  console.log('BUSINESS TRANSACTION', transaction);
  console.log('TRANSACTION JOURNALS', transactionJournals);
</script>

<MainLayout>
  <h3>Edit business transaction</h3>
  <h3>Edit transaction journal by transaction template id</h3>
</MainLayout>

<style>
</style>

