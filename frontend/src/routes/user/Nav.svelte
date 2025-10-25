

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
</nav>


<style>
    nav {
        display: flex;
        justify-content: space-evenly;
        gap: 1rem;
        width: 50%;
    }

    button {
        padding: 1rem 4rem;
        font-size: 15px;
        background: var(--color-surface);
        border: none;
        border-radius: 10px;
        color: var(--text-color)
    }
</style>