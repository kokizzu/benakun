<script>
  // @ts-nocheck
  import {UserLogout} from '../../jsApi.GEN.js';
  import {onMount} from 'svelte';
  import {isSideMenuOpen} from '../uiState.js';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import AiOutlineWarning from "svelte-icons-pack/ai/AiOutlineWarning";
  import FaUserCircle from "svelte-icons-pack/fa/FaUserCircle";
  import RiEditorOrganizationChart from "svelte-icons-pack/ri/RiEditorOrganizationChart";
  import FaSolidSlidersH from 'svelte-icons-pack/fa/FaSolidSlidersH';
  import HiSolidX from "svelte-icons-pack/hi/HiSolidX";
  import CgLogOut from "svelte-icons-pack/cg/CgLogOut";
  import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
  import AiOutlineWallet from "svelte-icons-pack/ai/AiOutlineWallet";
  import RiUserAdminLine from "svelte-icons-pack/ri/RiUserAdminLine";
  import AiOutlineHome from "svelte-icons-pack/ai/AiOutlineHome";
  import RiBuildingsCommunityLine from "svelte-icons-pack/ri/RiBuildingsCommunityLine";
  import SubMenuLink from "../SubMenuLink.svelte";

  export let doToggle = function() {
    isSideMenuOpen.set(!$isSideMenuOpen);
  };
  export let access = {
    'superAdmin': false,
    'tenantAdmin': false,
    'entryUser': false,
    'reportViewer': false,
    'guest': false,
    'user': false,
  };

  let segment1;
  onMount(() => {
    console.log('onMount.Menu');
    console.log('Access:',access);
    segment1 = window.location.pathname.split('/')[1];
    console.log('current path = ', window.location.pathname)
  });

  async function userLogout() {
    await UserLogout({}, function(o) {
      console.log(o);
      if (o.error) return alert(o.error);
      window.location = '/';
    });
  }
  
  function sel1(f) {
    return segment1 === f;
  }

  function solid1(f) {
    return segment1 === f ? 'icon_active' : 'icon_dark';
  }

  function selFull(f) {
    return window.location.pathname === f;
  }

  function solidFull(f) {
    return window.location.pathname === f ? 'icon_active' : 'icon_dark';
  }
</script>

{#if $isSideMenuOpen}
	<button class="backdrop" on:click={() => $isSideMenuOpen = !$isSideMenuOpen}></button>
{/if}
<aside class={$isSideMenuOpen ? `side_menu open` : `side_menu`}>
  <div class='side_menu_container'>
    <header>
      <h3>BenAkun</h3>
      <button class="close_btn" on:click|preventDefault={doToggle}>
        <Icon size={25} src={HiSolidX}/>
      </button>
    </header>
    <div class="menu_container">
      <nav class="menu_list">
        <a href="/" class:active={segment1 === ''}>
          <Icon size={22} src={AiOutlineHome} className={segment1 === ''  ? 'icon_active' : 'icon_dark'}/>
          <span>Home</span>
        </a>
        {#if access.reportViewer }
          <a href="/reportViewer/dashboard" class:active={segment1 === 'reportViewer'}>
            <Icon size={22} className={segment1 === 'reportViewer'  ? 'icon_active' : 'icon_dark'} src={AiOutlineWarning}/>
            <span>Report Viewer</span>
          </a>
        {/if}
        {#if access.entryUser }
          <a href='/entryUser/dashboard' class:active={segment1 === 'entryUser'}>
            <Icon size={22} className={segment1 === 'entryUser'  ? 'icon_active' : 'icon_dark'} src={FaSolidSlidersH}/>
            <span>Entry User</span>
          </a>
        {/if}
        {#if access.tenantAdmin}
          <a href='/tenantAdmin/dashboard' class:active={window.location.pathname === '/tenantAdmin/dashboard'}>
            <Icon size={22} className={window.location.pathname === '/tenantAdmin/dashboard'  ? 'icon_active' : 'icon_dark'} src={RiBuildingsCommunityLine}/>
            <span>Tenant Admin</span>
          </a>
          <div class="submenu">
            <SubMenuLink title='Organization' href='/tenantAdmin/organization' icon={RiEditorOrganizationChart}/>
            <SubMenuLink title='Budgeting' href='/tenantAdmin/budgeting' icon={AiOutlineWallet}/>
          </div>
        {/if}
        {#if access.superAdmin }
          <a href='/superAdmin/dashboard' class:active={segment1 === 'superAdmin'}>
            <Icon size={22} className={segment1 === 'superAdmin'  ? 'icon_active' : 'icon_dark'} src={RiUserAdminLine}/>
            <span>Super Admin</span>
          </a>
        {/if}
      </nav>
      <nav class="menu_list">
        {#if access.user}
          <a href='/user/profile' class:active={segment1 === 'user'}>
            <Icon size={22} className={segment1 === 'user' ? 'icon_active' : 'icon_dark'} src={FaUserCircle}/>
            <span>Profile</span>
          </a>
        {/if}
        {#if access.user || access.superAdmin}
          <button on:click={userLogout} class='logout'>
            <Icon size={22} className='icon_logout' src={CgLogOut}/>
            <span>Logout</span>
          </button>
        {/if}
      </nav>
    </div>
  </div>
</aside>

<style>
  :global(.icon_dark) {
    fill: var(--gray-006);
  }

  :global(.icon_active) {
    fill: var(--purple-002);
  }

  :global(.icon_logout) {
    fill: var(--red-002);
  }

  .backdrop {
    z-index: 8888;
    position: fixed;
    padding: 0;
    border: none;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: rgba(0 0 0 / 20%);
  }

  .side_menu {
    left             : -300px;
    display          : block;
    position         : fixed;
    z-index          : 9999;
    top              : 0;
    bottom           : 0;
    overflow-y       : auto;
    flex-direction   : row;
    flex-wrap        : nowrap;
    background-color : white;
    padding          : 16px 24px;
    width            : 300px;
    transition       : 0.3s;
  }

  .side_menu.open {
    left : 0;
  }

  .side_menu_container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: 100%;
    flex-wrap: nowrap;
    padding: 0;
    width: 100%;
    margin: 0;
  }

  .side_menu_container header {
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    align-items     : center;
    height: fit-content;
    flex-grow: 0;
  }

  .side_menu_container header h3 {
    font-size   : var(--font-xl);
    line-height : 1.5rem;
    font-weight: 800;
    padding     : 0;
    margin      : 0;
  }

  .side_menu_container header .close_btn {
    background-color: transparent;
    padding: 5px;
    border-radius: 50%;
    border: 1px solid transparent;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }

  .side_menu_container header .close_btn:hover {
    background-color: var(--gray-001);
    border: 1px solid var(--gray-002);
  }

  .side_menu_container header .close_btn:active {
    background-color: var(--gray-002);
  }

  .side_menu_container .menu_container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    flex-grow: 1;
  }

  .side_menu_container .menu_container .menu_list {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .side_menu_container .menu_container .menu_list {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .side_menu_container .menu_container .menu_list a {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-decoration: none;
    text-transform: capitalize;
    width: fit-content;
    color: var(--gray-007);
    font-size: var(--font-md);
  }

  .side_menu_container .menu_container .menu_list a:hover {
    color: var(--purple-002);
  }

  .side_menu_container .menu_container .menu_list a.active {
    font-weight: 600;
    color: var(--purple-002);
  }

  .side_menu_container .menu_container .menu_list .submenu {
    margin-left: 20px;
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .menu_list .logout {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-decoration: none;
    text-transform: capitalize;
    width: fit-content;
    color: var(--gray-007);
    font-size: var(--font-md);
    border: none;
    background-color: transparent;
    padding: 0;
    cursor: pointer;
    color: var(--red-002);
  }
</style>