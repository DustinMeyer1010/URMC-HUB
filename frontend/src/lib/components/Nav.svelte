<script lang="ts">

    import urmc from '$lib/assets/URMC.png';
	import ToggleTheme from './ToggleTheme.svelte';
    import menuIcon from "$lib/assets/menu-color-text-dark.png"
	import { onMount } from 'svelte';
	import { NavStateClass } from './NavState.svelte';



    let NavState: NavStateClass = new NavStateClass()
    let CurrentUser: string = $state("")
    let BookmarksRoute: string = $derived.by(() => {
        if (CurrentUser == "") {
            return "/bookmarks"
        }
        return `/bookmarks/${CurrentUser}`
    })
    onMount(() => {
        NavState.ResizeSetup()
        CurrentUser = localStorage.getItem("agent") ?? ""

        return () => window.removeEventListener("resize", NavState.HandleResize);
    });

</script>
<nav >
    <a href="/"><img  src={urmc} alt=""></a>
    {#if !NavState.Compact}
        {@render RegularNav()}
    {:else}
        {@render DropDownOnly()}
    {/if}
</nav>

{#snippet RegularNav()}
    <ul>
        <li>
            <a href="/">Search</a>
        </li>
        <li>
            <a href={BookmarksRoute}>Bookmarks</a>
        </li>
        <button bind:this={NavState.DropDownButton} onclick={NavState.OpenMenu}><img src={menuIcon} alt=""></button>
        <li class="dropdown">
            {#if NavState.ShowMenu}
                <div>
                    <a href="/bulk-lookup">Bulk User Lookup</a>
                    <a href="/api">API Docs</a>
                    <ToggleTheme/>
                </div>
                
            {/if}
        </li>
    </ul>
{/snippet}


{#snippet DropDownOnly()}
    <ul>
        <li class="dropdown">
            <button bind:this={NavState.DropDownButton} onclick={NavState.OpenMenu}><img src={menuIcon} alt=""></button>
            {#if NavState.ShowMenu}
                <div>
                    <a href="/">Search</a>
                    <a href={BookmarksRoute}>Bookmarks</a>
                    <a href="/bulk-lookup">Bulk User Lookup</a>
                    <a href="/api">API Docs</a>
                    <ToggleTheme/>
                </div>
            {/if}
        </li>
    </ul>
{/snippet}


<style>
    nav {
        display: flex;
        z-index: 10;
        justify-content: space-between;
        align-items: center;
        position: sticky;
        box-sizing: border-box;
        padding: 0.2rem 1rem;
        top: 0;
        left: 0;
        width: 100%;
        box-shadow: var(--shadow-soft);
        background-color: var(--color-surface);
    }

    div {
        position: absolute;
        display: flex;
        flex-direction: column;
        background: var(--color-surface);
        top: 50px;
        right: -50px;
        z-index: 10;
        min-width: 300px;
        box-sizing: border-box;
        box-shadow: var(--shadow-strong);
    }

    li.dropdown {
        position: relative;
    }

    div a:hover {
        background: var(--color-surface-lighter);
    }

    div a {
        padding: 1rem;
    }

    li:hover {
        color: var(--color-text-opacity-30);
    }



    button {
        background: transparent;
        border: none;
        padding: 0;
        margin: 0;
        order: 3;
    }

    button img {
        height: 25px;
    }

    img {
        height: 40px;
        border-radius: 5px;
    }

    ul {
        display: flex;
        gap: 2rem;
        list-style: none;
    }

    a {
        color: var(--color-text);
        text-decoration: none;
        font-size: 18px;
        transition: color 0.3s ease;
    }

    a:visited {
        color: var(--color-text);
    }

    a:hover {
        color: var(--primary)
    }

    @media (max-width: 500px) {
        div {
            right: 0;
        }
    }


</style>

