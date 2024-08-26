<script>
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { RiDesignLayoutLeftLine } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { IsShrinkMenu } from '../uiState.js';

  export let user = /** @type {import('../types/user.js').User}*/ ({});
</script>

<header class="navbar_container">
  <nav class="navbar">
    <div class="left">
      <button class="shrink_btn" on:click={() => {
        IsShrinkMenu.set(!$IsShrinkMenu)
        localStorage.setItem('IsShrinkMenu', JSON.stringify($IsShrinkMenu))
      }}>
        <Icon src={RiDesignLayoutLeftLine} size={20} />
      </button>
      <!-- <h2 class="display_title">BenAkun</h2> -->
    </div>
    <div class="right">
      <span class="display_email">{user.email}</span>
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
    /* background: linear-gradient(66deg, rgb(230, 251, 254) 0%, rgb(237, 221, 251) 100%); */
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

  /* .navbar_container .navbar .left .display_title {
    font-size: var(--font-lg);
    font-weight: 700;
  } */

  .navbar_container .navbar .right .display_email {
    padding: 5px 15px;
    border-radius: 8px;
    background-color: #0000000e;
    color: var(--gray-009);
    font-weight: 600;
  }

  /* Responsive to mobile device */
  @media only screen and (max-width : 768px) {}
</style>
