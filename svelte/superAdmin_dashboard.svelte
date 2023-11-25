<script>
  // @ts-nocheck
  import SideMenu from './_components/partials/SideMenu.svelte';
  import AdminSubMenu from './_components/AdminSubMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import { onMount } from 'svelte';

  let segments = {/* segments */};
  let user = {/* user */};

  onMount(()=>{
    console.log('User: ', user)
  })
</script>


<div class="root_layout">
  <SideMenu access={segments}/>
  <div class="root_container">
    <Navbar {user} />
    <div class="root_content">
      <AdminSubMenu />
      <p>TODO fill with proper menu</p>
    </div>
    <Footer />
  </div>
</div>

<style>
</style>