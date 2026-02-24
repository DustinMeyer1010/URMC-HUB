<!--
    TODO: Make an api endpoint that will return the most commonly used AD groups to add
-->

<script lang="ts">
	import Confirm from "./Confirm.svelte";
	import Message from "./Message.svelte";
	import { AddStateClass } from "./states/AddState.svelte";
	import Group from "../cards/Group.svelte";

    let {
        currentUser
    } : {
        currentUser: string
    } = $props()

    let AddState: AddStateClass = new AddStateClass()



</script>

<form onsubmit={AddState.SearchGroup}>
    <input oncontextmenu={(e: Event) => {e.preventDefault();AddState.Search=""}} bind:value={AddState.Search} type="text" placeholder="Search for Group to Add" >
    <button type="submit">Search</button>
</form>



<section>
    {#if AddState.Results}
        {#each AddState.Results as result}
            <Message {result}/>
        {/each}
    {/if}
    {#each AddState.Groups as group, idx (group.name)}
        <Group item={group} {idx}>
            <Confirm name={group.name} username={currentUser} title={"Add Group?"} value={"Add"} action={AddState.AddGroup}/>
        </Group>
    {:else}
        <h1>Lookup Group To Add</h1>
    {/each}
</section>



<style>

    h1 {
        text-align: center;
        font-size: 15px;
    }

    section {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        height: 100%;
        width: 90%;
    }

    form {
        position: fixed;
        bottom: 50px;
        left: 50%;
        transform: translateX(-50%);
        display: flex;
        justify-content: center;
        gap: 10px;
        z-index: 10;
        width: 50%;
    }


    input {
        border: none;
        background-color: var(--color-bg-opacity-80);
        border: 2px solid var(--color-primary);
        color: var(--color-text);
        width: 400px;
        padding: 1rem;
        border-radius: 10px;
    }

    input:focus,
    input:active {
        outline: none;
    }

    ul {
        display: flex;
        flex-direction: column;
        padding: 1rem;
        border-radius: 10px;
        gap: 10px;
        padding-right: 100px;
        word-break: break-all;
        list-style: none;
        background: var(--color-surface);
        animation: slideIn 0.5s ease forwards;
        animation-delay: var(--delay);
        opacity: 0;

    }

    button {
        background: var(----color-bg-opacity-30);
        backdrop-filter: blur(10px);
        border: 2px solid var(--color-primary);
        color: var(--color-primary);
        padding: 0.7rem 1rem;
        border-radius: 6px;
    }

    button:hover {
        background: var(--color-primary-hover-opacity-10);
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

    @media (max-width: 800px) {

        ul {
            padding-right: 1rem;
        }

        form {
            width: 75%;
        }
    }

</style>