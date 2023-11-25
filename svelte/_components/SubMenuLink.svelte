<script>
  //@ts-nocheck
  import Icon from 'svelte-icons-pack/Icon.svelte';

  export let href = '';
  export let icon = '';
  export let title = '';
</script>

<a href={href} class:active={window.location.pathname === href}>
  <Icon size={22} className={window.location.pathname === href  ? 'icon_active' : 'icon_dark'} src={icon}/>
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
    width: fit-content;
    color: var(--gray-007);
    font-size: var(--font-md);
  }

  a:hover {
    color: var(--purple-002);
  }

  a.active {
    font-weight: 600;
    color: var(--purple-002);
  }
</style>