<script>
  import { UserLogout } from '../../jsApi.GEN.js';
  import { onMount } from 'svelte';
  import { notifier } from '../notifier.js'
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import AiOutlineWarning from 'svelte-icons-pack/ai/AiOutlineWarning';
  import FaUserCircle from 'svelte-icons-pack/fa/FaUserCircle';
  import RiEditorOrganizationChart from 'svelte-icons-pack/ri/RiEditorOrganizationChart';
  import FaSolidSlidersH from 'svelte-icons-pack/fa/FaSolidSlidersH';
  import CgLogOut from 'svelte-icons-pack/cg/CgLogOut';
  import AiOutlineWallet from 'svelte-icons-pack/ai/AiOutlineWallet';
  import RiUserAdminLine from 'svelte-icons-pack/ri/RiUserAdminLine';
  import AiOutlineHome from 'svelte-icons-pack/ai/AiOutlineHome';
  import RiBuildingsCommunityLine from 'svelte-icons-pack/ri/RiBuildingsCommunityLine';
  import BsPostcard from 'svelte-icons-pack/bs/BsPostcard';
  import HiOutlineUserGroup from "svelte-icons-pack/ri/RiUserGroup2Line.js";
  import HiOutlineUserAdd from "svelte-icons-pack/ri/RiUserSpyLine.js";
  import SubMenuLink from './SubMenuLink.svelte';

  /** @typedef {import('../types/access.js').Access} Access*/

  /** @type Access */ // @ts-ignore
  export let access = {};

  let segment1;
  onMount(() => segment1 = window.location.pathname.split('/')[1] );

  async function userLogout() {
    await UserLogout({},
      /** @type import('../../jsApi.GEN.js').UserLogoutCallback */
      function(/** @type {any} */ o) {
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }
        window.document.location = '/';
      }
    );
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
          <Icon size="18" src={AiOutlineHome} className={segment1 === ''  ? 'icon_active' : 'icon_dark'}/>
          <span>Home</span>
        </a>
        {#if access.reportViewer }
          <a href="/reportViewer/dashboard" class:active={segment1 === 'reportViewer'}>
            <Icon size="18" className={segment1 === 'reportViewer'  ? 'icon_active' : 'icon_dark'} src={AiOutlineWarning}/>
            <span>Report Viewer</span>
          </a>
        {/if}
        {#if access.entryUser }
          <a href='/entryUser/dashboard' class:active={segment1 === 'entryUser'}>
            <Icon size="18" className={segment1 === 'entryUser'  ? 'icon_active' : 'icon_dark'} src={FaSolidSlidersH}/>
            <span>Entry User</span>
          </a>
        {/if}
        {#if access.tenantAdmin}
          <a href='/tenantAdmin/dashboard' class:active={window.location.pathname === '/tenantAdmin/dashboard'}>
            <Icon size="18" className={window.location.pathname === '/tenantAdmin/dashboard'  ? 'icon_active' : 'icon_dark'} src={RiBuildingsCommunityLine}/>
            <span>Tenant Admin</span>
          </a>
          <div class="submenu">
            <SubMenuLink title='Organization' href='/tenantAdmin/organization' icon={RiEditorOrganizationChart}/>
            <SubMenuLink title='Budgeting' href='/tenantAdmin/budgeting' icon={AiOutlineWallet}/>
            <SubMenuLink title='Chart of Acount' href='/tenantAdmin/coa' icon={BsPostcard}/>
          </div>
        {/if}
        {#if access.superAdmin }
          <a href='/superAdmin/dashboard' class:active={window.location.pathname === '/superAdmin/dashboard'}>
            <Icon size="18" className={window.location.pathname === '/superAdmin/dashboard'  ? 'icon_active' : 'icon_dark'} src={RiUserAdminLine}/>
            <span>Super Admin</span>
          </a>
          <div class="submenu">
            <SubMenuLink title='Users' href='/superAdmin/userManagement' icon={HiOutlineUserGroup}/>
            <SubMenuLink title='Tenants' href='/superAdmin/tenantManagement' icon={HiOutlineUserAdd}/>
          </div>
        {/if}
      </nav>
      <nav class="menu_list">
        {#if access.user}
          <a href='/user/profile' class:active={segment1 === 'user'}>
            <Icon size="18" className={segment1 === 'user' ? 'icon_active' : 'icon_dark'} src={FaUserCircle}/>
            <span>Profile</span>
          </a>
        {/if}
        {#if access.user || access.superAdmin}
          <button on:click={userLogout} class='logout'>
            <Icon size="18" className='icon_logout' src={CgLogOut}/>
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
    fill: var(--violet-005);
  }

  :global(.icon_logout) {
    fill: var(--red-002);
  }

  .side_menu {
    position: fixed;
    z-index: 800;
    bottom: 0;
    top: 0;
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
    gap: 15px;
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
    margin: 0 24px 15px 24px;
    user-select: none;
  }

  .side_menu_container header img {
    width : 30px;
    height: 30px;
  }

  .side_menu_container header h3 {
    font-size   : var(--font-lg);
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
    gap: 5px;
    padding: 0 10px;
  }

  .side_menu_container .menu_container .menu_list a {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-decoration: none;
    text-transform: capitalize;
    width: 100%;
    color: var(--gray-007);
    font-size: var(--font-md);
    padding: 10px 12px;
    border-radius: 8px;
  }

  .side_menu_container .menu_container .menu_list a:hover {
    background-color: var(--gray-001);
  }

  .side_menu_container .menu_container .menu_list a:active {
    background-color: var(--gray-002);
  }

  .side_menu_container .menu_container .menu_list a.active {
    color: var(--violet-005);
    background-color: var(--violet-transparent);
  }

  .side_menu_container .menu_container .menu_list .submenu {
    padding-left: 20px;
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .menu_list .logout {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-decoration: none;
    text-transform: capitalize;
    width: 100%;
    color: var(--gray-007);
    font-size: var(--font-md);
    border: none;
    border-radius: 8px;
    background-color: transparent;
    padding: 10px 12px;
    cursor: pointer;
    color: var(--red-005);
  }

  .menu_list .logout:hover {
    background-color: var(--red-transparent);
  }
</style>