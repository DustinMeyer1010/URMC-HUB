<script lang="ts">
	import Nav from "@components/User/Nav.svelte";
    import Profile from "@components/User/Profile.svelte"
	import Lockout from "@components/User/Lockout.svelte";
	import Drive from "@components/User/Drive.svelte";
	import Groups from "@components/User/Groups.svelte";
	import Add from "@components/User/Add.svelte";

	import { onMount } from "svelte";
	import { page } from "$app/state";
	import { goto } from "$app/navigation";
    import type { UserFullInfo } from "@t/user";

    import  { type Section, Sections } from "@t/section";

    let { data } : { data: UserFullInfo } = $props();
    let shownSection: Section = $state("PROFILE")
    let isMounted: boolean = $state(false)

    onMount(() => {
        let urlParams = page.url.searchParams;
        shownSection = urlParams.get("section")?.toUpperCase() as Section ?? "PROFILE"
        setURL()
        isMounted = true
    })

    const swapSection = (section: Section) => {
        shownSection = section
        setURL()
    }

    const setURL = () => {
        goto(`?section=${shownSection}`, { replaceState: true, keepFocus: true, noScroll: true })
    }

</script>



<nav>
    <h1>{data.name}</h1>
    <Nav sections={Sections} {swapSection}/>
</nav>

{#if isMounted}
    <main>
        {#if shownSection == "PROFILE"}
            <Profile user={data}/>
        {:else if shownSection == "LOCKOUT"}
            <Lockout username={data.username}/>
        {:else if shownSection == "DRIVES"}
            <Drive groups={data.member_of}/>
        {:else if shownSection == "GROUPS"}
            <Groups currentUser={data.username} groups={data.member_of}/>
        {:else if shownSection == "ADD"}
            <Add currentUser={data.username}/>
        {/if}
    </main>
{/if}





<style>
    main {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 1rem 5rem;
        flex-direction: column;
        gap: 1rem;
        margin-top: 150px;

    }

    nav {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        position: fixed;
        padding: 1rem 0rem;
        left: 0;
        top: 55px;
        gap: 1rem;
        width: 100%;
        margin-bottom: 1rem;
        z-index: 3;
        background: var(--color-bg);
    }

    @media (max-width: 800px) {
        main {
            margin-top: 175px;
            padding: 0 3rem;
        }
    }

    @media (max-width: 784px) {
        main {
            margin-top: 250px;
            padding: 0 2rem;
        }
    }

    @media (max-width: 465px) {
        main {
            margin-top: 300px;
            padding: 0 1rem;
        }
    }




</style>