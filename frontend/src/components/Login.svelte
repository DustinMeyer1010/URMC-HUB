<script lang="ts">
	import { LoginStateClass } from "./LoginState.svelte";

    let {
        login
    } : {
        login: () => void
    } = $props()

    let LoginState: LoginStateClass = new LoginStateClass(login)

</script>


<form onsubmit={(e) => LoginState.onsubmit(e)} action="login">
    <h1>Login</h1>

    <span>
        <label for="username">Username</label>
        <input class:error={LoginState.Error} type="text" id="username" bind:value={LoginState.Form.username}>
    </span>
    <span>
        <label for="password">Password</label>
        <input class:error={LoginState.Error} type="password" id="password" bind:value={LoginState.Form.password}>

    </span>

    {@render Error()}
    <button type="submit">Login</button>
</form>


{#snippet Error()}
    {#if LoginState.Error}
        <span class="error">Invalid Username or password</span>
    {/if}
{/snippet}


<style>

    h1 {
        align-self: center;
        text-align: center;
        margin: 0;
        padding: 0;
    }

    button {
        color: var(--color-primary);
        font-weight: bold;
        border: 2px solid var(--color-primary);
        padding: 0.5rem;
        border-radius: 10px;
        background-color: transparent;
    }

    button:hover {
        background-color: var(--color-primary-hover-opacity-20);
    }

    form {
        position: absolute;
        display: flex;
        flex-direction: column;
        background: var(--color-surface);
        padding: 1.5rem;
        gap: 1rem;
        border-radius: 10px;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

    span.error {
        font-size: 12px;
        font-weight: bold;
        color: var(--color-error);
    }

    input.error {
        border: 1px solid var(--color-error);
    }

    span {
        justify-content: center;
        align-items: center;
        display: flex;
        gap: 10px;
    }

    input {
        background: var(--color-bg-opacity-30);
        border: none;
        border-radius: 5px;
        padding: 0.5em;
        color: var(--color-text);
    }

    input:focus {
        outline:none
    }



    
</style>