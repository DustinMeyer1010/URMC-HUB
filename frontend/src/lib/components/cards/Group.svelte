<script lang="ts">
	import type { Group } from "$lib/types/group";
	import CopyButton from "../buttons/CopyButton.svelte";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
	import { GroupStateClass } from "./states/GroupState.svelte";
	import Card from "./Card.svelte";
	import { Icons } from "$lib/managers/icons";
	import type { Snippet } from "svelte";


    let {
        item,
        idx,
        children,
    } : {
        item: Group.CardInfo
        idx: number
        children?: Snippet
    } = $props()

    let GroupState: GroupStateClass = new GroupStateClass(item)


</script>

<Card {children} {idx} >
    {#snippet header()}
        {@render Link()}
        <CopyAllButton icon={Icons.COPY} copiedIcon={Icons.COPY_SUCCESSFUL} copyTemplate={GroupState.copyTemplate}/>
        {#if GroupState.sigGroup}
            <CopyAllButton icon={Icons.ADMIN} copyTemplate={""} dataTitle={"SIG Group"}/>
        {/if}
    {/snippet}
    {#snippet body()}
        {@render Content()}
    {/snippet}
</Card>

<!-- * Creates the link to the group page -->
{#snippet Link()}
    <a href={GroupState.pageLink} data-title={"Go to Group Page"}> 
        <img src={Icons.GROUP} alt="Group Icon Needed">
    </a>
{/snippet}

<!-- * Creates the content of the card -->
{#snippet Content()}
    <div id="content">
        <CopyButton value={GroupState.cn} category={"title"}/>
        <div id="attributes">
            {#if GroupState.information}
                <CopyButton value={GroupState.information} label={"INFORMATION"}/>
            {/if}
            {#if GroupState.description}
                <CopyButton value={GroupState.description} label={"DESCRIPTION"}/>
            {/if}
            {#if GroupState.readableOU}
                <CopyButton value={GroupState.readableOU} label={"OU"}/>
            {/if}
        </div>
    </div>
{/snippet}


<style >


    img {
        width: 25px;
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

    div#content {
        padding-left: 1.5rem;
        padding-top: 0.5rem;
        padding-bottom: 1.5rem;
        padding-right: 0.5rem;
    }

    div#attributes {
        padding-left: 5px;
        margin-top: 4px;
        display: flex;
        flex-direction: column;
        gap: 2px;
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
