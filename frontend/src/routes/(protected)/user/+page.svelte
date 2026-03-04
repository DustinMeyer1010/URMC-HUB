<script lang="ts">
	import { page } from "$app/state";
	import CopyAllButton from "$lib/components/buttons/CopyAllButton.svelte";
	import CopyButton from "$lib/components/buttons/CopyButton.svelte";
	import PageLoading from "$lib/components/loading/PageLoading.svelte";
	import { convertToReadableTime }  from "$lib/parsers/time";
	import { readableOU } from "$lib/parsers/ou";
	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
    import { browser } from "$app/environment";
	import Groups from "./pages/Groups.svelte";

    let ou: string = $state("")
    let section: "PROFILE" | "GROUPS" | "ADD" | "LOCKOUT" | "DRIVE" = $state("PROFILE")
    const attributes = ["cn","username","phone","department","dn","urid","netid","relationship","pwdlastset","email","description","title"]
    let userPromise: Promise<any> = $state(new Promise(() => {}));
    let timer: number = $state(3)

    function changeSection(newSection: "PROFILE" | "GROUPS" | "ADD" | "LOCKOUT" | "DRIVE") {
        section = newSection
    }
    
    function startTimer() {
        setInterval(() => {
            timer -= 1;
        }, 1000);
    }

    async function fetchUser(): Promise<any> {
        ou = page.url.searchParams.get("dn") ?? ""
        if (ou == "") goto("/search"); 
        const res = await fetch(`http://localhost:8000/api/user?dn=${encodeURIComponent(ou)}&attributes=${attributes.join(",")}`);
        if (!res.ok && browser) {
            setTimeout(() => {
                goto("/search")
            }, 3000);

            startTimer()

            throw new Error(`Failed to find User with ${ou}`);
        }
        return res.json();
    }

    onMount(async () => {
        userPromise = fetchUser()
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
    {#if section == "PROFILE"}
        {#await userPromise}
            <PageLoading/>
        {:then user} 
            {@render Content(user)}
        {:catch error}
            <h1>{error.message} Redirecting in {timer}s</h1>
        {/await}
    {:else if section == "GROUPS"}
    <Groups dn={ou}></Groups>
    {/if}
</main>

{#snippet Content(u: any)}
    <section>
        <div>
            <CopyAllButton copyTemplate={"test"} />
            <CopyButton value={u.cn[0]} category="title"/>
            <CopyButton label="USERNAME" value={u.username[0] ?? "NA"} category="list-item"/>
            <CopyButton label="NETID" value={u.netid[0] ?? "NA"} category="list-item"/>
            <CopyButton label="URID" value={u.urid[0] ?? "NA"} category="list-item"/>
            <CopyButton label="PHONE" value={u.phone[0] ?? "NA"} category="list-item"/>
            <CopyButton label="DEPARTMENT" value={u.department[0] ?? "NA"} category="list-item"/>
            <CopyButton label="RELATIONSHIP" value={u.relationship[0] ?? "NA"} category="list-item"/>
            <CopyButton label="LASTPASSWORDSET" value={convertToReadableTime(u.pwdlastset[0] ?? "0")} category="list-item"/>    
            <CopyButton label="OU" value={readableOU(u.dn[0]) ?? "NA"} category="list-item"/>
        </div>
    </section>
{/snippet}

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
