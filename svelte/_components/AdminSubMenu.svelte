<script>
  import { onMount } from 'svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import HiOutlinePresentationChartBar from "svelte-icons-pack/hi/HiOutlinePresentationChartBar";
  import HiOutlineUserGroup from "svelte-icons-pack/hi/HiOutlineUserGroup";
  import HiOutlineUserAdd from "svelte-icons-pack/hi/HiOutlineUserAdd";

  let segment2;
  onMount(() => {
    console.log('onMount.AdminSubMenu')
    segment2 = window.location.pathname.split('/').pop() || '';
  });
</script>

<nav class="admin_submenu">
  <a href="/superAdmin/dashboard" class:active={segment2==='dashboard'}>
    <Icon className={segment2 === 'dashboard'  ? 'nav_icon_active' : 'nav_icon_dark'} size="20" src={HiOutlinePresentationChartBar}/>
    <span>Dashboard</span>
  </a>
  <a href="/superAdmin/userManagement" class:active={segment2==='userManagement'}>
    <Icon className={segment2 === 'userManagement'  ? 'nav_icon_active' : 'nav_icon_dark'} size="20" src={HiOutlineUserGroup}/>
    <span>Users</span>
  </a>
  <a href="/superAdmin/tenantManagement" class:active={segment2==='tenantManagement'}>
    <Icon className={segment2 === 'tenantManagement'  ? 'nav_icon_active' : 'nav_icon_dark'} size="20" src={HiOutlineUserAdd}/>
    <span>Tenants</span>
  </a>
</nav>

<style>
  :global(.admin_submenu a .nav_icon_dark path) {
    stroke: var(--gray-006) !important;
  }

  :global(.admin_submenu a:hover .nav_icon_dark path) {
    stroke: var(--blue-006) !important;
  }

  :global(.admin_submenu a .nav_icon_active path) {
    stroke: var(--blue-006) !important;
  }
  .admin_submenu {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    gap: 10px;
  }

  .admin_submenu a {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 5px;
    padding: 5px 15px;
    text-decoration: none;
    color: var(--gray-007);
    font-weight: 600;
    border: 1px solid var(--gray-002);
    border-radius: 999px;
    background-color: var(--gray-002);
  }

  .admin_submenu a:hover {
    color: var(--blue-005);
  }

  .admin_submenu a.active {
    color: var(--blue-006);
    border: 1px solid var(--blue-005);
    background-color: #7474ec30;
  }
</style>
