<script>
  // @ts-nocheck
  import {isSideMenuOpen} from '../uiState.js';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import CgMenuLeft from "svelte-icons-pack/cg/CgMenuLeft";
  import {onMount} from 'svelte';

  export let user = {};

  function openSideMenu() {
    isSideMenuOpen.set(!$isSideMenuOpen);
  }

  onMount(() => {
    console.log('onMount.ProfileHeader')
  });
</script>

<header class="navbar_container">
  <nav class="navbar">
    <div class="left">
      <button class="menu_btn" on:click|preventDefault={openSideMenu}>
        <Icon color="#FFF" size={25} src={CgMenuLeft}/>
      </button>
      <span class="display_title">#{title}</span>
    </div>
    <div class="right">
      <span class="display_email">{user.email}</span>
    </div>
  </nav>
</header>

<style>
  .navbar_container {
    background-color: var(--purple-003);
    height: 70px;
    padding: 18px 40px;
    display: flex;
    color: #FFF;
  }

  .navbar_container .navbar {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  .navbar_container .navbar .left,
  .navbar_container .navbar .right {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    gap: 20px;
  }

  .navbar_container .navbar .left .menu_btn {
    background-color: transparent;
    padding: 5px;
    border-radius: 50%;
    border: 1px solid transparent;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    color: #FFF;
  }

  .navbar_container .navbar .left .menu_btn:hover {
    background-color: var(--purple-002);
    border: 1px solid var(--purple-001);
  }

  .navbar_container .navbar .left .menu_btn:active {
    background-color: var(--purple-001);
  }

  .navbar_container .navbar .left .display_title {
    font-size: var(--font-lg);
    font-weight: 700;
  }

  .navbar_container .navbar .right .display_email {
    padding: 5px 15px;
    background-color: var(--purple-002);
    border-radius: 9999px;
    border: 1px solid var(--purple-001);
  }

  /* Responsive to mobile device */
  @media only screen and (max-width : 768px) {}
</style>
