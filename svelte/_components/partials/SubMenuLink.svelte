<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';

  export let href = '';
  export let icon = {};
  export let title = '';
</script>

<a href={href} class:active={window.location.pathname === href}>
  <Icon size="18" className={window.location.pathname === href  ? 'icon_active' : 'icon_dark'} src={icon}/>
  <span>{title}</span>
</a>

<style>
  a {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-decoration: none;
    text-transform: capitalize;
    width: 100%;
    color: var(--gray-007);
    font-size: var(--font-md);
    padding: 10px 12px;
    border-radius: 8px;
  }

  a:hover {
    background-color: var(--gray-001);
  }

  a:active {
    background-color: var(--gray-002);
  }

  a.active {
    color: var(--violet-005);
    background-color: var(--violet-transparent);
  }
</style>