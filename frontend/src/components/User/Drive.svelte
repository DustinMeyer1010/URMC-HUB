<script lang="ts">
	import type { DriveSimpleInfo } from "@t/drive";
	import type { GroupSimpleInfo } from "@t/group";
	import { onMount } from "svelte";



    let data: DriveSimpleInfo[] | undefined = $state(undefined)
    let loading: boolean = $state(true)
    let {
        groups
    } : {
        groups: GroupSimpleInfo[]
    } = $props();
    
    onMount(async () => {
        loading = true
        const names = groups.map((group) => group.name)

        const res = await fetch("http://localhost:8000/api/drive/access", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(names)
        })

        data = await res.json()

        loading = false
    })


</script>


{#if !loading && data == undefined || data?.length == 0 }
    <h1 class="no-access">No drive access</h1>
{/if}
<section>
    {#if !loading}

        {#each data as access, idx}
            <div style="--delay: {idx * 50}ms">
                <h1 class="title">{access.drive}</h1>
                <ul>
                    {#each access.groups as group}
                        <li>{group}</li>
                    {/each}
                </ul>
            </div>
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
        padding: 2rem 4rem;
        border-radius: 10px;
        width: 90%;
        display: grid;
        grid-template-columns: 1fr;
        gap: 1rem;
        align-items: stretch;
    }

    h1.title {
        font-weight: bolder;
    }

    h1.no-access {
        width: 100%;
    }

    h1 {
        font-size: 18px;
        height: 30px;
        margin: 1rem;
        padding: 0;
        text-align: left;
    }

    div {
        height: fit-content;
        padding: 1rem;
        border-radius: 10px;
        background: var(--background-surface);
        animation: slideIn 0.3s forwards;
        animation-delay: var(--delay);
        opacity: 0;
    }

    div.loading {
        opacity: 1;
        animation: pulse 3s infinite;
        animation-delay: var(--delay);

    }

    ul {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        margin: 0;
        padding: 0;
        gap: 1rem;
        list-style: none;
    }

    li {
        border-radius: 3px;
        padding: 0.2rem;
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

    @media (max-width: 1010px) {



    }

    @media (max-width: 800px) {
        section {
            grid-template-columns: 1fr;
        }

    }


</style>