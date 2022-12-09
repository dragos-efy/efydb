<Header />
<slot />
<Footer />
<script>
	import Footer from "../components/Footer.svelte";
	import Header from "../components/Header.svelte";
    import { onMount } from "svelte";
    onMount(() => {
        console.log("import efy");
        import ("../../efy/efy.js");
    })
</script>