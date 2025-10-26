<script lang="ts">
	import Nav from "../Nav.svelte";
	import { slide } from "svelte/transition";
    import type { Section } from "../types"
	import type { UserFullInfo } from "@t/user";
    import Profile from "./Profile.svelte"
	import Lockout from "./Lockout.svelte";

    let { data } : { data: UserFullInfo } = $props();
    const duration = 200;
    const delay = 250;



    let shownSection: Section = $state("PROFILE")

    function swapSection(section: Section)  {
        shownSection = section
    }
</script>


<main>
    <h1>{data.name}</h1>
    <Nav {swapSection}/>
    {#if shownSection == "PROFILE"}
        <Profile user={data}/>
    {:else if shownSection == "LOCKOUT"}
        <Lockout username={data.username}/>
    {:else if shownSection == "DRIVES"}
    <section in:slide={{delay: delay, duration: duration}} out:slide={{duration: duration}}>
        DRIVES
    </section>
    {:else if shownSection == "GROUPS"}
    <section in:slide={{delay: delay, duration: duration}} out:slide={{duration: duration}}>
        GROUPS
    </section>
    {/if}
</main>




<style>
    main {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0 5rem;
        flex-direction: column;
        gap: 1rem;
    }

    @media (max-width: 850px) {
        main {
            padding: 0 1rem;
        }
    }




</style>