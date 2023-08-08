<Header />
<slot />
<Footer />
<script>
	import Footer from "../components/Footer.svelte";
	import Header from "../components/Header.svelte";
    import { onMount } from "svelte";
    onMount(() => {
        import ("../efy_config.css");
        import ("../../static/efy/efy.js");
    })
</script>