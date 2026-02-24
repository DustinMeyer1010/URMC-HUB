<script lang="ts">
	import type { Group } from "$lib/types/group";
    import Icon from '$lib/assets/double-left-arrow-primary.png';
	import CopyButton from "../CopyButton.svelte";
	import CopyAllButton from "../CopyAllButton.svelte";
	import type { Snippet } from "svelte";
	import { GroupStateClass } from "./GroupState.svelte";

    let {
        item,
        idx,
        children
    } : {
        item: Group.CardInfo
        idx: number
        children?: Snippet
    } = $props()

    let GroupState: GroupStateClass = new GroupStateClass(item, idx)

</script>

<!-- * Renders the content of the Card -->
<ul class:disabled={GroupState.CurrentGroup.ou.toLowerCase().includes("disabled")} style="--delay: {Math.min(GroupState.Idx * 50, 2000)}ms">
    <a href={`/group/${GroupState.CurrentGroup.name}`}> <img src={Icon} alt=""></a>
    <CopyButton value={GroupState.CurrentGroup.name} fontSize={18} marginBottom={15}/>
    <CopyAllButton copyText={GroupState.CopyText}/>
    {@render ObjectContent()}
    {#if children}
        {@render children()}
    {/if}
</ul>

<!-- * Renders the object information in copy format -->
{#snippet ObjectContent()}
    {#each Object.entries(GroupState.CurrentGroup).slice(1) as key}
        {#if key[1]}
            <CopyButton value={key[1]} label={key[0].toUpperCase()}/>
        {/if}
    {/each}
{/snippet}


<style >
    ul.disabled {
        color: var(--color-ad-disabled)
    }

    img {
        width: 25px;
    }


    ul {
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

        ul {
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
