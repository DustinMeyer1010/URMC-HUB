
<script lang="ts">
	import type { User } from "$lib/types/user";
	import CopyButton from "../buttons/CopyButton.svelte";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
	import { UserStateClass } from "./states/User.svelte";
	import Card from "./Card.svelte";
	import { Icons } from "$lib/managers/icons";

    let {
        item,
        idx,
    } : {
        item: User.CardInfo
        idx: number
    } = $props();


    let UserState: UserStateClass = new UserStateClass(item)


</script>

<!-- * Renders the main content -->

<Card {idx}>
    {#snippet header()}
        {@render Link()}
        <CopyAllButton 
            icon={Icons.COPY} 
            copiedIcon={Icons.COPY_SUCCESSFUL} 
            copyTemplate={UserState.copyTemplate} />
        <CopyAllButton 
            icon={Icons.SIMPLE_COPY} 
            copiedIcon={Icons.SIMPLE_COPY_SUCCESSFUL} 
            dataTitle={"Copy: Name (Username)"} 
            copyTemplate={UserState.simpleCopyTemplate} />
        {#if UserState.idleAccount} 
            <CopyAllButton 
                icon={Icons.IDLE} 
                copyTemplate={""} 
                dataTitle={"User has IDM_IdleAccounts_URMC"}/>
        {/if}
        {#if UserState.expiredPassword && UserState.username != ""}
            <CopyAllButton 
                icon={Icons.BAD_PASSWORD} 
                copyTemplate={""} 
                dataTitle={"Expired Password"}/>
        {/if}
    {/snippet}

    {#snippet body()}
        {@render Content()}
    {/snippet}
</Card>


{#snippet Link()}
    <a 
        href={UserState.username != "" ? UserState.pageLink : null}
        data-title={UserState.username == "" 
        ? "No URMC Active Directory Object" 
        : `Go to User Page ${UserState.disabled ? "(Disabled Account)" : ""}`}>  
        <img src={UserState.disabled ? Icons.DISABLED_PROFILE : Icons.PROFILE} alt="">
    </a>
{/snippet}

<!-- * Renders object information with each value being able to be copied -->
{#snippet Content()}
    <div id="content">
        {#if UserState.cn}
            <CopyButton value={UserState.cn} category={"title"}/>
        {/if}
        <div id="attributes">
            {#if UserState.username}
                <CopyButton value={UserState.username} label={"USERNAME"}/>
            {/if}
            {#if UserState.netid}
                <CopyButton value={UserState.netid} label={"NETID"}/>
            {/if}
            {#if UserState.urid}
                <CopyButton value={UserState.urid} label={"URID"}/>
            {/if}
            {#if UserState.email}
                <CopyButton  value={UserState.email} label={"EMAIL"}/>
            {/if}
            {#if UserState.readableDN}
                <CopyButton  value={UserState.readableDN} label={"DN"}/>
            {/if}
        </div>
    </div>

{/snippet}


<style >

    div#attributes {
        padding-left: 5px;
        margin-top: 4px;
        display: flex;
        flex-direction: column;
        gap: 2px;
    }


    a {
        color: var(--color-text);
    }

    a:hover::after:visited {
        color: var(--color-text);
    }

    a:hover::after {
        content: attr(data-title);
        position: absolute;
        top: -20px;
        left: 20px;
        background-color: #333;
        color: var(--color-text);
        padding: 6px 10px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        z-index: 100;
    }


    img {
        width: 25px;
    }


    div#content {
        padding-left: 1.5rem;
        padding-top: 0.5rem;
        padding-bottom: 1.5rem;
        padding-right: 0.5rem;
    }

    a {
        opacity: 0.8;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    @media (max-width: 750px) {
        div#content {
            padding-left: 1rem;
            padding-bottom: 1rem;
        }


    }




    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

</style>