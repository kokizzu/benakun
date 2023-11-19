<script>
    import Icon from 'svelte-icons-pack/Icon.svelte';

    export let href = '';
    export let icon = '';
    export let title = '';
</script>
<a href={href} class:active={window.location.pathname === href}>
    <Icon size={20} className={window.location.pathname === href  ? 'icon_active' : 'icon_dark'} src={icon}/>
    <span>{title}</span>
</a>
<style>

    a {
        display          : block;
        padding          : 0.3rem 0 0.1rem 0.5rem;
        color            : #475569;
        text-decoration  : none;
        margin           : 0;
        font-size        : 0.875rem !important;
        line-height      : 1.25rem;
        font-weight      : 700;
        text-transform   : uppercase;
        text-align       : left;
        background-color : transparent;
        border           : none;
        flex-direction   : row;
        align-items      : center;
        gap              : 15px;
    }

    a:hover {
        color : #64748B;
    }

    a span {
        display : inline;
    }

    a span {
        position : absolute;
        padding  : 5px;
    }

    .active { /*ACTIVE Navigation*/
        color : #4444EF !important;
    }
</style>