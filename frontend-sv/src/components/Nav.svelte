<script>

    import urmc from '$lib/assets/URMC.png';

</script>


<nav>
    <a href="/"><img src={urmc} alt=""></a>
    <ul>
        <li>
            <a href="/">Search</a>
        </li>
        <li>
            <a href="/bookmarks">Bookmarks</a>
        </li>
        <li>
            <a href="/Bulk">Bulk Search</a>
        </li>
    </ul>
</nav>


<style>
    nav {
        display: flex;
        justify-content: space-between;
        align-items: center;
        position: fixed;
        box-sizing: border-box;
        padding: 0.2rem 1rem;
        top: 0;
        left: 0;
        width: 100%;
        background-color: var(--background-surface);
    }

    img {
        height: 40px;
        border-radius: 5px;
    }

    ul {
        display: flex;
        gap: 2rem;
        list-style: none;
    }

    a {
        color: var(--text);
        text-decoration: none;
        font-size: 18px;
        transition: color 0.3s ease;
    }

    a:visited {
        color: var(--text);
    }

    a:hover {
        color: var(--primary)
    }

</style>

