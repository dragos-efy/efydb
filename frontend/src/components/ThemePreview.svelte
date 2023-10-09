<script lang="ts">
	export let theme: any;

	const onClick = () => {
		window.location.href = '/theme?id=' + theme.ID;
	};
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="theme" efy_card on:click={onClick}>
	<img src={theme.screenshot} alt={theme.name} loading="lazy" />
	<span class="theme-info">
		<p class="title">{theme.title}</p>
		<p class="username">@{theme.username}</p>
	</span>
</div>

<style>
	.theme {
		background: var(--efy_bg1);
		border-radius: var(--efy_radius);
		padding: 0;
		cursor: pointer;
	}
	.theme img {
		aspect-ratio: 16 / 9;
		object-fit: contain;
		border-radius: var(--efy_radius) var(--efy_radius) 0 0;
		border-bottom: var(--efy_border);
	}
	.theme-info {
		padding: 5rem 10rem 10rem 10rem;
		display: flex;
		flex-direction: column;
	}
	.theme-info p {
		margin: 0;
	}
	.theme .title {
		font-weight: bold;
	}
</style>
