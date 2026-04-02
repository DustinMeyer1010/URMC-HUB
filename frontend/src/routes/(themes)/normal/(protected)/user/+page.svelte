<script lang="ts">
	import { page } from "$app/state";
	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
	import Groups from "./views/Groups.svelte";
	import Lockout from "./views/Lockout.svelte";
	import About from "./views/About.svelte";

    let userDN: string = $state("")
    let section: "PROFILE" | "GROUPS" | "ADD" | "LOCKOUT" | "DRIVE" = $state("PROFILE")
    let loading: boolean = $state(true)

    function changeSection(newSection: "PROFILE" | "GROUPS" | "ADD" | "LOCKOUT" | "DRIVE") {
        section = newSection
    }


    onMount(async () => {
        userDN = page.url.searchParams.get("dn") ?? ""
        if (userDN == "") goto("/normal/search"); 
        loading = false
    })

</script>

<!-- * Waits for the response for the server to render content -->
 <main>
    <nav>
        <button class:active={section == "PROFILE"} onclick={() => changeSection("PROFILE")}>PROFILE</button>
        <button class:active={section == "LOCKOUT"} onclick={() => changeSection("LOCKOUT")}>LOCKOUT</button>
        <button class:active={section == "DRIVE"} onclick={() => changeSection("DRIVE")}>DRIVE</button>
        <button class:active={section == "ADD"} onclick={() => changeSection("ADD")}>ADD</button>
        <button class:active={section == "GROUPS"} onclick={() => changeSection("GROUPS")}>GROUPS</button>
    </nav>
    {#if !loading}
        {#if section == "PROFILE"}
            <About userDN={userDN} />
        {:else if section == "GROUPS"}
            <Groups userDN={userDN}></Groups>
        {:else if section == "LOCKOUT"}
            <Lockout userDN={userDN}/>
        {/if}
    {/if}
</main>



<style>

    nav {
        display: flex;
        flex-direction: row;
        position: sticky;
        top: 75px;
        gap: 1rem;
        width: 90%;
        background: var(--color-surface);
        padding: 10px;
        align-self: center;
        border-radius: 20px;
        transition: 0.3s;
    }

    button {
        font-weight: bold;
        color: var(--text);
        font-size: 14px;
        padding: 10px;
        width: 200px;
        flex-grow: 1;
        background: var(--color-surface-hover);
        border-radius: 20px;
        border: none;
        box-shadow: 
        inset -5px -5px 10px 2px rgba(255,255,255,0.1),
        inset 5px 5px 10px 2px rgba(0,0,0,0.4);
        transition: 0.3s ease-in-out;
    }

    button.active {
        box-shadow: 
        inset 5px 5px 10px 2px rgba(255,255,255,0.05),
        inset -5px -5px 10px 2px rgba(0,0,0,0.05),
        5px 5px 10px 2px rgba(0,0,0,0.3);
    }

    main {
        width: 100%;
        color: var(--text);
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    @media (max-width: 700px) {
        nav {
            width: 100%;
        }

    }

</style>
