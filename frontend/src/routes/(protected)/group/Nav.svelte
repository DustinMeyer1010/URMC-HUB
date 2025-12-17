

<script lang="ts">
	import type { Section } from "./[name]/state.svelte";

      

    let {
        sections,
        currentSection,
        swapSection
    } : {
        sections: Section[]
        currentSection: string
        swapSection: (section: Section) => void
    } = $props()


</script>


<nav>
    {#each sections as section}
        <button class:active={currentSection == section} onclick={() => swapSection(section)}>{section}</button>
    {/each}
</nav>



<style>
    nav {
        display: flex;
        align-items: center;
        justify-content: space-evenly;
        gap: 1rem;
        padding: 0rem 1rem;
        width: 100%;
        margin-top: 20px;
        margin-bottom: 20px;
    }

    button {
        padding: 1rem;
        font-size: 15px;
        width: 200px;
        background: var(--color-surface);
        border: none;
        border-radius: 10px;
        color: var(--text-color);
        text-wrap: nowarp;
        word-break: keep-all;
        white-space: nowrap;
    }

    button.active {
        background: var(--color-surface-lighter);
    }

    @media (max-width: 1000px) {
        nav {
            width: 90%;
            justify-content: center;
            flex-wrap: wrap;
        }
        button {
            padding: 1rem 2rem;
            width: fit-content;
            flex-grow: 1;
        }
    }

    @media (max-width: 800px) {
        button {
            padding: 1rem;
        }
    } 
</style>