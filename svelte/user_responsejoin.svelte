<script>
  let joined = '#{joined}';
  let error = '#{error}';
  let email = '#{email}';
</script>

<section class="verify_email_container">
  <div class="main_content">
    <h1>/*! title */</h1>
    {#if joined === 'true'}
      <p>{email} is joined</p>
    {/if}
    {#if error}
      <p class="error">{error}</p>
    {/if}
    <a href="/">
      <p>Back to main page</p>
    </a>
  </div>
</section>