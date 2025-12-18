<script lang="ts">
	import type { GroupSimpleInfo } from "@t/group";
    import Icon from '$lib/assets/double-left-arrow-primary.png';
	import CopyButton from "@components/CopyButton.svelte";
	import CopyAllButton from "@components/CopyAllButton.svelte";

    let {
        item,
        idx
    } : {
        item: GroupSimpleInfo
        idx: number
    } = $props()

    const copyText: string = `Name: ${item.name}\nInformation: ${item.information !== "" ? item.information : "NA"}\nDescription: ${item.description !== "" ? item.description : "NA"}\nOU: ${item.ou}`

</script>

<!-- * Renders the content of the Card -->
<ul class:disabled={item.ou.toLowerCase().includes("disabled")} style="--delay: {Math.min(idx * 50, 2000)}ms">
    <a href={`/group/${item.name}`}> <img src={Icon} alt=""></a>
    <CopyButton value={item.name} fontSize={18} marginBottom={15}/>
    <CopyAllButton {copyText}/>
    {@render ObjectContent()}
</ul>

<!-- * Renders the object information in copy format -->
{#snippet ObjectContent()}
    {#each Object.entries(item).slice(1) as key}
        {#if key[1]}
            <CopyButton value={key[1]} label={key[0]}/>
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
