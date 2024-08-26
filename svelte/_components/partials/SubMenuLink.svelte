<script>
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { IsShrinkMenu } from '../uiState';

  export let href = '';
  export let icon = /** @type {import('svelte-icons-pack').IconType} */ ({});
  export let title = '';

  const pathname = window.location.pathname;
</script>

<a
  href={href}
  title={title}
  class:active={pathname === href || pathname.includes(href)}
  >
  <Icon
    size="18"
    className={pathname === href || pathname.includes(href)
      ? 'icon_active icon'
      : 'icon_dark icon'
    }
    src={icon}
  />
  {#if !$IsShrinkMenu}
    <span>{title}</span>
  {/if}
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
    font-size: 15px;
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

  a span {
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  :global(a .icon) {
    flex-shrink: 0;
  }
</style>