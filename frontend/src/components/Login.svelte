<script lang="ts">
	import { goto } from "$app/navigation";


    type LoginForm = {
        username: string,
        password: string,
    }

    let error: boolean = $state(false)



    let form: LoginForm = $state({username: "", password: ""})

    const onsubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        
        const res = await fetch('http://localhost:8000/api/user/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(form)
        });

        if (res.status == 200) {
            goto("/", {invalidateAll: true})
            return
        }

        error = true




    }
</script>


<form onsubmit={(e) => onsubmit(e)} action="login">
    <h1>Login</h1>

    <span>
        <label for="username">Username</label>
        <input class:error={error} type="text" id="username" bind:value={form.username}>
    </span>
    <span>
        <label for="password">Password</label>
        <input class:error={error} type="password" id="password" bind:value={form.password}>

    </span>

    {#if error}
        <span class="error">Invalid Username or password</span>
    {/if}
    <button type="submit">Login</button>
</form>

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