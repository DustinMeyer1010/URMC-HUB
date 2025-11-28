<script lang="ts">
	import type { GroupSimpleInfo } from "@t/group";
	import Confirm from "./Confirm.svelte";
	import { goto } from "$app/navigation";
	import type { ModifyResults } from "@t/resutls";
	import Message from "./Message.svelte";

    let {
        currentUser
    } : {
        currentUser: string
    } = $props()



    let search: string = $state("")
    let groups: GroupSimpleInfo[] = $state([])
    let results: ModifyResults[] = $state([])



    const searchGroups = async (e: SubmitEvent) => {
        e.preventDefault()
        const res = await fetch(`http://localhost:8000/api/search/groups/${search}`)

        groups = await res.json() as GroupSimpleInfo[]
    }

    const addGroup = async (group: string) => {
        results = []
        const res = await fetch(`http://localhost:8000/api/user/${currentUser}`, {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                groups: [group]
            })
        })

        results = await res.json()


        goto(location.href, { invalidateAll: true, replaceState: true });

    }



</script>

<form onsubmit={searchGroups}>
    <input oncontextmenu={(e: Event) => {e.preventDefault();search=""}} bind:value={search} type="text" placeholder="Search For Group" >
    <button type="submit">Search</button>
</form>



<section>
    {#if results}
        {#each results as result}
            <Message {result}/>
        {/each}
    {/if}
    {#each groups as group, idx (group.name)}
        <ul style="--delay: {Math.min(idx * 50, 2000)}ms">
            <li class="title">{group.name !== "" ? group.name : "NA"}</li>
            <li><b>Description:</b> {group.description !== "" ? group.description : "NA"}</li>
            <li><b>Information:</b> {group.information !== "" ? group.information : "NA"}</li>
            <li><b>OU:</b> {group.ou !== "" ? group.ou : "NA"}</li>
            <Confirm name={group.name} title={"Add Group?"} value={"Add"} action={addGroup}/>
        </ul>
    {:else}
        <h1>Lookup Group To Add</h1>
    {/each}
</section>



<style>

    h1 {
        text-align: center;
    }

    section {
        height: 100%;
        width: 90%;
    }

    li.title {
        font-weight: bold;
        font-size: 18px;
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