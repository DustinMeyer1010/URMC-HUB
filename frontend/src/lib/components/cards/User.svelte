<!-- Add: This component should allow for children. Needed for adding extra buttons like add, remove, etc -->

<script lang="ts">
	import type { User } from "$lib/types/user";
    import Icon from '$lib/assets/double-left-arrow-primary.png';
    import disabledIcon from '$lib/assets/disabled-color-disabled.png'
	import CopyButton from "../buttons/CopyButton.svelte";
	import CopyAllButton from "../buttons/CopyAllButton.svelte";
	import type { Snippet } from "svelte";
	import { UserStateClass } from "./states/User.svelte";

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
<ul class:disabled={UserState.disabled} style="--delay: {Math.min(idx * 50, 2000)}ms">
    <CopyAllButton copyTemplate={UserState.copyTemplate} />
    {@render Disabled()}
    {@render Link()}
    {@render Content()}
    {@render children?.()}
</ul>

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
    {#if UserState.username != ""}
        <a href={UserState.pageLink}> <img src={Icon} alt=""></a>
    {/if}
{/snippet}

<!-- * Renders object information with each value being able to be copied -->
{#snippet Content()}
    {#if UserState.name}
        <CopyButton value={UserState.name} category={"title"}/>
    {/if}
    {#if UserState.username}
        <CopyButton value={UserState.username} label={"USERNAME"}/>
    {/if}
    {#if UserState.net_id}
        <CopyButton value={UserState.net_id} label={"NETID"}/>
    {/if}
    {#if UserState.urid}
        <CopyButton value={UserState.urid} label={"URID"}/>
    {/if}
    {#if UserState.email}
        <CopyButton value={UserState.email} label={"EMAIL"}/>
    {/if}
    {#if UserState.readableOU}
        <CopyButton value={UserState.readableOU} label={"OU"}/>
    {/if}

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


    img {
        width: 25px;
    }

    ul {
        display: flex;
        flex-direction: column;
        word-break: break-all;
        position: relative;
        gap: 0.3rem;
        border-radius: 10px;
        padding: 1.5rem;
        padding-left: 1.5rem;
        padding-right: 3rem;
        box-sizing: border-box;
        background: var(--color-surface);
        color: var(--color-text);
        opacity: 0;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
        margin: 0;
        list-style: none;
    }

    a {
        opacity: 0.8;
        position: absolute;
        display: flex;
        justify-content: center;
        align-items: center;
        top: 0.3rem;
        right: -0.5rem;
        width: 50px;
    }

    a img {
        transition: var(--transition-fast);
        transform: rotate(130deg);
    }

    a:hover img {
        transform: rotate(130deg) translate(-3px, -1px);
    }




    @media (max-width: 1000px) {
        ul.disabled {
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

        ul {
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