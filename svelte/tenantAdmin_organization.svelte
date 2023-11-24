<script>
    // @ts-nocheck
    import Menu from './_components/partials/SideMenu.svelte';
    import ProfileHeader from './_components/partials/Navbar.svelte';
    import Footer from './_components/Footer.svelte';

    let segments = {/* segments */};
    let user = {/* user */};

</script>


<section class='dashboard'>
    <Menu access={segments} />
    <div class='dashboard_main_content'>
        <ProfileHeader {user}></ProfileHeader>
        <div class='content'>
            <section class='tableview_container'>
                tenantAdmin organization
            </section>
        </div>
        <Footer></Footer>
    </div>
</section>

<style>
</style>