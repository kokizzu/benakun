<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';

  let segments = {/* segments */};
  let user = {/* user */};
  let orgs = [/* orgs */];

  onMount(() => {
    console.log('Segments:', segments);
    console.log('User:', user);
    console.log('Orgs:', orgs);
  })
</script>

<MainLayout>
  <div>
    <p>TenantAdmin Organization</p>
  </div>
</MainLayout>

<style>
</style> 