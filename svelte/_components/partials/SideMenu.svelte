<script>
  /** @typedef {import('../types/access.js').Access} Access */

  import { UserLogout } from '../../jsApi.GEN.js';
  import { onMount } from 'svelte';
  import { notifier } from '../notifier.js'
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import {
    AiOutlineWarning,
    AiOutlineWallet,
    AiOutlineHome
  } from '../../node_modules/svelte-icons-pack/dist/ai';
  import { BsDatabaseAdd } from '../../node_modules/svelte-icons-pack/dist/bs';
  import { OiNote16 } from '../../node_modules/svelte-icons-pack/dist/oi';
  import { FaCircleUser } from '../../node_modules/svelte-icons-pack/dist/fa';
  import {
    RiEditorOrganizationChart, RiUserFacesUserSettingsLine,
    RiBuildingsCommunityLine, RiBuildingsBankLine,
    RiUserFacesGroupLine, RiUserFacesContactsLine,
    RiMapMap2Line, RiFinanceSwapBoxLine, RiDocumentStickyNoteAddLine,
    RiDocumentFileCopy2Line, RiUserFacesUserLocationLine,
    RiDocumentBookReadLine, RiFinanceCurrencyLine,
    RiBusinessBarChartBoxLine
  } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { CgLogOut, CgBox } from '../../node_modules/svelte-icons-pack/dist/cg';
  import { BsPostcard } from '../../node_modules/svelte-icons-pack/dist/bs';
  import { SlBookOpen } from '../../node_modules/svelte-icons-pack/dist/sl';
  import { VscLaw } from '../../node_modules/svelte-icons-pack/dist/vsc';
  import SubMenuLink from './SubMenuLink.svelte';
  import { IsShrinkMenu } from '../uiState.js';

  export let access =  /** @type Access */ ({});

  let segment;
  onMount(() => segment = window.location.pathname.split('/')[1] );

  async function userLogout() {
    await UserLogout({}, /** @type import('../../jsApi.GEN.js').UserLogoutCallback */
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

<aside class="side_menu {$IsShrinkMenu ? 'shrink' : 'expand'}">
  <div class='side_menu_container'>
    <header>
      <img src="/assets/icons/benakun-logo.png" alt="" />
      {#if !$IsShrinkMenu}
        <h3>BenAkun</h3>
      {/if}
    </header>
    <div class="menu_container">
      <nav class="menu_list">
        <a
          href="/"
          title="Home"
          class:active={segment === ''}
          >
          <Icon
            size="18"
            src={AiOutlineHome}
            className="{segment === ''  ? 'icon_active' : 'icon_dark'} icon"
          />
          {#if !$IsShrinkMenu}
            <div class="title">
              <span class="text">Home</span>
              <span class="subtext">Beranda</span>
            </div>
          {/if}
        </a>
        {#if access.reportViewer }
          <hr />
          <a
            href="/reportViewer/dashboard"
            title="Report Viewer"
            class:active={window.location.pathname === '/reportViewer/dashboard'}
          >
            <Icon
              size="18"
              className="{window.location.pathname === '/reportViewer/dashboard' ? 'icon_active' : 'icon_dark'} icon"
              src={AiOutlineWarning}
            />
            {#if !$IsShrinkMenu}
              <div class="title">
                <span class="text">Report Viewer</span>
                <span class="subtext">Laporan</span>
              </div>
            {/if}
          </a>
          <div class="submenu">
            <SubMenuLink
              title='General Ledger'
              subtitle='Buku Besar'
              href='/reportViewer/generalLedger'
              icon={RiDocumentBookReadLine}
            />
            <SubMenuLink
              title='Trial Balance'
              subtitle='Neraca Saldo'
              href='/reportViewer/trialBalance'
              icon={VscLaw}
            />
            <SubMenuLink
              title='Financial Posistion'
              subtitle='Laporan Posisi Keuangan'
              href='/reportViewer/financialPosition'
              icon={RiFinanceCurrencyLine}
            />
            <SubMenuLink
              title='Loss / Income Statements'
              subtitle='Laporan Laba / Rugi'
              href='/reportViewer/lossIncomeStatements'
              icon={RiBusinessBarChartBoxLine}
            />
          </div>
        {/if}
        {#if access.fieldSupervisor }
          <hr />
          <a
            href="/fieldSupervisor/dashboard"
            title="Field Supervisor"
            class:active={segment === 'fieldSupervisor'}
          >
            <Icon
              size="18"
              className="{segment === 'fieldSupervisor'  ? 'icon_active' : 'icon_dark'} icon"
              src={RiUserFacesUserLocationLine}
            />
            {#if !$IsShrinkMenu}
              <div class="title">
                <span class="text">Field Supervisor</span>
                <span class="subtext">Pengawas Lapangan</span>
              </div>
            {/if}
          </a>
        {/if}
        {#if access.dataEntry }
          <hr />
          <a
            href='/dataEntry/dashboard'
            title="Data Entry"
            class:active={window.location.pathname === '/dataEntry/dashboard'}
          >
            <Icon
              size="18"
              className={window.location.pathname === '/dataEntry/dashboard'
                ? 'icon_active' : 'icon_dark'
              }
              src={BsDatabaseAdd}
            />
            {#if !$IsShrinkMenu}
              <div class="title">
                <span class="text">Data Entry</span>
                <span class="subtext">Entri Data</span>
              </div>
            {/if}
          </a>
          <div class="submenu">
            <SubMenuLink
              title='Transaction Entry'
              subtitle='Entri Transaksi'
              href='/dataEntry/template'
              icon={OiNote16}
            />
          </div>
        {/if}
        {#if access.tenantAdmin}
          <hr />
          <a
            href="/tenantAdmin/dashboard"
            title="Tenant Admin Dashboard"
            class:active={window.location.pathname === '/tenantAdmin/dashboard'}
          >
            <Icon
              size="18"
              className="{window.location.pathname === '/tenantAdmin/dashboard'  ? 'icon_active' : 'icon_dark'} icon"
              src={RiBuildingsCommunityLine}
            />
            {#if !$IsShrinkMenu}
              <div class="title">
                <span class="text">Tenant Admin</span>
                <span class="subtext">Admin Penyewa</span>
              </div>
            {/if}
          </a>
          <div class="submenu">
            <SubMenuLink
              title='Organization'
              subtitle='Organisasi'
              href='/tenantAdmin/organization'
              icon={RiEditorOrganizationChart}
            />
            <SubMenuLink
              title='Budgeting'
              subtitle='Pengelolaan Anggaran'
              href='/tenantAdmin/budgeting'
              icon={AiOutlineWallet}
            />
            <SubMenuLink
              title='Chart of Account'
              subtitle='Bagan Akun'
              href='/tenantAdmin/coa'
              icon={BsPostcard}
            />
            <SubMenuLink
              title='Bank Accounts'
              subtitle='Rekening Bank'
              href='/tenantAdmin/bankAccounts'
              icon={RiBuildingsBankLine}
            />
            <SubMenuLink
              title='Products'
              subtitle='Produk'
              href='/tenantAdmin/products'
              icon={CgBox}
            />
            <SubMenuLink
              title='Locations'
              subtitle='Lokasi'
              href='/tenantAdmin/locations'
              icon={RiMapMap2Line}
            />
            <SubMenuLink
              title='Inventory Changes'
              subtitle='Perubahan Inventory'
              href='/tenantAdmin/inventoryChanges'
              icon={RiFinanceSwapBoxLine}
            />
            <SubMenuLink
              title='Transaction Template'
              subtitle='Template Transaksi'
              href='/tenantAdmin/transactionTemplate'
              icon={RiDocumentStickyNoteAddLine}
            />
            <SubMenuLink
              title='Manual Journal'
              subtitle='Jurnal Manual'
              href='/tenantAdmin/manualJournal'
              icon={SlBookOpen}
            />
          </div>
        {/if}
        {#if access.superAdmin }
          <hr />
          <a
            href="/superAdmin/dashboard"
            title="Super Admin Dashboard"
            class:active={window.location.pathname === '/superAdmin/dashboard'}
          >
            <Icon
              size="18"
              className="{window.location.pathname === '/superAdmin/dashboard'  ? 'icon_active' : 'icon_dark'} icon"
              src={RiUserFacesUserSettingsLine}
            />
            {#if !$IsShrinkMenu}
              <div class="title">
                <span class="text">Super Admin</span>
                <span class="subtext">Admin Super</span>
              </div>
            {/if}
          </a>
          <div class="submenu">
            <SubMenuLink
              title="Users"
              subtitle="Pengguna"
              href="/superAdmin/userManagement"
              icon={RiUserFacesGroupLine}
            />
            <SubMenuLink
              title="Tenants"
              subtitle="Penyewa"
              href="/superAdmin/tenantManagement"
              icon={RiUserFacesContactsLine}
            />
            <SubMenuLink
              title="Access Log"
              subtitle="Log Akses"
              href="/superAdmin/accessLog"
              icon={RiDocumentFileCopy2Line}
            />
          </div>
        {/if}
      </nav>
      <nav class="menu_list">
        {#if access.user}
          <a
            href="/user/profile"
            title="Profile"
            class:active={segment === 'user'}
          >
            <Icon
              size="18"
              className="{segment === 'user' ? 'icon_active' : 'icon_dark'} icon"
              src={FaCircleUser}
            />
            {#if !$IsShrinkMenu}
              <div class="title">
                <span class="text">Profile</span>
                <span class="subtext">Profil</span>
              </div>
            {/if}
          </a>
        {/if}
        {#if access.user || access.superAdmin}
          <button on:click={userLogout} class="logout" title="Logout">
            <Icon size="18" className="icon_logout" src={CgLogOut}/>
            {#if !$IsShrinkMenu}
              <span>Logout</span>
            {/if}
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
    overflow-y       : scroll;
    overflow-x: hidden;
    flex-direction   : row;
    flex-wrap        : nowrap;
    background-color : white;
    padding          : 16px 0;
    border-right: 1px solid var(--gray-002);
    user-select: none;
  }

  .side_menu::view-transition {
    position: fixed;
    inset: 0;
  }

  .side_menu.expand {
    width: var(--sidemenu-width);
  }

  .side_menu.shrink {
    width: var(--sidemenu-width-sm);
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

  .side_menu.expand::-webkit-scrollbar-thumb {
    background-color : var(--gray-003);
    border-radius    : 8px;
    visibility: hidden;
    position: fixed;
  }

  .side_menu.expand::-webkit-resizer,
  .side_menu.expand::-webkit-scrollbar-button,
  .side_menu.expand::-webkit-scrollbar-corner {
    display: none;
  }

  .side_menu.side_menu.expand::-webkit-scrollbar-track-piece {
    background:transparent;
  }

  .side_menu.expand::-webkit-scrollbar-thumb:hover {
    background-color : var(--gray-004);
    visibility: initial;
  }

  .side_menu.shrink::-webkit-scrollbar-thumb:hover {
    visibility: hidden;
  }

  .side_menu:hover::-webkit-scrollbar-thumb {
    visibility: initial;
  }

  .side_menu.expand::-webkit-scrollbar {
    width : 8px;
    visibility: hidden;
  }

  .side_menu.shrink::-webkit-scrollbar {
    width : 0;
    visibility: hidden;
  }

  .side_menu::-webkit-scrollbar:hover {
    visibility: initial;
  }

  .side_menu:hover::-webkit-scrollbar {
    visibility: initial;
  }

  .side_menu::-webkit-scrollbar-track {
    background-color : transparent;
  }

  .side_menu_container header {
    display         : flex;
    flex-direction  : row;
    align-items     : center;
    gap: 10px;
    height: fit-content;
    flex-grow: 0;
    user-select: none;
  }

  .side_menu.expand header {  
    margin: 0 24px 15px 24px;
  }

  .side_menu.shrink header {  
    margin: 0 0 15px 0;
    justify-content: center;
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
    gap: 30px;
    flex-grow: 1;
    padding-bottom: 10px;
  }

  .side_menu_container .menu_container .menu_list {
    display: flex;
    flex-direction: column;
    gap: 0;
    padding: 0 2px 0 10px;
  }

  .side_menu.shrink .side_menu_container .menu_container .menu_list {
    align-items: center;
  }

  .side_menu_container .menu_container .menu_list a {
    display: flex;
    flex-direction: row;
    align-items: start;
    gap: 10px;
    text-decoration: none;
    text-transform: capitalize;
    width: 100%;
    color: var(--gray-007);
    font-size: 15px;
    padding: 10px 12px;
    border-radius: 8px;
  }

  .side_menu.expand .menu_list hr {
    display: none;
  }

  .side_menu.shrink .menu_list hr {
    background-color: var(--gray-003);
    border-radius: 9999px;
    height: 2px;
    border: none;
    margin: 5px;
    width: 75%;
  }

  .side_menu.shrink .side_menu_container .menu_container .menu_list a {
    justify-content: center;
    align-items: start;
    width: fit-content;
  }

  :global(.side_menu_container .menu_container .menu_list a .icon) {
    flex-shrink: 0;
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

  .side_menu_container .menu_container .menu_list a .title {
    margin-top: -2px;
    display: flex;
    flex-direction: column;
    gap: 2px;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .side_menu_container .menu_container .menu_list a .subtext {
    font-size: 10px;
    color: var(--gray-006);
  }

  .side_menu_container .menu_container .menu_list a.active .subtext {
    color: var(--violet-005);
  }

  .side_menu_container .menu_container .menu_list a span {
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .side_menu_container .menu_container .menu_list .submenu {
    display: flex;
    flex-direction: column;
    gap: 0;
  }

  .side_menu.expand .side_menu_container .menu_container .menu_list .submenu {
    padding-left: 20px;
  }

  .side_menu.shrink .side_menu_container .menu_container .menu_list .submenu {
    padding-left: 0;
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
    font-size: 15px;
    border: none;
    border-radius: 8px;
    background-color: transparent;
    padding: 10px 12px;
    cursor: pointer;
    color: var(--red-005);
  }

  .side_menu.shrink .logout {
    justify-content: center;
    width: fit-content
  }

  .menu_list .logout:hover {
    background-color: var(--red-transparent);
  }
</style>