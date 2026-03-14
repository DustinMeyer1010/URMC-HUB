<!-- Add: This component should allow for children. Needed for adding extra buttons like add, remove, etc -->

<script lang="ts">
	import type { User } from "$lib/types/user";
    import Icon from '$lib/assets/profile.png';
    import ProfileDisabledIcon from "$lib/assets/profile-disabled.png"
    import CopyIcon from "$lib/assets/copy-color-text.png"
    import CopySuccessIcon from "$lib/assets/copy-success.png"
    import SimpleCopyIcon from "$lib/assets/simple-copy.png"
    import SimpleCopySuccessIcon from "$lib/assets/simple-copy-success.png"
    import disabledIcon from '$lib/assets/disabled-color-disabled.png'
	import CopyButton from "../buttons/CopyButton.svelte";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
	import type { Snippet } from "svelte";
	import { UserStateClass } from "./states/User.svelte";
	import Card from "./Card.svelte";

    let {
        item,
        idx,
        children
    } : {
        item: User.CardInfo
        idx: number
        children?: Snippet
    } = $props();


    let UserState: UserStateClass = new UserStateClass(item)


</script>

<!-- * Renders the main content -->

<Card {idx}>
    {#snippet header()}
        {@render Link()}
        <CopyAllButton icon={CopyIcon} copiedIcon={CopySuccessIcon} copyTemplate={UserState.copyTemplate} />
        <CopyAllButton icon={SimpleCopyIcon} copiedIcon={SimpleCopySuccessIcon} copyTemplate={UserState.simpleCopyTemplate} />
    {/snippet}

    {#snippet body()}
        {@render Disabled()}
        {@render Content()}
        {@render children?.()}
    {/snippet}
</Card>



<!-- * Renders the disabled contnent if object return is disabled -->
{#snippet Disabled()}
    {#if UserState.disabled}
        <span class="disabled">
            <img src={disabledIcon} alt="">
            Disabled Account
        </span>
    {/if}
{/snippet}



<!-- Refactor: This should be done based on the DN rather than a person username -->
<!-- * Renders a link to users page -->
<!-- Note: If user does not have username user page will not working so no link is generated -->
{#snippet Link()}
    <a href={UserState.username != "" ? UserState.pageLink : null} data-title={UserState.username == "" ? "No URMC Active Directory Object" : "Go to User Page"}>  <img src={UserState.disabled ? ProfileDisabledIcon : Icon} alt=""></a>
{/snippet}

<!-- * Renders object information with each value being able to be copied -->
{#snippet Content()}
    <div id="content">
        {#if UserState.cn}
            <CopyButton value={UserState.cn} category={"title"}/>
        {/if}
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
            <CopyButton value={UserState.email} label={"EMAIL"}/>
        {/if}
        {#if UserState.readableDN}
            <CopyButton value={UserState.readableDN} label={"DN"}/>
        {/if}
    </div>

{/snippet}


<style >


    span.disabled {
        display: flex;
        justify-content: center;
        align-items: center;
        position: absolute;
        color: var(--color-ad-disabled);
        right: 1rem;
        bottom: 0.5rem;
        font-weight: bold;
        font-size: 12px;
    }

    span.disabled img {
        width: 20px;
        margin-right: 2px;
    }

    a:hover::after {
        content: attr(data-title);
        position: absolute;
        top: -20px;
        left: 20px;
        background-color: #333;
        color: var(--text);
        padding: 6px 10px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        z-index: 100;
    }

    a:hover::before {
        content: "";
        position: absolute;
        bottom: 110%;
        border: 6px solid transparent;
        border-top-color: #333;
        z-index: 100;
    }


    img {
        width: 20px;
    }


    div#content {
        padding-left: 1.5rem;
        padding-top: 0.5rem;
        padding-bottom: 1.5rem;
    }

    a {
        opacity: 0.8;
        display: flex;
        justify-content: center;
        align-items: center;
    }





    @media (max-width: 1000px) {
        div.disabled {
            padding-bottom: 3rem;
        }

        span.disabled {
            left: 50%;
            right: 0;
            transform: translateX(-50%);
            text-wrap: nowrap;
        }
    } 

    @media (max-width: 400px) {

        div#container {
            padding-right: 1rem;
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