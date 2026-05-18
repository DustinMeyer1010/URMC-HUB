<script lang="ts">
	import Groups from "./views/Groups.svelte";
	import Lockout from "./views/Lockout.svelte";
	import About from "./views/About.svelte";
	import Drives from "./views/Drives.svelte";
	import type { PageData } from "./$types";
	import Add from "./views/Add.svelte";
	import { goto } from "$app/navigation";

    let section: "GROUPS" | "ADD" | "LOCKOUT" | "DRIVE" = $state("GROUPS")

    let { data } : { data: PageData } = $props();

    const userDN = data.userDN
    section = data.section as "GROUPS" | "ADD" | "LOCKOUT" | "DRIVE"
    $inspect(userDN)
    function changeSection(newSection: "GROUPS" | "ADD" | "LOCKOUT" | "DRIVE") {
        updateFilterParms(newSection)
        section = newSection
    }

    function updateFilterParms(filter: string) {
        goto(`?dn=${encodeURIComponent(userDN)}&filter=${filter}`, { replaceState: true, keepFocus: true, noScroll: true })
    }

    $effect(() =>{ 
        document.body.style.overflow = "hidden"

        return () => {
            document.body.style.overflow = 'auto';
        };
    })

</script>



<!-- * Waits for the response for the server to render content -->
 <main>
    <nav>
        <button class:active={section == "LOCKOUT"} onclick={() => changeSection("LOCKOUT")}>LOCKOUT</button>
        <button class:active={section == "DRIVE"} onclick={() => changeSection("DRIVE")}>DRIVE</button>
        <button class:active={section == "ADD"} onclick={() => changeSection("ADD")}>ADD</button>
        <button class:active={section == "GROUPS"} onclick={() => changeSection("GROUPS")}>GROUPS</button>
    </nav>
    <section>
        <div id="left">
            <About {userDN} />
        </div>
        <div id="right">
            {#if section == "GROUPS"}
                <Groups {userDN} />
            {:else if section == "LOCKOUT"}
                <Lockout {userDN} />
            {:else if section == "DRIVE"}
                <Drives {userDN} />
            {:else if section == "ADD"}
                <Add {userDN} />
            {/if}
        </div>
    </section>

</main>



<style>

    nav {
        display: flex;
        flex-direction: row;
        position: sticky;
        gap: 1rem;
        width: 90%;
        z-index: 100;
        background: var(--color-surface);
        padding: 10px;
        align-self: center;
        border-radius: 20px;
        transition: 0.3s;
    }

    section {
        display: flex;
        flex-direction: row;
        padding: 1rem;
        gap: 1rem;
        width: 100%;
        overflow: hidden;
        flex-grow: 1;
        justify-items: center;
        align-items: center;
    }

    div#right {
        display: flex;
        flex-direction: column;
        overflow-y: scroll;
        height: 100%;
        flex-basis: 80%;
    }

    div#left {
        flex-basis: 40%;
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
        width: 100vw;
        color: var(--text);
        display: flex;
        flex-direction: column;
        align-items: center;
        height: calc(100vh - 40px);
        overflow: hidden
    }

    @media (max-width: 700px) {
        nav {
            width: 100%;
        }

        section {
            flex-direction: column;
            padding: 0.3rem;
        }

        div#left {
            flex-basis: 20%;
        }

        div#right {
            width: 100%;
        }

    }

</style>
