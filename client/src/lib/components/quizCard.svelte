<script lang="ts">
	export let id: string;
	export let title: string;
	export let totalQuestions: number;

	function handleNav() {
		window.location.assign(`/quiz/${id}`);
	}
</script>

<button class="bg-white/20 px-6 py-4 rounded-md" on:click|preventDefault={handleNav}>
	<h2>{title}</h2>
	<p>{totalQuestions} questions</p>
</button>
