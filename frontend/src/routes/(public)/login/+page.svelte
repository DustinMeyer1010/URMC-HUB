<script lang="ts">
	import { goto } from "$app/navigation";
	import { onMount } from "svelte";


    type LoginForm = {
        username: string,
        password: string,
    }

    onMount(() => {
        goto("/login", {replaceState: false, keepFocus: true})
    })



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
        }


    }
</script>


<form onsubmit={(e) => onsubmit(e)} action="login">
    <label for="username">username</label>
    <input type="text" id="username" bind:value={form.username}>
    <label for="password">password</label>
    <input type="password" id="password" bind:value={form.password}>
    <button type="submit">Login</button>
</form>

<style></style>