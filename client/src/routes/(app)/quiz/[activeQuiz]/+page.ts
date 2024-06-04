// import { api } from '$lib/api/base.js';
import type { QuizQuestion } from '$lib/types/quizCard';

async function loadQuizData(quizId: string) {
	// return api.get(`/quiz/${quizId}`);
	return {
		id: quizId,
		name: `Quiz ${quizId}`,
		questions: [
			{
				id: '1',
				question: 'What is the capital of France?',
				options: [
					{ id: '1', text: 'New York' },
					{ id: '2', text: 'London' },
					{ id: '3', text: 'Paris', isCorrect: true },
					{ id: '4', text: 'Dublin' }
				]
			}
		] as QuizQuestion[]
	};
}

export async function load({ params }) {
	const quizData = await loadQuizData(params.activeQuiz);

	return {
		activeQuiz: params.activeQuiz,
		quizData
	};
}
