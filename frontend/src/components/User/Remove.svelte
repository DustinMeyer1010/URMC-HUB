<script lang="ts">


    let needsToBeConfirmed: boolean = $state(false)

    let {
        group,
        removeGroup
    } : {
        group: string;
        removeGroup: (group: string) => void;
    } = $props()

    const OpenConfirmMenu = () => {
        needsToBeConfirmed = true
    }

    const yes = () => {
        removeGroup(group)
        needsToBeConfirmed = false
    }

    const no = () => {
        needsToBeConfirmed = false
    }



</script>

<div>
    
    {#if !needsToBeConfirmed}
        <button title={`REMOVE ${group}`} class="remove" onclick={OpenConfirmMenu}>Remove</button>
    {:else}
        <span>Remove Group?</span>
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
        font-size: 12px;
    }

    button.remove,
    button.no {
        color: var(--color-ad-disabled);
    }

    button.remove:hover,
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
            width: 100px;
            text-align: center;
        }
    }




</style>