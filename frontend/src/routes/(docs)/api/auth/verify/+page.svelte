<script lang="ts">
	import Code from "@components/Docs/Code.svelte";
    import Prism from "prismjs"
	import { endpoint, information, response, responses } from "./details";
	import Description from "@components/Docs/Description.svelte";
	import Response from "@components/Docs/Response.svelte";


</script>



<main>
    <section class="all">
        <h1 class="title">Verify</h1>
        <Description header={"Description"} information={information}/>
        <Code header={"Endpoint"} code={endpoint} prism={Prism.languages.http} language={"http"} />
        <Response responses={responses}/>
        <Code header={"Response"} code={response} prism={Prism.languages.http} language={"http"} />
    </section>

</main>


<style>

    h1.title {
        text-align: center;
    }

    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;

    }

    section.all {
        width: 50%;
        display: flex;
        flex-direction: column;
        gap: 10px;
    }


</style>