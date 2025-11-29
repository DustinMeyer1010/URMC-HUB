<script lang="ts">
	import { copyToClip, type CopyState } from "$lib/helper/copy.svelte";
	import type { GroupSimpleInfo } from "@t/group";
	import { fly } from "svelte/transition";
	import Remove from "./Remove.svelte";
	import { goto } from "$app/navigation";
	import type { ModifyResults } from "@t/resutls";
	import Message from "./Message.svelte";

    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })


    let {
        groups,
        currentUser
    } : {
        currentUser: string
        groups: GroupSimpleInfo[]
    } = $props();



    let filter: string = $state("")
    let filteredGroups: GroupSimpleInfo[] = $state(groups.sort((a, b) => a.name.localeCompare(b.name)))
    let results: ModifyResults[] = $state([])
    

    $effect(() => {
        if (filter == "") {
            filteredGroups = groups
            return
        }

        filteredGroups = groups.filter((group) => 
        group.name.toLowerCase().includes(filter.toLowerCase()))

    })

    const removeGroup = async (group: string) => {
        results = []
        const res = await fetch(`http://localhost:8000/api/user/${currentUser}/memberof`, {
            method: "DELETE",
            mode: "cors",
            body: JSON.stringify({
                groups: [group]
            })
        })

        results = await res.json()


        goto(location.href, { invalidateAll: true, replaceState: true });
    }

</script>



<section>
    {#each results as result}
        <Message {result}/>
    {/each}
    <input oncontextmenu={(e: Event) => {e.preventDefault();filter=""}} bind:value={filter} placeholder="Search For Group"/>
    {#each filteredGroups as group, idx }
        <ul style="--delay: {Math.min(idx * 50, 2000)}ms" out:fly={{x: 100}}>
            
            <button onclick={() => copyToClip(group.name, copyState)} class="name"><li>{group.name === copyState.copied ? "Copied" : group.name}</li></button>
            {#if group.description != ""}
                <button onclick={() => copyToClip(group.description, copyState)}>
                    <b>Description: </b> 
                    <li>{group.description === copyState.copied ? "Copied" : group.description}</li>
                </button>
            {/if}
            {#if group.information != ""}
                <button onclick={() => copyToClip(group.information, copyState)}>
                    <b>Information: </b> 
                    <li>{group.information === copyState.copied ? "Copied" : group.information}</li>
                </button>
            {/if}
            {#if group.ou != ""}
                <button onclick={() => copyToClip(group.ou, copyState)}>
                    <b>OU: </b> 
                    <li>{group.ou === copyState.copied ? "Copied" : group.ou}</li>
                </button>
            {/if}
            <Remove group={group.name} removeGroup={removeGroup}/>
        </ul>
    {/each}

</section>



<style>

    section{ 
        position: relative;
        display: flex;
        gap: 1rem;
        width: 90%;
        flex-direction: column;
        justify-content: center;
    }

    input {
        position: fixed;
        bottom: 50px;
        left: 50%;
        transform: translateX(-50%);
        width: 50%;
        padding: 1rem;
        color: var(--color-text);
        z-index: 1;
        background: var(--color-bg-opacity-80);
        border: 1px solid var(--color-primary-focus);
        border-radius: 10px;
        backdrop-filter: blur(10px);
    }

    input:focus {
        outline: none;
    }

    ul {
        display: flex;
        flex-grow: 1;
        flex-direction: column;
        gap: 0.5rem;
        list-style: none;
        animation: slideIn 0.5s forwards var(--delay);
        opacity: 0;
        padding: 1rem 1rem;
        padding-right: 120px;
        border-radius: 10px;
        background: var(--color-surface);
        margin: 0;
    }

    li {
        word-break: break-all;
    }

    button.name {
        font-weight: bold;
        font-size: 18px;
    }


    button {
        background: transparent;
        display: flex;
        flex-direction: row;
        border: none;
        gap: 1rem;
        padding-left: 1rem;
        font-size: 15px;
        color: var(--color-text)
    }

    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(100%);
        }
        to {
            opacity: 1;
            transform: translateY(0px);
        }
    }


    @media (max-width: 850px) {


        button {
            flex-direction: column;
            justify-content: flex-start;
            text-align: left;
        }

        ul {
            padding-right: 1rem;
        }


    }


</style>