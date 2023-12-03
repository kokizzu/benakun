<script>
    import { onMount } from "svelte";

  let ok = '#{ok}';
  let error = '#{error}';
  let message = '#{message}';
  let email = '#{email}';

  onMount(() => {
    console.log('ok', ok)
    console.log('message', message)
  })
</script>

<section class="verify_email_container">
  <div class="main_content">
    <h1>/*! title */</h1>
    {#if ok === 'true'}
      <p>{email}</p>
    {/if}
    {#if error}
      <p class="error">{error}</p>
    {/if}
    <a href="/">
      <p>Back to main page</p>
    </a>
  </div>
</section>