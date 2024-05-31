<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PoUpForms from './_components/PoUpForms.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminLocations } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/location.js').Location} Location */

  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let locations = /** @type any[][] */ ([/* locations */]);

  let isPopUpFormsReady = false;
  let popUpForms = null;
  onMount(() => {
    isPopUpFormsReady = true;
  });

  let isSubmitAddLocation = false;
  async function OnAddLocation(/** @type any[] */ payloads) {
    isSubmitAddLocation = true;
    /** @type Location */ //@ts-ignore
    const location = {
      tenantCode: user.tenantCode,
      name: payloads[1],
      country: payloads[2],
      stateProvice: payloads[3],
      cityRegency: payloads[4],
      subdistrict: payloads[5],
      village: payloads[6],
      rwBanjar: payloads[7],
      rtNeigb: payloads[8],
      address: payloads[9],
      lat: payloads[10],
      lng: payloads[11],
    }
    const i = {
      pager, location,
      cmd: 'upsert'
    }
    await TenantAdminLocations( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminLocationsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmitAddLocation = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        locations = o.locations;

        notifier.showSuccess('office location added')
        popUpForms.Reset();
      }
    );
    popUpForms.Hide();
  }
</script>

{#if isPopUpFormsReady}
  <PoUpForms
    bind:this={popUpForms}
    heading="Add office location"
    FIELDS={fields}
    bind:isSubmitted={isSubmitAddLocation}
    OnSubmit={OnAddLocation}
  />
{/if}

<MainLayout>
  <div>
    <MasterTable
      ACCESS={segments}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={locations}
     
      CAN_SHOW_INFO
      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
    >
      {#if user.tenantCode !== ''}
        <button
          class="action_btn"
          on:click={() => popUpForms.Show()}
          title="add account"
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
