<script>
  // @ts-nocheck
  import SideMenu from './_components/partials/SideMenu.svelte';
  import AdminSubMenu from './_components/AdminSubMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import TableView from './_components/TableView.svelte';
  import { SuperAdminUserManagement } from './jsApi.GEN';
  import ModalForm from './_components/ModalForm.svelte';
  import { notifier } from './_components/notifier.js';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from 'svelte-icons-pack/fa/FaSolidPlusCircle';

  let segments = {/* segments */};
  let fields = [/* fields */];
  let users = [/* users */];
  let pager = {/* pager */};
  let user = {/* user */};

  // return true if got error
  function handleResponse(res) {
    console.log(res);
    if (res.error) {
      notifier.showError(res.error);
      return true;
    }
    if (res.users && res.users.length) users = res.users;
    if (res.pager && res.pager.page) pager = res.pager;
  }

  async function refreshTableView(pagerIn) {
    // console.log( 'pagerIn=',pagerIn );
    await SuperAdminUserManagement(
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
    await SuperAdminUserManagement(
      {
        user: { id },
        cmd: 'form',
      },
      function (res) {
        if (!handleResponse(res)) form.showModal(res.user);
      }
    );
  }

  function addRow() {
    form.showModal({ id: '' });
  }

  async function saveRow(action, row) {
    let user = { ...row };
    if (!user.id) user.id = '0';
    await SuperAdminUserManagement(
      {
        user: user,
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
  <div class="root_container">
    <SideMenu access={segments} />
    <div class="root_content">
      <Navbar {user} />
      <div class="content">
        <AdminSubMenu />
        <div class="user_management">
          <ModalForm {fields} rowType="User" bind:this={form} onConfirm={saveRow}></ModalForm>
          <section class="tableview_container">
            <TableView {fields} bind:pager rows={users} onRefreshTableView={refreshTableView} onEditRow={editRow}>
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
</div>

<style>
  .user_management {
    margin-top: 30px;
  }
</style>