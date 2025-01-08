<script>
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import {
    RiDesignLayoutLeftLine,
    RiBusinessBarChartHorizontalLine
  } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { IsShrinkMenu } from '../uiState.js';
  import { IsUnixTimeExpired } from '../xHelper';
  // import { UserPurchaseSupport } from '../../jsApi.GEN';

  export let user = /** @type {import('../types/user.js').User}*/ ({});
  let viewportWidth = window.innerWidth;

  // async function purchaseSupport() {
  //   await UserPurchaseSupport({}, async (res) => {
  //     console.log(res);
  //   });
  // }
</script>

<header class="navbar_container">
  <nav class="navbar">
    <div class="left">
      <button class="shrink_btn" on:click={() => {
        IsShrinkMenu.set(!$IsShrinkMenu);
        if (viewportWidth > 768) {
          localStorage.setItem('IsShrinkMenu', JSON.stringify($IsShrinkMenu))
        }
      }}>
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
      <!-- <button on:click={purchaseSupport}>purchase</button> -->
      <picture>
        <img src="/assets/icons/benakun-logo.png" alt="" />
      </picture>
      <span class="display-email { IsUnixTimeExpired(user.supportExpiredAt) ? 'free' : 'premium'} ">
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

  /* .navbar_container .navbar .left .display_title {
    font-size: var(--font-lg);
    font-weight: 700;
  } */

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
