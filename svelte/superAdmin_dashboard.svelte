<script>
    // @ts-nocheck
    import Menu from './_components/Menu.svelte';
    import AdminSubMenu from './_components/AdminSubMenu.svelte';
    import ProfileHeader from './_components/ProfileHeader.svelte';
    import Footer from './_components/Footer.svelte';

    let segments = {/* segments */};
    let user = {/* user */};

</script>


<section class='dashboard'>
    <Menu access={segments} />
    <div class='dashboard_main_content'>
        <ProfileHeader {user}></ProfileHeader>
        <AdminSubMenu></AdminSubMenu>
        <div class='content'>
            <section class='tableview_container'>
                superAdmin dashboard
            </section>
        </div>
        <Footer></Footer>
    </div>
</section>

<style>
</style>