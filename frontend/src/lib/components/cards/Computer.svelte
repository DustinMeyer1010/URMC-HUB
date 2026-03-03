<script lang="ts">
    import { Computer } from "$lib/types/computer";
    import goToIcon from '$lib/assets/double-left-arrow-primary.png';
    import disabledIcon from '$lib/assets/disabled-computer.png'
	import { ComputerStateClass } from "./states/ComputerState.svelte";
	import CopyButton from "../buttons/CopyButton.svelte";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
    
    let {
        item,
        idx
    } : {
        item: Computer.CardInfo
        idx: number
    } = $props()

    let ComputerState: ComputerStateClass = new ComputerStateClass(item)

</script>

<div class:disabled={ComputerState.disabled} style="--delay: {Math.min(idx * 50, 2000)}ms">
    <CopyAllButton copyTemplate={ComputerState.copyTemplate} />
    {@render DisabledTag()}
    {@render Link()}
    {@render Content()}
</div>

<!-- TODO: Change over to query link rather than endpoint link-->
{#snippet Link()}
    <a href={ComputerState.pageLink}> 
        <img src={goToIcon} alt="">
    </a>
{/snippet}


{#snippet DisabledTag()} 
    {#if ComputerState.disabled}
        <span class="disabled">
            <img src={disabledIcon} alt="">
            Computer Disabled
        </span>
    {/if}
{/snippet}


{#snippet Content()}
    {#if ComputerState.name}
        <CopyButton value={ComputerState.name} category={"title"}/>
    {/if}
    {#if ComputerState.ou}
        <CopyButton value={ComputerState.readableOU} label={"OU"}/>
    {/if}
    {#if ComputerState.operating_system}
        <CopyButton value={ComputerState.operating_system} label={"OPERATING_SYSTEM"}/>
    {/if}
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


    div {
        display: flex;
        flex-direction: column;
        word-break: break-all;
        position: relative;
        gap: 0.3rem;
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
        box-shadow: 0px 0px 10px 5px rgba(0,0,0,0.2);
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
        div.disabled {
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
