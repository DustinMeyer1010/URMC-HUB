<script>

</script>

<form>
    <input type="text">
    <button>Search</button>
</form>

<style>
    form {
        display: flex;
        flex-wrap: nowrap;
        gap: 1rem;
    }

    input {
        border: none;
        background: var(--background-surface);
        padding: 0.5rem;
        border-radius: 5px;
        font-size: 18px;
        width: 300px;
        color: var(--text)
    }

    input:focus {
        outline: none;
    }

    button {
        background: transparent;
        border: 2px solid var(--primary);
        border-radius: 5px;
        padding: 0 1rem;
        font-weight: bold;
        color: var(--primary);
        transition: 0.3s ease;
    }

    button:hover {
        background: var(--primary);
        color: var(--text);
    }

    @media (max-width: 601px) {
        input {
            width: 200px;
        }
    }
</style>