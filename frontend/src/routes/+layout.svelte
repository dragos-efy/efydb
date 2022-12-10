<Header />
<slot />
<Footer />
<script>
	import Footer from "../components/Footer.svelte";
	import Header from "../components/Header.svelte";
    import { onMount } from "svelte";
    onMount(() => {
        import ("../../efy/efy.js");
    })
</script>