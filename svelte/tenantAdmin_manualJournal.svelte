<script>
  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/product.js').Product} Product */
  /** @typedef {import('./_components/types/coa.js').CoA} CoA */

  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PopUpManualJournal from './_components/PopUpManualJournal.svelte';
  import { onMount } from 'svelte';

  let segments = /** @type Access */ ({/* segments */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let user = /** @type User */ ({/* user */});
  let transactionJournals = /** @type any[][] */ ([/* transactionJournals */]);
  let coas = /** @type CoA[] */ ([/* coas */]);

  let isPopUpManualJournalReady = false;
  let popUpManualJournal = null;
  onMount(() => {
    isPopUpManualJournalReady = true;
  })
</script>

<PopUpManualJournal
  bind:this={popUpManualJournal}
  {coas}
/>

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      REFS={{
        'coaId': coas
      }}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={transactionJournals}
      
      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
    >
      {#if user.tenantCode !== ''}
        <button
          class="action_btn"
          title="add journal"
          on:click={() => popUpManualJournal.Show()}
        >
          <Icon
            color="var(--gray-007)"
            size="16"
            src={RiSystemAddBoxLine}
          />
        </button>
      {/if}
    </MasterTable>
  </div>
</MainLayout>

<style>
</style>
