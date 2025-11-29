<script lang="ts">
	import Nav from "@components/User/Nav.svelte";
    import Profile from "@components/User/Profile.svelte"
	import Lockout from "@components/User/Lockout.svelte";
	import Drive from "@components/User/Drive.svelte";
	import Groups from "@components/User/Groups.svelte";
	import Add from "@components/User/Add.svelte";

	import { onMount } from "svelte";
	import { page } from "$app/state";

    import  { type Section, Sections } from "@t/section";
	import { UserStateClass } from "./state.svelte";
	import PageLoading from "@components/Loading-Animations/PageLoading.svelte";

    let { data } : { data: {username: string} } = $props();
    let UserPageState = new UserStateClass()

    onMount(async () => {
        UserPageState.pageData = await UserPageState.UserFullInfo(data.username)
        let urlParams = page.url.searchParams;
        UserPageState.currentSection = urlParams.get("section")?.toUpperCase() as Section ?? "PROFILE"
        UserPageState.SetURL()
        UserPageState.loading = false
    })

</script>


{#if UserPageState.pageData != null}
    <nav>
        <h1>{UserPageState.pageData.name}</h1>
        <Nav sections={Sections} swapSection={UserPageState.SwapSections}/>
    </nav>

    {#if !UserPageState.loading}
        <main>
            {#if UserPageState.currentSection == "PROFILE"}
                <Profile user={UserPageState.pageData}/>
            {:else if UserPageState.currentSection == "LOCKOUT"}
                <Lockout username={UserPageState.pageData.username}/>
            {:else if UserPageState.currentSection == "DRIVES"}
                <Drive groups={UserPageState.pageData.member_of}/>
            {:else if UserPageState.currentSection == "GROUPS"}
                <Groups currentUser={UserPageState.pageData.username} groups={UserPageState.pageData.member_of}/>
            {:else if UserPageState.currentSection == "ADD"}
                <Add currentUser={UserPageState.pageData.username}/>
            {/if}
        </main>
    {:else}
        <PageLoading/>
    {/if}
{/if}




<style>
    
    h1 {
        padding: 10px;
        margin: 0;
    }

    main {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 1rem 2rem;
        flex-direction: column;
        gap: 1rem;
        margin-top: 150px;

    }

    nav {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        position: fixed;
        padding: 1rem 0rem;
        left: 0;
        top: 55px;
        gap: 1rem;
        width: 100%;
        margin-bottom: 1rem;
        z-index: 3;
        background: var(--color-bg);
    }

    @media (max-width: 700px) {
        main {
            padding: 0 1rem;
        }
    }



    @media (max-width: 465px) {
        main {
            margin-top: 200px;
        }
    }




</style>