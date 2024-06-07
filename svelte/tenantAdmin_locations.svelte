<script>
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import { IoClose } from './node_modules/svelte-icons-pack/dist/io';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import PoUpForms from './_components/PoUpForms.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminLocations } from './jsApi.GEN';
  import { notifier } from './_components/notifier';
  import Map from './_components/Map.svelte';

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


  let location = /** @type location */ ({});

  let isPopUpFormsReady = false;
  let popUpForms = null;
  onMount(() => {
    isPopUpFormsReady = true;
    if (locations && locations.length > 0) location = locations[0];
  });

  async function OnRefresh(/** @type PagerIn */ pagerIn) {
    const i = { pager: pagerIn, cmd: 'list' }
    await TenantAdminLocations( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminLocationsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        locations = o.locations;
      }
    );
  }

  async function OnEdit(/** @type any */ id, /** @type any[]*/ payloads) {
    /** @type Location */ // @ts-ignore
    const loc = {
      id: payloads[0],
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
      pager,
      location: loc,
      cmd: 'upsert'
    }
    await TenantAdminLocations( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminLocationsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        locations = o.locations;

        notifier.showSuccess('location updated');
      }
    );
  }

  async function OnDelete(/** @type any[] */ row) {
    const i = {
      pager,
      location: {
        id: row[0]
      },
      cmd: 'delete'
    }
    await TenantAdminLocations( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminLocationsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        locations = o.locations;

        notifier.showSuccess('location deleted');
      }
    );
  }

  async function OnRestore(/** @type any */ id) {
    const i = {
      pager,
      location: {
        id: id
      },
      cmd: 'restore'
    }
    await TenantAdminLocations( // @ts-ignore
      i, /** @type {import('../jsApi.GEN').TenantAdminLocationsCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        pager = o.pager;
        locations = o.locations;

        notifier.showSuccess('location restored');
      }
    );
  }

  let isSubmitAddLocation = false;
  async function OnAddLocation(/** @type any[] */ payloads) {
    isSubmitAddLocation = true;
    /** @type Location */ //@ts-ignore
    const loc = {
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
      pager,
      location: loc,
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
        location = o.location;

        notifier.showSuccess('office location added')
        popUpForms.Reset();
      }
    );
    popUpForms.Hide();
  }

  let isShowInfoLocation = false;
  let headingInfo = 'Location for '+location.name;
  let mapElm;
  let isMapReady;
  let Coord = {
    lat: location.lat,
    lng: location.lng,
    zoom: 8
  };

  async function OnInfo(/** @type any[] */ row) {
    Coord.lat = row[10];
    Coord.lng = row[11];
    headingInfo = 'Location for '+row[1];
    isMapReady = true;
    isShowInfoLocation = true;
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

{#if isShowInfoLocation}
  <div class={`popup_container ${isShowInfoLocation ? 'show' : ''}`}>
    <div class="popup">
      <header class="header">
        <h2>{headingInfo}</h2>
        <button on:click={() => isShowInfoLocation = false}>
          <Icon size="22" color="var(--red-005)" src={IoClose}/>
        </button>
      </header>
      <div class="forms">
        <div class="map">
          {#if isMapReady}
            <Map
              bind:this={mapElm}
              {Coord}
            />
          {/if}
        </div>
      </div>
    </div>
  </div>
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

      {OnInfo}
      {OnEdit}
      {OnRefresh}
      {OnDelete}
      {OnRestore}
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
  .popup_container {
    display: none;
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 2000;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

  .popup_container.show {
    display: flex;
  }

	.popup_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 700px;
		display: flex;
		flex-direction: column;
	}

  .popup_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 10px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup_container .popup header h2 {
		margin: 0;
	}

	.popup_container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup_container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup_container .popup header button:active {
		background-color: #ef444430;
	}

	.popup_container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

  .popup_container .popup .forms .map {
    width: 100%;
    height: 400px;
    border-radius: 8px;
    overflow: hidden;
  }
</style>