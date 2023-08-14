<script>
	import { page } from '$app/stores';
</script>

<section>
	<h1>{$page.status}</h1>
	<p>{$page?.error?.message}</p>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		min-height: 70vh;
	}

	section h1 {
		font-size: 120rem;
	}

	section p {
		font-size: 35rem;
	}
</style>
