<script lang="ts">
    let needsToBeConfirmed = $state(false)

    let {
        name,
        title,
        value,
        username,
        action
    } : {
        name: string
        title: string
        value: string
        username: string
        action: (usernaem: string, name: string) => void
    } = $props()

    const yes = () => {
        needsToBeConfirmed = false
        action(username, name)
    }

    const no = () => {
        needsToBeConfirmed = false  
    }

</script>


<div>
    {#if !needsToBeConfirmed}
        <button title={`add ${name}`} class="add" onclick={() => needsToBeConfirmed = true}>{value}</button>
    {:else}
        <span>{title}</span>
        <button class="yes" onclick={yes}>Yes</button>
        <button class="no" onclick={no}>No</button>
    {/if}
</div>


<style>
    
    div {
        position: absolute;
        bottom: 5px;
        right: 5px;
        display: flex;
        flex-direction: column;
        box-sizing: border-box;
        gap: 3px;
    }

    button {
        border: none;
        background: none;
        border-radius: 10px;
        padding: 0.5rem;
        height: 100%;
        min-width: 30px;
        text-align: center;
        color: var(--color-success);
        font-size: 12px;
    }

    span {
        word-break: keep-all;
    }

    button.no {
        color: var(--color-ad-disabled);
    }

    button.no:hover {
        background: var(--color-ad-disabled-opacity-20);
    }

    button.yes {
        color: var(--color-success)
    }
    
    button.yes:hover {
        background: var(--color-success-opacity-20);
    }

    @media (max-width: 800px) {
        div {
            flex-direction: row;
            position: relative;
            align-self: center;
            text-align: center;
            align-items: center;
            justify-content: center;
            width: 100%;
            gap: 10px;
            padding: 0 1rem;
        }
        button {
            flex-grow: 1;
        }

    }




</style>