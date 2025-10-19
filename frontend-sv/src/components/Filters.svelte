<script lang="ts">

    type Section = 'COMPUTERS' | 'USERS' | "GROUPS" | 'PRINTERS' | 'DRIVES'
    let selected: Section = $state<Section>('USERS')

    function toggleActive(name: Section) {
        selected = name;
    }

</script>


<section>
    <button style="order: 4;" class:active={selected === 'COMPUTERS'} onclick={() => toggleActive("COMPUTERS")}>Computers</button>
    <button style="order: 5;" class:active={selected === 'USERS'} onclick={() => toggleActive("USERS")}>Users</button>
    <button style="order: 3;" class:active={selected === 'GROUPS'} onclick={() => toggleActive("GROUPS")}>Groups</button>
    <button style="order: 2;" class:active={selected === 'PRINTERS'} onclick={() => toggleActive("PRINTERS")}>Printers</button>
    <button style="order: 1;" class:active={selected === 'DRIVES'} onclick={() => toggleActive("DRIVES")}>Share Drives</button>
</section>



<style>
    section {
        display: flex;
        gap: 3rem;
    }

    button {
        padding: 0.5rem 1rem;
        width: 150px;
        border: 2px solid var(--complement);
        border-radius: 20px;
        background: transparent;
        color: var(--complement);
        transition: 0.3s ease;
        font-weight: bold;
    }

    button.active {
        background: var(--complement);
        color: var(--text)
    }

    button:hover {
        background: var(--complement);
        color: var(--text)
    }

    @media (max-width: 800px) {
        button {
            width: 100px;
            padding: 5px;
            font-size: 12px;
        }

        section {
            flex-wrap: wrap-reverse;
            justify-content: center;
            align-items: center;
            gap: 1rem;
        }
    }

</style>