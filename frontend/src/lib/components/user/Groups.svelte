<!-- REFACTOR -->

<script lang="ts">
	import Remove from "./Remove.svelte";

	import Message from "./Message.svelte";
	import { GroupStateClass } from "./states/GroupState.svelte";
	import { onMount } from "svelte";
	import Group from "../cards/Group.svelte";


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
    {#each GroupState.FilteredGroups as group, idx (group.name) }
        <Group item={group} idx={idx} >
            <Remove group={group.name} username={username} removeGroup={GroupState.RemoveGroup}/> 
        </Group>
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
        padding-bottom: 100px;
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

</style>