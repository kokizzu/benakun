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
        <main><slot /></main>
        <Footer />
      </div>
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
    height: 100%;
		width: 100%;
		display: flex;
  }

  .root_layout .root_container .root_content {
		margin-left: var(--sidemenu-width);
		display: flex;
		flex-direction: column;
		-webkit-box-orient: vertical;
		-webkit-box-direction: normal;
		overflow-y: auto;
		min-height: calc(100vh - var(--navbar-height));
		transition: 0.3s;
		width: 100%;
  }

  .root_layout .root_container .root_content .content {
    width: 100%;
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 30px;
    min-height: calc(100vh - var(--navbar-height));
    max-height: fit-content;
    overflow: inherit;
  }
</style>