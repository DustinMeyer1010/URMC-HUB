<script lang="ts">
	import type { UserSimpleInfo } from "@t/user";


    let searchValue: string = $state("")
    let users: UserSimpleInfo[] = $state([])

    const onsubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        await fetch(`http://localhost:8000/api/search/users/${searchValue}`)
        .then(async (res) => users = await res.json())

    }


</script>


<section>
{#each users as user }
    <div>
        <span>{user.name}</span>
        <span>{user.email}</span>
        <span>{user.ou}</span> 
    </div>
{/each}
</section>



<form action="" onsubmit={onsubmit}>
    <input type="text" placeholder="Search for user to Add" bind:value={searchValue}>
</form>


<style>
    section {
        display: flex;
        flex-direction: column;
        gap: 10px;
    }

    div {
        background: var(--color-surface);
        display: flex;
        flex-direction: column;
    }


    form {
        position: fixed;
        bottom: 10px;
        left: 50%;
        transform: translateX(-50%);
    }
</style>

