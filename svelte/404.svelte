<script>
  import BackHomeBtn from "./_components/BackHomeBtn.svelte";

  let title = '#{title}';
  let description = '#{description}';
</script>

<section class="root_container">
  <div class="container">
    <h1>{title}</h1>
    <p>{description}</p>
    <BackHomeBtn />
  </div>
</section>

<style>
  .root_container {
    display: flex;
    justify-content: center;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: var(--gray-002);
    height: 100vh;
    width: 100vw;
  }

  .root_container .container {
    margin-top: 60px;
    background-color: white;
    border-radius: 10px;
    border: 1px solid var(--gray-003);
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 20px;
    height: fit-content;
    width: 400px;
    text-align: center;
  }

  .root_container .container h1 {
    font-size: var(--font-xxl);
    margin: 0;
  }

  .root_container .container p {
    margin: 0;
    font-size: var(--font-lg);
  }
</style>