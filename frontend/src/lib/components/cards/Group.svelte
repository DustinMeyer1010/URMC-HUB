<script lang="ts">
	import type { Group } from "$lib/types/group";
    import Icon from '$lib/assets/double-left-arrow-primary.png';
	import CopyButton from "../buttons/CopyButton.svelte";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
	import type { Snippet } from "svelte";
	import { GroupStateClass } from "./states/GroupState.svelte";


    let {
        item,
        idx,
        children
    } : {
        item: Group.CardInfo
        idx: number
        children?: Snippet
    } = $props()

    let GroupState: GroupStateClass = new GroupStateClass(item)


</script>

<!-- * Renders the content of the Card -->
<div class:disabled={GroupState.ou.toLowerCase().includes("disabled")} style="--delay: {Math.min(idx * 50, 2000)}ms">
    <CopyAllButton copyTemplate={GroupState.copyTemplate}/>
    {@render Link()}
    {@render Content()}
    {@render children?.()}
</div>

<!-- * Creates the link to the group page -->
{#snippet Link()}
    <a href={GroupState.pageLink}> 
        <img src={Icon} alt="">
    </a>
{/snippet}

<!-- * Creates the content of the card -->
{#snippet Content()}
    <CopyButton value={GroupState.name} category={"title"}/>
    {#if GroupState.information}
        <CopyButton value={GroupState.information} label={"INFORMATION"}/>
    {/if}
    {#if GroupState.description}
        <CopyButton value={GroupState.description} label={"DESCRIPTION"}/>
    {/if}
    {#if GroupState.readableOU}
        <CopyButton value={GroupState.readableOU} label={"OU"}/>
    {/if}
{/snippet}


<style >
    div.disabled {
        color: var(--color-ad-disabled)
    }

    img {
        width: 25px;
    }


    div {
        display: flex;
        flex-direction: column;
        word-break: break-all;
        position: relative;
        gap: 0.5rem;
        border-radius: 10px;
        padding: 1.5rem;
        padding-left: 1.5rem;
        padding-right: 3rem;
        box-sizing: border-box;
        background: var(--color-surface);
        color: var(--color-text);
        opacity: 0;
        overflow-x: hidden;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
        margin: 0;
        list-style: none;
    }

    a {
        opacity: 0.8;
        position: absolute;
        display: flex;
        justify-content: center;
        align-items: center;
        top: 0.3rem;
        right: -0.5rem;
        width: 50px;
    }

    a img {
        transition: var(--transition-fast);
        transform: rotate(130deg);
    }

    a:hover img {
        transform: rotate(130deg) translate(-3px, -1px);
    }


    @media (max-width: 400px) {

        div {
            padding-right: 1rem;

        }
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
