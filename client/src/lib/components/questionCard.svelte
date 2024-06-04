<script lang="ts">
	import type { ButtonStyle } from '$lib/types/button';
	import type { QuizOption } from '$lib/types/quizCard';
	import Button from '$lib/ui/button.svelte';

	export let question: string;
	export let options: QuizOption[];

	$: selectedOption = '';

	function getSelectedStyleOption(option: QuizOption): ButtonStyle {
		return option.isCorrect ? 'success' : 'danger';
	}

	function handleChoice(option: QuizOption) {
		selectedOption = option.id;
	}
</script>

<section>
	<div class="bg-white/20 rounded-lg px-4 py-6 space-y-6">
		<h2 class="text-lg">{question}</h2>
		<div class="grid grid-cols-2 gap-4 w-full">
			{#each options as option}
				<Button
					class="w-full disabled:opacity-75"
					disabled={!!selectedOption}
					variant={selectedOption == option.id ? getSelectedStyleOption(option) : 'default'}
					on:click={() => handleChoice(option)}
				>
					{option.text}
				</Button>
			{/each}
		</div>
	</div>
</section>
