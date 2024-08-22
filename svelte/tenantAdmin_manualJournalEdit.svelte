<script>
  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/product.js').Product} Product */
  /** @typedef {import('./_components/types/coa.js').CoA} CoA */
  /** @typedef {import('./_components/types/transaction.js').TransactionTemplate} TransactionTemplate */
  /** @typedef {import('./_components/types/transaction.js').TransactionJournal} TransactionJournal */
  /** @typedef {import('./_components/types/transaction.js').DetailObjectTransaction} DetailObjectTransaction */
  /** @typedef {import('./_components/types/transaction.js').TransactionTplDetail} TransactionTplDetail */
  /** @typedef {import('./_components/types/organization.js').Org} Org */

  import MainLayout from './_layouts/mainLayout.svelte';

  let segments              = /** @type Access */ ({/* segments */});
  let user                  = /** @type User */ ({/* user */});
  let transactionJournal    = /** @type TransactionJournal */ ({/* transactionJournal */});

  console.log(transactionJournal);
</script>

<MainLayout>
  <div>
    <h1>Edit journal</h1>
  </div>
</MainLayout>

