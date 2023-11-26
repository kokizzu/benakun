<script>
  export let id;
  export let value;
  export let label;
  export let type = 'text';

  function typeAction(node) {
    node.type = type;
  }
</script>

<div class="input_box">
  <label for={id}>{label}</label>
  <input bind:value {id} use:typeAction />
</div>

<style>
  .input_box {
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input_box label {
    font-size: var(--font-base);
    margin-left: 10px;
  }

  .input_box input {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 10px 12px;
  }

  .input_box input:focus {
    border-color: var(--blue-005);
    outline: 1px solid var(--blue-005);
  }
</style>
