<script>
  import TableView from './_components/TableView.svelte';
  import { SuperAdminUserManagement } from './jsApi.GEN';
  import ModalForm from './_components/ModalForm.svelte';
  import { notifier } from './_components/notifier.js';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from 'svelte-icons-pack/fa/FaSolidPlusCircle';
    import MainLayout from './_layouts/mainLayout.svelte';

  let segments = {/* segments */};
  let fields = [/* fields */];
  let users = [/* users */];
  let pager = {/* pager */};
  let user = {/* user */};

  console.log('Fields:', fields);
  console.log('Users:', users);
  console.log('Pager:', pager);

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

<MainLayout>
  <div class="user_management">
    <ModalForm {fields} rowType="User" bind:this={form} onConfirm={saveRow}></ModalForm>
    <section class="tableview_container">
      <TableView {fields} bind:pager rows={users} onRefreshTableView={refreshTableView} onEditRow={editRow}>
        <button on:click={addRow} class="action_btn">
          <Icon size="17" color="#FFF" src={FaSolidPlusCircle} />
          <span>Add</span>
        </button>
      </TableView>
    </section>
  </div>
</MainLayout>

<style>
  .user_management {
    margin-top: 30px;
  }
</style>