<script lang="ts">
	import Nav from "../Nav.svelte";
    import type { Section } from "../types"
	import type { UserFullInfo } from "@t/user";
    import Profile from "./Profile.svelte"
	import Lockout from "./Lockout.svelte";
	import Drive from "./Drive.svelte";
	import Groups from "./Groups.svelte";

    let { data } : { data: UserFullInfo } = $props();

    let shownSection: Section = $state("PROFILE")

    function swapSection(section: Section)  {
        shownSection = section
    }

</script>



<nav>
    <h1>{data.name}</h1>
    <Nav {swapSection}/>
</nav>
<main>
    {#if shownSection == "PROFILE"}
        <Profile user={data}/>
    {:else if shownSection == "LOCKOUT"}
        <Lockout username={data.username}/>
    {:else if shownSection == "DRIVES"}
        <Drive groups={data.member_of}/>
    {:else if shownSection == "GROUPS"}
        <Groups groups={data.member_of}/>
    {/if}
</main>




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
        z-index: 10;
        background: var(--color-bg);
    }

    @media (max-width: 850px) {
        main {
            padding: 0 1rem;
        }
    }




</style>