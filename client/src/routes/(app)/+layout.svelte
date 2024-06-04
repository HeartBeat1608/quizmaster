<script lang="ts">
	import Button from '$lib/ui/button.svelte';

	function handleNavigate(event: CustomEvent) {
		event.preventDefault();
		window.location.assign('/account');
	}
</script>

<main class="max-w-[1440px] mx-auto px-6">
	<nav class="flex flex-row justify-between items-center py-4">
		<a class="flex gap-1" href="/quiz">
			<h1 class="text-xl text-indigo-500">Quiz</h1>
			<h2 class="text-lg">Master</h2>
		</a>

		<div>
			<Button variant="action" on:click={handleNavigate}>Account</Button>
		</div>
	</nav>

	<section class="mt-4">
		<slot />
	</section>
</main>
