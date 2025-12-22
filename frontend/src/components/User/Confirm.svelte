<script lang="ts">

    import Icon from "$lib/assets/Add-Green.png"

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
        <button title={`ADD ${name}`} class="add" onclick={() => needsToBeConfirmed = true}>
            <img src={Icon} alt="add"/>
        </button>
    {:else}
        <span>{title}</span>
        <button class="yes" onclick={yes}>Yes</button>
        <button class="no" onclick={no}>No</button>
    {/if}
</div>


<style>

    img {
        height: 15px;
        width: 15px;
        transition: 0.3s ease;
        
    }
    
    div {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: center;
        position: relative;
        align-self: flex-end;
        text-align: center;
        gap: 10px;
        box-sizing: border-box;
        margin-right: -35px;
        margin-bottom: -15px;
    }

    button {
        display: flex;
        padding: 5px 0;
        justify-content: center;
        min-width: 50px;
        border: none;
        background: none;
        border-radius: 5px;
        height: 100%;
        text-align: center;
        font-size: 12px;

    }

    button.add {
        justify-content: flex-end;
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

    button:hover img {
        transform: scale(1.5);
    }

    @media (max-width: 400px) {
        div {
            flex-wrap: nowrap;
            margin-right: -5px;
        }
        button {
            flex-grow: 1;
        }

    }




</style>