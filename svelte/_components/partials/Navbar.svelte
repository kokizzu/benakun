<script>
  // @ts-nocheck
  import {onMount} from 'svelte';

  export let user = {};
  onMount(() => {
    console.log('onMount.ProfileHeader =', user);
  });
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
    top: 0;
    position: sticky;
    z-index: 800;
    background-color: #FFF;
    height: 60px;
    padding: 0 40px;
    display: flex;
    align-items: center;
    border-bottom: 1px solid var(--gray-002);
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
    border-radius: 9999px;
    border: 1px solid var(--gray-003);
  }

  /* Responsive to mobile device */
  @media only screen and (max-width : 768px) {}
</style>
