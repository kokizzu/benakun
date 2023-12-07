<script>
  // @ts-nocheck
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import CgTrash from "svelte-icons-pack/cg/CgTrash";
  import HiOutlinePencilAlt from "svelte-icons-pack/hi/HiOutlinePencilAlt";
  import { onMount } from 'svelte';

  let segments = {/* segments */};
  let user = {/* user */};
  let staffs = [/* staffs */];

  let staffStatus = [];
  let tableReady = false;

  onMount(()=> {
    console.log('Staff =', staffs)
    if (staffs && staffs.length) {
      for (let i = 0; i < staffs.length; i++) {
        const invSts = staffs[i]['invitationState'].split(' ');
        if (invSts.length > 1) {
          invSts.forEach((invSt) => {
            const st = invSt.split(':');
            if (st[1] === user.tenantCode) {
              staffStatus.push(st[2]);
            }
          })
        } else {
          const st = invSts[0].split(':');
          staffStatus.push(st[2]);
        }
      }
      tableReady = true;
    }
  })
</script>

<div class="root_layout">
  <div class="root_container">
    <SideMenu access={segments} />
    <div class="root_content">
      <Navbar {user} />
      <div class="content">
        <div class="staff_table">
          <table>
            <thead>
              <tr>
                <th class="primary" scope="col">User ID</th>
                <th scope="col">Email</th>
                <th scope="col">Full Name</th>
                <th scope="col">Role</th>
                <th scope="col">Status</th>
                <th scope="col">Action</th>
              </tr>
            </thead>
            <tbody>
              {#if tableReady}
              {#if staffs && staffs.length}
                {#each staffs as staff, idx}
                  <tr>
                    <td class="text-center">{staff.id}</td>
                    <td>{staff.email}</td>
                    <td>{staff.fullName || '--'}</td>
                    <td class="text-capitalize">{staff.role}</td>
                    <td>{staffStatus[idx]}</td>
                    <td>
                      <div class="action_btns">
                      <button class="btn delete_btn">
                        <Icon size={15} src={CgTrash} />
                      </button>
                      <button class="btn edit_btn">
                        <Icon size={15} src={HiOutlinePencilAlt} />
                      </button>
                      </div>
                    </td>
                  </tr>
                {/each}
              {/if}
              {/if}
            </tbody>
            <tfoot>
              <tr>
                {#if staffs && staffs.length}
                <td colspan="6"><i>Total Staff: {staffs.length}</i></td>
                {:else}
                  <td colspan="6"><i>Total Staff: 0</i></td>
                {/if}
              </tr>
            </tfoot>
          </table>
        </div>
      </div>
      <Footer />
    </div>
  </div>
</div>

<style>
  .staff_table {
    width: fit-content;
    height: fit-content;
    border-radius: 10px;
    overflow: hidden;
  }

  .staff_table .action_btns {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 0;
  }

  .staff_table .btn {
    padding: 6px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: transparent;
    border: none;
    cursor: pointer;
    border-radius: 9999px;
    color: var(--gray-008);
  }

  .staff_table .btn:hover {
    background-color: var(--gray-002);
  }

  :global(.staff_table .delete_btn:hover svg) {
    color: var(--red-002);
  }

  :global(.staff_table .edit_btn:hover svg path) {
    stroke: var(--yellow-002);
  }

  table {
    text-align: left;
    position: relative;
    border-collapse: collapse;
  }

  .staff_table table thead tr th {
    border: 1px solid var(--gray-008);
  }
  
  td, th {
    border: 1px solid var(--gray-002);
    padding: 10px;
  }
  th {
    background: var(--gray-009);
    color: white;
    border-radius: 0;
    position: sticky;
    top: 0;
    padding: 10px;
  }
  
  tfoot > tr {
    background-color: var(--gray-002);
    font-weight: 700;
    color: var(--blue-006);
  }

</style>
