<!-- 
 REFACTOR: GetDrives() should be call inside of GetGroups
 REFACTOR: Fix the sizing and colors of groups to drives
-->
<script lang="ts">
	import { onMount } from "svelte";
	import { DriveStateClass } from "./states/DriveState.svelte";
	import CopyButton from "$lib/components/buttons/CopyButton.svelte";


    let {
        username
    } : {
        username: string
    } = $props();

    let DriveState: DriveStateClass = new DriveStateClass(username)
    
    onMount(async () => {
        await DriveState.GetGroups()
        await DriveState.GetDrives()
    })


</script>

<section>
    {#if !DriveState.Loading}
        {#each DriveState.DrivesAccess as access, idx}
            <div style="--delay: {idx * 50}ms">
                <CopyButton value={access.drive} label="" fontSize={18} marginBottom={10}/>
                <ul>
                    {#each access.groups as group}
                        <li><CopyButton value={group} label="" fontSize={12}/></li>
                    {/each}
                </ul>
            </div>
        {:else}
            <h1 class="no-access">No drive access</h1>
        {/each}
    {:else}
        {#each Array(8).fill(0) as _,x}
            <div class="loading" style="--delay: {x * 250}ms">
            </div>
        {/each}
    {/if}

</section>


<style>

    section {
        border-radius: 10px;
        width: 90%;
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        gap: 1rem;
        align-items: stretch;
    }

    h1.no-access {
        width: 100%;
    }

    h1 {
        font-size: 18px;
        margin: 1rem;
        padding: 0;
        text-align: left;
    }

    div {
        padding: 1rem;
        border-radius: 10px;
        background: var(--color-surface);
        animation: slideIn 0.3s forwards;
        animation-delay: var(--delay);
        display: flex;
        flex-direction: column;
        word-break: break-all;
        flex-wrap: wrap;
        opacity: 0;
    }

    div.loading {
        height: 100px;
        opacity: 1;
        animation: pulse 3s infinite;
        animation-delay: var(--delay);

    }

    ul {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        margin-top: 20px;
        padding: 0;
        gap: 1rem;
        list-style: none;
    }

    li {
        border-radius: 10px;
        padding: 0.2rem 1rem;
        background: var(--color-bg-opacity-30);
    }


    @keyframes pulse {
        0% {
            opacity: 1;
        }
        50% {
            opacity: 0.5;
        }
        100% {
            opacity: 1;
        }
    }

    @keyframes slideIn {
        from {
            transform: translateY(100%);
            opacity: 0;
        }
        to {
            transform: translateY(0);
            opacity: 1;
        }
    }


    @media (max-width: 800px) {
        section {
            grid-template-columns: 1fr;
        }

    }


</style>