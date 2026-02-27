<script lang="ts">
    import { Copy } from '$lib/types/copy';
    import Icon from '$lib/assets/double-left-arrow-primary.png';
	import { Drive } from "$lib/types/drive";
	import { blur } from 'svelte/transition';
	import CopyAllButton from '../buttons/CopyAllButton.svelte';
	import CopyButton from "../buttons/CopyButton.svelte";
	import { DriveStateClass } from './states/DriveState.svelte';

    let {
        item,
        idx
    } : {
        item: Drive.CardInfo
        idx: number
    } = $props()

    let DriveState: DriveStateClass = new DriveStateClass(idx, item)

</script>

<!-- * Renders the main content for the card-->
<ul style="--delay: {Math.min(idx * 50, 2000)}ms">
    <CopyAllButton copyTemplate={DriveState.CopyText} />
    <a href={`/drive?name=${DriveState.SerializedSearch}`}> <img src={Icon} alt=""></a>
    <CopyButton value={DriveState.CurrentDrive.drive} category={"title"}/>
    <CopyButton value={DriveState.CurrentDrive.local_path} label={"LOCAL_PATH"}/>
    {@render GroupsContent()}
</ul>


<!-- * Renders the Groups for the drive-->
{#snippet GroupsContent()}
    <li>
        GROUPS: 
        <!-- Note: If more than 10 groups then you can search for specific group by name-->
        {#if DriveState.CurrentDrive.groups.length >= 10}
            <input placeholder="Search For Group" bind:value={DriveState.SearchValue}/> 
        {/if}
        <div >
            <!-- Note: Each group is transitioned in and can be copied -->
            {#each DriveState.CurrentDrive.groups as group,i }
                <button 
                    class:active={DriveState.CopyState.copied === group} 
                    onclick={() => Copy.ToClipboard(group, DriveState.CopyState)} 
                    out:blur={{ duration: 200}} 
                    class="group-button" 
                    style="--delay: {Math.min(i * 20, 2000)}ms">
                        {DriveState.CopyState.copied === group ? `Copied` : group}
                </button>
            {/each}
        </div>
    </li>
{/snippet}

<style >

    input {
        background: var(--color-bg-opacity-30);
        border: none;
        color: var(--color-text);
        padding: 0.2rem;
        border-radius: 5px;
    }


    button.active {
        color: var(--color-primary-focus);

    }

    input:focus {
        outline: none;
    }

    div {
        display: flex;
        gap: 0.5rem;
        transition: 0.5s ease;
        margin-top: 0.5rem;
        padding: 0.5rem;
        flex-wrap: wrap;
        max-height: 150px;
        overflow-y: auto;
        overflow-x: hidden;
        scrollbar-width: thin; 
        scrollbar-color: rgba(255,255,255,0.3) transparent;
    }

    img {
        width: 25px;
    }



    button.group-button {
        padding: 0.5rem 0.6rem;
        background: var(--color-bg-opacity-50);
        border-radius: 20px;
        font-size: 12px;
        animation: 0.3s fillin forwards;
        animation-delay: var(--delay);
        opacity: 0;
    }

    
    button {
        background: none;
        border: none;
        padding: 0;
        font: inherit;
        color: inherit;
        cursor: pointer;
    }

    button:focus,
    button:active{
        outline: none;
    }



    ul {
        display: flex;
        flex-direction: column;
        word-break: break-all;
        position: relative;
        gap: 0.5rem;
        border-radius: 10px;
        padding: 1.5rem;
        padding-bottom: 2rem;
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


    li {
        text-align: left;
    }

    button {
        background: none;
        border: none;
        padding: 0;
        font: inherit;
        color: inherit;
        cursor: pointer;
    }

    button:active,
    button:focus { 
        border:none;
        outline: none;
    }

    @media (max-width: 400px) {
        ul {
            padding-right: 1rem;

        }
    }

    @media (max-width: 800px) {

        div {
            overflow-y: scroll;
        }
    }

    @keyframes fillin {
        from {
            opacity: 0;
            transform: translateY(100%);
        }
        to {
            opacity: 1;
            transform: translate(0,0);
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

