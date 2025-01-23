<script>
    import { onMount } from 'svelte';
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import {
    RiDesignLayoutLeftLine,
    RiBusinessBarChartHorizontalLine
  } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { IsShrinkMenu } from '../uiState.js';
  import { IsUnixTimeExpired } from '../xHelper';

  export let user = /** @type {import('../types/user.js').User}*/ ({});

  let isUserBoughtSupport = false;

  onMount(() => {
    console.log('user:', user);
    if (!IsUnixTimeExpired(user.supportExpiredAt)) {
      isUserBoughtSupport = true;
    }
  })

  function handleShrinkMenu() {
    console.log('handleShrinkMenu', $IsShrinkMenu);
    IsShrinkMenu.set(!$IsShrinkMenu);
    if (window.innerWidth > 768) {
      localStorage.setItem('IsShrinkMenu', JSON.stringify($IsShrinkMenu))
    }
  }

  function handleUpgrade() {
    window.location.href = '/user/purchaseSupport';
  }
</script>

<header class="navbar_container">
  <nav class="navbar">
    <div class="left">
      <button class="shrink_btn" on:click={handleShrinkMenu}>
        <Icon
          src={RiDesignLayoutLeftLine}
          size={20}
          className="icon_wide"
        />
        <Icon
          src={RiBusinessBarChartHorizontalLine}
          size={20}
          className="icon_small"
        />
      </button>
    </div>
    <div class="right">
      <picture>
        <img src="/assets/icons/benakun-logo.png" alt="" />
      </picture>
      {#if !isUserBoughtSupport}
        <button class="upgrade" on:click={handleUpgrade}>
          <span>Upgrade</span>
        </button>
      {/if}
      <span class="display-email { isUserBoughtSupport ? 'premium' : 'free'} ">
        <div class="wrapper">{user.email}</div>
      </span>
    </div>
  </nav>
</header>

<style>
  .navbar_container {
    display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 0 20px;
    background-color: #FFF;
		position: sticky;
    top: 0;
    left: 0;
		height: var(--navbar-height);
		border-bottom: 1px solid var(--gray-003);
		width: 100%;
		z-index: 1200;
		top: 0;
  }

  .navbar_container .navbar {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    height: var(--navbar-height);
  }

  .navbar_container .navbar .left,
  .navbar_container .navbar .right {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    gap: 20px;
  }

  .navbar_container .navbar .left .shrink_btn {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 10px;
    border-radius: 8px;
    background-color: transparent;
    border: none;
    cursor: pointer;
  }

  .navbar_container .navbar .left .shrink_btn:hover {
    background-color: var(--gray-001);
  }

  :global(.navbar_container .navbar .left .shrink_btn .icon_wide) {
    display: block;
  }

  :global(.navbar_container .navbar .left .shrink_btn .icon_small) {
    display: none;
  }

  .navbar_container .navbar .right .display-email.free {
    padding: 5px 15px;
    border-radius: 8px;
    background-color: #0000000e;
    color: var(--gray-009);
    font-weight: 600;
  }

  .navbar_container .navbar .right .display-email.premium {
    border: none;
    background: linear-gradient(135deg, var(--blue-004), var(--green-004), var(--violet-005));
    padding: 2px;
    border-radius: 8px;
  }

  .navbar_container .navbar .right .display-email.premium .wrapper {
    background-color: #FFF;
    border: none;
    padding: 5px 15px;
    color: var(--gray-009);
    border-radius: 6px;
  }

  .navbar_container .navbar .right button.upgrade {
    padding: 5px 15px;
    border-radius: 8px;
    background-color: var(--violet-006);
    color: #FFF;
    font-weight: 600;
    border: none;
    cursor: pointer;
  }

  .navbar_container .navbar .right button.upgrade:hover {
    background-color: var(--violet-005);
  }

  .navbar_container .navbar .right picture {
    display: none;
  }

  @media only screen and (max-width : 768px) {
    .navbar_container {
      padding: 0 15px;
    }

    .navbar_container .navbar {
      flex-direction: row-reverse;      
    }

    .navbar_container .navbar .left .shrink_btn {
      padding: 0;
    }

    .navbar_container .navbar .left .shrink_btn:hover {
      background-color: transparent;
    }

    :global(.navbar_container .navbar .left .shrink_btn .icon_wide) {
      display: none;
    }

    :global(.navbar_container .navbar .left .shrink_btn .icon_small) {
      display: block;
      transform: rotate(180deg);
    }

    .navbar_container .navbar .right {
      display: flex;
      flex-direction: row;
      gap: 20px;
      align-items: center;
    }

    .navbar_container .navbar .right picture {
      display: block;
    }

    .navbar_container .navbar .right picture img {
      width: 30px;
      height: 30px
    }
  }
</style>
