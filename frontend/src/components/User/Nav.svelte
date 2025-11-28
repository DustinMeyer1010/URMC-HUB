

<script lang="ts">  
	import type { Section } from "../../types/section";

    let {
        sections,
        swapSection
    } : {
        sections: Section[]
        swapSection: (section: Section) => void
    } = $props()


</script>

<nav>
    {#each sections as section}
        <button onclick={() => swapSection(section)}>{section}</button>
    {/each}
</nav>


<style>
    nav {
        display: flex;
        align-items: center;
        justify-content: space-evenly;
        gap: 1rem;
        width: 90%;
    }

    button {
        padding: 1rem;
        font-size: 12px;
        max-width: 200px;
        width: 200px;
        background: var(--color-surface);
        border: none;
        border-radius: 10px;
        color: var(--text-color);
        text-wrap: nowarp;
        word-break: keep-all;
        white-space: nowrap;
    }

    button:hover {
        background: var(--color-surface-hover);
    }

    @media (max-width: 1000px) {
        nav {
            width: 90%;
            justify-content: center;
            box-sizing: border-box;
            flex-wrap: wrap;
        }
        button {
            padding: 1rem 1rem;
            width: fit-content;
            flex-grow: 1;
        }
    }

    @media (max-width: 800px) {
        button {
            padding: 1rem;
            font-size: 12px;
        }
    } 
</style>