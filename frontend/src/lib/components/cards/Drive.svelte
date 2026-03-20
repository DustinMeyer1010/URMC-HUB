<script lang="ts">
	import { Icons } from "$lib/managers/icons";
	import { Drive } from "$lib/types/drive";
	import CopyAllButton from '../buttons/CopyAllButton.svelte';
	import CopyButton from "../buttons/CopyButton.svelte";
	import Card from "./Card.svelte";
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

<Card {idx} >
    {#snippet header()}
        {@render Link()}
        <CopyAllButton  icon={Icons.COPY} copiedIcon={Icons.COPY_SUCCESSFUL} copyTemplate={DriveState.copyTemplate} />
    {/snippet}
    {#snippet body()}
        {@render Content()}
    {/snippet}
</Card>


{#snippet Link()}
        <a href={DriveState.pageLink} data-title={"Go to Drive Page"}> 
            <img src={Icons.DRIVE} alt="Drive Icon Needed">
        </a>
{/snippet}

{#snippet Content()}
    <div id="content">
        <CopyButton value={DriveState.drive} category={"title"}/>
        <CopyButton value={DriveState.local_path} label={"LOCAL_PATH"}/>
        {@render Groups()}
    </div>
{/snippet}


<!-- * Renders the Groups for the drive-->
{#snippet Groups()}
    <section id="groups">
        GROUPS: 
        <!-- Note: If more than 10 groups then you can search for specific group by name-->
        {#if DriveState.groups.length >= 10}
            <input placeholder="Search For Group" bind:value={DriveState.searchValue}/> 
        {/if}
        <span>{DriveState.groupCount}</span>
        <div id="groups" >
            <!-- Note: Each group is transitioned in and can be copied -->
            {#each DriveState.filteredGroups as group (group) }
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

    div#content {
        padding-left: 1.5rem;
        padding-top: 0.5rem;
        padding-bottom: 1.5rem;
        padding-right: 0.5rem;
    }

    div#groups {
        display: flex;
        gap: 0.3rem;
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
        width: 20px;
    }


    a {
        color: var(--color-text);
    }

    a:hover::after:visited {
        color: var(--color-text);
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


    section#groups {
        text-align: left;
    }
    
    a {
        opacity: 0.8;
        display: flex;
        justify-content: center;
        align-items: center;
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

