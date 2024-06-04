<script lang="ts">
	import QuizCard from '$lib/components/quizCard.svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	let quizData = [...data.quizData, ...data.quizData];
</script>

<div class="space-y-8">
	<div>
		<h1 class="w-full text-center text-3xl font-normal text-gray-500">Quizzes</h1>
		<h3 class="w-full text-center text-lg font-normal text-gray-700">
			Follow your instincts and get on a quiz.
		</h3>
	</div>

	<div class="grid grid-cols-6 gap-3">
		{#each quizData as quiz}
			<QuizCard id={quiz.id} totalQuestions={quiz.totalQuestions} title={quiz.title} />
		{/each}
	</div>
</div>
