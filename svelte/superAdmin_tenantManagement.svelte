<script>
  // @ts-nocheck
  import SideMenu from './_components/partials/SideMenu.svelte';
  import AdminSubMenu from './_components/AdminSubMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import TableView from './_components/TableView.svelte';
  import { SuperAdminTenantManagement } from './jsApi.GEN';
  import ModalForm from './_components/ModalForm.svelte';
  import { notifier } from './_components/notifier.js';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from 'svelte-icons-pack/fa/FaSolidPlusCircle';
  import { onMount } from 'svelte';

  let segments = {
    /* segments */
  };
  let fields = [
    /* fields */
  ];
  let tenants =
    [
      /* tenants */
    ] || [];
  let pager = {
    /* pager */
  };
  let user = {
    /* user */
  };

  onMount(() => {
    tenants = tenants || [];
    console.log('tenants', tenants);
  });

  // return true if got error
  function handleResponse(res) {
    console.log(res);
    if (res.error) {
      notifier.showError(res.error);
      return true;
    }
    if (res.tenants && res.tenants.length) tenants = res.tenants;
    if (res.pager && res.pager.page) pager = res.pager;
  }

  async function refreshTableView(pagerIn) {
    // console.log( 'pagerIn=',pagerIn );
    await SuperAdminTenantManagement(
      {
        pager: pagerIn,
        cmd: 'list',
      },
      function (res) {
        handleResponse(res);
      }
    );
  }

  let form = ModalForm; // for lookup

  async function editRow(id, row) {
    await SuperAdminTenantManagement(
      {
        tenant: { id },
        cmd: 'form',
      },
      function (res) {
        if (!handleResponse(res)) form.showModal(res.tenant);
      }
    );
  }

  function addRow() {
    form.showModal({ id: '' });
  }

  async function saveRow(action, row) {
    let tenant = { ...row };
    if (!tenant.id) tenant.id = '0';
    await SuperAdminTenantManagement(
      {
        tenant: tenant,
        cmd: action,
        pager: pager, // force refresh page, will be slow
      },
      function (res) {
        if (handleResponse(res)) {
          return form.setLoading(false); // has error
        }
        form.hideModal(); // success
      }
    );
  }
</script>

<div class="root_layout">
  <SideMenu access={segments} />
  <div class="root_container">
    <Navbar {user} />
    <div class="root_content">
      <AdminSubMenu />
      <div>
        <ModalForm {fields} rowType="Tenant" bind:this={form} onConfirm={saveRow}></ModalForm>
        <section class="tableview_container">
          <TableView {fields} bind:pager rows={tenants} onRefreshTableView={refreshTableView} onEditRow={editRow}>
            <button on:click={addRow} class="action_btn">
              <Icon size={17} color="#FFF" src={FaSolidPlusCircle} />
              <span>Add</span>
            </button>
          </TableView>
        </section>
      </div>
    </div>
    <Footer />
  </div>
</div>
