<script>
  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/master.js').ExtendedAction} ExtendedAction */
  /** @typedef {import('./_components/types/transaction.js').BusinessTransaction} BusinessTransaction */

  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import { RiDesignBallPenLine } from './node_modules/svelte-icons-pack/dist/ri';

  let segments = /** @type Access */ ({/* segments */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let user = /** @type User */ ({/* user */});
  let transaction = /** @type BusinessTransaction */ ({/* transaction */});
  let transactions = /** @type any[][] */ ([/* transactions */]);
  let transactionTemplates = /** @type Record<string|number, string> */ ({/* transactionTemplates */});

   /** @type {ExtendedAction[]} */
   const LINKS = [
    {
      icon: RiDesignBallPenLine,
      isTargetBlank: false,
      link: (/** @type any */ row) => `/fieldSupervisor/businessTransaction/${row[0]}`,
      tooltip: 'Edit',
    }
  ]
</script>

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:PAGER={pager}
      bind:FIELDS={fields}
      bind:MASTER_ROWS={transactions}
      REFS={{
        'transactionTemplateId': transactionTemplates
      }}
      
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
      CAN_OPEN_LINK
      IS_CUSTOM_EDIT

      LINKS={LINKS}
    >
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>
