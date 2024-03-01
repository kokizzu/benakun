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
        <div class="coa_levels">
          <div class="coa">
            <div class="parent">
              <span>Assets</span>
              <div class="options">
                <button class="btn">Add</button>
                <button class="btn">Edit</button>
              </div>
            </div>
            <div class="childs">
              <span>1.1 Aktiva tetap</span>
              <span>1.2 xxxxxxxxx</span>
              <span>1.3 xxxxxxxx</span>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  </div>
</div>

<style>
</style>