<script lang="ts">
	import type { GroupSimpleInfo } from "@t/group";



    let search: string = $state("")
    let groups: GroupSimpleInfo[] = $state([])


    const searchGroups = async (e: SubmitEvent) => {
        e.preventDefault()
        const res = await fetch(`http://localhost:8000/api/search/groups/${search}`)

        groups = await res.json() as GroupSimpleInfo[]

    }

    const addGroup = async () => {
        const res = await fetch("http://localhost:8000/api/user/group/add", {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
                users: ["dmeyer20"],
                groups: ["ISDG_CTX_eRecord2"]
            })
        })

        const data = await res.json()

    }


</script>

<form onsubmit={searchGroups}>
    <input bind:value={search} type="text" placeholder="Search For Group" >
    <button type="submit">Search</button>
</form>

<section>
    {#each groups as group, idx}
        {group.name}
    {/each}
</section>



<style>

    section {
        height: 100%;
    }

    input {
        padding: 0.5rem 1rem;
        border: none;
        background-color: var(--color-surface);
        color: var(--color-text);
        font-size: 15px;
        border-radius: 5px;
        text-align: center;
    }

    input:focus,
    input:active {
        outline: none;
    }

</style>