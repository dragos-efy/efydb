<script lang="ts">
	export let title: string;
	export let message: string;
	export let onClose: () => void;
	export let onConfirm: (() => void) | null = null;

	const onEvent = (event: Event) => {
		if (event.target !== event.currentTarget) return;
		onClose();
	};

	let onConfirmed = () => {
		if (onConfirm) onConfirm();
		onClose();
	};

	const handleKeydown = (event: any) => {
		if (event.key !== 'Escape') return;
		onClose();
	};
</script>

<svelte:window on:keydown={handleKeydown} />

<!-- svelte-ignore a11y-click-events-have-key-events -->
<section id="container" on:click={onEvent}>
	<div id="dialog">
		<h1>{title}</h1>
		<p>{message}</p>
		<div id="buttons">
			{#if onConfirm}
				<button on:click={onClose}>Cancel</button>
			{/if}
			<button on:click={onConfirmed}>Ok</button>
		</div>
	</div>
</section>

<style>
	#container {
		backdrop-filter: blur(20rem);
		position: absolute;
		height: 100vh;
		width: 100vw;
		max-width: 100%;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	#dialog {
		background: var(--efy_bg);
		border-radius: var(--efy_radius);
		border: var(--efy_border);
		padding: 30rem;
		display: flex;
		flex-direction: column;
		min-width: 25vw;
	}
	#buttons {
		align-self: flex-end;
		margin-top: 20rem;
	}
</style>
