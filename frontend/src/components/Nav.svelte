<script lang="ts">

    import urmc from '$lib/assets/URMC.png';
	import ToggleTheme from './ToggleTheme.svelte';
    import menuIcon from "$lib/assets/menu-color-text-dark.png"
	import { onMount } from 'svelte';

    let showMenu: boolean = $state(false)

    let width = $state(0);
    

    onMount(() => {
        width = window.innerWidth;

        const handleResize = () => {
            width = window.innerWidth;
        }

        window.addEventListener("resize", handleResize);

        return () => window.removeEventListener("resize", handleResize);
    });



    let breakDown: boolean = $derived.by(() => {
        return width <= 500
    })

</script>

{#if !breakDown}
<nav>
    <a href="/"><img src={urmc} alt=""></a>

    <ul>
        <li>
            <a href="/">Search</a>
        </li>
        <li>
            <a href="/bookmarks">Bookmarks</a>
        </li>
        <li class="dropdown">
            <button onclick={() => showMenu = !showMenu}><img src={menuIcon} alt=""></button>
            {#if showMenu}
                <div>
                    <a href="/api">API Docs</a>
                    <a href="/bulk_add">Bulk User Add</a>
                    <a href="/bulk_remove">Bulk User Remove</a>
                    <a href="/bulk_lookup">Bulk User Lookup</a>
                    <ToggleTheme/>
                </div>
                
            {/if}
        </li>
    </ul>
</nav>
{:else}
<nav>
    <a href="/"><img src={urmc} alt=""></a>
    
    <ul>
        <li class="dropdown">
            <button onclick={() => showMenu = !showMenu}><img src={menuIcon} alt=""></button>
            {#if showMenu}
                <div>
                    <a href="/">Search</a>
                    <a href="/bookmarks">Bookmarks</a>
                    <a href="/api">API Docs</a>
                    <a href="/bulk_add">Bulk User Add</a>
                    <a href="/bulk_remove">Bulk User Remove</a>
                    <a href="/bulk_lookup">Bulk User Lookup</a>
                    <ToggleTheme/>
                </div>
                
            {/if}
        </li>
    </ul>
</nav>
{/if}


<style>
    nav {
        display: flex;
        z-index: 10;
        justify-content: space-between;
        align-items: center;
        position: fixed;
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
        right: 0px;
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

</style>

