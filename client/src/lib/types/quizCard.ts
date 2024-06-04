export type QuizOption = {
	id: string;
	text: string;
	isCorrect?: boolean;
};

export type QuizQuestion = {
	id: string;
	question: string;
	options: QuizOption[];
};
