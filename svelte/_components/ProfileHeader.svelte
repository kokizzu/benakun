<script>
    // @ts-nocheck
    import {isSideMenuOpen} from './uiState.js';
    import Icon from 'svelte-icons-pack/Icon.svelte';
    import FaSolidBars from 'svelte-icons-pack/fa/FaSolidBars';
    import {onMount} from 'svelte';

    export let user = {};

    function openSideMenu() {
        isSideMenuOpen.set(!$isSideMenuOpen);
    }

    onMount(() => {
        console.log('onMount.ProfileHeader')
    });


</script>

<header class='profile_header'>
    <nav class='navbar'>
        <div class='label_menu'>
            <button on:click|preventDefault={openSideMenu}>
                <Icon color='#FFF' size={20} src={FaSolidBars}/>
            </button>
            <p>#{title}</p>
        </div>
        <div class='right_nav'>
            {user.email}
        </div>
    </nav>
</header>

<style>
    .profile_header {
        background-color : #4444EF;
        height           : fit-content;
        position         : relative;
        padding          : 18px 50px 80px 50px;
        display          : flex;
        flex-direction   : column;
    }

    .profile_header .navbar {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
    }

    .profile_header .navbar .label_menu {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        color          : white;
    }

    .profile_header .navbar .label_menu button {
        padding       : 10px;
        border        : none;
        background    : none;
        border-radius : 5px;
        font-size     : 14px;
        color         : white;
        cursor        : pointer;
    }

    .profile_header .navbar .label_menu button:hover {
        background-color : rgba(255, 255, 255, 0.3);
    }

    .profile_header .navbar .label_menu p {
        font-size   : 17px;
        font-weight : 600;
        line-height : 1.5rem;
        padding     : 0;
        margin      : 0 0 0 15px;
    }

    .profile_header .navbar .right_nav {
        display        : flex;
        flex-direction : row;
        gap            : 30px;
        align-items    : center;
    }

    /* Responsive to mobile device */
    @media only screen and (max-width : 768px) {
        .profile_header {
            padding : 15px 20px 80px 20px !important;
        }

        .profile_header .navbar .label_menu button {
            padding : 0;
        }

        .profile_header .navbar .right_nav {
            gap : 10px;
        }
    }
</style>
