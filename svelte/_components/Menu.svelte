<script>
    // @ts-nocheck
    import {UserLogout} from '../jsApi.GEN.js';
    import {onMount} from 'svelte';
    import {isSideMenuOpen} from './uiState.js';

    import Icon from 'svelte-icons-pack/Icon.svelte';
    import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
    import FaSolidShoppingBag from 'svelte-icons-pack/fa/FaSolidShoppingBag';
    import FaSolidBuilding from 'svelte-icons-pack/fa/FaSolidBuilding';
    import FaSolidSlidersH from 'svelte-icons-pack/fa/FaSolidSlidersH';
    import FaSolidUserCircle from 'svelte-icons-pack/fa/FaSolidUserCircle';
    import FaSolidSignInAlt from 'svelte-icons-pack/fa/FaSolidSignInAlt';
    import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
    import FaSolidList from "svelte-icons-pack/fa/FaSolidList";
    import FaSolidTree from "svelte-icons-pack/fa/FaSolidTree";
    import SubMenuLink from "./SubMenuLink.svelte";

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
        console.log(access);
        segment1 = window.location.pathname.split('/')[1];
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
    <aside class='side_menu_admin'>
        <div class='side_menu_admin_container'>
            <header>
                <h3>BenAkun</h3>
                <button on:click|preventDefault={doToggle}>
                    <Icon size={20} color='#475569' src={FaSolidTimes}/>
                </button>
            </header>
            <div class='menu_container'>
                <!-- PAGES -->
                <hr/>
                <h6>MENU</h6>
                <nav class='menu'>
                    <a href='/' class:active={sel1('')}>
                        <Icon size={22} className={solid1('')} src={FaSolidHome}/>
                        <span>HOME</span>
                    </a>
                    {#if access.reportViewer }
                        <a href='/reportViewer/dashboard' class:active={sel1('reportViewer')}>
                            <Icon size={20} className={solid1('reportViewer')} src={FaSolidSlidersH}/>
                            <span>REPORT VIEWER</span>
                        </a>
                    {/if}
                    {#if access.entryUser }
                        <a href='/entryUser/dashboard' class:active={sel1('entryUser')}>
                            <Icon size={20} className={solid1('entryUser')} src={FaSolidSlidersH}/>
                            <span>ENTRY USER</span>
                        </a>
                    {/if}
                    {#if access.tenantAdmin}
                        <a href='/tenantAdmin/dashboard' class:active={sel1('tenantAdmin')}>
                            <Icon size={20} className={solid1('tenantAdmin')} src={FaSolidBuilding}/>
                            <span>TENANT ADMIN</span>
                        </a>
                        <div class="submenu">
                            <SubMenuLink title='Organization' href='/tenantAdmin/organization' icon={FaSolidTree}/>
                            <SubMenuLink title='Budgeting' href='/tenantAdmin/budgeting' icon={FaSolidList}/>
                        </div>
                    {/if}
                    {#if access.superAdmin }
                        <a href='/superAdmin/dashboard' class:active={sel1('superAdmin')}>
                            <Icon size={22} className={solid1('superAdmin')} src={FaSolidShoppingBag}/>
                            <span>SUPER ADMIN</span>
                        </a>
                    {/if}
                </nav>
                <!-- SETTING -->
                <hr/>
                <h6>SETTING</h6>
                <nav class='menu'>
                    {#if access.user}
                        <a href='/user/profile' class:active={segment1 === 'user'}>
                            <Icon size={22} className={segment1 === 'user' ? 'icon_active' : 'icon_dark'} src={FaSolidUserCircle}/>
                            <span>PROFILE</span>
                        </a>
                    {/if}
                    {#if access.user || access.superAdmin}
                        <button on:click={userLogout} class='logout'>
                            <Icon size={22} className='icon_dark' src={FaSolidSignInAlt}/>
                            <span>LOGOUT</span>
                        </button>
                    {/if}
                </nav>
            </div>
        </div>
    </aside>
{/if}
<style>

    .side_menu_admin {
        left             : 0;
        display          : block;
        position         : fixed;
        z-index          : 9999;
        top              : 0;
        bottom           : 0;
        overflow-y       : auto;
        flex-direction   : row;
        flex-wrap        : nowrap;
        overflow         : auto;
        background-color : white;
        color            : #475569;
        padding          : 16px 24px;
        width            : 300px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .side_menu_admin_container {
        flex-direction : column;
        align-items    : stretch;
        min-height     : 100%;
        flex-wrap      : nowrap;
        padding        : 0;
        display        : flex;
        width          : 100%;
        margin         : 0 auto;
    }

    .side_menu_admin_container header {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
    }

    .side_menu_admin_container header h3 {
        font-size   : 16px;
        line-height : 1.5rem;
        padding     : 0;
        margin      : 0;
    }

    .side_menu_admin_container header button {
        padding       : 5px;
        border        : none;
        background    : none;
        border-radius : 5px;
        font-size     : 14px;
        cursor        : pointer;
    }

    .side_menu_admin_container header button:hover {
        background-color : rgb(0 0 0 / 0.07);
        color            : #4444EF;
    }

    .menu_container {
        margin-top     : 1rem;
        align-items    : stretch;
        flex-direction : column;
        display        : flex;
    }

    .menu_container hr {
        margin : 1rem 0;
    }

    .menu_container h6 {
        font-size : 15px;
        margin    : 12px 0;
    }

    .menu_container .menu {
        display        : flex;
        flex-direction : column;
        margin-bottom  : 10px;
    }

    .menu_container .menu a, .menu .logout { /*MENU LISTS*/
        color            : #475569;
        text-decoration  : none;
        margin           : 0;
        padding          : 0.3rem 0 0.1rem 0.5rem;
        font-size        : 0.875rem !important;
        line-height      : 1.25rem;
        font-weight      : 700;
        text-transform   : uppercase;
        text-align       : left;
        background-color : transparent;
        border           : none;
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        gap              : 15px;
    }

    .menu_container .menu .logout {
        cursor        : pointer;
        margin-top    : 0;
        margin-bottom : 0;
        margin-right  : 0;
    }


    .menu_container .menu a:hover, .menu_container .menu .logout:hover { /*HOVER*/
        color : #64748B;
    }

    .active { /*ACTIVE Navigation*/
        color : #4444EF !important;
    }

    .submenu {
        margin-left : 2.5rem;
    }
</style>