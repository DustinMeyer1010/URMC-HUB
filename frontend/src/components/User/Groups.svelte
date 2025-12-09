<!-- REFACTOR -->

<script lang="ts">
	import { fly } from "svelte/transition";
	import Remove from "./Remove.svelte";

	import Message from "./Message.svelte";
	import { GroupStateClass } from "./GroupState.svelte";
	import CopyButton from "@components/CopyButton.svelte";
	import { onMount } from "svelte";


    let {
        username
    } : {
        username: string
    } = $props();

    let GroupState: GroupStateClass = new GroupStateClass()


    onMount(async () => {
        await GroupState.GetGroupForUser(username)
        GroupState.Groups = GroupState.Groups.sort((a, b) => a.name.localeCompare(b.name))
    })
</script>


{#each GroupState.Results as result}
    <Message {result}/>
{/each}
<section>
    <input oncontextmenu={(e: Event) => {e.preventDefault();GroupState.Filter=""}} bind:value={GroupState.Filter} placeholder="Search For Group"/>
    {#each GroupState.FilteredGroups as group, idx }
        <ul style="--delay: {Math.min(idx * 50, 2000)}ms" out:fly={{x: 100}}>
            <CopyButton label={""} value={group.name}/>
            {#if group.description != ""}
                <CopyButton label={"Description"} value={group.description}/>
            {/if}
            {#if group.information != ""}
                <CopyButton label={"Information"} value={group.information}/>
            {/if}
            {#if group.ou != ""}
                <CopyButton label={"OU"} value={group.ou}/>
            {/if}
            <Remove group={group.name} username={username} removeGroup={GroupState.RemoveGroup}/>
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


        ul {
            padding-right: 1rem;
        }


    }


</style>