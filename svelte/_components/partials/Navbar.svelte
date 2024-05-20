<script>
  /** @type {import('../types/user.js').User}*/ // @ts-ignore
  export let user = {};
</script>

<header class="navbar_container">
  <nav class="navbar">
    <div class="left">
      <h2 class="display_title">BenAkun</h2>
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

  .navbar_container .navbar .left .display_title {
    font-size: var(--font-lg);
    font-weight: 700;
  }

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
