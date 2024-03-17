<script>
  // @ts-nocheck
  import {UserLogout} from '../../jsApi.GEN.js';
  import {onMount} from 'svelte';
  import {notifier} from '../notifier.js'

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import AiOutlineWarning from "svelte-icons-pack/ai/AiOutlineWarning";
  import FaUserCircle from "svelte-icons-pack/fa/FaUserCircle";
  import RiEditorOrganizationChart from "svelte-icons-pack/ri/RiEditorOrganizationChart";
  import FaSolidSlidersH from 'svelte-icons-pack/fa/FaSolidSlidersH';
  import CgLogOut from "svelte-icons-pack/cg/CgLogOut";
  import AiOutlineWallet from "svelte-icons-pack/ai/AiOutlineWallet";
  import RiUserAdminLine from "svelte-icons-pack/ri/RiUserAdminLine";
  import AiOutlineHome from "svelte-icons-pack/ai/AiOutlineHome";
  import RiBuildingsCommunityLine from "svelte-icons-pack/ri/RiBuildingsCommunityLine";
  import BsPostcard from 'svelte-icons-pack/bs/BsPostcard';
  import SubMenuLink from "../SubMenuLink.svelte";

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
      if (o.error) return notifier.showError(o.error);
      window.document.location = '/';
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

<aside class="side_menu">
  <div class='side_menu_container'>
    <header>
      <img src="/assets/icons/benakun-logo.png" alt="" />
      <h3>BenAkun</h3>
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
            <SubMenuLink title='Chart of Acount' href='/tenantAdmin/coa' icon={BsPostcard}/>
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


  .side_menu {
    display          : block;
    overflow-y       : auto;
    flex-direction   : row;
    flex-wrap        : nowrap;
    background-color : white;
    padding          : 16px 0;
    width            : 260px;
    border-right: 1px solid var(--gray-002);
    user-select: none;
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
    align-items     : center;
    gap: 10px;
    height: fit-content;
    flex-grow: 0;
    margin: 0 24px 20px 24px;
    user-select: none;
  }

  .side_menu_container header img {
    width : 40px;
    height: 40px;
  }

  .side_menu_container header h3 {
    font-size   : var(--font-xl);
    line-height : 1.5rem;
    font-weight: 800;
    padding     : 0;
    margin      : 0;
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
    gap: 7px;
    padding: 0 10px;
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
    padding: 10px 12px;
    border-radius: 999px;
  }

  .side_menu_container .menu_container .menu_list a:hover {
    font-weight: 700;
    background-color: var(--gray-002);
  }

  .side_menu_container .menu_container .menu_list a.active {
    font-weight: 600;
    color: var(--purple-002);
    background-color: #7474ec30;
  }

  .side_menu_container .menu_container .menu_list .submenu {
    padding-left: 20px;
    display: flex;
    flex-direction: column;
    gap: 7px;
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
    border-radius: 999px;
    background-color: transparent;
    padding: 10px 12px;
    cursor: pointer;
    color: var(--red-002);
  }

  .menu_list .logout:hover {
    background-color: #ef444430;
  }
</style>