<script>
  import SideMenu from '../_components/partials/SideMenu.svelte';
  import Navbar from '../_components/partials/Navbar.svelte';
  import Footer from '../_components/partials/Footer.svelte';

  /** @type {import('../_components/types/access').Access} */ // @ts-ignore
  let segments = {/* segments */};
  /** @type {import('../_components/types/user').User} */ // @ts-ignore
  let user = {/* user */}
</script>

<div class="root_layout">
  <div class="root_container">
    <SideMenu access={segments} />
    <div class="root_content">
      <Navbar {user} />
      <div class="content">
        <slot />
      </div>
      <Footer />
    </div>
  </div>
</div>

<style>
  .root_layout {
    display: block;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    height: 100vh;
    width: 100vw;
  }

  .root_layout .root_container {
    display: flex;
    flex-direction: row;
    height: 100%;
    width: 100%;
  }

  .root_layout .root_container .root_content {
    flex-grow: 1;
    overflow-y: auto;
    height: 100%;
    position: relative;
    background-color: var(--gray-001);
  }

  .root_layout .root_container .root_content .content {
    width: 100%;
    min-height: 80dvh;
    padding: 20px;
  }
</style>