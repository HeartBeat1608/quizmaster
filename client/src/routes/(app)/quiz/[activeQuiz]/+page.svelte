<script lang="ts">
	import type { PageData } from './$types';
	import QuestionCard from '$lib/components/questionCard.svelte';
	import Button from '$lib/ui/button.svelte';

	export let data: PageData;

	$: activeQuestion = 0;
	$: actionQuestionObject = data.quizData.questions[activeQuestion];

	function handleChoice() {
		if (activeQuestion < data.quizData.questions.length - 1) activeQuestion += 1;
		else {
			alert('Quiz Completed');
			window.location.assign('/quiz');
		}
	}
</script>

<div class="space-y-6">
	<div class="bg-white/20 rounded-lg px-4 py-6">
		Active Quiz <b>{data.quizData.name}</b>.
	</div>

	<QuestionCard
		question={actionQuestionObject.question}
		options={actionQuestionObject.options}
		on:choice={handleChoice}
	/>

	<div class="w-full flex items-center justify-center gap-4">
		<Button variant="action" on:click={handleChoice}>Next Question</Button>
	</div>
</div>
