

<script lang="ts">
    import type { Section } from "./types"    

    let {
        swapSection
    } : {
        swapSection: (section: Section) => void
    } = $props()


</script>

<nav>
    <button onclick={() => swapSection("PROFILE")}>Profile</button>
    <button onclick={() => swapSection("LOCKOUT")}>Lockout</button>
    <button onclick={() => swapSection("DRIVES")}>Drives</button>
    <button onclick={() => swapSection("GROUPS")}>Groups</button>
    <button onclick={() => swapSection("GROUPS")}>Add/Remove</button>
</nav>


<style>
    nav {
        display: flex;
        align-items: center;
        justify-content: space-evenly;
        gap: 1rem;
        padding: 0rem 1rem;
        width: 100%;
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