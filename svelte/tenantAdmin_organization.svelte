<script>
  // @ts-nocheck
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';

  let segments = {/* segments */};
  let user = {/* user */};
</script>

<div class="root_layout">
  <div class="root_container">
    <SideMenu access={segments} />
    <div class="root_content">
      <Navbar {user} />
      <div class="content">
        <p>Tenant Admin Organization</p>
      </div>
      <Footer />
    </div>
  </div>
</div>

<style>
</style>
