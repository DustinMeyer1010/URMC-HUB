<script lang="ts">
    import { page } from '$app/state';
    import { DocAuth, DocRoutes, DocSearchRoutes, DocUser } from "../../../routes/(docs)/docRoutes"

</script>



<a href="/api">
    <h1>URMC-HUB <sub>API</sub></h1>
</a>
{@render rootRoutes()}
{@render subRoutes("Search", DocSearchRoutes)}
{@render subRoutes("User", DocUser)}
{@render subRoutes("Auth", DocAuth)}


{#snippet rootRoutes()}
    <section>
        {#each DocRoutes as route }
            <a href={route.path} class={page.url.pathname === route.path ? 'active' : ''}>
                {route.name}
            </a>
        {/each}
    </section>
{/snippet}

{#snippet subRoutes(SubName: string, routes: any)}
    <section class="sub">
		<hr/>
        <h2>{SubName}</h2>

        {#each routes as route }
            <a href={route.path} class={page.url.pathname === route.path ? 'active' : ''}>
                {route.name}
            </a>
        {/each}
    </section>
{/snippet}


<style>
	section {
		background: var(--color-surface);
		display: flex;
		flex-direction: column;
        gap: 3px;
        width: 100%;
        margin-bottom: 10px;
	}

    hr {
        display: inline;
        width: 100%;
        align-self: flex-start;
        justify-self: flex-start;
    }



	h1 {
		padding: 0;
		margin: 0;
		font-size: 18px;
		text-align: center;

	}

	h2 {
		font-size: 15px;
		margin: 0;
		padding-left: 1rem;
	}

	a {
		color: var(--color-text);
		text-decoration: none;
		width: 100%;
		padding: 0.5rem 1rem;
	}


	a.active {
		color: var(--color-primary) !important;
	}

	a:hover {
		color: var(--color-text-opacity-80)

	}

	a:active {
		color: var(--color-primary)
	}

	a:visited {
		color: var(--color-text);
	}

	sub {
		color: var(--color-primary);
		font-size: 12px;
	}

</style>