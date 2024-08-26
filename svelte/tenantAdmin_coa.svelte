<script>
  /** @typedef {import('./_components/types/user').User} User */
  /** @typedef {import('./_components/types/access').Access} Access */
  /** @typedef {import('./_components/types/coa').CoA} CoA */
  /** @typedef {import('./_components/types/tenant.js').Tenant} Tenant */

  import { onMount } from 'svelte';
  import { TenantAdminCoa, TenantAdminSyncCoa } from './jsApi.GEN.js';
  import { notifier } from './_components/notifier.js';
  import CoaTree from './_components/CoaTree.svelte';
  import PopUpCoA from './_components/PopUpCoa.svelte';
  import MainLayout from './_layouts/mainLayout.svelte';
  import InputCustom from './_components/InputCustom.svelte';

  let segments  = /** @type Access */ ({/* segments */});
  let user      = /** @type User */   ({/* user */});
  let coas      = /** @type CoA[] */  ([/* coas */]);
  let tenant    = /** @type Tenant */   ({/* tenant */});
  let coasChoices = /** @type Record<string|number, string> */   ({/* coaChoices */});
  
  let REFORMAT_COAS = /** @type CoA[] */ ([]);

  let coaVisited = {};
  function coaMaker(/** @type {string|number} */ id) {
    if(coaVisited[id]) return null;
    coaVisited[id] = true;
    /** @type CoA */
    let coaFormatted = {
      id: 0,
      name: '',
      parentId: 0,
      children: [],
      createdAt: 0,
      createdBy: '',
      updatedAt: 0,
      updatedBy: '',
      deletedAt: 0,
      tenantCode: '',
      label: '',
      deletedBy: '',
      restoredBy: '',
      tenantLabel: '',
    }
    for (let i in coas) {
      if (coas[i].id == Number(id)) { // @ts-ignore
        const children = /** @type number[] */ (coas[i].children);
        if (children && children.length) {
          for (let j in children) {
            const childId = children[j]
            const child = coaMaker(childId); // @ts-ignore
            if(!child) continue;
            coaFormatted.children = [...coaFormatted.children, child];
          }
        }
        coaFormatted.id = coas[i].id
        coaFormatted.tenantCode = coas[i].tenantCode
        coaFormatted.label = coas[i].label
        coaFormatted.name = coas[i].name
        coaFormatted.parentId = coas[i].parentId
        coaFormatted.createdAt = coas[i].createdAt
        coaFormatted.createdBy = coas[i].createdBy
        coaFormatted.updatedAt = coas[i].updatedAt
        coaFormatted.updatedBy = coas[i].updatedBy
        coaFormatted.deletedAt = coas[i].deletedAt

        switch (coas[i].id) {
          case tenant.productsCoaId:
            coaFormatted.tenantLabel = 'Products'
            break;
          case tenant.suppliersCoaId:
            coaFormatted.tenantLabel = 'Suppliers'
            break;
          case tenant.customersCoaId:
            coaFormatted.tenantLabel = 'Customers'
            break;
          case tenant.staffsCoaId:
            coaFormatted.tenantLabel = 'Staffs'
            break;
          case tenant.banksCoaId:
            coaFormatted.tenantLabel = 'Banks'
            break;
          case tenant.customerReceivablesCoaId:
            coaFormatted.tenantLabel = 'Customers Receivables'
            break;
        }
      }
    }

    return coaFormatted;
  }

  function reformatCoas() {
    coaVisited = {}
    let toCoas = /** @type CoA[] */ ([]);

    if (coas && coas.length) {
      for (let i in coas) {
        const id = coas[i].id;
        const coa = coaMaker(id);
        if (!coa) continue;
        toCoas = [...toCoas, coa];
      }
    }

    if (toCoas && toCoas.length) {
      toCoas = toCoas.filter(obj => obj.parentId <= 0);
    }

    return toCoas;
  }

  onMount(() => {
    REFORMAT_COAS = reformatCoas();
  });

  function updateEventHandler(/** @type {CustomEvent<{ coas: CoA[] }>} */ e) {
    coas = e.detail.coas;
    REFORMAT_COAS = [];
    REFORMAT_COAS = reformatCoas();
  }

  let popUpCoa;
  let id        = 0;
  let parentId  = 0;
  let name      = ''
  let label     = '';
  let isSubmittedCoA = false;

  async function submitUpsertCoa() {
    isSubmittedCoA = true;
    /** @type CoA */ //@ts-ignore
    const coa = { id, parentId, name, label }
    const i = {
      cmd: 'upsert',
      coa
    }
    await TenantAdminCoa( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminCoaCallback} */
      /** @returns {Promise<void>} */
      function(/** @type any */ o) {
        isSubmittedCoA = false;
        id = 0, parentId = 0, name = '', label = '';
        popUpCoa.hide();
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        
        coas = o.coas;
        REFORMAT_COAS = [];
        REFORMAT_COAS = reformatCoas();

        notifier.showSuccess('coa updated');
      }
    )
  }

  function toggleShowPopUpCoa() {
    id        = 0;
    parentId  = 0;
    name      = '';
    label     = '';
    popUpCoa.show();
  }

  
  let coaMoving = /** @type {CoA|null} */ (null);

  async function coaMovingHandler(/** @type {CustomEvent<{ coa: CoA }>} */ e) {
    coaMoving = e.detail.coa;
  }
  
  async function coaMovedHandler(/** @type {CustomEvent<{ coa: CoA }>} */ e) {
    const parentCoa = /** @type {CoA} */ e.detail.coa;
    if (!coaMoving) return;
    if (parentCoa.id == coaMoving.parentId) return;
    if (parentCoa.id === coaMoving.id) return;

    const i = {
      cmd: 'move',
      coa: {
        id: Number(coaMoving.id),
        parentId: Number(coaMoving.parentId)
      },
      moveToIdx: 0,
      toParentId: Number(parentCoa.id)
    }
    await TenantAdminCoa( // @ts-ignore
      i, /** @returns {Promise<void>}*/
      function (/** @type {any} */o) {
        if (o.error) {
          notifier.showError(o.error);
          console.log(o);
          return;
        }
        coas = o.coas;
        REFORMAT_COAS = [];
        REFORMAT_COAS = reformatCoas();

        notifier.showSuccess(`coa "` + coasChoices[coaMoving.id] + `" moved to coa "` + coasChoices[parentCoa.id] + `"`);
      }
    );
  }

  let productsCoaId = tenant.productsCoaId || 0;
  let suppliersCoaId = tenant.suppliersCoaId || 0;
  let customersCoaId = tenant.customersCoaId || 0;
  let customerReceivablesCoaId = tenant.customerReceivablesCoaId || 0;
  let staffsCoaId = tenant.staffsCoaId || 0;
  let banksCoaId = tenant.banksCoaId || 0;

  let isSubmitSyncCoA = false;
  async function SubmitSyncCoA() {
    isSubmitSyncCoA = true;
    const i = {
      tenant: {
        productsCoaId,
        suppliersCoaId,
        customersCoaId,
        customerReceivablesCoaId,
        staffsCoaId,
        banksCoaId
      }
    }
    await TenantAdminSyncCoa( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').TenantAdminSyncCoaCallback} */
      function(/** @type any */ o) {
        isSubmitSyncCoA = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        tenant = o.tenant
        productsCoaId = tenant.productsCoaId;
        suppliersCoaId = tenant.suppliersCoaId;
        customersCoaId = tenant.customersCoaId;
        staffsCoaId = tenant.staffsCoaId;
        banksCoaId = tenant.banksCoaId;
        customerReceivablesCoaId = tenant.customerReceivablesCoaId;

        REFORMAT_COAS = [];
        REFORMAT_COAS = reformatCoas();

        notifier.showSuccess('coa synced');
      }
    )
  }
</script>

<PopUpCoA
  bind:this={popUpCoa}
  bind:name
  bind:label
  bind:isSubmitted={isSubmittedCoA}
  onSubmit={submitUpsertCoa}
/>

<MainLayout>
  <div class="coa_levels shadow">
    <div class="tenants_sync_coa">
      <div class="inputs_container">
        <InputCustom
          type="combobox"
          label="Products CoA"
          id="productCoaId"
          values={coasChoices}
          bind:value={productsCoaId}
          isObject
        />
        <InputCustom
          type="combobox"
          label="Suppliers CoA"
          id="suppliersCoaId"
          values={coasChoices}
          bind:value={suppliersCoaId}
          isObject
        />
        <InputCustom
          type="combobox"
          label="Customers CoA"
          id="customersCoaId"
          values={coasChoices}
          bind:value={customersCoaId}
          isObject
        />
        <InputCustom
          type="combobox"
          label="Customers Receivables CoA"
          id="customersReceivablesCoaId"
          values={coasChoices}
          bind:value={customerReceivablesCoaId}
          isObject
        />
        <InputCustom
          type="combobox"
          label="Staffs CoA"
          id="staffsCoaId"
          values={coasChoices}
          bind:value={staffsCoaId}
          isObject
        />
        <InputCustom
          type="combobox"
          label="Banks CoA"
          id="banksCoaId"
          values={coasChoices}
          bind:value={banksCoaId}
          isObject
        />
      </div>
      <button
        disabled={isSubmitSyncCoA}
        class="sync_btn"
        on:click={SubmitSyncCoA}
      >Sync CoA</button>
    </div>
    <div class="header_options">
      <button
        class="add"
        on:click={toggleShowPopUpCoa}
      >Add new CoA</button>
    </div>
    {#each (REFORMAT_COAS || []) as c, idxParent (c.id)}
      <div class="coa">
        <CoaTree
          coa={c}
          num={`${idxParent+1}`}
          parentNum={idxParent+1}
          isRootCoA
          indent={0}
          on:update={updateEventHandler}
          on:moving={coaMovingHandler}
          on:moved={coaMovedHandler}
        />
      </div>
    {/each}
  </div>
</MainLayout>

<style>
  .coa_levels {
    display: grid;
    grid-template-columns: 1fr;
    width: 100%;
    background-color: #FFF;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    overflow: hidden;
    padding: 15px 15px 20px;
  }

  .tenants_sync_coa {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .tenants_sync_coa .inputs_container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
  }

  .coa_levels .header_options {
    display: flex;
    flex-direction: row-reverse;
    padding-bottom: 10px;
  }

  .coa_levels .tenants_sync_coa .sync_btn,
  .coa_levels .header_options .add {
    border: none;
    padding: 10px 20px;
    border-radius: 8px;
    background-color: var(--blue-006);
    color: #FFF;
    font-size: var(--font-lg);
    cursor: pointer;
    width: fit-content;
  }

  .coa_levels .tenants_sync_coa .sync_btn:hover,
  .coa_levels .header_options .add:hover {
    background-color: var(--blue-005);
  }

  .coa_levels .tenants_sync_coa .sync_btn:disabled,
  .coa_levels .header_options .add:disabled {
    background-color: var(--gray-004);
    cursor: not-allowed;
    color: var(--gray-008);
  }

  .coa_levels .coa {
    display: flex;
    flex-direction: column;
    overflow: hidden;
    gap: 10px;
  }
</style>