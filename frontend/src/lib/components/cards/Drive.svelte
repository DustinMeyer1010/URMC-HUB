<script lang="ts">
    import Icon from '$lib/assets/double-left-arrow-primary.png';
	import { Drive } from "$lib/types/drive";
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

    let DriveState: DriveStateClass = new DriveStateClass(item)

</script>

<!-- * Renders the main content for the card-->
<div id="card" style="--delay: {Math.min(idx * 50, 2000)}ms">
    <CopyAllButton copyTemplate={DriveState.copyTemplate} />
    {@render Link()}
    {@render Content()}
</div>


{#snippet Link()}
        <a href={DriveState.pageLink}> 
            <img src={Icon} alt="">
        </a>
{/snippet}

{#snippet Content()}
    <CopyButton value={DriveState.drive} category={"title"}/>
    <CopyButton value={DriveState.local_path} label={"LOCAL_PATH"}/>
    {@render Groups()}
{/snippet}


<!-- * Renders the Groups for the drive-->
{#snippet Groups()}
    <section id="groups">
        GROUPS: 
        <!-- Note: If more than 10 groups then you can search for specific group by name-->
        {#if DriveState.groups.length >= 10}
            <input placeholder="Search For Group" bind:value={DriveState.searchValue}/> 
        {/if}
        <div id="groups" >
            <!-- Note: Each group is transitioned in and can be copied -->
            {#each DriveState.filteredGroups as group,i }
                <CopyButton value={group} category={"drive-group"}></CopyButton>
            {/each}
        </div>
    </section>
{/snippet}

<style >

    input {
        background: var(--color-bg-opacity-30);
        border: none;
        color: var(--color-text);
        padding: 0.2rem;
        border-radius: 5px;
    }



    input:focus {
        outline: none;
    }

    div#groups {
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


    div#card {
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


    section#groups {
        text-align: left;
    }

    @media (max-width: 400px) {
        div {
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

