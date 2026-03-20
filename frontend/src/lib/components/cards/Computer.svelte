<script lang="ts">
    import { Computer } from "$lib/types/computer";
	import { ComputerStateClass } from "./states/ComputerState.svelte";
	import CopyButton from "../buttons/CopyButton.svelte";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
	import Card from "./Card.svelte";
	import { Icons } from "$lib/managers/icons";
    
    let {
        item,
        idx
    } : {
        item: Computer.CardInfo
        idx: number
    } = $props()

    let ComputerState: ComputerStateClass = new ComputerStateClass(item)

</script>

<Card {idx} >
    {#snippet header()}
        {@render Link()}
        <CopyAllButton icon={Icons.COPY} copiedIcon={Icons.COPY_SUCCESSFUL} copyTemplate={ComputerState.copyTemplate} />
    {/snippet}

    {#snippet body()}
        {@render Content()}
    {/snippet}
</Card>


<!-- TODO: Change over to query link rather than endpoint link-->
{#snippet Link()}
    <a href={ComputerState.pageLink} data-title={ComputerState.disabled ? "Go to Computer Page (Disabled)" : "Go to Computer Page"}> 
        <img 
        src={ComputerState.disabled ? Icons.DISABLED_COMPUTER : Icons.COMPUTER} 
        alt="Computer Icon Needed"
        >
    </a>
{/snippet}


{#snippet Content()}
    <div id="content">
        {#if ComputerState.cn}
            <CopyButton value={ComputerState.cn} category={"title"}/>
        {/if}
        <div id="attributes">
            {#if ComputerState.operatingSystem}
                <CopyButton value={ComputerState.operatingSystem} label={"OPERATING_SYSTEM"}/>
            {/if}
            {#if ComputerState.information}
                <CopyButton value={ComputerState.information} label={"information"}/>
            {/if}
            {#if ComputerState.description}
                <CopyButton value={ComputerState.description} label={"description"}/>
            {/if}
                {#if ComputerState.dn}
                <CopyButton value={ComputerState.readableOU} label={"OU"}/>
            {/if}
        </div>
    </div>
{/snippet}

<style >

    div#content {
        padding-left: 1.5rem;
        padding-top: 0.5rem;
        padding-bottom: 1.5rem;
        padding-right: 0.5rem;
    }

    div#attributes {
        display: flex;
        flex-direction: column;
        gap: 4px;
        padding-left: 10px;
        margin-top: 4px;
    }

    a {
        opacity: 0.8;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    a:hover::after {
        content: attr(data-title);
        position: absolute;
        top: -20px;
        left: 20px;
        background-color: #333;
        color: var(--color-text);
        padding: 6px 10px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        z-index: 100;
    }


    img {
        width: 20px;
    }

    
    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

</style>
