<script lang="ts">
    import Icon from "$lib/assets/Remove-Red.png"

    let needsToBeConfirmed: boolean = $state(false)

    let {
        group,
        username,
        removeGroup
    } : {
        group: string;
        username: string;
        removeGroup: (username: string, group: string) => Promise<void>;
    } = $props()

    const OpenConfirmMenu = () => {
        needsToBeConfirmed = true
    }

    const yes = async () => {
        removeGroup(username, group)
        needsToBeConfirmed = false
    }

    const no = () => {
        needsToBeConfirmed = false
    }



</script>


    {#if !needsToBeConfirmed}
    <div class="content" class:active={!needsToBeConfirmed}>
        <button  title={`REMOVE ${group}`} class="remove" onclick={OpenConfirmMenu}>
            <img src={Icon} alt="Remove" title="Remove"/>
        </button>
    </div>
    {:else}
    <div class="content" class:active={needsToBeConfirmed}>
        <div class="confirm">
            <span>Remove Group?</span>
            <button class="yes" onclick={yes}>Yes</button>
            <button class="no" onclick={no}>No</button>
        </div>
    </div>
    {/if}




<style>

    img {
        height: 20px;
        width: 20px;
        transition: 0.3s ease;
    }

    div.confirm {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        text-align: center;
    }

    div.content {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: center;
        position: relative;
        align-self: flex-end;
        text-align: center;
        gap: 10px;
        min-height: 40px;
        box-sizing: border-box;
        margin-right: -30px;
        margin-bottom: -15px;
    }
    
    div.content.active {
        animation: 0.1s swap;
    }

    button {
        display: flex;
        justify-content: center;
        min-width: 40px;
        padding: 0 5px;
        border: none;
        background: none;
        border-radius: 5px;
        padding: 0.5rem;
        height: 100%;
        text-align: center;
        font-size: 12px;
    }

    button.remove {
        justify-content: flex-end;
    }


    button:hover img {
        transform: scale(1.2);
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

    @keyframes swap{
        from {
            transform: translateX(100px);
        } to {
            transform: translateX(0px);
        }
    }


    @media (max-width: 400px) {
        div.content {
            flex-wrap: nowrap;
            margin-right: 0px;
        }

        button {
            max-width: 100px;
            flex-grow: 1;
        }


    }



</style>