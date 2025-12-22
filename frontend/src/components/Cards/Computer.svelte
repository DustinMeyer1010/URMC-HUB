<script lang="ts">
    import type { Computer } from "@t/computer";
    import goToIcon from '$lib/assets/double-left-arrow-primary.png';
    import disabledIcon from '$lib/assets/disabled-computer.png'
	import { ComputerStateClass } from "./ComputerState.svelte";
	import CopyButton from "@components/CopyButton.svelte";
	import CopyAllButton from "@components/CopyAllButton.svelte";
    
    let {
        item,
        idx
    } : {
        item: Computer.CardInfo
        idx: number
    } = $props()

    let ComputerState: ComputerStateClass = new ComputerStateClass(
        item.ou.toLowerCase().includes("disabled"),
        `Name: ${item.name}\nOU: ${item.ou}\nOS: ${item.operating_system}`
    )

</script>

<ul class:disabled={ComputerState.Disabled} style="--delay: {Math.min(idx * 50, 2000)}ms">
    {@render DisabledContent()}
    <CopyButton value={item.name} fontSize={18} marginBottom={15}/>
    <CopyAllButton copyText={ComputerState.AllText} />
    <a href={`/computer/${item.name}`}> <img src={goToIcon} alt=""></a>
    {@render ObjectContent()}
</ul>


{#snippet DisabledContent()}
    {#if ComputerState.Disabled}
        <span class="disabled"><img src={disabledIcon} alt="">Computer Disabled</span>
    {/if}
{/snippet}


{#snippet ObjectContent()}
    {#each Object.entries(item).slice(1) as key}
        {#if key[1]}
            <CopyButton value={key[1]} label={key[0]}/>
        {/if}
    {/each}
{/snippet}

<style >
    span.disabled {
        display: flex;
        justify-content: center;
        align-items: center;
        position: absolute;
        color: var(--color-ad-disabled);
        right: 1rem;
        bottom: 0.5rem;
        font-weight: bold;
        font-size: 12px;
    }

    span.disabled img {
        width: 20px;
        margin-right: 2px;
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

    @media (max-width: 800px) {
        ul.disabled {
            padding-bottom: 3rem;
        }

        span.disabled {
            left: 50%;
            right: 0;
            transform: translateX(-50%);
            text-wrap: nowrap;
        }
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
